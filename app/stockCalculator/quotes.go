package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Quote struct {
	Symbol      string
	Name        string
	Price       float64
	PriceDigits int
}

type QuoteProvider interface {
	Fetch(symbols []string) (map[string]Quote, error)
}

type ExchangeRateProvider interface {
	FetchHKDCNY() (float64, error)
	FetchUSDCNY() (float64, error)
}

type EastMoneyQuoteProvider struct {
	client   *http.Client
	baseURLs []string
}

type FrankfurterExchangeRateProvider struct {
	client *http.Client
}

func NewEastMoneyQuoteProvider() *EastMoneyQuoteProvider {
	return &EastMoneyQuoteProvider{
		client: &http.Client{Timeout: 6 * time.Second},
		baseURLs: []string{
			"https://push2.eastmoney.com/api/qt/ulist.np/get",
			"http://push2.eastmoney.com/api/qt/ulist.np/get",
			"https://push2his.eastmoney.com/api/qt/ulist.np/get",
		},
	}
}

func NewFrankfurterExchangeRateProvider() *FrankfurterExchangeRateProvider {
	return &FrankfurterExchangeRateProvider{
		client: &http.Client{Timeout: 6 * time.Second},
	}
}

func (p *EastMoneyQuoteProvider) Fetch(symbols []string) (map[string]Quote, error) {
	if len(symbols) == 0 {
		return map[string]Quote{}, nil
	}

	secIDs, symbolBySecID, normalizedSymbols, usSymbols, failures := splitQuoteSymbols(symbols)

	quotes := make(map[string]Quote, len(symbols))
	var fetchFailures []string
	if len(secIDs) > 0 {
		eastMoneyQuotes, err := p.fetchEastMoneyQuotes(secIDs, symbolBySecID)
		if err != nil {
			fetchFailures = append(fetchFailures, err.Error())
			eastMoneyQuotes = make(map[string]Quote, len(secIDs))
		}
		p.fillMissingAStockPrices(normalizedSymbols, eastMoneyQuotes)
		for symbol, quote := range eastMoneyQuotes {
			quotes[symbol] = quote
		}
	}
	if len(usSymbols) > 0 {
		usQuotes, err := p.fetchUSQuotes(usSymbols)
		if err != nil {
			fetchFailures = append(fetchFailures, err.Error())
		}
		for symbol, quote := range usQuotes {
			quotes[symbol] = quote
		}
	}

	if len(quotes) == 0 {
		if len(fetchFailures) > 0 {
			return nil, fmt.Errorf("行情请求失败: %s", strings.Join(fetchFailures, "; "))
		}
		if len(failures) > 0 {
			return nil, fmt.Errorf("没有解析到有效行情: %s", strings.Join(failures, "; "))
		}
		return nil, errors.New("没有解析到有效行情，可能是接口不可用")
	}
	return quotes, nil
}

func splitQuoteSymbols(symbols []string) ([]string, map[string]string, []string, []string, []string) {
	secIDs := make([]string, 0, len(symbols))
	symbolBySecID := make(map[string]string, len(symbols))
	normalizedSymbols := make([]string, 0, len(symbols))
	usSymbols := make([]string, 0, len(symbols))
	var failures []string
	for _, symbol := range symbols {
		normalized := normalizeQuoteSymbol(symbol)
		if normalized == "" {
			continue
		}
		if strings.HasPrefix(normalized, "us") {
			usSymbols = append(usSymbols, normalized)
			continue
		}
		secID, err := toEastMoneySecID(normalized)
		if err != nil {
			failures = append(failures, fmt.Sprintf("%s: %v", symbol, err))
			continue
		}
		secIDs = append(secIDs, secID)
		normalizedSymbols = append(normalizedSymbols, normalized)
		symbolBySecID[secID] = normalized
	}
	return secIDs, symbolBySecID, normalizedSymbols, usSymbols, failures
}

func (p *EastMoneyQuoteProvider) fetchEastMoneyQuotes(secIDs []string, symbolBySecID map[string]string) (map[string]Quote, error) {
	body, err := p.fetchEastMoneyBody(secIDs)
	if err == nil {
		return parseEastMoneyQuotes(body, symbolBySecID)
	}
	if len(secIDs) == 1 {
		return nil, err
	}

	quotes := make(map[string]Quote, len(secIDs))
	var failures []string
	for _, secID := range secIDs {
		body, itemErr := p.fetchEastMoneyBody([]string{secID})
		if itemErr != nil {
			failures = append(failures, fmt.Sprintf("%s: %v", secID, itemErr))
			continue
		}
		itemQuotes, itemErr := parseEastMoneyQuotes(body, symbolBySecID)
		if itemErr != nil {
			failures = append(failures, fmt.Sprintf("%s: %v", secID, itemErr))
			continue
		}
		for symbol, quote := range itemQuotes {
			quotes[symbol] = quote
		}
	}
	if len(quotes) > 0 {
		return quotes, nil
	}
	if len(failures) > 0 {
		return nil, fmt.Errorf("行情请求失败: %s", strings.Join(failures, "; "))
	}
	return nil, err
}

