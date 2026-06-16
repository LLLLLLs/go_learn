package stockcalculator

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Quote struct {
	Symbol               string
	Name                 string
	Price                float64
	PriceDigits          int
	OpenPrice            float64
	OpenPriceDigits      int
	ExtendedPrice        float64
	ExtendedPriceDigits  int
	ExtendedPriceChecked bool
}

type QuoteProvider interface {
	Fetch(symbols []string) (map[string]Quote, error)
}

type ExchangeRateProvider interface {
	FetchHKDCNY() (float64, error)
	FetchUSDCNY() (float64, error)
	FetchKRWCNY() (float64, error)
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

func (p *EastMoneyQuoteProvider) Fetch(symbols []string) (result map[string]Quote, err error) {
	start := time.Now()
	log.Printf("[quotes] fetch start symbols=%v", symbols)
	defer func() {
		log.Printf("[quotes] fetch done elapsed=%s count=%d err=%v", time.Since(start), len(result), err)
	}()

	if len(symbols) == 0 {
		return map[string]Quote{}, nil
	}

	normalizedInput := normalizeQuoteSymbols(symbols)
	var fetchFailures []string

	quotes := make(map[string]Quote, len(symbols))
	tencentSymbols := filterTencentFallbackSymbols(normalizedInput)
	if len(tencentSymbols) > 0 {
		tencentQuotes, err := p.fetchTencentFallbackQuotes(tencentSymbols)
		if err != nil {
			fetchFailures = append(fetchFailures, err.Error())
		}
		mergeValidQuotes(quotes, tencentQuotes)
		mergeValidQuotes(quotes, p.fetchYahooSupplementalUSQuotes(validUSQuoteSymbols(normalizedInput, quotes)))
		if hasValidQuotesForSymbols(normalizedInput, quotes) {
			return quotes, nil
		}
		log.Printf("[quotes] tencent unavailable or incomplete, switching providers symbols=%v missing=%v", normalizedInput, missingQuoteSymbols(normalizedInput, quotes))
	}

	eastMoneySymbols := missingQuoteSymbols(normalizedInput, quotes)
	secIDs, symbolBySecID, failures := p.splitQuoteSymbols(eastMoneySymbols)
	if len(secIDs) > 0 {
		eastMoneyQuotes, err := p.fetchEastMoneyQuotes(secIDs, symbolBySecID)
		if err != nil {
			fetchFailures = append(fetchFailures, err.Error())
		}
		mergeValidQuotes(quotes, eastMoneyQuotes)
	}
	mergeValidQuotes(quotes, p.fetchYahooSupplementalUSQuotes(validUSQuoteSymbols(normalizedInput, quotes)))

	yahooSymbols := filterYahooFallbackSymbols(missingQuoteSymbols(normalizedInput, quotes))
	if len(yahooSymbols) > 0 {
		yahooQuotes, err := p.fetchYahooQuotes(yahooSymbols)
		if err != nil {
			fetchFailures = append(fetchFailures, err.Error())
		}
		mergeValidQuotes(quotes, yahooQuotes)
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

func mergeValidQuotes(target map[string]Quote, source map[string]Quote) {
	for symbol, quote := range source {
		if quote.Price <= 0 {
			continue
		}
		target[symbol] = mergeQuote(target[symbol], quote)
	}
}

func mergeQuote(base Quote, next Quote) Quote {
	if next.Symbol != "" {
		base.Symbol = next.Symbol
	}
	if next.Name != "" {
		base.Name = next.Name
	}
	if next.Price > 0 {
		base.Price = next.Price
		base.PriceDigits = next.PriceDigits
	}
	if next.OpenPrice > 0 {
		base.OpenPrice = next.OpenPrice
		base.OpenPriceDigits = next.OpenPriceDigits
	}
	if next.ExtendedPriceChecked {
		base.ExtendedPrice = next.ExtendedPrice
		base.ExtendedPriceDigits = next.ExtendedPriceDigits
		base.ExtendedPriceChecked = true
	}
	return base
}

func hasValidQuotesForSymbols(symbols []string, quotes map[string]Quote) bool {
	if len(quotes) == 0 {
		return false
	}
	for _, symbol := range symbols {
		if quotes[symbol].Price <= 0 {
			return false
		}
	}
	return true
}

func validUSQuoteSymbols(symbols []string, quotes map[string]Quote) []string {
	filtered := make([]string, 0, len(symbols))
	for _, symbol := range symbols {
		if !strings.HasPrefix(symbol, "us") {
			continue
		}
		if quotes[symbol].Price > 0 {
			filtered = append(filtered, symbol)
		}
	}
	return filtered
}

func normalizeQuoteSymbols(symbols []string) []string {
	normalized := make([]string, 0, len(symbols))
	for _, symbol := range symbols {
		item := normalizeQuoteSymbol(symbol)
		if item != "" {
			normalized = append(normalized, item)
		}
	}
	return normalized
}

func missingQuoteSymbols(symbols []string, quotes map[string]Quote) []string {
	missing := make([]string, 0, len(symbols))
	for _, symbol := range symbols {
		if quotes[symbol].Price <= 0 {
			missing = append(missing, symbol)
		}
	}
	return missing
}

func filterTencentFallbackSymbols(symbols []string) []string {
	filtered := make([]string, 0, len(symbols))
	for _, symbol := range symbols {
		if strings.HasPrefix(symbol, "sh") || strings.HasPrefix(symbol, "sz") ||
			strings.HasPrefix(symbol, "hk") || strings.HasPrefix(symbol, "us") {
			filtered = append(filtered, symbol)
		}
	}
	return filtered
}

func filterYahooFallbackSymbols(symbols []string) []string {
	filtered := make([]string, 0, len(symbols))
	for _, symbol := range symbols {
		if strings.HasPrefix(symbol, "hk") {
			continue
		}
		if strings.HasPrefix(symbol, "sh") || strings.HasPrefix(symbol, "sz") || strings.HasPrefix(symbol, "us") ||
			strings.HasPrefix(symbol, "ks") || strings.HasPrefix(symbol, "kq") {
			filtered = append(filtered, symbol)
		}
	}
	return filtered
}

func filterUSSymbols(symbols []string) []string {
	filtered := make([]string, 0, len(symbols))
	for _, symbol := range symbols {
		if strings.HasPrefix(symbol, "us") {
			filtered = append(filtered, symbol)
		}
	}
	return filtered
}

func (p *EastMoneyQuoteProvider) fetchYahooSupplementalUSQuotes(symbols []string) map[string]Quote {
	usSymbols := filterUSSymbols(symbols)
	if len(usSymbols) == 0 {
		return nil
	}
	quotes, err := p.fetchYahooQuotes(usSymbols)
	if err != nil {
		log.Printf("[quotes] yahoo supplemental us failed err=%v", err)
		return nil
	}
	return quotes
}

func (p *EastMoneyQuoteProvider) splitQuoteSymbols(symbols []string) ([]string, map[string]string, []string) {
	secIDs := make([]string, 0, len(symbols))
	symbolBySecID := make(map[string]string, len(symbols))
	var failures []string
	for _, symbol := range symbols {
		normalized := normalizeQuoteSymbol(symbol)
		if normalized == "" {
			continue
		}
		if strings.HasPrefix(normalized, "us") {
			secIDsForUS, err := eastMoneyUSSecIDCandidates(normalized)
			if err != nil {
				failures = append(failures, fmt.Sprintf("%s: %v", symbol, err))
				continue
			}
			secIDs = append(secIDs, secIDsForUS...)
			for _, secID := range secIDsForUS {
				symbolBySecID[secID] = normalized
			}
			continue
		}
		secID, err := toEastMoneySecID(normalized)
		if err != nil {
			failures = append(failures, fmt.Sprintf("%s: %v", symbol, err))
			continue
		}
		secIDs = append(secIDs, secID)
		symbolBySecID[secID] = normalized
	}
	return secIDs, symbolBySecID, failures
}

func (p *EastMoneyQuoteProvider) fetchEastMoneyQuotes(secIDs []string, symbolBySecID map[string]string) (map[string]Quote, error) {
	start := time.Now()
	log.Printf("[quotes] eastmoney batch start secids=%s", strings.Join(secIDs, ","))
	body, err := p.fetchEastMoneyBody(secIDs)
	if err == nil {
		quotes, parseErr := parseEastMoneyQuotes(body, symbolBySecID)
		log.Printf("[quotes] eastmoney batch done elapsed=%s count=%d err=%v", time.Since(start), len(quotes), parseErr)
		return quotes, parseErr
	}
	if len(secIDs) == 1 {
		log.Printf("[quotes] eastmoney batch failed elapsed=%s err=%v", time.Since(start), err)
		return nil, err
	}

	log.Printf("[quotes] eastmoney batch failed elapsed=%s err=%v", time.Since(start), err)
	return nil, err
}

func normalizeQuoteSymbol(symbol string) string {
	return strings.ToLower(strings.TrimSpace(symbol))
}

func (p *EastMoneyQuoteProvider) fetchTencentFallbackQuotes(symbols []string) (map[string]Quote, error) {
	start := time.Now()
	requestSymbols, symbolByRequest := toTencentSymbols(symbols)
	if len(requestSymbols) == 0 {
		return map[string]Quote{}, nil
	}

	log.Printf("[quotes] tencent fallback start symbols=%v requests=%v", symbols, requestSymbols)
	rawQuotes, err := p.fetchTencentQuotes(requestSymbols)
	if err != nil {
		log.Printf("[quotes] tencent fallback failed elapsed=%s err=%v", time.Since(start), err)
		return nil, err
	}

	quotes := make(map[string]Quote, len(rawQuotes))
	for requestSymbol, quote := range rawQuotes {
		symbol := symbolByRequest[normalizeQuoteSymbol(requestSymbol)]
		if symbol == "" {
			symbol = normalizeQuoteSymbol(requestSymbol)
		}
		quote.Symbol = symbol
		quotes[symbol] = quote
	}
	log.Printf("[quotes] tencent fallback done elapsed=%s count=%d", time.Since(start), len(quotes))
	return quotes, nil
}

func toTencentSymbols(symbols []string) ([]string, map[string]string) {
	requestSymbols := make([]string, 0, len(symbols))
	symbolByRequest := make(map[string]string, len(symbols))
	for _, symbol := range symbols {
		requestSymbol, err := toTencentSymbol(symbol)
		if err != nil {
			continue
		}
		requestSymbols = append(requestSymbols, requestSymbol)
		symbolByRequest[normalizeQuoteSymbol(requestSymbol)] = symbol
	}
	return requestSymbols, symbolByRequest
}

func toTencentSymbol(symbol string) (string, error) {
	symbol = normalizeQuoteSymbol(symbol)
	if len(symbol) < 3 {
		return "", errors.New("股票代码无效")
	}
	prefix := symbol[:2]
	code := symbol[2:]
	switch prefix {
	case "sh", "sz", "hk":
		return symbol, nil
	case "us":
		ticker := strings.ToUpper(code)
		if ticker == "" {
			return "", errors.New("美股代码无效")
		}
		return "us" + ticker, nil
	default:
		return "", errors.New("腾讯暂不支持该行情代码")
	}
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
			symbol = normalizeQuoteSymbol(symbol)
			quote.Symbol = symbol
			quotes[symbol] = quote
		}
	}
	return quotes, nil
}

func (p *EastMoneyQuoteProvider) fetchEastMoneyBody(secIDs []string) ([]byte, error) {
	if len(p.baseURLs) == 0 {
		return nil, errors.New("东方财富行情地址未配置")
	}
	baseURL := p.baseURLs[0]
	start := time.Now()
	log.Printf("[quotes] eastmoney request start base=%s secids=%s", baseURL, strings.Join(secIDs, ","))
	body, err := p.fetchEastMoneyBodyOnce(baseURL, secIDs)
	if err == nil {
		log.Printf("[quotes] eastmoney request done base=%s elapsed=%s bytes=%d", baseURL, time.Since(start), len(body))
		return body, nil
	}
	log.Printf("[quotes] eastmoney request failed base=%s elapsed=%s err=%v", baseURL, time.Since(start), err)
	return nil, fmt.Errorf("行情请求失败: %w", err)
}

func (p *EastMoneyQuoteProvider) fetchEastMoneyBodyOnce(baseURL string, secIDs []string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set("fltt", "2")
	q.Set("fields", "f12,f13,f14,f2,f18")
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

func eastMoneyUSSecIDCandidates(symbol string) ([]string, error) {
	symbol = normalizeQuoteSymbol(symbol)
	if !strings.HasPrefix(symbol, "us") || len(symbol) <= 2 {
		return nil, errors.New("美股代码无效")
	}
	ticker := strings.ToUpper(symbol[2:])
	return []string{
		"105." + ticker,
		"106." + ticker,
		"107." + ticker,
	}, nil
}

func (p *EastMoneyQuoteProvider) fetchYahooQuotes(symbols []string) (map[string]Quote, error) {
	start := time.Now()
	log.Printf("[quotes] yahoo start symbols=%v", symbols)
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
		log.Printf("[quotes] yahoo done elapsed=%s count=%d failures=%d", time.Since(start), len(quotes), len(failures))
		return quotes, nil
	}
	if len(failures) > 0 {
		err := fmt.Errorf("Yahoo 美股行情请求失败: %s", strings.Join(failures, "; "))
		log.Printf("[quotes] yahoo failed elapsed=%s err=%v", time.Since(start), err)
		return nil, err
	}
	err := errors.New("Yahoo 美股行情请求失败")
	log.Printf("[quotes] yahoo failed elapsed=%s err=%v", time.Since(start), err)
	return nil, err
}

func (p *EastMoneyQuoteProvider) fetchSingleYahooQuote(symbol string) (Quote, error) {
	start := time.Now()
	yahooSymbol, err := toYahooSymbol(symbol)
	if err != nil {
		return Quote{}, err
	}
	req, err := http.NewRequest(http.MethodGet, "https://query1.finance.yahoo.com/v8/finance/chart/"+yahooSymbol, nil)
	if err != nil {
		return Quote{}, err
	}
	q := req.URL.Query()
	q.Set("range", "1d")
	q.Set("interval", "1m")
	if strings.HasPrefix(normalizeQuoteSymbol(symbol), "us") {
		q.Set("range", "5d")
		q.Set("includePrePost", "true")
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Referer", "https://finance.yahoo.com/")
	req.Close = true

	client := *p.client
	if client.Timeout == 0 || client.Timeout > 2*time.Second {
		client.Timeout = 2 * time.Second
	}
	log.Printf("[quotes] yahoo item start symbol=%s url=%s", symbol, req.URL.String())
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[quotes] yahoo item failed symbol=%s elapsed=%s err=%v", symbol, time.Since(start), err)
		return Quote{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("Yahoo 接口返回状态码 %d", resp.StatusCode)
		log.Printf("[quotes] yahoo item failed symbol=%s elapsed=%s err=%v", symbol, time.Since(start), err)
		return Quote{}, err
	}
	var payload struct {
		Chart struct {
			Result []struct {
				Meta struct {
					Symbol                     string              `json:"symbol"`
					LongName                   string              `json:"longName"`
					ShortName                  string              `json:"shortName"`
					RegularMarketPrice         float64             `json:"regularMarketPrice"`
					RegularMarketPreviousClose float64             `json:"regularMarketPreviousClose"`
					ChartPreviousClose         float64             `json:"chartPreviousClose"`
					PreviousClose              float64             `json:"previousClose"`
					CurrentTradingPeriod       yahooTradingPeriods `json:"currentTradingPeriod"`
				} `json:"meta"`
				Timestamp  []int64 `json:"timestamp"`
				Indicators struct {
					Quote []struct {
						Close []float64 `json:"close"`
					} `json:"quote"`
				} `json:"indicators"`
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
	isUSSymbol := strings.HasPrefix(normalizeQuoteSymbol(symbol), "us")
	extendedPrice := 0.0
	if len(payload.Chart.Result[0].Indicators.Quote) > 0 {
		closes := payload.Chart.Result[0].Indicators.Quote[0].Close
		periods := payload.Chart.Result[0].Meta.CurrentTradingPeriod
		extendedPrice = latestExtendedMarketPrice(payload.Chart.Result[0].Timestamp, closes, periods)
	}
	openPrice := payload.Chart.Result[0].Meta.RegularMarketPreviousClose
	if openPrice <= 0 {
		openPrice = payload.Chart.Result[0].Meta.PreviousClose
	}
	if openPrice <= 0 {
		openPrice = payload.Chart.Result[0].Meta.ChartPreviousClose
	}
	name := yahooDisplayName(symbol, payload.Chart.Result[0].Meta.LongName, payload.Chart.Result[0].Meta.ShortName, payload.Chart.Result[0].Meta.Symbol)
	if name == "" {
		name = yahooSymbol
	}
	quote := Quote{
		Symbol:      symbol,
		Name:        name,
		Price:       price,
		PriceDigits: decimalDigits(strconv.FormatFloat(price, 'f', -1, 64)),
	}
	if openPrice > 0 {
		quote.OpenPrice = openPrice
		quote.OpenPriceDigits = decimalDigits(strconv.FormatFloat(openPrice, 'f', -1, 64))
	}
	quote.ExtendedPriceChecked = isUSSymbol
	if extendedPrice > 0 {
		quote.ExtendedPrice = extendedPrice
		quote.ExtendedPriceDigits = decimalDigits(strconv.FormatFloat(extendedPrice, 'f', -1, 64))
	}
	log.Printf("[quotes] yahoo item done symbol=%s elapsed=%s price=%v", symbol, time.Since(start), quote.Price)
	return quote, nil
}

type yahooPeriod struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

type yahooTradingPeriods struct {
	Pre     yahooPeriod `json:"pre"`
	Regular yahooPeriod `json:"regular"`
	Post    yahooPeriod `json:"post"`
}

func priceInYahooPeriod(timestamps []int64, prices []float64, period yahooPeriod) float64 {
	if period.Start <= 0 || period.End <= 0 || len(timestamps) == 0 || len(prices) == 0 {
		return 0
	}
	limit := len(timestamps)
	if len(prices) < limit {
		limit = len(prices)
	}
	for i := limit - 1; i >= 0; i-- {
		ts := timestamps[i]
		if ts < period.Start || ts > period.End {
			continue
		}
		if prices[i] > 0 {
			return prices[i]
		}
	}
	return 0
}

func latestExtendedMarketPrice(timestamps []int64, prices []float64, periods yahooTradingPeriods) float64 {
	if periods.Regular.Start <= 0 || len(timestamps) == 0 || len(prices) == 0 {
		return 0
	}
	limit := len(timestamps)
	if len(prices) < limit {
		limit = len(prices)
	}
	for i := limit - 1; i >= 0; i-- {
		if prices[i] <= 0 {
			continue
		}
		switch yahooSession(timestamps[i], periods) {
		case "pre", "post":
			return prices[i]
		case "regular":
			return 0
		}
	}
	return 0
}

func yahooSession(timestamp int64, periods yahooTradingPeriods) string {
	if inYahooPeriodOffset(timestamp, periods.Regular, periods.Regular.Start) {
		return "regular"
	}
	if inYahooPeriodOffset(timestamp, periods.Pre, periods.Regular.Start) {
		return "pre"
	}
	if inYahooPeriodOffset(timestamp, periods.Post, periods.Regular.Start) {
		return "post"
	}
	return ""
}

func inYahooPeriodOffset(timestamp int64, period yahooPeriod, anchor int64) bool {
	if period.Start <= 0 || period.End <= 0 || period.End < period.Start || anchor <= 0 {
		return false
	}
	const daySeconds = 24 * 60 * 60
	offset := positiveModulo(timestamp-anchor, daySeconds)
	start := positiveModulo(period.Start-anchor, daySeconds)
	end := positiveModulo(period.End-anchor, daySeconds)
	if start <= end {
		return offset >= start && offset <= end
	}
	return offset >= start || offset <= end
}

func positiveModulo(value int64, mod int64) int64 {
	if mod <= 0 {
		return 0
	}
	value %= mod
	if value < 0 {
		value += mod
	}
	return value
}

func yahooDisplayName(symbol string, names ...string) string {
	if name := localizedQuoteName(symbol); name != "" {
		return name
	}
	for _, name := range names {
		name = strings.TrimSpace(name)
		if name != "" {
			return name
		}
	}
	return ""
}

func localizedQuoteName(symbol string) string {
	switch normalizeQuoteSymbol(symbol) {
	case "usaaoi":
		return "应用光电"
	case "usmrvl":
		return "迈威尔科技"
	case "usmu":
		return "美光科技"
	case "usnok":
		return "诺基亚"
	case "usspcx":
		return "SpaceX"
	case "ks000660":
		return "SK海力士"
	case "ks005930":
		return "三星电子"
	case "ks373220":
		return "LG新能源"
	case "ks207940":
		return "三星生物制剂"
	case "ks005380":
		return "现代汽车"
	case "ks000270":
		return "起亚"
	case "kq035720":
		return "Kakao"
	case "ks035420":
		return "NAVER"
	default:
		return ""
	}
}

func toYahooSymbol(symbol string) (string, error) {
	symbol = normalizeQuoteSymbol(symbol)
	if len(symbol) < 3 {
		return "", errors.New("股票代码无效")
	}
	prefix := symbol[:2]
	code := symbol[2:]
	switch prefix {
	case "us":
		ticker := strings.ToUpper(code)
		if ticker == "" {
			return "", errors.New("美股代码无效")
		}
		return strings.ReplaceAll(ticker, ".", "-"), nil
	case "hk":
		if len(code) == 0 || len(code) > 5 || !isDigits(code) {
			return "", errors.New("港股代码无效")
		}
		return fmt.Sprintf("%05s.HK", code), nil
	case "sh":
		if len(code) != 6 || !isDigits(code) {
			return "", errors.New("A 股代码无效")
		}
		return code + ".SS", nil
	case "sz":
		if len(code) != 6 || !isDigits(code) {
			return "", errors.New("A 股代码无效")
		}
		return code + ".SZ", nil
	case "ks":
		if len(code) != 6 || !isDigits(code) {
			return "", errors.New("韩股代码无效")
		}
		return code + ".KS", nil
	case "kq":
		if len(code) != 6 || !isDigits(code) {
			return "", errors.New("韩股代码无效")
		}
		return code + ".KQ", nil
	default:
		return "", errors.New("Yahoo 暂不支持该行情代码")
	}
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
				Open   json.RawMessage `json:"f18"`
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
		openText := strings.Trim(string(item.Open), `"`)
		openPrice, err := strconv.ParseFloat(openText, 64)
		if err != nil {
			openPrice = 0
		}
		quotes[symbol] = Quote{
			Symbol:          symbol,
			Name:            strings.TrimSpace(item.Name),
			Price:           price,
			PriceDigits:     decimalDigits(priceText),
			OpenPrice:       openPrice,
			OpenPriceDigits: decimalDigits(openText),
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
	openPrice := 0.0
	openDigits := 0
	if len(fields) > 4 {
		if value, err := strconv.ParseFloat(fields[4], 64); err == nil && value > 0 {
			openPrice = value
			openDigits = decimalDigits(fields[4])
		}
	}

	name := fields[1]
	if !utf8.ValidString(name) {
		name = ""
	}

	return symbol, Quote{
		Symbol:          symbol,
		Name:            name,
		Price:           price,
		PriceDigits:     decimalDigits(priceText),
		OpenPrice:       openPrice,
		OpenPriceDigits: openDigits,
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

func (p *FrankfurterExchangeRateProvider) FetchKRWCNY() (float64, error) {
	return p.fetchCNYRate("KRW", "未获取到有效 KRW/CNY 汇率")
}

func (p *FrankfurterExchangeRateProvider) fetchCNYRate(base string, emptyMessage string) (float64, error) {
	start := time.Now()
	log.Printf("[fx] %s/CNY start", base)
	req, err := http.NewRequest(http.MethodGet, "https://api.frankfurter.dev/v1/latest?base="+base+"&symbols=CNY", nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := p.client.Do(req)
	if err != nil {
		log.Printf("[fx] %s/CNY failed elapsed=%s err=%v", base, time.Since(start), err)
		return 0, fmt.Errorf("汇率请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("汇率接口返回状态码 %d", resp.StatusCode)
		log.Printf("[fx] %s/CNY failed elapsed=%s err=%v", base, time.Since(start), err)
		return 0, err
	}

	var payload struct {
		Rates map[string]float64 `json:"rates"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return 0, fmt.Errorf("解析汇率响应失败: %w", err)
	}
	rate := payload.Rates["CNY"]
	if rate <= 0 {
		err := errors.New(emptyMessage)
		log.Printf("[fx] %s/CNY failed elapsed=%s err=%v", base, time.Since(start), err)
		return 0, err
	}
	log.Printf("[fx] %s/CNY done elapsed=%s rate=%v", base, time.Since(start), rate)
	return rate, nil
}
