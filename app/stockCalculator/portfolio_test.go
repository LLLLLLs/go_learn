package main

import (
	"encoding/json"
	"testing"
	"time"
)

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

func TestPortfolioPersistsAndLoadsQuoteCache(t *testing.T) {
	p, err := LoadPortfolio("")
	if err != nil {
		t.Fatalf("LoadPortfolio returned error: %v", err)
	}
	if err := p.AdjustBalanceByCurrency(1000, "美元", 1); err != nil {
		t.Fatalf("AdjustBalanceByCurrency returned error: %v", err)
	}
	if err := p.AddRecord(TradeTypeBuy, "AAPL", 2, 100, "美元"); err != nil {
		t.Fatalf("AddRecord returned error: %v", err)
	}
	if err := p.RefreshQuotes(staticQuoteProvider{
		"usaapl": {Symbol: "usaapl", Name: "APPLE INC", Price: 312.51, PriceDigits: 2},
	}); err != nil {
		t.Fatalf("RefreshQuotes returned error: %v", err)
	}

	payload, err := p.MarshalState()
	if err != nil {
		t.Fatalf("MarshalState returned error: %v", err)
	}
	var state persistedState
	if err := json.Unmarshal(payload, &state); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if state.QuoteCache["usaapl"].Price != 312.51 {
		t.Fatalf("QuoteCache price = %v, want %v", state.QuoteCache["usaapl"].Price, 312.51)
	}

	loaded, err := LoadPortfolio(string(payload))
	if err != nil {
		t.Fatalf("LoadPortfolio returned error: %v", err)
	}
	holding := loaded.Summary().Holdings[0]
	if holding.Name != "APPLE INC" || holding.CurrentPrice != 312.51 {
		t.Fatalf("Loaded holding = %+v", holding)
	}
}

func TestPortfolioKeepsCachedQuoteAfterRebuild(t *testing.T) {
	p, err := LoadPortfolio("")
	if err != nil {
		t.Fatalf("LoadPortfolio returned error: %v", err)
	}
	if err := p.AdjustBalanceByCurrency(1000, "美元", 1); err != nil {
		t.Fatalf("AdjustBalanceByCurrency returned error: %v", err)
	}
	if err := p.AddRecord(TradeTypeBuy, "AAPL", 1, 100, "美元"); err != nil {
		t.Fatalf("AddRecord returned error: %v", err)
	}
	if err := p.RefreshQuotes(staticQuoteProvider{
		"usaapl": {Symbol: "usaapl", Name: "APPLE INC", Price: 312.51, PriceDigits: 2},
	}); err != nil {
		t.Fatalf("RefreshQuotes returned error: %v", err)
	}
	if err := p.AddRecord(TradeTypeBuy, "AAPL", 1, 120, "美元"); err != nil {
		t.Fatalf("AddRecord returned error: %v", err)
	}

	holding := p.Summary().Holdings[0]
	if holding.CurrentPrice != 312.51 {
		t.Fatalf("CurrentPrice = %v, want %v", holding.CurrentPrice, 312.51)
	}
}

func TestPortfolioHasFreshQuoteCache(t *testing.T) {
	now := time.Date(2026, 6, 3, 11, 40, 0, 0, time.Local)
	state := persistedState{
		InitialCapitalUSD: 1000,
		USDRate:           7.2,
		Records: []TradeRecord{{
			ID:             "1",
			CreatedAt:      now.Add(-time.Hour),
			Type:           TradeTypeBuy,
			Code:           "AAPL",
			DisplayCode:    "US:AAPL",
			Symbol:         "usaapl",
			Market:         "美股",
			Currency:       "USD",
			MarketCurrency: "USD",
			Quantity:       1,
			Price:          100,
			MarketPrice:    100,
			FXRate:         7.2,
		}},
		QuoteCache: quoteCache{
			"usaapl": {Name: "APPLE INC", Price: 312.51, PriceDigits: 2, UpdatedAt: now.Add(-30 * time.Second)},
		},
	}
	payload, err := json.Marshal(state)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	p, err := LoadPortfolio(string(payload))
	if err != nil {
		t.Fatalf("LoadPortfolio returned error: %v", err)
	}
	if !p.HasFreshQuoteCache(time.Minute, now) {
		t.Fatal("HasFreshQuoteCache = false, want true")
	}
	if p.HasFreshQuoteCache(time.Minute, now.Add(31*time.Second)) {
		t.Fatal("HasFreshQuoteCache = true for stale cache, want false")
	}
}

func TestPortfolioLoadsCachedLastRefreshAt(t *testing.T) {
	updatedAt := time.Date(2026, 6, 3, 12, 1, 20, 0, time.Local)
	state := persistedState{
		InitialCapitalUSD: 1000,
		USDRate:           7.2,
		Records: []TradeRecord{{
			ID:             "1",
			CreatedAt:      updatedAt.Add(-time.Hour),
			Type:           TradeTypeBuy,
			Code:           "AAPL",
			DisplayCode:    "US:AAPL",
			Symbol:         "usaapl",
			Market:         "美股",
			Currency:       "USD",
			MarketCurrency: "USD",
			Quantity:       1,
			Price:          100,
			MarketPrice:    100,
			FXRate:         7.2,
		}},
		QuoteCache: quoteCache{
			"usaapl": {Name: "APPLE INC", Price: 312.51, PriceDigits: 2, UpdatedAt: updatedAt},
		},
	}
	payload, err := json.Marshal(state)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	p, err := LoadPortfolio(string(payload))
	if err != nil {
		t.Fatalf("LoadPortfolio returned error: %v", err)
	}
	if got := p.Summary().LastRefreshAt; !got.Equal(updatedAt) {
		t.Fatalf("LastRefreshAt = %s, want %s", got, updatedAt)
	}
}