func (p *EastMoneyQuoteProvider) fillMissingAStockPrices(symbols []string, quotes map[string]Quote) {
	fallbackSymbols := make([]string, 0, len(symbols))
	for _, symbol := range symbols {
		normalized := normalizeQuoteSymbol(symbol)
		if !strings.HasPrefix(normalized, "sh") && !strings.HasPrefix(normalized, "sz") {
			continue
		}
		if quotes[normalized].Price > 0 {
			continue
		}
		fallbackSymbols = append(fallbackSymbols, normalized)
	}
	if len(fallbackSymbols) == 0 {
		return
	}

	fallbackQuotes, err := p.fetchTencentQuotes(fallbackSymbols)
	if err != nil {
		return
	}
	for symbol, fallback := range fallbackQuotes {
		if fallback.Price <= 0 {
			continue
		}
		quote := quotes[symbol]
		if quote.Symbol == "" {
			quote.Symbol = symbol
		}
		if quote.Name == "" {
			quote.Name = fallback.Name
		}
		quote.Price = fallback.Price
		quote.PriceDigits = fallback.PriceDigits
		quotes[symbol] = quote
	}
}

func normalizeQuoteSymbol(symbol string) string {
	return strings.ToLower(strings.TrimSpace(symbol))
}

func (p *EastMoneyQuoteProvider) fetchTencentQuotes(symbols []string) (map[string]Quote, error) {
	req, err := http.NewRequest(http.MethodGet, "https://qt.gtimg.cn/q="+strings.Join(symbols, ","), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Referer", "https://gu.qq.com/")
	req.Close = true

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("腾讯行情接口返回状态码 %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取腾讯行情响应失败: %w", err)
	}

	lines := strings.Split(decodeGBK(body), ";")
	quotes := make(map[string]Quote, len(symbols))
	for _, line := range lines {
		symbol, quote, ok := parseTencentLine(line)
		if ok {
			quotes[symbol] = quote
		}
	}
	return quotes, nil
}

func (p *EastMoneyQuoteProvider) fetchEastMoneyBody(secIDs []string) ([]byte, error) {
	var lastErr error
	for _, baseURL := range p.baseURLs {
		for attempt := 0; attempt < 2; attempt++ {
			body, err := p.fetchEastMoneyBodyOnce(baseURL, secIDs)
			if err == nil {
				return body, nil
			}
			lastErr = err
			time.Sleep(time.Duration(100*(attempt+1)) * time.Millisecond)
		}
	}
	return nil, fmt.Errorf("行情请求失败: %w", lastErr)
}

func (p *EastMoneyQuoteProvider) fetchEastMoneyBodyOnce(baseURL string, secIDs []string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set("fltt", "2")
	q.Set("fields", "f12,f13,f14,f2")
	q.Set("secids", strings.Join(secIDs, ","))
	req.URL.RawQuery = q.Encode()
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Referer", "https://quote.eastmoney.com/")
	req.Close = true
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("行情接口返回状态码 %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取行情响应失败: %w", err)
	}
	return body, nil
}

func toEastMoneySecID(symbol string) (string, error) {
	symbol = strings.ToLower(strings.TrimSpace(symbol))
	if len(symbol) < 3 {
		return "", errors.New("股票代码无效")
	}

	prefix := symbol[:2]
	code := symbol[2:]
	switch prefix {
	case "sh", "sz", "bj":
		if len(code) != 6 || !isDigits(code) {
			return "", errors.New("A 股代码无效")
		}
		market := "0"
		if prefix == "sh" {
			market = "1"
		}
		return market + "." + code, nil
	case "hk":
		if len(code) == 0 || len(code) > 5 || !isDigits(code) {
			return "", errors.New("港股代码无效")
		}
		return "116." + fmt.Sprintf("%05s", code), nil
	default:
		return "", errors.New("仅支持 A 股、港股和美股行情")
	}
}

func (p *EastMoneyQuoteProvider) fetchUSQuotes(symbols []string) (map[string]Quote, error) {
	quotes, err := p.fetchStooqQuotes(symbols)
	if err == nil && len(quotes) == len(symbols) {
		return quotes, nil
	}

	yahooQuotes, yahooErr := p.fetchYahooQuotes(symbols)
	if len(quotes) == 0 {
		quotes = yahooQuotes
	} else {
		for symbol, quote := range yahooQuotes {
			quotes[symbol] = quote
		}
	}
	if len(quotes) > 0 {
		return quotes, nil
	}
	if yahooErr != nil {
		return nil, yahooErr
	}
	return nil, err
}

