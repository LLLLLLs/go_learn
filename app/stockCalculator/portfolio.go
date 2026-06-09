package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	statePreferenceKey = "portfolio_state_v1"
	defaultHKDRate     = 0.92
	defaultUSDRate     = 7.20
	defaultRefreshSec  = 60
	moneyEpsilon       = 1e-6
)

type TradeType string

const (
	TradeTypeBuy  TradeType = "buy"
	TradeTypeSell TradeType = "sell"
)

type TradeRecord struct {
	ID             string       `json:"id"`
	CreatedAt      time.Time    `json:"created_at"`
	Type           TradeType    `json:"type"`
	Code           string       `json:"code"`
	DisplayCode    string       `json:"display_code"`
	Symbol         string       `json:"symbol"`
	Market         string       `json:"market"`
	Currency       string       `json:"currency"`
	MarketCurrency string       `json:"market_currency,omitempty"`
	Quantity       int          `json:"quantity"`
	Price          float64      `json:"price"`
	MarketPrice    float64      `json:"market_price,omitempty"`
	FXRate         float64      `json:"fx_rate"`
	Fees           FeeBreakdown `json:"fees,omitempty"`
	FeeTotal       float64      `json:"fee_total,omitempty"`
	FeeTotalBase   float64      `json:"fee_total_base,omitempty"`
}

type FeeBreakdown struct {
	PlatformFee         float64 `json:"platform_fee,omitempty"`
	StampDuty           float64 `json:"stamp_duty,omitempty"`
	SFCTransactionLevy  float64 `json:"sfc_transaction_levy,omitempty"`
	HKEXSettlementFee   float64 `json:"hkex_settlement_fee,omitempty"`
	AFRCTransactionLevy float64 `json:"afrc_transaction_levy,omitempty"`
	HKEXTradingFee      float64 `json:"hkex_trading_fee,omitempty"`
}

type persistedState struct {
	InitialCapital     float64       `json:"initial_capital,omitempty"`
	InitialCapitalCNY  float64       `json:"initial_capital_cny,omitempty"`
	InitialCapitalHKD  float64       `json:"initial_capital_hkd,omitempty"`
	InitialCapitalUSD  float64       `json:"initial_capital_usd,omitempty"`
	HKDRate            float64       `json:"hkd_rate"`
	USDRate            float64       `json:"usd_rate,omitempty"`
	RefreshIntervalSec int           `json:"refresh_interval_sec,omitempty"`
	QuoteCache         quoteCache    `json:"quote_cache,omitempty"`
	FXCache            fxCache       `json:"fx_cache,omitempty"`
	Records            []TradeRecord `json:"records"`
}

type persistedQuote struct {
	Name        string    `json:"name,omitempty"`
	Price       float64   `json:"price,omitempty"`
	PriceDigits int       `json:"price_digits,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type quoteCache map[string]persistedQuote

type persistedFXRate struct {
	Rate      float64   `json:"rate,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type fxCache struct {
	HKDCNY persistedFXRate `json:"hkd_cny,omitempty"`
	USDCNY persistedFXRate `json:"usd_cny,omitempty"`
}

type Position struct {
	Code         string
	DisplayCode  string
	Symbol       string
	Market       string
	Currency     string
	Name         string
	Quantity     int
	AvgCostLocal float64
	AvgCostBase  float64
	LastPrice    float64
	PriceDigits  int
}

type Portfolio struct {
	mu sync.RWMutex

	initialCapitalCNY float64
	initialCapitalHKD float64
	initialCapitalUSD float64
	hkdRate           float64
	usdRate           float64
	refreshInterval   time.Duration
	records           []TradeRecord
	quoteCache        quoteCache
	fxCache           fxCache

	positions      map[string]*Position
	cashCNY        float64
	cashHKD        float64
	cashUSD        float64
	realizedPnL    float64
	lastRefreshAt  time.Time
	lastRefreshErr string
}

type Summary struct {
	InitialCapital    float64
	InitialCapitalCNY float64
	InitialCapitalHKD float64
	InitialCapitalUSD float64
	HKDRate           float64
	USDRate           float64
	RefreshInterval   time.Duration
	Cash              float64
	CashCNY           float64
	CashHKD           float64
	CashUSD           float64
	MarketValue       float64
	TotalAssets       float64
	RealizedPnL       float64
	UnrealizedPnL     float64
	TotalReturn       float64
	FeeTotals         FeeBreakdown
	FeeTotal          float64
	FeeMarketTotals   []FeeMarketTotal
	LastRefreshAt     time.Time
	LastRefreshErr    string
	Holdings          []HoldingSummary
	Records           []TradeSummary
}

