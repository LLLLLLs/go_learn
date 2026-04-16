package main

import (
	"bytes"
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
	Symbol string
	Name   string
	Price  float64
}

type QuoteProvider interface {
	Fetch(symbols []string) (map[string]Quote, error)
}

type ExchangeRateProvider interface {
	FetchHKDCNY() (float64, error)
}

type TencentQuoteProvider struct {
	client *http.Client
}

type FrankfurterExchangeRateProvider struct {
	client *http.Client
}

func NewTencentQuoteProvider() *TencentQuoteProvider {
	return &TencentQuoteProvider{
		client: &http.Client{Timeout: 6 * time.Second},
	}
}

func NewFrankfurterExchangeRateProvider() *FrankfurterExchangeRateProvider {
	return &FrankfurterExchangeRateProvider{
		client: &http.Client{Timeout: 6 * time.Second},
	}
}

func (p *TencentQuoteProvider) Fetch(symbols []string) (map[string]Quote, error) {
	if len(symbols) == 0 {
		return map[string]Quote{}, nil
	}

	req, err := http.NewRequest(http.MethodGet, "https://qt.gtimg.cn/q="+strings.Join(symbols, ","), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Referer", "https://gu.qq.com/")

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("行情请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("行情接口返回状态码 %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取行情响应失败: %w", err)
	}

	decoded := decodeGBK(body)
	lines := strings.Split(decoded, ";")
	quotes := make(map[string]Quote, len(symbols))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		symbol, quote, ok := parseTencentLine(line)
		if !ok {
			continue
		}
		quotes[symbol] = quote
	}

	if len(quotes) == 0 {
		return nil, errors.New("没有解析到有效行情，可能是接口不可用")
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

	price, err := strconv.ParseFloat(fields[3], 64)
	if err != nil {
		return "", Quote{}, false
	}

	name := fields[1]
	if !utf8.ValidString(name) {
		name = ""
	}

	return symbol, Quote{
		Symbol: symbol,
		Name:   name,
		Price:  price,
	}, true
}

func (p *FrankfurterExchangeRateProvider) FetchHKDCNY() (float64, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.frankfurter.dev/v1/latest?base=HKD&symbols=CNY", nil)
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
		return 0, errors.New("未获取到有效 HKD/CNY 汇率")
	}
	return rate, nil
}