func TestPortfolioNextQuoteRefreshFromCache(t *testing.T) {
	updatedAt := time.Date(2026, 6, 3, 12, 1, 20, 0, time.Local)
	now := time.Date(2026, 6, 3, 12, 1, 50, 0, time.Local)
	state := persistedState{
		InitialCapitalUSD: 1000,
		USDRate:           7.2,
		Records: []TradeRecord{{
			ID:             "1",
			CreatedAt:      updatedAt.Add(-time.Hour),
			Type:           TradeTypeBuy,
			Code:           "AAPL",
			DisplayCode:    "US:AAPL",
			Symbol:         "usaapl",
			Market:         "美股",
			Currency:       "USD",
			MarketCurrency: "USD",
			Quantity:       1,
			Price:          100,
			MarketPrice:    100,
			FXRate:         7.2,
		}},
		QuoteCache: quoteCache{
			"usaapl": {Name: "APPLE INC", Price: 312.51, PriceDigits: 2, UpdatedAt: updatedAt},
		},
	}
	payload, err := json.Marshal(state)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	p, err := LoadPortfolio(string(payload))
	if err != nil {
		t.Fatalf("LoadPortfolio returned error: %v", err)
	}

	next, ok := p.NextQuoteRefreshFromCache(time.Minute, now)
	if !ok {
		t.Fatal("NextQuoteRefreshFromCache ok = false, want true")
	}
	want := time.Date(2026, 6, 3, 12, 2, 20, 0, time.Local)
	if !next.Equal(want) {
		t.Fatalf("next = %s, want %s", next, want)
	}
	if remaining := timeUntil(next, now); remaining != 30*time.Second {
		t.Fatalf("remaining = %s, want 30s", remaining)
	}
}

func TestPortfolioPersistsAndLoadsFXCache(t *testing.T) {
	updatedAt := time.Date(2026, 6, 3, 12, 1, 20, 0, time.Local)
	p, err := LoadPortfolio("")
	if err != nil {
		t.Fatalf("LoadPortfolio returned error: %v", err)
	}
	if err := p.SetExchangeRateWithTime(0.93, updatedAt); err != nil {
		t.Fatalf("SetExchangeRateWithTime returned error: %v", err)
	}
	if err := p.SetUSDExchangeRateWithTime(7.25, updatedAt.Add(10*time.Second)); err != nil {
		t.Fatalf("SetUSDExchangeRateWithTime returned error: %v", err)
	}

	payload, err := p.MarshalState()
	if err != nil {
		t.Fatalf("MarshalState returned error: %v", err)
	}
	var state persistedState
	if err := json.Unmarshal(payload, &state); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if state.FXCache.HKDCNY.Rate != 0.93 || !state.FXCache.HKDCNY.UpdatedAt.Equal(updatedAt) {
		t.Fatalf("HKD cache = %+v", state.FXCache.HKDCNY)
	}
	if state.FXCache.USDCNY.Rate != 7.25 || !state.FXCache.USDCNY.UpdatedAt.Equal(updatedAt.Add(10*time.Second)) {
		t.Fatalf("USD cache = %+v", state.FXCache.USDCNY)
	}

	loaded, err := LoadPortfolio(string(payload))
	if err != nil {
		t.Fatalf("LoadPortfolio returned error: %v", err)
	}
	summary := loaded.Summary()
	if summary.HKDRate != 0.93 || summary.USDRate != 7.25 {
		t.Fatalf("rates = %.4f/%.4f, want 0.93/7.25", summary.HKDRate, summary.USDRate)
	}
}

func TestPortfolioNextFXRefreshFromCache(t *testing.T) {
	updatedAt := time.Date(2026, 6, 3, 12, 1, 20, 0, time.Local)
	now := time.Date(2026, 6, 3, 12, 1, 50, 0, time.Local)
	state := persistedState{
		FXCache: fxCache{
			HKDCNY: persistedFXRate{Rate: 0.93, UpdatedAt: updatedAt},
			USDCNY: persistedFXRate{Rate: 7.25, UpdatedAt: updatedAt.Add(10 * time.Second)},
		},
	}
	payload, err := json.Marshal(state)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	p, err := LoadPortfolio(string(payload))
	if err != nil {
		t.Fatalf("LoadPortfolio returned error: %v", err)
	}

	next, ok := p.NextFXRefreshFromCache(time.Minute, now)
	if !ok {
		t.Fatal("NextFXRefreshFromCache ok = false, want true")
	}
	want := time.Date(2026, 6, 3, 12, 2, 20, 0, time.Local)
	if !next.Equal(want) {
		t.Fatalf("next = %s, want %s", next, want)
	}
	if _, ok := p.NextFXRefreshFromCache(time.Minute, now.Add(31*time.Second)); ok {
		t.Fatal("NextFXRefreshFromCache ok = true for stale cache, want false")
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
