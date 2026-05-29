package main

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

func TestParseStooqQuote(t *testing.T) {
	raw := []byte("Symbol,Date,Time,Name,Close\nAAPL.US,2026-05-28,22:00:19,APPLE INC,312.51\n")

	quote, err := parseStooqQuote(raw, "usaapl", "AAPL")
	if err != nil {
		t.Fatalf("parseStooqQuote returned error: %v", err)
	}
	if quote.Symbol != "usaapl" {
		t.Fatalf("Symbol = %q, want %q", quote.Symbol, "usaapl")
	}
	if quote.Name != "APPLE INC" {
		t.Fatalf("Name = %q, want %q", quote.Name, "APPLE INC")
	}
	if quote.Price != 312.51 {
		t.Fatalf("Price = %v, want %v", quote.Price, 312.51)
	}
}

func TestFetchUSQuotesUsesYahooFallback(t *testing.T) {
	provider := &EastMoneyQuoteProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if req.URL.Host == "stooq.com" {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader("Symbol,Date,Time,Name,Close\nAAPL.US,N/D,N/D,AAPL.US,N/D\n")),
					}, nil
				}
				if req.URL.Host == "query1.finance.yahoo.com" {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{"chart":{"result":[{"meta":{"symbol":"AAPL","regularMarketPrice":312.51}}],"error":null}}`)),
					}, nil
				}
				t.Fatalf("unexpected host: %s", req.URL.Host)
				return nil, nil
			}),
		},
	}

	quotes, err := provider.Fetch([]string{"usaapl"})
	if err != nil {
		t.Fatalf("Fetch returned error: %v", err)
	}
	if quotes["usaapl"].Price != 312.51 {
		t.Fatalf("Price = %v, want %v", quotes["usaapl"].Price, 312.51)
	}
}

func TestEastMoneyFetchFallsBackToNextBaseURL(t *testing.T) {
	provider := &EastMoneyQuoteProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if strings.Contains(req.URL.Path, "fail") {
					return &http.Response{
						StatusCode: http.StatusBadGateway,
						Body:       io.NopCloser(strings.NewReader("temporary failure")),
					}, nil
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"rc":0,"data":{"diff":[{"f2":154.9,"f12":"09992","f13":116,"f14":"泡泡玛特"}]}}`)),
				}, nil
			}),
		},
		baseURLs: []string{
			"https://example.invalid/fail",
			"https://example.invalid/ok",
		},
	}

	quotes, err := provider.Fetch([]string{"hk09992"})
	if err != nil {
		t.Fatalf("Fetch returned error: %v", err)
	}
	if quotes["hk09992"].Price != 154.9 {
		t.Fatalf("Price = %v, want %v", quotes["hk09992"].Price, 154.9)
	}
}

func TestEastMoneyFetchUsesTencentFallbackForZeroAStockPrice(t *testing.T) {
	provider := &EastMoneyQuoteProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if req.URL.Host == "qt.gtimg.cn" {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`v_sh513310="1~fallback name~513310~4.227~4.227~0.000~0~0~0~0.000";`)),
					}, nil
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"rc":0,"data":{"diff":[{"f2":0.0,"f12":"513310","f13":1,"f14":"中韩半导体ETF华泰柏瑞"}]}}`)),
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
	if quote.Name != "中韩半导体ETF华泰柏瑞" {
		t.Fatalf("Name = %q, want %q", quote.Name, "中韩半导体ETF华泰柏瑞")
	}
	if quote.Price != 4.227 {
		t.Fatalf("Price = %v, want %v", quote.Price, 4.227)
	}
	if quote.PriceDigits != 3 {
		t.Fatalf("PriceDigits = %v, want %v", quote.PriceDigits, 3)
	}
}

func TestEastMoneyFetchFallsBackToSingleRequestsWhenBatchFails(t *testing.T) {
	provider := &EastMoneyQuoteProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				secIDs := req.URL.Query().Get("secids")
				if strings.Contains(secIDs, ",") {
					return nil, errors.New("EOF")
				}
				body := `{"rc":0,"data":{"diff":[{"f2":158.5,"f12":"09992","f13":116,"f14":"泡泡玛特"}]}}`
				if secIDs == "1.513350" {
					body = `{"rc":0,"data":{"diff":[{"f2":1.278,"f12":"513350","f13":1,"f14":"标普油气ETF富国"}]}}`
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(body)),
				}, nil
			}),
		},
		baseURLs: []string{"https://push2.eastmoney.com/api/qt/ulist.np/get"},
	}

	quotes, err := provider.Fetch([]string{"hk09992", "sh513350"})
	if err != nil {
		t.Fatalf("Fetch returned error: %v", err)
	}
	if quotes["hk09992"].Price != 158.5 {
		t.Fatalf("hk09992 Price = %v, want %v", quotes["hk09992"].Price, 158.5)
	}
	if quotes["sh513350"].Price != 1.278 {
		t.Fatalf("sh513350 Price = %v, want %v", quotes["sh513350"].Price, 1.278)
	}
}

func TestEastMoneyFetchUsesAStockFallbackWhenEastMoneyUnavailable(t *testing.T) {
	provider := &EastMoneyQuoteProvider{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if req.URL.Host == "qt.gtimg.cn" {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`v_sh513350="1~oil gas etf~513350~1.278~1.280~0.000~0~0~0~0.000";`)),
					}, nil
				}
				return nil, errors.New("EOF")
			}),
		},
		baseURLs: []string{"https://push2.eastmoney.com/api/qt/ulist.np/get"},
	}

	quotes, err := provider.Fetch([]string{"hk09992", "sh513350"})
	if err != nil {
		t.Fatalf("Fetch returned error: %v", err)
	}
	if _, ok := quotes["hk09992"]; ok {
		t.Fatal("hk09992 should be left unchanged when EastMoney is unavailable")
	}
	if quotes["sh513350"].Price != 1.278 {
		t.Fatalf("sh513350 Price = %v, want %v", quotes["sh513350"].Price, 1.278)
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
