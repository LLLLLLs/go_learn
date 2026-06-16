package stockcalculator

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestToEastMoneySecID(t *testing.T) {
	tests := map[string]string{
		"sh513310": "1.513310",
		"sz000001": "0.000001",
		"bj920001": "0.920001",
		"hk00700":  "116.00700",
		"hk01022":  "116.01022",
		"hk09992":  "116.09992",
	}

	for input, want := range tests {
		got, err := toEastMoneySecID(input)
		if err != nil {
			t.Fatalf("toEastMoneySecID(%q) returned error: %v", input, err)
		}
		if got != want {
			t.Fatalf("toEastMoneySecID(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestParseEastMoneyQuotes(t *testing.T) {
	raw := []byte(`{"rc":0,"data":{"diff":[{"f2":154.9,"f18":152.3,"f12":"09992","f13":116,"f14":"泡泡玛特"},{"f2":4.122,"f18":4.08,"f12":"513310","f13":1,"f14":"中韩半导体ETF华泰柏瑞"}]}}`)

	quotes, err := parseEastMoneyQuotes(raw, map[string]string{
		"116.09992": "hk09992",
		"1.513310":  "sh513310",
	})
	if err != nil {
		t.Fatalf("parseEastMoneyQuotes returned error: %v", err)
	}

	hkQuote := quotes["hk09992"]
	if hkQuote.Name != "泡泡玛特" {
		t.Fatalf("Name = %q, want %q", hkQuote.Name, "泡泡玛特")
	}
	if hkQuote.Price != 154.9 {
		t.Fatalf("Price = %v, want %v", hkQuote.Price, 154.9)
	}
	if hkQuote.OpenPrice != 152.3 {
		t.Fatalf("OpenPrice = %v, want previous close %v", hkQuote.OpenPrice, 152.3)
	}

	etfQuote := quotes["sh513310"]
	if etfQuote.Price != 4.122 {
		t.Fatalf("Price = %v, want %v", etfQuote.Price, 4.122)
	}
	if etfQuote.PriceDigits != 3 {
		t.Fatalf("PriceDigits = %v, want %v", etfQuote.PriceDigits, 3)
	}
}

func TestEastMoneyUSSecIDCandidates(t *testing.T) {
	candidates, err := eastMoneyUSSecIDCandidates("usmrvl")
	if err != nil {
		t.Fatalf("eastMoneyUSSecIDCandidates returned error: %v", err)
	}
	want := []string{"105.MRVL", "106.MRVL", "107.MRVL"}
	if strings.Join(candidates, ",") != strings.Join(want, ",") {
		t.Fatalf("candidates = %v, want %v", candidates, want)
	}
}

func TestToYahooSymbol(t *testing.T) {
	tests := map[string]string{
		"usmrvl":   "MRVL",
		"usbrk.b":  "BRK-B",
		"hk09992":  "09992.HK",
		"sh513310": "513310.SS",
		"sz160644": "160644.SZ",
		"ks005930": "005930.KS",
		"kq035720": "035720.KQ",
	}
	for input, want := range tests {
		got, err := toYahooSymbol(input)
		if err != nil {
			t.Fatalf("toYahooSymbol(%q) returned error: %v", input, err)
		}
		if got != want {
			t.Fatalf("toYahooSymbol(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestToTencentSymbol(t *testing.T) {
	tests := map[string]string{
		"usmrvl":   "usMRVL",
		"usbrk.b":  "usBRK.B",
		"hk09992":  "hk09992",
		"sh513310": "sh513310",
		"sz160644": "sz160644",
	}
	for input, want := range tests {
		got, err := toTencentSymbol(input)
		if err != nil {
			t.Fatalf("toTencentSymbol(%q) returned error: %v", input, err)
		}
		if got != want {
			t.Fatalf("toTencentSymbol(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestFetchUsesTencentFirstWhenSuccessful(t *testing.T) {
	var hosts []string
	provider := &EastMoneyQuoteProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				hosts = append(hosts, req.URL.Host)
				switch req.URL.Host {
				case "qt.gtimg.cn":
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`v_usMRVL="200~Marvell~MRVL.OQ~290.790~288.120~291.000";`)),
					}, nil
				case "query1.finance.yahoo.com":
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(yahooUSChartBody("MRVL", 290.79, 288.12, 291.2, 292.3))),
					}, nil
				default:
					t.Fatalf("unexpected request before tencent success: %s", req.URL.String())
					return nil, nil
				}
			}),
		},
		baseURLs: []string{"https://push2.eastmoney.com/api/qt/ulist.np/get"},
	}

	quotes, err := provider.Fetch([]string{"usmrvl"})
	if err != nil {
		t.Fatalf("Fetch returned error: %v", err)
	}
	quote := quotes["usmrvl"]
	if quote.Price != 290.79 {
		t.Fatalf("Price = %v, want %v", quote.Price, 290.79)
	}
	if quote.OpenPrice != 288.12 {
		t.Fatalf("OpenPrice = %v, want previous close %v", quote.OpenPrice, 288.12)
	}
	if quote.ExtendedPrice != 292.3 {
		t.Fatalf("ExtendedPrice = %v, want Yahoo latest extended price", quote.ExtendedPrice)
	}
	if strings.Join(hosts, ",") != "qt.gtimg.cn,query1.finance.yahoo.com" {
		t.Fatalf("hosts = %v, want tencent plus yahoo supplemental", hosts)
	}
}

