package main

import (
	"crypto/rand"
	"embed"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	stockcalculator "stockCalculator"
)

//go:embed templates/*.html static/*.css
var webFS embed.FS

type server struct {
	mu            sync.Mutex
	store         *stockcalculator.UserStore
	portfolio     *stockcalculator.Portfolio
	statePath     string
	quotes        stockcalculator.QuoteProvider
	rates         stockcalculator.ExchangeRateProvider
	templates     *template.Template
	lastNotice    notice
	nextRefreshAt time.Time
	sessions      map[string]string
}

type notice struct {
	Type    string
	Message string
}

type pageData struct {
	Notice           notice
	CurrentUser      stockcalculator.UserProfile
	Summary          stockcalculator.Summary
	Status           string
	NextRefresh      string
	NextRefreshUnix  int64
	Holdings         []holdingView
	Records          []recordView
	FeeTotals        string
	RefreshSec       int
	RecentCodes      []string
	SellOptions      []sellOption
	HKDInput         string
	USDInput         string
	CNYBalance       string
	HKDBalance       string
	USDBalance       string
	CashBreakdown    string
	CapitalBreakdown string
}

type authPageData struct {
	Mode    string
	Notice  notice
	Account string
	Name    string
}

type holdingView struct {
	DisplayCode        string
	Name               string
	Market             string
	Currency           string
	Quantity           int
	AvgCostLocal       string
	CurrentPrice       string
	MarketValueLocal   string
	MarketValue        string
	UnrealizedPnLLocal string
	UnrealizedPnL      string
	ReturnPercent      string
}

type recordView struct {
	Time        string
	TypeLabel   string
	DisplayCode string
	Market      string
	Currency    string
	Quantity    int
	Price       string
	FXRate      string
	AmountBase  string
	Fee         string
	RealizedPnL string
}

type sellOption struct {
	Code     string
	Label    string
	Quantity int
}

type refreshResponse struct {
	Status          string        `json:"status"`
	NextRefresh     string        `json:"next_refresh"`
	NextRefreshUnix int64         `json:"next_refresh_unix"`
	Metrics         metricsView   `json:"metrics"`
	Holdings        []holdingView `json:"holdings"`
}

type metricsView struct {
	CapitalBreakdown string `json:"capital_breakdown"`
	CashBreakdown    string `json:"cash_breakdown"`
	MarketValue      string `json:"market_value"`
	TotalAssets      string `json:"total_assets"`
	RealizedPnL      string `json:"realized_pnl"`
	RealizedPnLClass string `json:"realized_pnl_class"`
	FeeTotals        string `json:"fee_totals"`
	UnrealizedPnL    string `json:"unrealized_pnl"`
	UnrealizedClass  string `json:"unrealized_class"`
	TotalReturn      string `json:"total_return"`
	TotalReturnClass string `json:"total_return_class"`
}

func main() {
	addr := flag.String("addr", defaultListenAddr(), "HTTP listen address")
	flag.Parse()

	srv, err := newServer()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("GET /static/", noCache(http.FileServer(http.FS(webFS))))
	mux.HandleFunc("GET /", srv.handleIndex)
	mux.HandleFunc("GET /login", srv.handleLoginPage)
	mux.HandleFunc("POST /login", srv.handleLogin)
	mux.HandleFunc("GET /register", srv.handleRegisterPage)
	mux.HandleFunc("POST /register", srv.handleRegister)
	mux.HandleFunc("POST /logout", srv.handleLogout)
	mux.HandleFunc("POST /account", srv.handleAccount)
	mux.HandleFunc("POST /trade", srv.handleTrade)
	mux.HandleFunc("POST /balance", srv.handleBalance)
	mux.HandleFunc("POST /settings", srv.handleSettings)
	mux.HandleFunc("POST /refresh", srv.handleRefresh)

	log.Printf("webapp listening on http://localhost%s", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux))
}

func defaultListenAddr() string {
	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		port = "8080"
	}
	if strings.Contains(port, ":") {
		return port
	}
	return ":" + port
}

