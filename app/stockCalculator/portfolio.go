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
	moneyEpsilon       = 1e-6
)

type TradeType string

const (
	TradeTypeBuy  TradeType = "buy"
	TradeTypeSell TradeType = "sell"
)

type TradeRecord struct {
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	Type           TradeType `json:"type"`
	Code           string    `json:"code"`
	DisplayCode    string    `json:"display_code"`
	Symbol         string    `json:"symbol"`
	Market         string    `json:"market"`
	Currency       string    `json:"currency"`
	MarketCurrency string    `json:"market_currency,omitempty"`
	Quantity       int       `json:"quantity"`
	Price          float64   `json:"price"`
	MarketPrice    float64   `json:"market_price,omitempty"`
	FXRate         float64   `json:"fx_rate"`
}

type persistedState struct {
	InitialCapital    float64       `json:"initial_capital,omitempty"`
	InitialCapitalCNY float64       `json:"initial_capital_cny,omitempty"`
	InitialCapitalHKD float64       `json:"initial_capital_hkd,omitempty"`
	HKDRate           float64       `json:"hkd_rate"`
	Records           []TradeRecord `json:"records"`
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
}

type Portfolio struct {
	mu sync.RWMutex

	initialCapitalCNY float64
	initialCapitalHKD float64
	hkdRate           float64
	records           []TradeRecord

	positions      map[string]*Position
	cashCNY        float64
	cashHKD        float64
	realizedPnL    float64
	lastRefreshAt  time.Time
	lastRefreshErr string
}

type Summary struct {
	InitialCapital    float64
	InitialCapitalCNY float64
	InitialCapitalHKD float64
	HKDRate           float64
	Cash              float64
	CashCNY           float64
	CashHKD           float64
	MarketValue       float64
	TotalAssets       float64
	RealizedPnL       float64
	UnrealizedPnL     float64
	TotalReturn       float64
	LastRefreshAt     time.Time
	LastRefreshErr    string
	Holdings          []HoldingSummary
	Records           []TradeSummary
}

type HoldingSummary struct {
	DisplayCode        string
	Name               string
	Market             string
	Currency           string
	Quantity           int
	AvgCostLocal       float64
	CurrentPrice       float64
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
}

func LoadPortfolio(raw string) (*Portfolio, error) {
	p := &Portfolio{
		hkdRate:   defaultHKDRate,
		positions: make(map[string]*Position),
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
	if p.initialCapitalCNY == 0 && p.initialCapitalHKD == 0 && state.InitialCapital != 0 {
		p.initialCapitalCNY = state.InitialCapital
	}
	if state.HKDRate > 0 {
		p.hkdRate = state.HKDRate
	}
	p.records = append(p.records, state.Records...)

	if err := p.rebuildLocked(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Portfolio) MarshalState() ([]byte, error) {
	p.mu.RLock()
	state := persistedState{
		InitialCapitalCNY: p.initialCapitalCNY,
		InitialCapitalHKD: p.initialCapitalHKD,
		HKDRate:           p.hkdRate,
		Records:           append([]TradeRecord(nil), p.records...),
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
		return errors.New("余额币种仅支持 RMB 或 港币")
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
		return errors.New("股价币种仅支持 CNY 或 HKD")
	}

	security, err := normalizeSecurity(rawCode)
	if err != nil {
		return err
	}
	marketPrice, err := convertPriceToMarketCurrency(price, priceCurrency, security.Currency, p.hkdRate)
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
		FXRate:         fxRateForCurrency(priceCurrency, p.hkdRate),
	}

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
		pos.LastPrice = quote.Price
		if quote.Name != "" {
			pos.Name = quote.Name
		}
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
		})
	}

	initialCapital := p.initialCapitalCNY + p.initialCapitalHKD*p.hkdRate
	cash := p.cashCNY + p.cashHKD*p.hkdRate
	totalAssets := cash + marketValue
	return Summary{
		InitialCapital:    initialCapital,
		InitialCapitalCNY: p.initialCapitalCNY,
		InitialCapitalHKD: p.initialCapitalHKD,
		HKDRate:           p.hkdRate,
		Cash:              cash,
		CashCNY:           p.cashCNY,
		CashHKD:           p.cashHKD,
		MarketValue:       marketValue,
		TotalAssets:       totalAssets,
		RealizedPnL:       p.realizedPnL,
		UnrealizedPnL:     unrealized,
		TotalReturn:       totalAssets - initialCapital,
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

	return nil
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
	p.adjustCashByCurrencyLocked(record.marketCurrency(), -record.marketPrice()*float64(record.Quantity))
}