func TestFetchFallsBackToEastMoneyForAStockAndUSAfterTencentFailure(t *testing.T) {
	var eastMoneySecIDs string
	provider := &EastMoneyQuoteProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if req.URL.Host == "qt.gtimg.cn" {
					return nil, errors.New("EOF")
				}
				if req.URL.Host == "push2.eastmoney.com" {
					eastMoneySecIDs = req.URL.Query().Get("secids")
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{"rc":0,"data":{"diff":[{"f2":290.79,"f12":"MRVL","f13":105,"f14":"迈威尔科技"},{"f2":4.227,"f12":"513310","f13":1,"f14":"中韩半导体ETF华泰柏瑞"}]}}`)),
					}, nil
				}
				if req.URL.Host == "query1.finance.yahoo.com" {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(yahooUSChartBody("MRVL", 290.79, 288.12, 291.2, 292.3))),
					}, nil
				}
				t.Fatalf("unexpected request: %s", req.URL.String())
				return nil, nil
			}),
		},
		baseURLs: []string{"https://push2.eastmoney.com/api/qt/ulist.np/get"},
	}

	quotes, err := provider.Fetch([]string{"usmrvl", "sh513310"})
	if err != nil {
		t.Fatalf("Fetch returned error: %v", err)
	}
	if quotes["usmrvl"].Price != 290.79 {
		t.Fatalf("usmrvl Price = %v, want %v", quotes["usmrvl"].Price, 290.79)
	}
	if quotes["usmrvl"].ExtendedPrice != 292.3 {
		t.Fatalf("usmrvl ExtendedPrice = %v, want %v", quotes["usmrvl"].ExtendedPrice, 292.3)
	}
	if quotes["sh513310"].Price != 4.227 {
		t.Fatalf("sh513310 Price = %v, want %v", quotes["sh513310"].Price, 4.227)
	}
	if !strings.Contains(eastMoneySecIDs, "105.MRVL") || !strings.Contains(eastMoneySecIDs, "1.513310") {
		t.Fatalf("eastmoney secids = %q, want US and A-stock requests", eastMoneySecIDs)
	}
}

func TestEastMoneyBatchDoesNotTryNextBaseURL(t *testing.T) {
	var urls []string
	provider := &EastMoneyQuoteProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				urls = append(urls, req.URL.String())
				if req.URL.Host == "query1.finance.yahoo.com" {
					t.Fatalf("unexpected yahoo request for HK fallback")
				}
				if req.URL.Host == "qt.gtimg.cn" {
					return nil, errors.New("EOF")
				}
				return &http.Response{
					StatusCode: http.StatusBadGateway,
					Body:       io.NopCloser(strings.NewReader("temporary failure")),
				}, nil
			}),
		},
		baseURLs: []string{
			"https://example.invalid/fail",
			"https://example.invalid/ok",
		},
	}

	quotes, err := provider.Fetch([]string{"hk09992"})
	if err == nil {
		t.Fatalf("Fetch returned nil error with quotes=%v", quotes)
	}
	if strings.Contains(strings.Join(urls, ","), "/ok") {
		t.Fatalf("urls = %v, should not try next EastMoney base URL", urls)
	}
}

func TestFetchUsesEastMoneyWhenTencentReturnsNoPrice(t *testing.T) {
	provider := &EastMoneyQuoteProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if req.URL.Host == "query1.finance.yahoo.com" {
					t.Fatalf("unexpected yahoo request before eastmoney fallback")
				}
				if req.URL.Host == "qt.gtimg.cn" {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`v_sh513310="";`)),
					}, nil
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"rc":0,"data":{"diff":[{"f2":4.227,"f12":"513310","f13":1,"f14":"中韩半导体ETF华泰柏瑞"}]}}`)),
				}, nil
			}),
		},
		baseURLs: []string{"https://push2.eastmoney.com/api/qt/ulist.np/get"},
	}

	quotes, err := provider.Fetch([]string{"sh513310"})
	if err != nil {
		t.Fatalf("Fetch returned error: %v", err)
	}
	quote := quotes["sh513310"]
	if quote.Price != 4.227 {
		t.Fatalf("Price = %v, want %v", quote.Price, 4.227)
	}
	if quote.PriceDigits != 3 {
		t.Fatalf("PriceDigits = %v, want %v", quote.PriceDigits, 3)
	}
}