type FeeMarketTotal struct {
	Market   string
	Currency string
	Amount   float64
}

type HoldingSummary struct {
	DisplayCode        string
	Name               string
	Market             string
	Currency           string
	Quantity           int
	AvgCostLocal       float64
	CurrentPrice       float64
	CurrentPriceDigits int
	MarketValue        float64
	MarketValueLocal   float64
	UnrealizedPnL      float64
	UnrealizedPnLLocal float64
}

type TradeSummary struct {
	Time        time.Time
	TypeLabel   string
	DisplayCode string
	Market      string
	Currency    string
	Quantity    int
	Price       float64
	FXRate      float64
	AmountBase  float64
	Fee         float64
	FeeBase     float64
}

func LoadPortfolio(raw string) (*Portfolio, error) {
	p := &Portfolio{
		hkdRate:         defaultHKDRate,
		usdRate:         defaultUSDRate,
		refreshInterval: defaultRefreshSec * time.Second,
		quoteCache:      make(quoteCache),
		positions:       make(map[string]*Position),
	}

	if strings.TrimSpace(raw) == "" {
		return p, nil
	}

	var state persistedState
	if err := json.Unmarshal([]byte(raw), &state); err != nil {
		return nil, fmt.Errorf("加载本地数据失败: %w", err)
	}

	p.initialCapitalCNY = state.InitialCapitalCNY
	p.initialCapitalHKD = state.InitialCapitalHKD
	p.initialCapitalUSD = state.InitialCapitalUSD
	if p.initialCapitalCNY == 0 && p.initialCapitalHKD == 0 && p.initialCapitalUSD == 0 && state.InitialCapital != 0 {
		p.initialCapitalCNY = state.InitialCapital
	}
	if state.HKDRate > 0 {
		p.hkdRate = state.HKDRate
	}
	if state.USDRate > 0 {
		p.usdRate = state.USDRate
	}
	if state.FXCache.HKDCNY.Rate > 0 {
		p.hkdRate = state.FXCache.HKDCNY.Rate
		p.fxCache.HKDCNY = state.FXCache.HKDCNY
	}
	if state.FXCache.USDCNY.Rate > 0 {
		p.usdRate = state.FXCache.USDCNY.Rate
		p.fxCache.USDCNY = state.FXCache.USDCNY
	}
	if state.RefreshIntervalSec > 0 {
		p.refreshInterval = time.Duration(state.RefreshIntervalSec) * time.Second
	}
	for symbol, quote := range state.QuoteCache {
		normalized := normalizeQuoteSymbol(symbol)
		if normalized == "" || quote.Price <= 0 {
			continue
		}
		p.quoteCache[normalized] = quote
	}
	p.records = append(p.records, withCalculatedFees(state.Records)...)

	if err := p.rebuildLocked(); err != nil {
		return nil, err
	}
	p.applyCachedLastRefreshLocked()
	return p, nil
}

func (p *Portfolio) MarshalState() ([]byte, error) {
	p.mu.RLock()
	records := withCalculatedFees(p.records)
	state := persistedState{
		InitialCapitalCNY:  p.initialCapitalCNY,
		InitialCapitalHKD:  p.initialCapitalHKD,
		InitialCapitalUSD:  p.initialCapitalUSD,
		HKDRate:            p.hkdRate,
		USDRate:            p.usdRate,
		RefreshIntervalSec: int(p.refreshInterval / time.Second),
		QuoteCache:         p.quoteCacheStateLocked(),
		FXCache:            p.fxCache,
		Records:            records,
	}
	p.mu.RUnlock()

	payload, err := json.Marshal(state)
	if err != nil {
		return nil, fmt.Errorf("序列化本地数据失败: %w", err)
	}
	return payload, nil
}

func (p *Portfolio) AdjustBalance(delta float64) error {
	return p.adjustBalanceByNormalizedCurrency(delta, "CNY")
}

func (p *Portfolio) AdjustBalanceByCurrency(amount float64, currency string, direction float64) error {
	if amount <= 0 {
		return errors.New("调整金额必须大于 0")
	}
	currency = normalizePriceCurrency(currency)
	if currency == "" {
		return errors.New("余额币种仅支持 RMB、港币或美元")
	}
	return p.adjustBalanceByNormalizedCurrency(direction*amount, currency)
}