func (p *EastMoneyQuoteProvider) fetchStooqQuotes(symbols []string) (map[string]Quote, error) {
	quotes := make(map[string]Quote, len(symbols))
	var failures []string
	for _, symbol := range symbols {
		quote, err := p.fetchSingleStooqQuote(symbol)
		if err != nil {
			failures = append(failures, fmt.Sprintf("%s: %v", symbol, err))
			continue
		}
		quotes[symbol] = quote
	}
	if len(quotes) > 0 {
		return quotes, nil
	}
	if len(failures) > 0 {
		return nil, fmt.Errorf("Stooq 美股行情请求失败: %s", strings.Join(failures, "; "))
	}
	return nil, errors.New("Stooq 美股行情请求失败")
}

func (p *EastMoneyQuoteProvider) fetchSingleStooqQuote(symbol string) (Quote, error) {
	ticker := usTickerFromSymbol(symbol)
	if ticker == "" {
		return Quote{}, errors.New("美股代码无效")
	}

	req, err := http.NewRequest(http.MethodGet, "https://stooq.com/q/l/", nil)
	if err != nil {
		return Quote{}, err
	}
	q := req.URL.Query()
	q.Set("s", strings.ToLower(strings.ReplaceAll(ticker, ".", "-"))+".us")
	q.Set("f", "sd2t2ncl1")
	q.Set("h", "")
	q.Set("e", "csv")
	req.URL.RawQuery = q.Encode()
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Close = true

	resp, err := p.client.Do(req)
	if err != nil {
		return Quote{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Quote{}, fmt.Errorf("Stooq 接口返回状态码 %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Quote{}, fmt.Errorf("读取 Stooq 响应失败: %w", err)
	}
	return parseStooqQuote(body, symbol, ticker)
}

func parseStooqQuote(body []byte, symbol string, ticker string) (Quote, error) {
	rows, err := csv.NewReader(strings.NewReader(string(body))).ReadAll()
	if err != nil {
		return Quote{}, fmt.Errorf("解析 Stooq 响应失败: %w", err)
	}
	if len(rows) < 2 || len(rows[1]) < 5 {
		return Quote{}, errors.New("Stooq 响应缺少有效数据")
	}
	priceText := strings.TrimSpace(rows[1][4])
	price, err := strconv.ParseFloat(priceText, 64)
	if err != nil || price <= 0 {
		return Quote{}, errors.New("Stooq 未返回有效价格")
	}
	name := strings.TrimSpace(rows[1][3])
	if name == "" || strings.EqualFold(name, "N/D") {
		name = ticker
	}
	return Quote{
		Symbol:      symbol,
		Name:        name,
		Price:       price,
		PriceDigits: decimalDigits(priceText),
	}, nil
}

func (p *EastMoneyQuoteProvider) fetchYahooQuotes(symbols []string) (map[string]Quote, error) {
	quotes := make(map[string]Quote, len(symbols))
	var failures []string
	for _, symbol := range symbols {
		quote, err := p.fetchSingleYahooQuote(symbol)
		if err != nil {
			failures = append(failures, fmt.Sprintf("%s: %v", symbol, err))
			continue
		}
		quotes[symbol] = quote
	}
	if len(quotes) > 0 {
		return quotes, nil
	}
	if len(failures) > 0 {
		return nil, fmt.Errorf("Yahoo 美股行情请求失败: %s", strings.Join(failures, "; "))
	}
	return nil, errors.New("Yahoo 美股行情请求失败")
}

func (p *EastMoneyQuoteProvider) fetchSingleYahooQuote(symbol string) (Quote, error) {
	ticker := usTickerFromSymbol(symbol)
	if ticker == "" {
		return Quote{}, errors.New("美股代码无效")
	}
	req, err := http.NewRequest(http.MethodGet, "https://query1.finance.yahoo.com/v8/finance/chart/"+strings.ReplaceAll(ticker, ".", "-"), nil)
	if err != nil {
		return Quote{}, err
	}
	q := req.URL.Query()
	q.Set("range", "1d")
	q.Set("interval", "1m")
	req.URL.RawQuery = q.Encode()
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Referer", "https://finance.yahoo.com/")
	req.Close = true

	resp, err := p.client.Do(req)
	if err != nil {
		return Quote{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Quote{}, fmt.Errorf("Yahoo 接口返回状态码 %d", resp.StatusCode)
	}
	var payload struct {
		Chart struct {
			Result []struct {
				Meta struct {
					Symbol             string  `json:"symbol"`
					RegularMarketPrice float64 `json:"regularMarketPrice"`
				} `json:"meta"`
			} `json:"result"`
			Error any `json:"error"`
		} `json:"chart"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return Quote{}, fmt.Errorf("解析 Yahoo 响应失败: %w", err)
	}
	if len(payload.Chart.Result) == 0 {
		return Quote{}, errors.New("Yahoo 响应缺少有效数据")
	}
	price := payload.Chart.Result[0].Meta.RegularMarketPrice
	if price <= 0 {
		return Quote{}, errors.New("Yahoo 未返回有效价格")
	}
	name := strings.TrimSpace(payload.Chart.Result[0].Meta.Symbol)
	if name == "" {
		name = ticker
	}
	return Quote{
		Symbol:      symbol,
		Name:        name,
		Price:       price,
		PriceDigits: decimalDigits(strconv.FormatFloat(price, 'f', -1, 64)),
	}, nil
}

func usTickerFromSymbol(symbol string) string {
	symbol = normalizeQuoteSymbol(symbol)
	if !strings.HasPrefix(symbol, "us") || len(symbol) <= 2 {
		return ""
	}
	return strings.ToUpper(symbol[2:])
}

func isDigits(value string) bool {
	for _, r := range value {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func parseEastMoneyQuotes(body []byte, symbolBySecID map[string]string) (map[string]Quote, error) {
	var payload struct {
		RC   int `json:"rc"`
		Data *struct {
			Diff []struct {
				Price  json.RawMessage `json:"f2"`
				Code   string          `json:"f12"`
				Market int             `json:"f13"`
				Name   string          `json:"f14"`
			} `json:"diff"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, fmt.Errorf("解析行情响应失败: %w", err)
	}
	if payload.RC != 0 || payload.Data == nil {
		return nil, errors.New("行情响应缺少有效数据")
	}

	quotes := make(map[string]Quote, len(payload.Data.Diff))
	for _, item := range payload.Data.Diff {
		secID := fmt.Sprintf("%d.%s", item.Market, item.Code)
		symbol := symbolBySecID[secID]
		if symbol == "" {
			continue
		}
		priceText := strings.Trim(string(item.Price), `"`)
		price, err := strconv.ParseFloat(priceText, 64)
		if err != nil {
			price = 0
		}
		quotes[symbol] = Quote{
			Symbol:      symbol,
			Name:        strings.TrimSpace(item.Name),
			Price:       price,
			PriceDigits: decimalDigits(priceText),
		}
	}
	return quotes, nil
}

func decodeGBK(body []byte) string {
	reader := transform.NewReader(bytes.NewReader(body), simplifiedchinese.GBK.NewDecoder())
	decoded, err := io.ReadAll(reader)
	if err != nil {
		return string(body)
	}
	return string(decoded)
}

func parseTencentLine(line string) (string, Quote, bool) {
	line = strings.TrimSpace(line)
	eq := strings.Index(line, "=")
	if eq <= 2 || !strings.HasPrefix(line, "v_") {
		return "", Quote{}, false
	}

	symbol := line[2:eq]
	firstQuote := strings.Index(line, "\"")
	lastQuote := strings.LastIndex(line, "\"")
	if firstQuote < 0 || lastQuote <= firstQuote {
		return "", Quote{}, false
	}

	payload := line[firstQuote+1 : lastQuote]
	fields := strings.Split(payload, "~")
	if len(fields) < 4 {
		return "", Quote{}, false
	}

	priceText := fields[3]
	price, err := strconv.ParseFloat(priceText, 64)
	if err != nil || price <= 0 {
		return "", Quote{}, false
	}

	name := fields[1]
	if !utf8.ValidString(name) {
		name = ""
	}

	return symbol, Quote{
		Symbol:      symbol,
		Name:        name,
		Price:       price,
		PriceDigits: decimalDigits(priceText),
	}, true
}

func decimalDigits(value string) int {
	value = strings.TrimSpace(value)
	dot := strings.Index(value, ".")
	if dot < 0 {
		return 2
	}
	digits := len(value) - dot - 1
	if digits < 2 {
		return 2
	}
	if digits > 4 {
		return 4
	}
	return digits
}

func (p *FrankfurterExchangeRateProvider) FetchHKDCNY() (float64, error) {
	return p.fetchCNYRate("HKD", "未获取到有效 HKD/CNY 汇率")
}

func (p *FrankfurterExchangeRateProvider) FetchUSDCNY() (float64, error) {
	return p.fetchCNYRate("USD", "未获取到有效 USD/CNY 汇率")
}

func (p *FrankfurterExchangeRateProvider) fetchCNYRate(base string, emptyMessage string) (float64, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.frankfurter.dev/v1/latest?base="+base+"&symbols=CNY", nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := p.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("汇率请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("汇率接口返回状态码 %d", resp.StatusCode)
	}

	var payload struct {
		Rates map[string]float64 `json:"rates"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return 0, fmt.Errorf("解析汇率响应失败: %w", err)
	}
	rate := payload.Rates["CNY"]
	if rate <= 0 {
		return 0, errors.New(emptyMessage)
	}
	return rate, nil
}