func TestFetchUsesTencentForHKBeforeEastMoney(t *testing.T) {
	provider := &EastMoneyQuoteProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if req.URL.Host == "query1.finance.yahoo.com" {
					t.Fatalf("unexpected yahoo request before tencent hk success")
				}
				if req.URL.Host == "qt.gtimg.cn" {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`v_hk09992="100~POP MART~09992~174.400~179.200~178.100~11022500.0~0~0~174.400";`)),
					}, nil
				}
				t.Fatalf("unexpected request before tencent hk success: %s", req.URL.String())
				return nil, nil
			}),
		},
		baseURLs: []string{"https://push2.eastmoney.com/api/qt/ulist.np/get"},
	}

	quotes, err := provider.Fetch([]string{"hk09992"})
	if err != nil {
		t.Fatalf("Fetch returned error: %v", err)
	}
	quote := quotes["hk09992"]
	if quote.Price != 174.4 {
		t.Fatalf("Price = %v, want %v", quote.Price, 174.4)
	}
	if quote.Name != "POP MART" {
		t.Fatalf("Name = %q, want %q", quote.Name, "POP MART")
	}
}

func TestFetchUsesTencentByMarketBeforeEastMoney(t *testing.T) {
	provider := &EastMoneyQuoteProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if req.URL.Host == "qt.gtimg.cn" {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`v_hk09992="100~POP MART~09992~158.500~0~0";v_sh513350="1~标普油气ETF富国~513350~1.278~0~0";v_usMRVL="200~Marvell~MRVL.OQ~290.790~0~0";`)),
					}, nil
				}
				if req.URL.Host == "query1.finance.yahoo.com" {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(yahooUSChartBody("MRVL", 290.79, 288.12, 291.2, 292.3))),
					}, nil
				}
				t.Fatalf("unexpected request before tencent success: %s", req.URL.String())
				return nil, nil
			}),
		},
		baseURLs: []string{"https://push2.eastmoney.com/api/qt/ulist.np/get"},
	}

	quotes, err := provider.Fetch([]string{"hk09992", "sh513350", "usmrvl"})
	if err != nil {
		t.Fatalf("Fetch returned error: %v", err)
	}
	if quotes["hk09992"].Price != 158.5 {
		t.Fatalf("hk09992 Price = %v, want %v", quotes["hk09992"].Price, 158.5)
	}
	if quotes["sh513350"].Price != 1.278 {
		t.Fatalf("sh513350 Price = %v, want %v", quotes["sh513350"].Price, 1.278)
	}
	if quotes["usmrvl"].Price != 290.79 {
		t.Fatalf("usmrvl Price = %v, want %v", quotes["usmrvl"].Price, 290.79)
	}
	if quotes["usmrvl"].ExtendedPrice != 292.3 {
		t.Fatalf("usmrvl extended = %+v", quotes["usmrvl"])
	}
}