func (p *Portfolio) SetExchangeRate(hkdRate float64) error {
	if hkdRate <= 0 {
		return errors.New("港币汇率必须大于 0")
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	p.hkdRate = hkdRate
	return nil
}

func (p *Portfolio) SetUSDExchangeRate(usdRate float64) error {
	if usdRate <= 0 {
		return errors.New("美元汇率必须大于 0")
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	p.usdRate = usdRate
	return nil
}

func (p *Portfolio) SetExchangeRateWithTime(hkdRate float64, updatedAt time.Time) error {
	if hkdRate <= 0 {
		return errors.New("港币汇率必须大于 0")
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	p.hkdRate = hkdRate
	p.fxCache.HKDCNY = persistedFXRate{Rate: hkdRate, UpdatedAt: updatedAt}
	return nil
}

func (p *Portfolio) SetUSDExchangeRateWithTime(usdRate float64, updatedAt time.Time) error {
	if usdRate <= 0 {
		return errors.New("美元汇率必须大于 0")
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	p.usdRate = usdRate
	p.fxCache.USDCNY = persistedFXRate{Rate: usdRate, UpdatedAt: updatedAt}
	return nil
}

func (p *Portfolio) RefreshInterval() time.Duration {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.refreshInterval <= 0 {
		return defaultRefreshSec * time.Second
	}
	return p.refreshInterval
}

func (p *Portfolio) SetRefreshInterval(interval time.Duration) error {
	if interval <= 0 {
		return errors.New("刷新间隔必须大于 0")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.refreshInterval = interval
	return nil
}

func (p *Portfolio) HasFreshQuoteCache(maxAge time.Duration, now time.Time) bool {
	_, ok := p.NextQuoteRefreshFromCache(maxAge, now)
	return ok
}

func (p *Portfolio) NextQuoteRefreshFromCache(maxAge time.Duration, now time.Time) (time.Time, bool) {
	if maxAge <= 0 {
		return time.Time{}, false
	}
	p.mu.RLock()
	defer p.mu.RUnlock()
	if len(p.positions) == 0 {
		return time.Time{}, false
	}
	var next time.Time
	for symbol := range p.positions {
		quote := p.quoteCache[symbol]
		if quote.Price <= 0 || quote.UpdatedAt.IsZero() {
			return time.Time{}, false
		}
		nextForQuote := quote.UpdatedAt.Add(maxAge)
		if !nextForQuote.After(now) {
			return time.Time{}, false
		}
		if next.IsZero() || nextForQuote.Before(next) {
			next = nextForQuote
		}
	}
	return next, true
}

func (p *Portfolio) NextFXRefreshFromCache(maxAge time.Duration, now time.Time) (time.Time, bool) {
	if maxAge <= 0 {
		return time.Time{}, false
	}
	p.mu.RLock()
	defer p.mu.RUnlock()

	rates := []persistedFXRate{p.fxCache.HKDCNY, p.fxCache.USDCNY}
	var next time.Time
	for _, item := range rates {
		if item.Rate <= 0 || item.UpdatedAt.IsZero() {
			return time.Time{}, false
		}
		nextForRate := item.UpdatedAt.Add(maxAge)
		if !nextForRate.After(now) {
			return time.Time{}, false
		}
		if next.IsZero() || nextForRate.Before(next) {
			next = nextForRate
		}
	}
	return next, true
}

func (p *Portfolio) AddRecord(kind TradeType, rawCode string, quantity int, price float64, priceCurrency string) error {
	if kind != TradeTypeBuy && kind != TradeTypeSell {
		return errors.New("交易类型无效")
	}
	if quantity <= 0 {
		return errors.New("股数必须大于 0")
	}
	if price <= 0 {
		return errors.New("价格必须大于 0")
	}
	priceCurrency = normalizePriceCurrency(priceCurrency)
	if priceCurrency == "" {
		return errors.New("股价币种仅支持 CNY、HKD 或 USD")
	}

	security, err := normalizeSecurity(rawCode)
	if err != nil {
		return err
	}
	marketPrice, err := convertPriceToMarketCurrency(price, priceCurrency, security.Currency, p.hkdRate, p.usdRate)
	if err != nil {
		return err
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if kind == TradeTypeSell {
		pos := p.positions[security.Symbol]
		if pos == nil || pos.Quantity < quantity {
			return errors.New("平仓股数超过当前持仓")
		}
	}

	record := TradeRecord{
		ID:             strconv.FormatInt(time.Now().UnixNano(), 10),
		CreatedAt:      time.Now(),
		Type:           kind,
		Code:           security.Code,
		DisplayCode:    security.DisplayCode,
		Symbol:         security.Symbol,
		Market:         security.Market,
		Currency:       priceCurrency,
		MarketCurrency: security.Currency,
		Quantity:       quantity,
		Price:          price,
		MarketPrice:    marketPrice,
		FXRate:         fxRateForCurrency(priceCurrency, p.hkdRate, p.usdRate),
	}
	record = withCalculatedFee(record)

	p.records = append(p.records, record)
	return p.rebuildLocked()
}

func (p *Portfolio) RefreshQuotes(provider QuoteProvider) error {
	p.mu.RLock()
	symbols := make([]string, 0, len(p.positions))
	for symbol := range p.positions {
		symbols = append(symbols, symbol)
	}
	p.mu.RUnlock()

	now := time.Now()
	if len(symbols) == 0 {
		p.mu.Lock()
		p.lastRefreshAt = now
		p.lastRefreshErr = ""
		p.mu.Unlock()
		return nil
	}

	quotes, err := provider.Fetch(symbols)

	p.mu.Lock()
	defer p.mu.Unlock()

	p.lastRefreshAt = now
	if err != nil {
		p.lastRefreshErr = err.Error()
		return err
	}

	p.lastRefreshErr = ""
	for symbol, quote := range quotes {
		pos := p.positions[symbol]
		if pos == nil {
			continue
		}
		if quote.Price > 0 {
			pos.LastPrice = quote.Price
			pos.PriceDigits = quote.PriceDigits
		}
		if quote.Name != "" {
			pos.Name = quote.Name
		}
		p.updateQuoteCacheLocked(symbol, quote, now)
	}
	return nil
}

func (p *Portfolio) Summary() Summary {
	p.mu.RLock()
	defer p.mu.RUnlock()

	holdings := make([]HoldingSummary, 0, len(p.positions))
	var marketValue float64
	var unrealized float64

	for _, pos := range p.sortedPositionsLocked() {
		currentLocal := pos.LastPrice * float64(pos.Quantity)
		avgCostLocal := pos.AvgCostLocal * float64(pos.Quantity)
		currentBase := pos.LastPrice * float64(pos.Quantity) * p.fxRateForPositionLocked(pos)
		avgCostBase := pos.AvgCostBase * float64(pos.Quantity)
		rowPnL := currentBase - avgCostBase
		rowPnLLocal := currentLocal - avgCostLocal

		holdings = append(holdings, HoldingSummary{
			DisplayCode:        pos.DisplayCode,
			Name:               pos.Name,
			Market:             pos.Market,
			Currency:           pos.Currency,
			Quantity:           pos.Quantity,
			AvgCostLocal:       pos.AvgCostLocal,
			CurrentPrice:       pos.LastPrice,
			CurrentPriceDigits: pos.PriceDigits,
			MarketValue:        currentBase,
			MarketValueLocal:   currentLocal,
			UnrealizedPnL:      rowPnL,
			UnrealizedPnLLocal: rowPnLLocal,
		})
		marketValue += currentBase
		unrealized += rowPnL
	}

	records := make([]TradeSummary, 0, len(p.records))
	for i := len(p.records) - 1; i >= 0; i-- {
		record := p.records[i]
		records = append(records, TradeSummary{
			Time:        record.CreatedAt,
			TypeLabel:   record.typeLabel(),
			DisplayCode: record.DisplayCode,
			Market:      record.Market,
			Currency:    record.Currency,
			Quantity:    record.Quantity,
			Price:       record.Price,
			FXRate:      record.FXRate,
			AmountBase:  float64(record.Quantity) * record.Price * record.FXRate,
			Fee:         record.FeeTotal,
			FeeBase:     record.FeeTotalBase,
		})
	}

	initialCapital := p.initialCapitalCNY + p.initialCapitalHKD*p.hkdRate + p.initialCapitalUSD*p.usdRate
	cash := p.cashCNY + p.cashHKD*p.hkdRate + p.cashUSD*p.usdRate
	totalAssets := cash + marketValue
	feeTotals := p.feeTotalsLocked()
	return Summary{
		InitialCapital:    initialCapital,
		InitialCapitalCNY: p.initialCapitalCNY,
		InitialCapitalHKD: p.initialCapitalHKD,
		InitialCapitalUSD: p.initialCapitalUSD,
		HKDRate:           p.hkdRate,
		USDRate:           p.usdRate,
		RefreshInterval:   p.refreshInterval,
		Cash:              cash,
		CashCNY:           p.cashCNY,
		CashHKD:           p.cashHKD,
		CashUSD:           p.cashUSD,
		MarketValue:       marketValue,
		TotalAssets:       totalAssets,
		RealizedPnL:       p.realizedPnL,
		UnrealizedPnL:     unrealized,
		TotalReturn:       totalAssets - initialCapital,
		FeeTotals:         feeTotals,
		FeeTotal:          feeTotals.Total(),
		FeeMarketTotals:   p.feeMarketTotalsLocked(),
		LastRefreshAt:     p.lastRefreshAt,
		LastRefreshErr:    p.lastRefreshErr,
		Holdings:          holdings,
		Records:           records,
	}
}

func (p *Portfolio) rebuildLocked() error {
	p.positions = make(map[string]*Position)
	p.cashCNY = p.initialCapitalCNY
	p.cashHKD = p.initialCapitalHKD
	p.cashUSD = p.initialCapitalUSD
	p.realizedPnL = 0

	for _, record := range p.records {
		switch record.Type {
		case TradeTypeBuy:
			p.applyBuyLocked(record)
		case TradeTypeSell:
			if err := p.applySellLocked(record); err != nil {
				return err
			}
		default:
			return fmt.Errorf("未知交易类型: %s", record.Type)
		}
	}

	p.applyQuoteCacheLocked()
	return nil
}

func (p *Portfolio) applyQuoteCacheLocked() {
	for symbol, pos := range p.positions {
		quote := p.quoteCache[symbol]
		if quote.Price <= 0 {
			continue
		}
		pos.LastPrice = quote.Price
		pos.PriceDigits = quote.PriceDigits
		if quote.Name != "" {
			pos.Name = quote.Name
		}
	}
}

func (p *Portfolio) applyCachedLastRefreshLocked() {
	if len(p.positions) == 0 {
		return
	}
	var lastRefreshAt time.Time
	for symbol := range p.positions {
		quote := p.quoteCache[symbol]
		if quote.Price <= 0 || quote.UpdatedAt.IsZero() {
			return
		}
		if lastRefreshAt.IsZero() || quote.UpdatedAt.Before(lastRefreshAt) {
			lastRefreshAt = quote.UpdatedAt
		}
	}
	p.lastRefreshAt = lastRefreshAt
	p.lastRefreshErr = ""
}

func (p *Portfolio) updateQuoteCacheLocked(symbol string, quote Quote, updatedAt time.Time) {
	if p.quoteCache == nil {
		p.quoteCache = make(quoteCache)
	}

	cached := p.quoteCache[symbol]
	if quote.Price > 0 {
		cached.Price = quote.Price
		cached.PriceDigits = quote.PriceDigits
		cached.UpdatedAt = updatedAt
	}
	if quote.Name != "" {
		cached.Name = quote.Name
	}
	if cached.Price > 0 {
		p.quoteCache[symbol] = cached
	}
}

func (p *Portfolio) quoteCacheStateLocked() quoteCache {
	if len(p.positions) == 0 || len(p.quoteCache) == 0 {
		return nil
	}
	cache := make(quoteCache, len(p.positions))
	for symbol := range p.positions {
		quote := p.quoteCache[symbol]
		if quote.Price <= 0 {
			continue
		}
		cache[symbol] = quote
	}
	if len(cache) == 0 {
		return nil
	}
	return cache
}

func (p *Portfolio) applyBuyLocked(record TradeRecord) {
	pos := p.positions[record.Symbol]
	if pos == nil {
		pos = &Position{
			Code:        record.Code,
			DisplayCode: record.DisplayCode,
			Symbol:      record.Symbol,
			Market:      record.Market,
			Currency:    record.marketCurrency(),
		}
		p.positions[record.Symbol] = pos
	}

	newQty := pos.Quantity + record.Quantity
	totalLocalCost := pos.AvgCostLocal*float64(pos.Quantity) + record.marketPrice()*float64(record.Quantity)
	totalBaseCost := pos.AvgCostBase*float64(pos.Quantity) + record.Price*record.FXRate*float64(record.Quantity)

	pos.Quantity = newQty
	pos.AvgCostLocal = totalLocalCost / float64(newQty)
	pos.AvgCostBase = totalBaseCost / float64(newQty)
	p.adjustCashByCurrencyLocked(record.marketCurrency(), -(record.marketPrice()*float64(record.Quantity) + record.FeeTotal))
}

func (p *Portfolio) applySellLocked(record TradeRecord) error {
	pos := p.positions[record.Symbol]
	if pos == nil || pos.Quantity < record.Quantity {
		return fmt.Errorf("交易记录无效，%s 卖出股数超过持仓", record.DisplayCode)
	}

	p.adjustCashByCurrencyLocked(record.marketCurrency(), record.marketPrice()*float64(record.Quantity)-record.FeeTotal)
	p.realizedPnL += (record.Price*record.FXRate - pos.AvgCostBase) * float64(record.Quantity)

	pos.Quantity -= record.Quantity
	if pos.Quantity == 0 {
		delete(p.positions, record.Symbol)
		return nil
	}

	return nil
}

func withCalculatedFees(records []TradeRecord) []TradeRecord {
	items := make([]TradeRecord, len(records))
	for i, record := range records {
		items[i] = withCalculatedFee(record)
	}
	return items
}

func withCalculatedFee(record TradeRecord) TradeRecord {
	fees := calculateTradeFees(record)
	record.Fees = fees
	record.FeeTotal = fees.Total()
	record.FeeTotalBase = record.FeeTotal * recordFeeFXRate(record)
	return record
}

func calculateTradeFees(record TradeRecord) FeeBreakdown {
	amount := record.marketPrice() * float64(record.Quantity)
	if amount <= 0 || record.Quantity <= 0 {
		return FeeBreakdown{}
	}

	switch record.Market {
	case "美股":
		return FeeBreakdown{
			PlatformFee: minFloat64(maxFloat64(0.0099*float64(record.Quantity), 1.99), amount*0.015),
		}
	case "港股":
		return FeeBreakdown{
			PlatformFee:         maxFloat64(amount*0.0005, 18),
			StampDuty:           ceilToDollar(amount * 0.001),
			SFCTransactionLevy:  amount * 0.000027,
			HKEXSettlementFee:   amount * 0.000042,
			AFRCTransactionLevy: amount * 0.0000015,
			HKEXTradingFee:      amount * 0.0000565,
		}
	default:
		return FeeBreakdown{}
	}
}

func (f FeeBreakdown) Total() float64 {
	return f.PlatformFee + f.StampDuty + f.SFCTransactionLevy + f.HKEXSettlementFee + f.AFRCTransactionLevy + f.HKEXTradingFee
}

func (f FeeBreakdown) Add(other FeeBreakdown) FeeBreakdown {
	return FeeBreakdown{
		PlatformFee:         f.PlatformFee + other.PlatformFee,
		StampDuty:           f.StampDuty + other.StampDuty,
		SFCTransactionLevy:  f.SFCTransactionLevy + other.SFCTransactionLevy,
		HKEXSettlementFee:   f.HKEXSettlementFee + other.HKEXSettlementFee,
		AFRCTransactionLevy: f.AFRCTransactionLevy + other.AFRCTransactionLevy,
		HKEXTradingFee:      f.HKEXTradingFee + other.HKEXTradingFee,
	}
}

func (p *Portfolio) feeTotalsLocked() FeeBreakdown {
	var totals FeeBreakdown
	for _, record := range p.records {
		totals = totals.Add(record.Fees.ToBase(recordFeeFXRate(record)))
	}
	return totals
}

func (p *Portfolio) feeMarketTotalsLocked() []FeeMarketTotal {
	type key struct {
		market   string
		currency string
	}
	amounts := make(map[key]float64)
	for _, record := range p.records {
		if record.FeeTotal <= moneyEpsilon {
			continue
		}
		item := key{market: record.Market, currency: record.marketCurrency()}
		amounts[item] += record.FeeTotal
	}

	items := make([]FeeMarketTotal, 0, len(amounts))
	for item, amount := range amounts {
		items = append(items, FeeMarketTotal{
			Market:   item.market,
			Currency: item.currency,
			Amount:   amount,
		})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Market == items[j].Market {
			return items[i].Currency < items[j].Currency
		}
		return feeMarketSortKey(items[i].Market) < feeMarketSortKey(items[j].Market)
	})
	return items
}

func feeMarketSortKey(market string) string {
	switch market {
	case "美股":
		return "1"
	case "港股":
		return "2"
	case "A股":
		return "3"
	default:
		return "9" + market
	}
}

func (f FeeBreakdown) ToBase(rate float64) FeeBreakdown {
	if rate <= 0 {
		rate = 1
	}
	return FeeBreakdown{
		PlatformFee:         f.PlatformFee * rate,
		StampDuty:           f.StampDuty * rate,
		SFCTransactionLevy:  f.SFCTransactionLevy * rate,
		HKEXSettlementFee:   f.HKEXSettlementFee * rate,
		AFRCTransactionLevy: f.AFRCTransactionLevy * rate,
		HKEXTradingFee:      f.HKEXTradingFee * rate,
	}
}

func recordFeeFXRate(record TradeRecord) float64 {
	if normalizePriceCurrency(record.marketCurrency()) == "CNY" || record.FXRate <= 0 {
		return 1
	}
	return record.FXRate
}

func ceilToDollar(value float64) float64 {
	if value <= 0 {
		return 0
	}
	return float64(int(value + 0.999999999))
}

func maxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func minFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func (p *Portfolio) adjustBalanceByNormalizedCurrency(delta float64, currency string) error {
	if delta == 0 {
		return errors.New("调整金额不能为 0")
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	switch currency {
	case "CNY":
		nextCash := p.cashCNY + delta
		if nextCash < -moneyEpsilon {
			return errors.New("减少金额超过当前 RMB 余额")
		}
		p.initialCapitalCNY += delta
		if absFloat64(p.initialCapitalCNY) < moneyEpsilon {
			p.initialCapitalCNY = 0
		}
	case "HKD":
		nextCash := p.cashHKD + delta
		if nextCash < -moneyEpsilon {
			return errors.New("减少金额超过当前港币余额")
		}
		p.initialCapitalHKD += delta
		if absFloat64(p.initialCapitalHKD) < moneyEpsilon {
			p.initialCapitalHKD = 0
		}
	case "USD":
		nextCash := p.cashUSD + delta
		if nextCash < -moneyEpsilon {
			return errors.New("减少金额超过当前美元余额")
		}
		p.initialCapitalUSD += delta
		if absFloat64(p.initialCapitalUSD) < moneyEpsilon {
			p.initialCapitalUSD = 0
		}
	default:
		return errors.New("余额币种仅支持 RMB、港币或美元")
	}

	return p.rebuildLocked()
}

func (p *Portfolio) adjustCashByCurrencyLocked(currency string, delta float64) {
	switch normalizePriceCurrency(currency) {
	case "HKD":
		p.cashHKD += delta
		if absFloat64(p.cashHKD) < moneyEpsilon {
			p.cashHKD = 0
		}
	case "USD":
		p.cashUSD += delta
		if absFloat64(p.cashUSD) < moneyEpsilon {
			p.cashUSD = 0
		}
	default:
		p.cashCNY += delta
		if absFloat64(p.cashCNY) < moneyEpsilon {
			p.cashCNY = 0
		}
	}
}

func absFloat64(value float64) float64 {
	if value < 0 {
		return -value
	}
	return value
}

func (p *Portfolio) sortedPositionsLocked() []*Position {
	items := make([]*Position, 0, len(p.positions))
	for _, pos := range p.positions {
		copyPos := *pos
		items = append(items, &copyPos)
	}

	sort.Slice(items, func(i, j int) bool {
		valueI := items[i].LastPrice * float64(items[i].Quantity) * p.fxRateForPositionLocked(items[i])
		valueJ := items[j].LastPrice * float64(items[j].Quantity) * p.fxRateForPositionLocked(items[j])
		if valueI == valueJ {
			return items[i].DisplayCode < items[j].DisplayCode
		}
		return valueI > valueJ
	})
	return items
}

func (p *Portfolio) fxRateForPositionLocked(pos *Position) float64 {
	return fxRateForCurrency(pos.Currency, p.hkdRate, p.usdRate)
}

func (r TradeRecord) typeLabel() string {
	if r.Type == TradeTypeSell {
		return "平仓"
	}
	return "买入"
}

func (r TradeRecord) marketCurrency() string {
	if r.MarketCurrency != "" {
		return r.MarketCurrency
	}
	return r.Currency
}

func (r TradeRecord) marketPrice() float64 {
	if r.MarketPrice > 0 {
		return r.MarketPrice
	}
	return r.Price
}

type Security struct {
	Code        string
	DisplayCode string
	Symbol      string
	Market      string
	Currency    string
}

func (s Security) FXRate(hkdRate float64) float64 {
	return fxRateForCurrency(s.Currency, hkdRate, defaultUSDRate)
}

func normalizePriceCurrency(raw string) string {
	switch strings.ToUpper(strings.TrimSpace(raw)) {
	case "CNY", "RMB", "人民币":
		return "CNY"
	case "HKD", "港币":
		return "HKD"
	case "USD", "美元", "美金":
		return "USD"
	default:
		return ""
	}
}

func fxRateForCurrency(currency string, hkdRate float64, usdRate float64) float64 {
	switch normalizePriceCurrency(currency) {
	case "HKD":
		return hkdRate
	case "USD":
		return usdRate
	default:
		return 1
	}
}

func convertPriceToMarketCurrency(price float64, inputCurrency, marketCurrency string, hkdRate float64, usdRate float64) (float64, error) {
	inputCurrency = normalizePriceCurrency(inputCurrency)
	marketCurrency = normalizePriceCurrency(marketCurrency)
	if inputCurrency == "" || marketCurrency == "" {
		return 0, errors.New("不支持的币种")
	}
	if inputCurrency == marketCurrency {
		return price, nil
	}
	inputRate := fxRateForCurrency(inputCurrency, hkdRate, usdRate)
	marketRate := fxRateForCurrency(marketCurrency, hkdRate, usdRate)
	if inputRate <= 0 || marketRate <= 0 {
		return 0, errors.New("汇率必须大于 0")
	}
	return price * inputRate / marketRate, nil
}

func normalizeSecurity(raw string) (Security, error) {
	code := strings.ToLower(strings.TrimSpace(raw))
	code = strings.ReplaceAll(code, " ", "")
	if code == "" {
		return Security{}, errors.New("股票代码不能为空")
	}

	prefix := ""
	switch {
	case strings.HasPrefix(code, "us:") || strings.HasPrefix(code, "us."):
		prefix = "us"
		code = code[3:]
	case strings.HasPrefix(code, "us") && len(code) > 2 && !isDigits(code[2:]):
		prefix = "us"
		code = code[2:]
	case strings.HasPrefix(code, "sh"), strings.HasPrefix(code, "sz"), strings.HasPrefix(code, "bj"), strings.HasPrefix(code, "hk"):
		prefix = code[:2]
		code = code[2:]
	}

	if prefix == "us" || containsLetter(code) {
		code = strings.TrimSpace(strings.ToUpper(code))
		code = strings.ReplaceAll(code, "-", ".")
		if !isValidUSTicker(code) {
			return Security{}, errors.New("美股代码只能包含字母、数字、点或连字符")
		}
		return Security{
			Code:        code,
			DisplayCode: "US:" + code,
			Symbol:      "us" + strings.ToLower(code),
			Market:      "美股",
			Currency:    "USD",
		}, nil
	}

	for _, r := range code {
		if r < '0' || r > '9' {
			return Security{}, errors.New("股票代码只能包含数字，或使用 sh/sz/bj/hk/us 前缀")
		}
	}

	if prefix == "hk" || (prefix == "" && len(code) <= 5) {
		if len(code) == 0 || len(code) > 5 {
			return Security{}, errors.New("港股代码长度应为 1 到 5 位")
		}
		display := fmt.Sprintf("%05s", code)
		return Security{
			Code:        display,
			DisplayCode: display,
			Symbol:      "hk" + display,
			Market:      "港股",
			Currency:    "HKD",
		}, nil
	}

	if len(code) != 6 {
		return Security{}, errors.New("A 股代码需要 6 位数字")
	}

	if prefix == "" {
		switch code[0] {
		case '5', '6', '9':
			prefix = "sh"
		case '4', '8':
			prefix = "bj"
		default:
			prefix = "sz"
		}
	}

	if prefix != "sh" && prefix != "sz" && prefix != "bj" {
		return Security{}, errors.New("仅支持 A 股、港股和美股代码")
	}

	return Security{
		Code:        code,
		DisplayCode: strings.ToUpper(prefix) + code,
		Symbol:      prefix + code,
		Market:      "A股",
		Currency:    "CNY",
	}, nil
}

func containsLetter(value string) bool {
	for _, r := range value {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return true
		}
	}
	return false
}

func isValidUSTicker(value string) bool {
	if value == "" || len(value) > 12 {
		return false
	}
	for _, r := range value {
		if (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '.' || r == '-' {
			continue
		}
		return false
	}
	return true
}