func newServer() (*server, error) {
	tmpl, err := template.New("layout.html").Funcs(template.FuncMap{
		"pnlClass": pnlClass,
	}).ParseFS(webFS, "templates/*.html")
	if err != nil {
		return nil, err
	}

	store, err := stockcalculator.LoadUserStore()
	if err != nil {
		return nil, err
	}
	portfolio, statePath, err := stockcalculator.LoadPortfolioFileForUser(store)
	if err != nil {
		return nil, err
	}

	srv := &server{
		store:         store,
		portfolio:     portfolio,
		statePath:     statePath,
		quotes:        stockcalculator.NewEastMoneyQuoteProvider(),
		rates:         stockcalculator.NewFrankfurterExchangeRateProvider(),
		templates:     tmpl,
		nextRefreshAt: time.Now().Add(portfolio.RefreshInterval()),
		sessions:      make(map[string]string),
	}
	return srv, nil
}

func (s *server) handleIndex(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	setNoCacheHeaders(w)
	if !s.ensureAuthenticatedLocked(w, r) {
		return
	}
	data := s.pageDataLocked()
	s.lastNotice = notice{}
	if err := s.templates.ExecuteTemplate(w, "layout.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *server) handleLoginPage(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	setNoCacheHeaders(w)
	if s.currentSessionUserLocked(r) != "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if err := s.templates.ExecuteTemplate(w, "auth.html", authPageData{Mode: "login"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *server) handleRegisterPage(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	setNoCacheHeaders(w)
	if s.currentSessionUserLocked(r) != "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if err := s.templates.ExecuteTemplate(w, "auth.html", authPageData{Mode: "register"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *server) handleLogin(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	setNoCacheHeaders(w)
	if err := r.ParseForm(); err != nil {
		s.renderAuthLocked(w, authPageData{Mode: "login", Notice: errorNotice(err)})
		return
	}
	account := r.FormValue("account")
	user, err := s.store.Authenticate(account, r.FormValue("password"))
	if err != nil {
		s.renderAuthLocked(w, authPageData{Mode: "login", Account: account, Notice: errorNotice(err)})
		return
	}
	if err := s.activateUserLocked(user.ID); err != nil {
		s.renderAuthLocked(w, authPageData{Mode: "login", Account: account, Notice: errorNotice(err)})
		return
	}
	s.createSessionLocked(w, user.ID)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) handleRegister(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	setNoCacheHeaders(w)
	if err := r.ParseForm(); err != nil {
		s.renderAuthLocked(w, authPageData{Mode: "register", Notice: errorNotice(err)})
		return
	}
	account := r.FormValue("account")
	name := r.FormValue("name")
	user, err := s.store.RegisterUser(account, r.FormValue("password"), r.FormValue("password_confirm"), name)
	if err != nil {
		s.renderAuthLocked(w, authPageData{Mode: "register", Account: account, Name: name, Notice: errorNotice(err)})
		return
	}
	if err := s.activateUserLocked(user.ID); err != nil {
		s.renderAuthLocked(w, authPageData{Mode: "register", Account: account, Name: name, Notice: errorNotice(err)})
		return
	}
	s.createSessionLocked(w, user.ID)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) handleLogout(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if cookie, err := r.Cookie("stockcalc_session"); err == nil {
		delete(s.sessions, cookie.Value)
	}
	http.SetCookie(w, &http.Cookie{Name: "stockcalc_session", Value: "", Path: "/", MaxAge: -1, HttpOnly: true, SameSite: http.SameSiteLaxMode})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (s *server) renderAuthLocked(w http.ResponseWriter, data authPageData) {
	if err := s.templates.ExecuteTemplate(w, "auth.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *server) ensureAuthenticatedLocked(w http.ResponseWriter, r *http.Request) bool {
	userID := s.currentSessionUserLocked(r)
	if userID == "" {
		if wantsJSON(r) {
			http.Error(w, "请先登录", http.StatusUnauthorized)
			return false
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return false
	}
	if s.store.CurrentUser().ID == userID {
		return true
	}
	if err := s.activateUserLocked(userID); err != nil {
		deleteSessionCookie(w, r, s.sessions)
		if wantsJSON(r) {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return false
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return false
	}
	return true
}

func (s *server) currentSessionUserLocked(r *http.Request) string {
	cookie, err := r.Cookie("stockcalc_session")
	if err != nil {
		return ""
	}
	return s.sessions[cookie.Value]
}

func (s *server) createSessionLocked(w http.ResponseWriter, userID string) {
	token, err := randomToken()
	if err != nil {
		token = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	s.sessions[token] = userID
	http.SetCookie(w, &http.Cookie{
		Name:     "stockcalc_session",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

func (s *server) activateUserLocked(userID string) error {
	if _, err := s.store.SetCurrentUserByID(userID); err != nil {
		return err
	}
	return s.reloadPortfolioLocked()
}

func deleteSessionCookie(w http.ResponseWriter, r *http.Request, sessions map[string]string) {
	if cookie, err := r.Cookie("stockcalc_session"); err == nil {
		delete(sessions, cookie.Value)
	}
	http.SetCookie(w, &http.Cookie{Name: "stockcalc_session", Value: "", Path: "/", MaxAge: -1, HttpOnly: true, SameSite: http.SameSiteLaxMode})
}

func randomToken() (string, error) {
	raw := make([]byte, 32)
	if _, err := rand.Read(raw); err != nil {
		return "", err
	}
	return hex.EncodeToString(raw), nil
}

func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setNoCacheHeaders(w)
		next.ServeHTTP(w, r)
	})
}

func setNoCacheHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

func (s *server) handleTrade(w http.ResponseWriter, r *http.Request) {
	s.withMutation(w, r, func() notice {
		if err := r.ParseForm(); err != nil {
			return errorNotice(err)
		}
		quantity, err := parseIntField(r.FormValue("quantity"), "股数")
		if err != nil {
			return errorNotice(err)
		}
		price, err := parseFloatField(r.FormValue("price"), "价格")
		if err != nil {
			return errorNotice(err)
		}
		kind := stockcalculator.TradeTypeBuy
		if r.FormValue("type") == string(stockcalculator.TradeTypeSell) {
			kind = stockcalculator.TradeTypeSell
		}
		if err := s.portfolio.AddRecord(kind, r.FormValue("code"), quantity, price, r.FormValue("currency")); err != nil {
			return errorNotice(err)
		}
		if err := s.saveLocked(); err != nil {
			return errorNotice(err)
		}
		s.refreshQuotesLocked()
		return successNotice("交易记录已保存")
	})
}

func (s *server) handleBalance(w http.ResponseWriter, r *http.Request) {
	s.withMutation(w, r, func() notice {
		if err := r.ParseForm(); err != nil {
			return errorNotice(err)
		}
		amount, err := parseFloatField(r.FormValue("amount"), "余额调整金额")
		if err != nil {
			return errorNotice(err)
		}
		direction := 1.0
		if r.FormValue("direction") == "out" {
			direction = -1
		}
		if err := s.portfolio.AdjustBalanceByCurrency(amount, r.FormValue("currency"), direction); err != nil {
			return errorNotice(err)
		}
		if err := s.saveLocked(); err != nil {
			return errorNotice(err)
		}
		return successNotice("余额已更新")
	})
}

func (s *server) handleSettings(w http.ResponseWriter, r *http.Request) {
	s.withMutation(w, r, func() notice {
		if err := r.ParseForm(); err != nil {
			return errorNotice(err)
		}
		if hkd := strings.TrimSpace(r.FormValue("hkd_rate")); hkd != "" {
			rate, err := parseFloatField(hkd, "港币汇率")
			if err != nil {
				return errorNotice(err)
			}
			if err := s.portfolio.SetExchangeRate(rate); err != nil {
				return errorNotice(err)
			}
		}
		if usd := strings.TrimSpace(r.FormValue("usd_rate")); usd != "" {
			rate, err := parseFloatField(usd, "美元汇率")
			if err != nil {
				return errorNotice(err)
			}
			if err := s.portfolio.SetUSDExchangeRate(rate); err != nil {
				return errorNotice(err)
			}
		}
		refreshSec, err := parseIntField(r.FormValue("refresh_sec"), "刷新间隔")
		if err != nil {
			return errorNotice(err)
		}
		if err := s.portfolio.SetRefreshInterval(time.Duration(refreshSec) * time.Second); err != nil {
			return errorNotice(err)
		}
		s.nextRefreshAt = time.Now().Add(s.portfolio.RefreshInterval())
		if err := s.saveLocked(); err != nil {
			return errorNotice(err)
		}
		return successNotice("设置已保存")
	})
}

func (s *server) handleRefresh(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		s.respondRefreshError(w, r, err)
		return
	}

	s.mu.Lock()
	if !s.ensureAuthenticatedLocked(w, r) {
		s.mu.Unlock()
		return
	}
	if r.FormValue("auto") == "1" && time.Now().Before(s.nextRefreshAt) {
		response := s.refreshResponseLocked()
		s.mu.Unlock()
		s.writeRefreshResponse(w, r, response, nil)
		return
	}

	target := r.FormValue("target")
	var errs []error
	if target == "quotes" || target == "all" {
		if err := s.refreshQuotesLocked(); err != nil {
			errs = append(errs, err)
		}
	}
	if target == "rates" || target == "all" {
		if err := s.refreshRatesLocked(); err != nil {
			errs = append(errs, err)
		}
	}
	err := joinErrors(errs)
	if err != nil {
		s.lastNotice = errorNotice(err)
	}
	response := s.refreshResponseLocked()
	s.mu.Unlock()

	s.writeRefreshResponse(w, r, response, err)
}

func (s *server) handleAccount(w http.ResponseWriter, r *http.Request) {
	s.withMutation(w, r, func() notice {
		if err := r.ParseForm(); err != nil {
			return errorNotice(err)
		}
		user := s.store.CurrentUser()
		updated, err := s.store.UpdateUserProfile(user.ID, r.FormValue("name"), r.FormValue("current_password"), r.FormValue("new_password"), r.FormValue("new_password_confirm"))
		if err != nil {
			return errorNotice(err)
		}
		return successNotice(fmt.Sprintf("已更新用户 %s", updated.Name))
	})
}

func (s *server) withMutation(w http.ResponseWriter, r *http.Request, fn func() notice) {
	s.mu.Lock()
	if !s.ensureAuthenticatedLocked(w, r) {
		s.mu.Unlock()
		return
	}
	s.lastNotice = fn()
	s.mu.Unlock()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) pageDataLocked() pageData {
	summary := s.portfolio.Summary()
	return pageData{
		Notice:           s.lastNotice,
		CurrentUser:      s.store.CurrentUser(),
		Summary:          summary,
		Status:           buildStatus(summary),
		NextRefresh:      formatNextRefreshCountdown(s.nextRefreshAt, time.Now()),
		NextRefreshUnix:  s.nextRefreshAt.Unix(),
		Holdings:         buildHoldingViews(summary.Holdings),
		Records:          buildRecordViews(summary.Records),
		FeeTotals:        formatFeeMarketTotals(summary.FeeMarketTotals),
		RefreshSec:       int(summary.RefreshInterval / time.Second),
		RecentCodes:      recentCodes(summary.Records),
		SellOptions:      sellOptions(summary.Holdings),
		HKDInput:         fmt.Sprintf("%.4f", summary.HKDRate),
		USDInput:         fmt.Sprintf("%.4f", summary.USDRate),
		CNYBalance:       formatMoneyWithUnit("RMB", summary.CashCNY),
		HKDBalance:       formatMoneyWithUnit("港币", summary.CashHKD),
		USDBalance:       formatMoneyWithUnit("美元", summary.CashUSD),
		CashBreakdown:    formatCurrencyBreakdown(summary.CashCNY, summary.CashHKD, summary.CashUSD, summary.Cash),
		CapitalBreakdown: formatCurrencyBreakdown(summary.InitialCapitalCNY, summary.InitialCapitalHKD, summary.InitialCapitalUSD, summary.InitialCapital),
	}
}

func (s *server) refreshResponseLocked() refreshResponse {
	summary := s.portfolio.Summary()
	return refreshResponse{
		Status:          buildStatus(summary),
		NextRefresh:     formatNextRefreshCountdown(s.nextRefreshAt, time.Now()),
		NextRefreshUnix: s.nextRefreshAt.Unix(),
		Metrics:         buildMetricsView(summary),
		Holdings:        buildHoldingViews(summary.Holdings),
	}
}

func (s *server) writeRefreshResponse(w http.ResponseWriter, r *http.Request, response refreshResponse, err error) {
	if wantsJSON(r) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
		}
		_ = json.NewEncoder(w).Encode(response)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) respondRefreshError(w http.ResponseWriter, r *http.Request, err error) {
	if wantsJSON(r) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.mu.Lock()
	s.lastNotice = errorNotice(err)
	s.mu.Unlock()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func wantsJSON(r *http.Request) bool {
	return r.FormValue("partial") == "1" || strings.Contains(r.Header.Get("Accept"), "application/json") || r.Header.Get("X-Requested-With") == "fetch"
}

func (s *server) reloadPortfolioLocked() error {
	portfolio, statePath, err := stockcalculator.LoadPortfolioFileForUser(s.store)
	if err != nil {
		return err
	}
	s.portfolio = portfolio
	s.statePath = statePath
	s.nextRefreshAt = time.Now().Add(portfolio.RefreshInterval())
	return nil
}

func (s *server) refreshMarketData(force bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.refreshMarketDataLocked(force)
}

func (s *server) refreshMarketDataLocked(force bool) {
	now := time.Now()
	if force {
		_ = s.refreshQuotesLocked()
		_ = s.refreshRatesLocked()
		return
	}
	if _, ok := s.portfolio.NextQuoteRefreshFromCache(s.portfolio.RefreshInterval(), now); !ok {
		_ = s.refreshQuotesLocked()
	}
	if _, ok := s.portfolio.NextFXRefreshFromCache(s.portfolio.RefreshInterval(), now); !ok {
		_ = s.refreshRatesLocked()
	}
}

func (s *server) refreshQuotesLocked() error {
	err := s.portfolio.RefreshQuotes(s.quotes)
	s.nextRefreshAt = time.Now().Add(s.portfolio.RefreshInterval())
	if saveErr := s.saveLocked(); saveErr != nil && err == nil {
		err = saveErr
	}
	return err
}

func (s *server) refreshRatesLocked() error {
	var errs []error
	now := time.Now()
	if rate, err := s.rates.FetchHKDCNY(); err == nil {
		if setErr := s.portfolio.SetExchangeRateWithTime(rate, now); setErr != nil {
			errs = append(errs, setErr)
		}
	} else {
		errs = append(errs, fmt.Errorf("HKD/CNY: %w", err))
	}
	if rate, err := s.rates.FetchUSDCNY(); err == nil {
		if setErr := s.portfolio.SetUSDExchangeRateWithTime(rate, now); setErr != nil {
			errs = append(errs, setErr)
		}
	} else {
		errs = append(errs, fmt.Errorf("USD/CNY: %w", err))
	}
	if saveErr := s.saveLocked(); saveErr != nil {
		errs = append(errs, saveErr)
	}
	return joinErrors(errs)
}

func (s *server) saveLocked() error {
	return s.portfolio.SaveToFile(s.statePath)
}

func buildHoldingViews(holdings []stockcalculator.HoldingSummary) []holdingView {
	items := make([]holdingView, 0, len(holdings))
	for _, holding := range holdings {
		name := strings.TrimSpace(holding.Name)
		if name == "" {
			name = "-"
		}
		items = append(items, holdingView{
			DisplayCode:        holding.DisplayCode,
			Name:               name,
			Market:             holding.Market,
			Currency:           holding.Currency,
			Quantity:           holding.Quantity,
			AvgCostLocal:       formatPrice(holding.AvgCostLocal, holding.CurrentPriceDigits),
			CurrentPrice:       formatPrice(holding.CurrentPrice, holding.CurrentPriceDigits),
			MarketValueLocal:   formatMoneyWithUnit(holding.Currency, holding.MarketValueLocal),
			MarketValue:        formatMoney(holding.MarketValue),
			UnrealizedPnLLocal: formatMoneyWithUnit(holding.Currency, holding.UnrealizedPnLLocal),
			UnrealizedPnL:      formatMoney(holding.UnrealizedPnL),
			ReturnPercent:      formatPercent(calculateHoldingReturn(holding)),
		})
	}
	return items
}

func buildMetricsView(summary stockcalculator.Summary) metricsView {
	realized := formatMoney(summary.RealizedPnL)
	unrealized := formatMoney(summary.UnrealizedPnL)
	totalReturn := formatMoney(summary.TotalReturn)
	return metricsView{
		CapitalBreakdown: formatCurrencyBreakdown(summary.InitialCapitalCNY, summary.InitialCapitalHKD, summary.InitialCapitalUSD, summary.InitialCapital),
		CashBreakdown:    formatCurrencyBreakdown(summary.CashCNY, summary.CashHKD, summary.CashUSD, summary.Cash),
		MarketValue:      formatMoney(summary.MarketValue),
		TotalAssets:      formatMoney(summary.TotalAssets),
		RealizedPnL:      realized,
		RealizedPnLClass: pnlClass(realized),
		FeeTotals:        formatFeeMarketTotals(summary.FeeMarketTotals),
		UnrealizedPnL:    unrealized,
		UnrealizedClass:  pnlClass(unrealized),
		TotalReturn:      totalReturn,
		TotalReturnClass: pnlClass(totalReturn),
	}
}

func buildRecordViews(records []stockcalculator.TradeSummary) []recordView {
	items := make([]recordView, 0, len(records))
	for _, record := range records {
		items = append(items, recordView{
			Time:        record.Time.Format("2006-01-02 15:04:05"),
			TypeLabel:   record.TypeLabel,
			DisplayCode: record.DisplayCode,
			Market:      record.Market,
			Currency:    record.Currency,
			Quantity:    record.Quantity,
			Price:       fmt.Sprintf("%.2f", record.Price),
			FXRate:      fmt.Sprintf("%.4f", record.FXRate),
			AmountBase:  formatMoney(record.AmountBase),
			Fee:         formatMoneyWithUnit(record.Currency, record.Fee),
			RealizedPnL: formatTradeRealizedPnL(record),
		})
	}
	return items
}

func recentCodes(records []stockcalculator.TradeSummary) []string {
	seen := make(map[string]bool, 5)
	items := make([]string, 0, 5)
	for _, record := range records {
		code := strings.TrimSpace(record.DisplayCode)
		if code == "" || seen[code] {
			continue
		}
		seen[code] = true
		items = append(items, code)
		if len(items) == 5 {
			break
		}
	}
	return items
}

func sellOptions(holdings []stockcalculator.HoldingSummary) []sellOption {
	items := make([]sellOption, 0, len(holdings))
	for _, holding := range holdings {
		if holding.Quantity <= 0 {
			continue
		}
		label := holding.DisplayCode
		if name := strings.TrimSpace(holding.Name); name != "" && name != "-" {
			label = fmt.Sprintf("%s | %s | %d股", holding.DisplayCode, name, holding.Quantity)
		} else {
			label = fmt.Sprintf("%s | %d股", holding.DisplayCode, holding.Quantity)
		}
		items = append(items, sellOption{Code: holding.DisplayCode, Label: label, Quantity: holding.Quantity})
	}
	return items
}

func calculateHoldingReturn(holding stockcalculator.HoldingSummary) float64 {
	cost := holding.AvgCostLocal * float64(holding.Quantity)
	if cost == 0 {
		return 0
	}
	return holding.UnrealizedPnLLocal / cost
}

func buildStatus(summary stockcalculator.Summary) string {
	if len(summary.Holdings) == 0 {
		return fmt.Sprintf("暂无持仓。港股按 %.4f、美股按 %.4f 折算为人民币。", summary.HKDRate, summary.USDRate)
	}
	if summary.LastRefreshAt.IsZero() {
		return "已存在持仓，等待首次行情刷新..."
	}
	if summary.LastRefreshErr != "" {
		return fmt.Sprintf("最近刷新: %s，失败原因: %s", summary.LastRefreshAt.Format("2006-01-02 15:04:05"), summary.LastRefreshErr)
	}
	return fmt.Sprintf("最近刷新: %s，每 %s 自动更新一次，HKD/CNY %.4f，USD/CNY %.4f",
		summary.LastRefreshAt.Format("2006-01-02 15:04:05"), formatRefreshInterval(summary.RefreshInterval), summary.HKDRate, summary.USDRate)
}

func parseFloatField(raw string, field string) (float64, error) {
	value, err := strconv.ParseFloat(strings.TrimSpace(raw), 64)
	if err != nil {
		return 0, fmt.Errorf("%s格式不正确", field)
	}
	return value, nil
}

func parseIntField(raw string, field string) (int, error) {
	value, err := strconv.Atoi(strings.TrimSpace(raw))
	if err != nil {
		return 0, fmt.Errorf("%s格式不正确", field)
	}
	return value, nil
}

func formatMoney(value float64) string {
	return fmt.Sprintf("¥%.2f", value)
}

func formatMoneyWithUnit(currency string, value float64) string {
	switch strings.ToUpper(strings.TrimSpace(currency)) {
	case "HKD", "港币":
		return fmt.Sprintf("HK$%.2f", value)
	case "USD", "美元":
		return fmt.Sprintf("$%.2f", value)
	default:
		return fmt.Sprintf("¥%.2f", value)
	}
}

func formatCurrencyBreakdown(cny, hkd, usd, total float64) string {
	return fmt.Sprintf("RMB %s / HKD %s / USD %s / 合计 %s", formatMoney(cny), formatMoneyWithUnit("HKD", hkd), formatMoneyWithUnit("USD", usd), formatMoney(total))
}

func formatTradeRealizedPnL(record stockcalculator.TradeSummary) string {
	if record.TypeLabel != "平仓" {
		return ""
	}
	return formatMoney(record.RealizedPnL)
}

func formatFeeMarketTotals(fees []stockcalculator.FeeMarketTotal) string {
	parts := make([]string, 0, len(fees))
	for _, fee := range fees {
		if fee.Amount > 1e-6 {
			parts = append(parts, fmt.Sprintf("%s: %s", fee.Market, formatMoneyWithUnit(fee.Currency, fee.Amount)))
		}
	}
	if len(parts) == 0 {
		return "--"
	}
	return strings.Join(parts, " / ")
}

func formatPercent(value float64) string {
	return fmt.Sprintf("%.2f%%", value*100)
}

func formatRefreshInterval(interval time.Duration) string {
	if interval <= 0 {
		interval = 60 * time.Second
	}
	if interval%time.Minute == 0 {
		return fmt.Sprintf("%d 分钟", int(interval/time.Minute))
	}
	return fmt.Sprintf("%d 秒", int(interval/time.Second))
}

func formatNextRefreshCountdown(next time.Time, now time.Time) string {
	if next.IsZero() {
		return "下次更新: --"
	}
	remaining := next.Sub(now)
	if remaining < 0 {
		remaining = 0
	}
	seconds := int(remaining.Round(time.Second) / time.Second)
	if seconds >= 60 {
		return fmt.Sprintf("下次更新: %d分%02d秒", seconds/60, seconds%60)
	}
	return fmt.Sprintf("下次更新: %d秒", seconds)
}

func formatPrice(value float64, digits int) string {
	if digits < 2 {
		digits = 2
	}
	if digits > 4 {
		digits = 4
	}
	return strconv.FormatFloat(value, 'f', digits, 64)
}

func pnlClass(value string) string {
	clean := strings.TrimSpace(value)
	if clean == "" {
		return ""
	}
	if strings.Contains(clean, "-") {
		return "loss"
	}
	normalized := strings.NewReplacer("HK", "", "¥", "", "$", "", "%", "", ",", "", " ", "").Replace(clean)
	amount, err := strconv.ParseFloat(normalized, 64)
	if err != nil || amount <= 0 {
		return ""
	}
	return "gain"
}

func successNotice(message string) notice {
	return notice{Type: "success", Message: message}
}

func errorNotice(err error) notice {
	if err == nil {
		return notice{}
	}
	return notice{Type: "error", Message: err.Error()}
}

func joinErrors(errs []error) error {
	filtered := make([]error, 0, len(errs))
	for _, err := range errs {
		if err != nil {
			filtered = append(filtered, err)
		}
	}
	if len(filtered) == 0 {
		return nil
	}
	return errors.Join(filtered...)
}