func TestFetchFallsBackToYahooForMissingTencentAStockAndUS(t *testing.T) {
	var yahooPaths []string
	provider := &EastMoneyQuoteProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if req.URL.Host == "qt.gtimg.cn" {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`v_sh513310="";v_usMRVL="";`)),
					}, nil
				}
				if req.URL.Host == "query1.finance.yahoo.com" {
					yahooPaths = append(yahooPaths, req.URL.Path)
					body := `{"chart":{"result":[{"meta":{"symbol":"MRVL","regularMarketPrice":290.79}}],"error":null}}`
					if strings.Contains(req.URL.Path, "/chart/513310.SS") {
						body = `{"chart":{"result":[{"meta":{"symbol":"513310.SS","regularMarketPrice":4.227}}],"error":null}}`
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}, nil
				}
				return nil, errors.New("EOF")
			}),
		},
		baseURLs: []string{"https://push2.eastmoney.com/api/qt/ulist.np/get"},
	}

	quotes, err := provider.Fetch([]string{"usmrvl", "sh513310"})
	if err != nil {
		t.Fatalf("Fetch returned error: %v", err)
	}
	if quotes["usmrvl"].Price != 290.79 {
		t.Fatalf("usmrvl Price = %v, want %v", quotes["usmrvl"].Price, 290.79)
	}
	if quotes["sh513310"].Price != 4.227 {
		t.Fatalf("sh513310 Price = %v, want %v", quotes["sh513310"].Price, 4.227)
	}
	joined := strings.Join(yahooPaths, ",")
	if !strings.Contains(joined, "/chart/MRVL") || !strings.Contains(joined, "/chart/513310.SS") {
		t.Fatalf("yahoo paths = %v, want US and A-stock requests", yahooPaths)
	}
}

func TestFetchFallsBackToYahooForKoreanStocks(t *testing.T) {
	var yahooPath string
	provider := &EastMoneyQuoteProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if req.URL.Host != "query1.finance.yahoo.com" {
					t.Fatalf("unexpected non-yahoo request for korean stock: %s", req.URL.String())
				}
				yahooPath = req.URL.Path
				body := `{"chart":{"result":[{"meta":{"symbol":"000660.KS","longName":"SK hynix Inc.","shortName":"SK hynix","regularMarketPrice":70000,"regularMarketOpen":69000,"chartPreviousClose":68000}}],"error":null}}`
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(body)),
				}, nil
			}),
		},
		baseURLs: []string{"https://push2.eastmoney.com/api/qt/ulist.np/get"},
	}

	quotes, err := provider.Fetch([]string{"ks000660"})
	if err != nil {
		t.Fatalf("Fetch returned error: %v", err)
	}
	quote := quotes["ks000660"]
	if quote.Price != 70000 || quote.OpenPrice != 68000 {
		t.Fatalf("quote = %+v", quote)
	}
	if quote.Name != "SK海力士" {
		t.Fatalf("Name = %q, want localized Korean name", quote.Name)
	}
	if yahooPath != "/v8/finance/chart/000660.KS" {
		t.Fatalf("yahooPath = %q, want korean yahoo symbol", yahooPath)
	}
}