func (p *Portfolio) applySellLocked(record TradeRecord) error {
	pos := p.positions[record.Symbol]
	if pos == nil || pos.Quantity < record.Quantity {
		return fmt.Errorf("交易记录无效，%s 卖出股数超过持仓", record.DisplayCode)
	}

	p.adjustCashByCurrencyLocked(record.marketCurrency(), record.marketPrice()*float64(record.Quantity))
	p.realizedPnL += (record.Price*record.FXRate - pos.AvgCostBase) * float64(record.Quantity)

	pos.Quantity -= record.Quantity
	if pos.Quantity == 0 {
		delete(p.positions, record.Symbol)
		return nil
	}

	return nil
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
	default:
		return errors.New("余额币种仅支持 RMB 或 港币")
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
	if pos.Currency == "HKD" {
		return p.hkdRate
	}
	return 1
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
	if s.Currency == "HKD" {
		return hkdRate
	}
	return 1
}

func normalizePriceCurrency(raw string) string {
	switch strings.ToUpper(strings.TrimSpace(raw)) {
	case "CNY", "RMB", "人民币":
		return "CNY"
	case "HKD", "港币":
		return "HKD"
	default:
		return ""
	}
}

func fxRateForCurrency(currency string, hkdRate float64) float64 {
	if normalizePriceCurrency(currency) == "HKD" {
		return hkdRate
	}
	return 1
}

func convertPriceToMarketCurrency(price float64, inputCurrency, marketCurrency string, hkdRate float64) (float64, error) {
	inputCurrency = normalizePriceCurrency(inputCurrency)
	marketCurrency = normalizePriceCurrency(marketCurrency)
	if inputCurrency == "" || marketCurrency == "" {
		return 0, errors.New("不支持的币种")
	}
	if inputCurrency == marketCurrency {
		return price, nil
	}
	if hkdRate <= 0 {
		return 0, errors.New("港币汇率必须大于 0")
	}
	if inputCurrency == "HKD" && marketCurrency == "CNY" {
		return price * hkdRate, nil
	}
	if inputCurrency == "CNY" && marketCurrency == "HKD" {
		return price / hkdRate, nil
	}
	return 0, errors.New("暂不支持该币种转换")
}

func normalizeSecurity(raw string) (Security, error) {
	code := strings.ToLower(strings.TrimSpace(raw))
	code = strings.ReplaceAll(code, " ", "")
	if code == "" {
		return Security{}, errors.New("股票代码不能为空")
	}

	prefix := ""
	switch {
	case strings.HasPrefix(code, "sh"), strings.HasPrefix(code, "sz"), strings.HasPrefix(code, "bj"), strings.HasPrefix(code, "hk"):
		prefix = code[:2]
		code = code[2:]
	}

	for _, r := range code {
		if r < '0' || r > '9' {
			return Security{}, errors.New("股票代码只能包含数字，或带 sh/sz/bj/hk 前缀")
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
		return Security{}, errors.New("仅支持 A 股和港股代码")
	}

	return Security{
		Code:        code,
		DisplayCode: strings.ToUpper(prefix) + code,
		Symbol:      prefix + code,
		Market:      "A股",
		Currency:    "CNY",
	}, nil
}
