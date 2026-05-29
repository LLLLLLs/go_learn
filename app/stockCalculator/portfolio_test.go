package main

import "testing"

func TestNormalizeSecuritySupportsUSStocks(t *testing.T) {
	tests := []struct {
		input       string
		displayCode string
		symbol      string
		market      string
		currency    string
	}{
		{"AAPL", "US:AAPL", "usaapl", "美股", "USD"},
		{"US:MSFT", "US:MSFT", "usmsft", "美股", "USD"},
		{"brk-b", "US:BRK.B", "usbrk.b", "美股", "USD"},
	}

	for _, tt := range tests {
		got, err := normalizeSecurity(tt.input)
		if err != nil {
			t.Fatalf("normalizeSecurity(%q) returned error: %v", tt.input, err)
		}
		if got.DisplayCode != tt.displayCode || got.Symbol != tt.symbol || got.Market != tt.market || got.Currency != tt.currency {
			t.Fatalf("normalizeSecurity(%q) = %+v", tt.input, got)
		}
	}
}

func TestConvertPriceToMarketCurrencySupportsUSD(t *testing.T) {
	price, err := convertPriceToMarketCurrency(720, "RMB", "美元", 0.9, 7.2)
	if err != nil {
		t.Fatalf("convertPriceToMarketCurrency returned error: %v", err)
	}
	if price != 100 {
		t.Fatalf("converted price = %v, want %v", price, 100.0)
	}

	price, err = convertPriceToMarketCurrency(10, "美元", "港币", 0.9, 7.2)
	if err != nil {
		t.Fatalf("convertPriceToMarketCurrency returned error: %v", err)
	}
	if price != 80 {
		t.Fatalf("converted price = %v, want %v", price, 80.0)
	}
}

func TestPortfolioUSDBalanceAndUSStockSummary(t *testing.T) {
	p, err := LoadPortfolio("")
	if err != nil {
		t.Fatalf("LoadPortfolio returned error: %v", err)
	}
	if err := p.SetUSDExchangeRate(7.2); err != nil {
		t.Fatalf("SetUSDExchangeRate returned error: %v", err)
	}
	if err := p.AdjustBalanceByCurrency(1000, "美元", 1); err != nil {
		t.Fatalf("AdjustBalanceByCurrency returned error: %v", err)
	}
	if err := p.AddRecord(TradeTypeBuy, "AAPL", 2, 100, "美元"); err != nil {
		t.Fatalf("AddRecord returned error: %v", err)
	}
	if err := p.RefreshQuotes(staticQuoteProvider{
		"usaapl": {Symbol: "usaapl", Name: "APPLE INC", Price: 100, PriceDigits: 2},
	}); err != nil {
		t.Fatalf("RefreshQuotes returned error: %v", err)
	}

	summary := p.Summary()
	if summary.InitialCapitalUSD != 1000 {
		t.Fatalf("InitialCapitalUSD = %v, want %v", summary.InitialCapitalUSD, 1000.0)
	}
	if summary.CashUSD != 800 {
		t.Fatalf("CashUSD = %v, want %v", summary.CashUSD, 800.0)
	}
	if summary.Cash != 5760 {
		t.Fatalf("Cash = %v, want %v", summary.Cash, 5760.0)
	}
	if len(summary.Holdings) != 1 {
		t.Fatalf("len(Holdings) = %d, want 1", len(summary.Holdings))
	}
	holding := summary.Holdings[0]
	if holding.DisplayCode != "US:AAPL" || holding.Currency != "USD" || holding.MarketValue != 1440 {
		t.Fatalf("Holding = %+v", holding)
	}
}

type staticQuoteProvider map[string]Quote

func (p staticQuoteProvider) Fetch(symbols []string) (map[string]Quote, error) {
	quotes := make(map[string]Quote, len(symbols))
	for _, symbol := range symbols {
		if quote, ok := p[symbol]; ok {
			quotes[symbol] = quote
		}
	}
	return quotes, nil
}