func TestFetchYahooDoesNotUseCurrentPriceAsBaseDuringExtendedSession(t *testing.T) {
	provider := &EastMoneyQuoteProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if req.URL.Host != "query1.finance.yahoo.com" {
					t.Fatalf("unexpected non-yahoo request: %s", req.URL.String())
				}
				body := `{"chart":{"result":[{"meta":{"symbol":"AAOI","regularMarketPrice":191.55,"regularMarketPreviousClose":169.05,"chartPreviousClose":196.64,"previousClose":168.9,"currentTradingPeriod":{"pre":{"start":8000,"end":9999},"regular":{"start":10000,"end":20000},"post":{"start":20001,"end":24000}}},"timestamp":[8010,8020],"indicators":{"quote":[{"close":[190.7,190.8]}]}}],"error":null}}`
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(body)),
				}, nil
			}),
		},
	}

	quotes, err := provider.fetchYahooQuotes([]string{"usaaoi"})
	if err != nil {
		t.Fatalf("fetchYahooQuotes returned error: %v", err)
	}
	if quotes["usaaoi"].OpenPrice != 169.05 {
		t.Fatalf("OpenPrice = %v, want regularMarketPreviousClose 169.05", quotes["usaaoi"].OpenPrice)
	}
	if quotes["usaaoi"].ExtendedPrice != 190.8 {
		t.Fatalf("ExtendedPrice = %v, want pre-market price", quotes["usaaoi"].ExtendedPrice)
	}
}

func TestYahooDisplayNameFallsBackToLongName(t *testing.T) {
	name := yahooDisplayName("ks123456", "Example Corp.", "Example")
	if name != "Example Corp." {
		t.Fatalf("name = %q, want longName fallback", name)
	}
}

func TestFetchKRWCNYUsesKoreanWonBase(t *testing.T) {
	var gotBase string
	provider := &FrankfurterExchangeRateProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				gotBase = req.URL.Query().Get("base")
				if symbols := req.URL.Query().Get("symbols"); symbols != "CNY" {
					t.Fatalf("symbols = %q, want CNY", symbols)
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"rates":{"CNY":0.0052}}`)),
				}, nil
			}),
		},
	}

	rate, err := provider.FetchKRWCNY()
	if err != nil {
		t.Fatalf("FetchKRWCNY returned error: %v", err)
	}
	if gotBase != "KRW" || rate != 0.0052 {
		t.Fatalf("base/rate = %q/%v, want KRW/0.0052", gotBase, rate)
	}
}

func TestLatestExtendedMarketPriceKeepsLastCloseBeforeNextPreMarket(t *testing.T) {
	periods := yahooTradingPeriods{
		Pre:     yahooPeriod{Start: 8000, End: 9999},
		Regular: yahooPeriod{Start: 10000, End: 20000},
		Post:    yahooPeriod{Start: 20001, End: 24000},
	}
	timestamps := []int64{20010, 23990, 25000}
	prices := []float64{292.1, 292.3, 0}

	price := latestExtendedMarketPrice(timestamps, prices, periods)
	if price != 292.3 {
		t.Fatalf("latestExtendedMarketPrice = %v, want last after-hours close", price)
	}
}

func TestLatestExtendedMarketPriceReturnsEmptyDuringRegularSession(t *testing.T) {
	periods := yahooTradingPeriods{
		Pre:     yahooPeriod{Start: 8000, End: 9999},
		Regular: yahooPeriod{Start: 10000, End: 20000},
		Post:    yahooPeriod{Start: 20001, End: 24000},
	}
	timestamps := []int64{8010, 10010, 19990}
	prices := []float64{291.2, 290.8, 290.79}

	price := latestExtendedMarketPrice(timestamps, prices, periods)
	if price != 0 {
		t.Fatalf("latestExtendedMarketPrice = %v, want empty during regular session", price)
	}
}

func yahooUSChartBody(symbol string, price, previousClose, preMarket, postMarket float64) string {
	return fmt.Sprintf(`{"chart":{"result":[{"meta":{"symbol":%q,"regularMarketPrice":%v,"chartPreviousClose":%v,"currentTradingPeriod":{"pre":{"start":8000,"end":9999},"regular":{"start":10000,"end":20000},"post":{"start":20001,"end":24000}}},"timestamp":[8010,19990,20010,23990,25000],"indicators":{"quote":[{"close":[%v,%v,%v,%v,null]}]}}],"error":null}}`, symbol, price, previousClose, preMarket, price, postMarket-0.1, postMarket)
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
