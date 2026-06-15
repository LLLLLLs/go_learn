package stockcalculator

import (
	"errors"
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
	raw := []byte(`{"rc":0,"data":{"diff":[{"f2":154.9,"f12":"09992","f13":116,"f14":"泡泡玛特"},{"f2":4.122,"f12":"513310","f13":1,"f14":"中韩半导体ETF华泰柏瑞"}]}}`)

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
				if req.URL.Host != "qt.gtimg.cn" {
					t.Fatalf("unexpected request before tencent success: %s", req.URL.String())
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`v_usMRVL="200~Marvell~MRVL.OQ~290.790~0~0";`)),
				}, nil
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
	if strings.Join(hosts, ",") != "qt.gtimg.cn" {
		t.Fatalf("hosts = %v, want tencent only", hosts)
	}
}

func TestFetchFallsBackToEastMoneyForAStockAndUSAfterTencentFailure(t *testing.T) {
	var eastMoneySecIDs string
	provider := &EastMoneyQuoteProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if req.URL.Host == "query1.finance.yahoo.com" {
					t.Fatalf("unexpected yahoo request before eastmoney fallback")
				}
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

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
