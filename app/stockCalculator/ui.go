package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

type UI struct {
	app          fyne.App
	window       fyne.Window
	portfolio    *Portfolio
	provider     QuoteProvider
	rateProvider ExchangeRateProvider
	userStore    *UserStore
	statePath    string
	currentUser  UserProfile

	userSelect            *widget.Select
	balanceAdjustEntry    *widget.Entry
	balanceCurrencySelect *widget.Select
	balanceCNYLabel       *widget.Label
	balanceHKDLabel       *widget.Label
	rateInfoLabel         *widget.Label
	codeEntry             *widget.Entry
	recentCodeSelect      *widget.Select
	recentCodeOptions     []string
	recentCodeMap         map[string]string
	sellHoldingOptions    []string
	sellHoldingMap        map[string]HoldingSummary
	quantityEntry         *widget.Entry
	quantitySlider        *widget.Slider
	quantityHintLabel     *widget.Label
	priceEntry            *widget.Entry
	priceCurrencySelect   *widget.Select
	tradeTypeSelect       *widget.Select
	statusLabel           *widget.Label

	summaryLabels map[string]*widget.Label
	holdingsTable *widget.Table
	recordsTable  *widget.Table

	tableMu       sync.RWMutex
	holdingsRows  [][]string
	recordsRows   [][]string
	stopRefreshCh chan struct{}

	suppressUserSelection bool
	suppressQuantitySync  bool
}

func NewUI(app fyne.App) (*UI, error) {
	userStore, err := LoadUserStore()
	if err != nil {
		return nil, err
	}

	portfolio, statePath, err := LoadPortfolioForUser(userStore, app.Preferences())
	if err != nil {
		return nil, err
	}

	ui := &UI{
		app:            app,
		window:         app.NewWindow("股票收益追踪计算器"),
		portfolio:      portfolio,
		provider:       NewTencentQuoteProvider(),
		rateProvider:   NewFrankfurterExchangeRateProvider(),
		userStore:      userStore,
		statePath:      statePath,
		currentUser:    userStore.CurrentUser(),
		recentCodeMap:  make(map[string]string),
		sellHoldingMap: make(map[string]HoldingSummary),
		summaryLabels:  make(map[string]*widget.Label),
		stopRefreshCh:  make(chan struct{}),
	}

	ui.build()
	ui.refreshView()
	return ui, nil
}

func (ui *UI) Run() {
	ui.window.Resize(fyne.NewSize(1320, 920))
	ui.window.CenterOnScreen()
	ui.window.SetMaster()

	go ui.startRefreshLoop()
	ui.window.ShowAndRun()
}

func (ui *UI) build() {
	ui.window.SetOnClosed(func() {
		select {
		case <-ui.stopRefreshCh:
		default:
			close(ui.stopRefreshCh)
		}
	})

	settingsPanel := ui.buildSettingsPanel()
	tradePanel := ui.buildTradePanel()
	summaryPanel := ui.buildSummaryPanel()
	tables := ui.buildTables()

	top := container.NewGridWithColumns(2, settingsPanel, tradePanel)
	split := container.NewVSplit(top, tables)
	split.SetOffset(0.33)
	content := container.NewBorder(summaryPanel, ui.statusLabelWidget(), nil, nil,
		split,
	)

	ui.window.SetContent(container.NewPadded(content))
}

func (ui *UI) buildSettingsPanel() fyne.CanvasObject {
	userPanel := ui.buildUserPanel()

	ui.balanceAdjustEntry = widget.NewEntry()
	ui.balanceAdjustEntry.SetPlaceHolder("输入调整金额")
	ui.balanceCurrencySelect = widget.NewSelect([]string{"RMB", "港币"}, nil)
	ui.balanceCurrencySelect.SetSelected("RMB")
	ui.balanceCNYLabel = widget.NewLabel("RMB: --")
	ui.balanceHKDLabel = widget.NewLabel("港币: --")
	ui.rateInfoLabel = widget.NewLabel("汇率: 加载中...")

	addButton := widget.NewButton("增加", func() {
		ui.adjustBalance(1)
	})
	reduceButton := widget.NewButton("减少", func() {
		ui.adjustBalance(-1)
	})
	adjustRow := container.NewBorder(nil, nil, nil, container.NewHBox(addButton, reduceButton), ui.balanceAdjustEntry)
	currencyRow := container.NewBorder(nil, nil, nil, ui.rateInfoLabel, ui.balanceCurrencySelect)

	form := widget.NewForm(
		widget.NewFormItem("余额调整", adjustRow),
		widget.NewFormItem("余额币种", currencyRow),
		widget.NewFormItem("RMB余额", ui.balanceCNYLabel),
		widget.NewFormItem("港币余额", ui.balanceHKDLabel),
	)

	card := widget.NewCard("余额管理", "汇率会在打开应用时联网获取，余额调整会影响累计收益基准",
		container.NewVBox(userPanel, form))
	return card
}

func (ui *UI) buildUserPanel() fyne.CanvasObject {
	ui.userSelect = widget.NewSelect(ui.userStore.UserOptions(), func(selected string) {
		if ui.suppressUserSelection {
			return
		}
		if selected == createUserOptionName {
			ui.showCreateUserDialog()
			return
		}
		if err := ui.switchUser(selected); err != nil {
			dialog.ShowError(err, ui.window)
			ui.refreshUserOptions(ui.currentUser.Name)
		}
	})
	ui.refreshUserOptions(ui.currentUser.Name)

	renameButton := widget.NewButton("改名", func() {
		ui.showRenameUserDialog()
	})
	userRow := container.NewBorder(nil, nil, nil, renameButton, ui.userSelect)

	form := widget.NewForm(widget.NewFormItem("用户列表", userRow))
	return widget.NewCard("用户切换", "每个用户的数据单独保存在 data/users/<user-id> 目录中",
		form)
}

func (ui *UI) buildTradePanel() fyne.CanvasObject {
	ui.tradeTypeSelect = widget.NewSelect([]string{"买入", "平仓"}, func(string) {
		ui.updateTradeMode()
	})
	ui.tradeTypeSelect.SetSelected("买入")

	ui.codeEntry = widget.NewEntry()
	ui.codeEntry.SetPlaceHolder("A股如 600519 / sz000001，港股如 00700")
	ui.codeEntry.OnChanged = func(value string) {
		ui.updatePriceCurrencyByCode(value)
	}

	ui.recentCodeSelect = widget.NewSelect(nil, func(selected string) {
		if strings.TrimSpace(selected) == "" {
			return
		}
		if ui.tradeTypeSelect.Selected == "平仓" {
			ui.applySellHoldingSelection(selected)
			return
		}
		code := ui.recentCodeMap[selected]
		if code == "" {
			code = selected
		}
		ui.codeEntry.SetText(code)
	})
	ui.recentCodeSelect.PlaceHolder = "最近使用"
	ui.recentCodeSelect.Disable()

	ui.quantityEntry = widget.NewEntry()
	ui.quantityEntry.SetPlaceHolder("例如 100")
	ui.quantityEntry.OnChanged = ui.syncQuantityFromEntry

	ui.quantitySlider = widget.NewSlider(0, 0)
	ui.quantitySlider.Step = 1
	ui.quantitySlider.OnChanged = ui.syncQuantityFromSlider
	ui.quantityHintLabel = widget.NewLabel("买入模式下不限制股数")

	ui.priceEntry = widget.NewEntry()
	ui.priceEntry.SetPlaceHolder("按本币输入，例如 1530.50 / 398.20")

	ui.priceCurrencySelect = widget.NewSelect([]string{"港币", "RMB"}, nil)
	ui.priceCurrencySelect.SetSelected("港币")

	submitButton := widget.NewButton("添加记录", func() {
		quantity, err := parseIntField(ui.quantityEntry.Text, "股数")
		if err != nil {
			dialog.ShowError(err, ui.window)
			return
		}
		price, err := parseFloatField(ui.priceEntry.Text, "价格")
		if err != nil {
			dialog.ShowError(err, ui.window)
			return
		}

		tradeType := TradeTypeBuy
		if ui.tradeTypeSelect.Selected == "平仓" {
			tradeType = TradeTypeSell
		}

		if err := ui.portfolio.AddRecord(tradeType, ui.codeEntry.Text, quantity, price, ui.priceCurrencySelect.Selected); err != nil {
			dialog.ShowError(err, ui.window)
			return
		}
		if err := ui.portfolio.SaveToFile(ui.statePath); err != nil {
			dialog.ShowError(err, ui.window)
			return
		}

		ui.codeEntry.SetText("")
		ui.quantityEntry.SetText("")
		ui.priceEntry.SetText("")
		ui.priceCurrencySelect.SetSelected("港币")
		ui.refreshView()
		go ui.refreshQuotes()
	})

	refreshButton := widget.NewButton("立即刷新行情", func() {
		go ui.refreshQuotes()
	})

	form := widget.NewForm(
		widget.NewFormItem("交易类型", ui.tradeTypeSelect),
		widget.NewFormItem("股票代码", container.NewBorder(nil, nil, nil, ui.recentCodeSelect, ui.codeEntry)),
		widget.NewFormItem("股数", ui.quantityEntry),
		widget.NewFormItem("股数进度", container.NewBorder(nil, nil, nil, ui.quantityHintLabel, ui.quantitySlider)),
		widget.NewFormItem("股价币种", ui.priceCurrencySelect),
		widget.NewFormItem("价格(本币)", ui.priceEntry),
	)

	card := widget.NewCard(
		"交易记录",
		"支持手动录入买入和平仓，A 股自动识别 sh/sz/bj，港股自动识别 hk",
		container.NewVBox(form, container.NewHBox(submitButton, refreshButton)),
	)
	return card
}

func (ui *UI) buildSummaryPanel() fyne.CanvasObject {
	fields := []struct {
		key   string
		title string
		unit  string
	}{
		{"initialCapital", "净入金", "RMB/HKD"},
		{"cash", "现金余额", "RMB/HKD"},
		{"marketValue", "持仓市值", "CNY"},
		{"totalAssets", "总资产", "CNY"},
		{"realizedPnL", "已实现盈亏", "CNY"},
		{"unrealizedPnL", "浮动盈亏", "CNY"},
		{"totalReturn", "累计收益", "CNY"},
	}

	cards := make([]fyne.CanvasObject, 0, len(fields))
	for _, item := range fields {
		value := widget.NewLabel("--")
		value.Wrapping = fyne.TextWrapWord
		ui.summaryLabels[item.key] = value
		cards = append(cards, widget.NewCard(item.title, item.unit, value))
	}

	return container.NewGridWithColumns(len(cards), cards...)
}

func (ui *UI) buildTables() fyne.CanvasObject {
	ui.holdingsTable = widget.NewTable(
		func() (int, int) {
			ui.tableMu.RLock()
			defer ui.tableMu.RUnlock()
			return len(ui.holdingsRows), ui.columnCount(ui.holdingsRows)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			ui.tableMu.RLock()
			defer ui.tableMu.RUnlock()
			label.SetText(ui.cellText(ui.holdingsRows, id.Row, id.Col))
		},
	)
	ui.holdingsTable.OnSelected = func(id widget.TableCellID) {
		ui.holdingsTable.Unselect(id)
	}

	ui.recordsTable = widget.NewTable(
		func() (int, int) {
			ui.tableMu.RLock()
			defer ui.tableMu.RUnlock()
			return len(ui.recordsRows), ui.columnCount(ui.recordsRows)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			ui.tableMu.RLock()
			defer ui.tableMu.RUnlock()
			label.SetText(ui.cellText(ui.recordsRows, id.Row, id.Col))
		},
	)

	ui.setColumnWidths()

	holdingsCard := widget.NewCard("当前持仓", "10 秒自动刷新现价，持仓表按股票本币展示，总资产仍按人民币汇总",
		container.NewMax(ui.holdingsTable))
	recordsCard := widget.NewCard("交易流水", "港股交易按录入当时的汇率折算到人民币",
		container.NewMax(ui.recordsTable))

	tabs := container.NewAppTabs(
		container.NewTabItem("持仓", holdingsCard),
		container.NewTabItem("流水", recordsCard),
	)
	return tabs
}

func (ui *UI) statusLabelWidget() fyne.CanvasObject {
	ui.statusLabel = widget.NewLabel("等待行情刷新...")
	return ui.statusLabel
}

func (ui *UI) setColumnWidths() {
	widths1 := []int{110, 180, 70, 60, 110, 110, 120, 120}
	for idx, width := range widths1 {
		ui.holdingsTable.SetColumnWidth(idx, width)
	}

	widths2 := []int{145, 70, 110, 70, 70, 60, 95, 70, 120}
	for idx, width := range widths2 {
		ui.recordsTable.SetColumnWidth(idx, width)
	}
}

func (ui *UI) refreshView() {
	summary := ui.portfolio.Summary()
	ui.window.SetTitle(fmt.Sprintf("股票收益追踪计算器 - %s", ui.currentUser.Name))

	ui.summaryLabels["initialCapital"].SetText(formatCurrencyBreakdown("RMB", summary.InitialCapitalCNY, "港币", summary.InitialCapitalHKD, summary.InitialCapital))
	ui.summaryLabels["cash"].SetText(formatCurrencyBreakdown("RMB", summary.CashCNY, "港币", summary.CashHKD, summary.Cash))
	ui.summaryLabels["marketValue"].SetText(formatMoney(summary.MarketValue))
	ui.summaryLabels["totalAssets"].SetText(formatMoney(summary.TotalAssets))
	ui.summaryLabels["realizedPnL"].SetText(formatMoney(summary.RealizedPnL))
	ui.summaryLabels["unrealizedPnL"].SetText(formatMoney(summary.UnrealizedPnL))
	ui.summaryLabels["totalReturn"].SetText(formatMoney(summary.TotalReturn))
	if ui.balanceCNYLabel != nil {
		ui.balanceCNYLabel.SetText(formatMoneyWithUnit("RMB", summary.CashCNY))
	}
	if ui.balanceHKDLabel != nil {
		ui.balanceHKDLabel.SetText(formatMoneyWithUnit("港币", summary.CashHKD))
	}

	ui.tableMu.Lock()
	ui.holdingsRows = buildHoldingsRows(summary.Holdings)
	ui.recordsRows = buildRecordRows(summary.Records)
	ui.tableMu.Unlock()
	ui.refreshRecentCodeOptions(summary.Records)
	ui.refreshSellHoldingOptions(summary.Holdings)
	ui.updateTradeMode()

	ui.holdingsTable.Refresh()
	ui.recordsTable.Refresh()
	ui.statusLabel.SetText(buildStatus(summary))
	if ui.rateInfoLabel != nil {
		ui.rateInfoLabel.SetText(fmt.Sprintf("汇率: %.4f", summary.HKDRate))
	}
}

func (ui *UI) startRefreshLoop() {
	ui.refreshMarketData()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ui.refreshQuotes()
		case <-ui.stopRefreshCh:
			return
		}
	}
}

func (ui *UI) refreshQuotes() {
	_ = ui.portfolio.RefreshQuotes(ui.provider)
	ui.refreshView()
}

func (ui *UI) refreshMarketData() {
	if ui.rateProvider != nil {
		if rate, err := ui.rateProvider.FetchHKDCNY(); err == nil {
			if err := ui.portfolio.SetExchangeRate(rate); err == nil {
				_ = ui.portfolio.SaveToFile(ui.statePath)
			}
		}
	}
	ui.refreshQuotes()
}

func (ui *UI) columnCount(rows [][]string) int {
	if len(rows) == 0 {
		return 0
	}
	return len(rows[0])
}

func (ui *UI) cellText(rows [][]string, row, col int) string {
	if row < 0 || row >= len(rows) {
		return ""
	}
	if col < 0 || col >= len(rows[row]) {
		return ""
	}
	return rows[row][col]
}

func buildHoldingsRows(holdings []HoldingSummary) [][]string {
	rows := [][]string{{
		"代码",
		"名称",
		"市场",
		"股数",
		"成本价",
		"现价",
		"市值",
		"浮盈亏",
	}}

	for _, holding := range holdings {
		name := holding.Name
		if strings.TrimSpace(name) == "" {
			name = "-"
		}
		rows = append(rows, []string{
			holding.DisplayCode,
			name,
			holding.Market,
			strconv.Itoa(holding.Quantity),
			fmt.Sprintf("%.2f %s", holding.AvgCostLocal, holding.Currency),
			fmt.Sprintf("%.2f %s", holding.CurrentPrice, holding.Currency),
			formatMoneyWithUnit(holding.Currency, holding.MarketValueLocal),
			formatMoneyWithUnit(holding.Currency, holding.UnrealizedPnLLocal),
		})
	}

	return rows
}

func buildRecordRows(records []TradeSummary) [][]string {
	rows := [][]string{{
		"时间",
		"类型",
		"代码",
		"市场",
		"币种",
		"股数",
		"价格",
		"汇率",
		"金额(CNY)",
	}}

	for _, record := range records {
		rows = append(rows, []string{
			record.Time.Format("2006-01-02 15:04:05"),
			record.TypeLabel,
			record.DisplayCode,
			record.Market,
			record.Currency,
			strconv.Itoa(record.Quantity),
			fmt.Sprintf("%.2f", record.Price),
			fmt.Sprintf("%.4f", record.FXRate),
			formatMoney(record.AmountBase),
		})
	}

	return rows
}

func buildStatus(summary Summary) string {
	if len(summary.Holdings) == 0 {
		return fmt.Sprintf("暂无持仓。港股按当前汇率 %.4f 折算为人民币。", summary.HKDRate)
	}
	if summary.LastRefreshAt.IsZero() {
		return "已存在持仓，等待首次行情刷新..."
	}
	if summary.LastRefreshErr != "" {
		return fmt.Sprintf("最近刷新: %s，失败原因: %s", summary.LastRefreshAt.Format("2006-01-02 15:04:05"), summary.LastRefreshErr)
	}
	return fmt.Sprintf("最近刷新: %s，每 10 秒自动更新一次，当前港币汇率 %.4f",
		summary.LastRefreshAt.Format("2006-01-02 15:04:05"), summary.HKDRate)
}

func (ui *UI) switchUser(name string) error {
	user, err := ui.userStore.SetCurrentUserByName(name)
	if err != nil {
		return err
	}

	portfolio, statePath, err := LoadPortfolioForUser(ui.userStore, ui.app.Preferences())
	if err != nil {
		return err
	}

	ui.currentUser = user
	ui.portfolio = portfolio
	ui.statePath = statePath
	ui.tradeTypeSelect.SetSelected("买入")
	ui.codeEntry.SetText("")
	if ui.recentCodeSelect != nil {
		ui.recentCodeSelect.ClearSelected()
	}
	ui.quantityEntry.SetText("")
	if ui.quantitySlider != nil {
		ui.quantitySlider.SetValue(0)
	}
	ui.priceEntry.SetText("")
	ui.priceCurrencySelect.SetSelected("港币")
	if ui.balanceAdjustEntry != nil {
		ui.balanceAdjustEntry.SetText("")
	}
	if ui.balanceCurrencySelect != nil {
		ui.balanceCurrencySelect.SetSelected("RMB")
	}
	ui.refreshUserOptions(user.Name)
	ui.refreshView()
	go ui.refreshMarketData()
	return nil
}

func (ui *UI) refreshUserOptions(selected string) {
	ui.suppressUserSelection = true
	ui.userSelect.Options = ui.userStore.UserOptions()
	ui.userSelect.Refresh()
	ui.userSelect.SetSelected(selected)
	ui.suppressUserSelection = false
}

func (ui *UI) showCreateUserDialog() {
	confirmed := false
	entryDialog := dialog.NewEntryDialog("创建新用户", "用户名", func(value string) {
		user, err := ui.userStore.CreateUser(value)
		if err != nil {
			dialog.ShowError(err, ui.window)
			ui.refreshUserOptions(ui.currentUser.Name)
			return
		}
		confirmed = true
		if err := ui.switchUser(user.Name); err != nil {
			dialog.ShowError(err, ui.window)
		}
	}, ui.window)
	entryDialog.SetPlaceholder("输入不重复的用户名")
	entryDialog.SetOnClosed(func() {
		if !confirmed {
			ui.refreshUserOptions(ui.currentUser.Name)
		}
	})
	entryDialog.Show()
}

func (ui *UI) showRenameUserDialog() {
	original := ui.currentUser
	entryDialog := dialog.NewEntryDialog("重命名用户", "新用户名", func(value string) {
		user, err := ui.userStore.RenameUser(original.ID, value)
		if err != nil {
			dialog.ShowError(err, ui.window)
			return
		}
		ui.currentUser = user
		ui.refreshUserOptions(user.Name)
		ui.refreshView()
	}, ui.window)
	entryDialog.SetText(original.Name)
	entryDialog.SetPlaceholder("输入不重复的用户名")
	entryDialog.Show()
}

func (ui *UI) updatePriceCurrencyByCode(rawCode string) {
	security, err := normalizeSecurity(rawCode)
	if err != nil {
		return
	}
	if security.Currency == "CNY" {
		ui.priceCurrencySelect.SetSelected("RMB")
		return
	}
	ui.priceCurrencySelect.SetSelected("港币")
}

func (ui *UI) refreshSellHoldingOptions(holdings []HoldingSummary) {
	options := make([]string, 0, len(holdings))
	ui.sellHoldingMap = make(map[string]HoldingSummary, len(holdings))
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
		options = append(options, label)
		ui.sellHoldingMap[label] = holding
	}
	ui.sellHoldingOptions = options
}

func (ui *UI) refreshRecentCodeOptions(records []TradeSummary) {
	seen := make(map[string]bool, 5)
	options := make([]string, 0, 5)
	ui.recentCodeMap = make(map[string]string, 5)
	for _, record := range records {
		code := strings.TrimSpace(record.DisplayCode)
		if code == "" || seen[code] {
			continue
		}
		seen[code] = true
		label := code
		if name := ui.lookupNameByCode(code); name != "" && name != "-" {
			label = fmt.Sprintf("%s | %s", code, name)
		}
		options = append(options, label)
		ui.recentCodeMap[label] = code
		if len(options) == 5 {
			break
		}
	}
	ui.recentCodeOptions = options
}

func (ui *UI) updateTradeMode() {
	if ui.tradeTypeSelect == nil || ui.recentCodeSelect == nil || ui.codeEntry == nil || ui.quantitySlider == nil || ui.quantityHintLabel == nil {
		return
	}

	if ui.tradeTypeSelect.Selected == "平仓" {
		currentSelected := strings.TrimSpace(ui.recentCodeSelect.Selected)
		currentCode := strings.TrimSpace(ui.codeEntry.Text)

		ui.codeEntry.Disable()
		ui.quantityEntry.Enable()
		ui.recentCodeSelect.PlaceHolder = "选择当前持仓"
		ui.recentCodeSelect.Options = ui.sellHoldingOptions
		ui.recentCodeSelect.Refresh()
		if len(ui.sellHoldingOptions) == 0 {
			ui.recentCodeSelect.Disable()
			ui.recentCodeSelect.ClearSelected()
			ui.quantityEntry.Disable()
			ui.codeEntry.SetText("")
			ui.quantityEntry.SetText("")
			ui.configureQuantitySlider(0)
			ui.quantityHintLabel.SetText("暂无可平仓持仓")
			return
		}

		ui.recentCodeSelect.Enable()
		ui.quantityHintLabel.SetText("选择持仓后可拖动，最多为当前持股数")
		if matched := ui.findSellHoldingOption(currentSelected, currentCode); matched != "" {
			ui.recentCodeSelect.SetSelected(matched)
		} else {
			ui.recentCodeSelect.SetSelected(ui.sellHoldingOptions[0])
		}
		return
	}

	ui.codeEntry.Enable()
	ui.quantityEntry.Enable()
	ui.recentCodeSelect.PlaceHolder = "最近使用"
	ui.recentCodeSelect.Options = ui.recentCodeOptions
	ui.recentCodeSelect.ClearSelected()
	if len(ui.recentCodeOptions) == 0 {
		ui.recentCodeSelect.Disable()
	} else {
		ui.recentCodeSelect.Enable()
	}
	ui.recentCodeSelect.Refresh()
	ui.configureQuantitySlider(0)
	ui.quantityHintLabel.SetText("买入模式下不限制股数")
}

func (ui *UI) findSellHoldingOption(currentSelected string, currentCode string) string {
	for _, option := range ui.sellHoldingOptions {
		if option == currentSelected {
			return option
		}
	}
	if currentCode == "" {
		return ""
	}
	for _, option := range ui.sellHoldingOptions {
		holding, ok := ui.sellHoldingMap[option]
		if !ok {
			continue
		}
		if strings.EqualFold(strings.TrimSpace(holding.DisplayCode), currentCode) {
			return option
		}
	}
	return ""
}

func (ui *UI) applySellHoldingSelection(selected string) {
	holding, ok := ui.sellHoldingMap[selected]
	if !ok {
		return
	}
	ui.codeEntry.SetText(holding.DisplayCode)
	ui.updatePriceCurrencyByCode(holding.DisplayCode)
	ui.configureQuantitySlider(holding.Quantity)
	ui.quantityHintLabel.SetText(fmt.Sprintf("最多可平仓 %d 股", holding.Quantity))
	if strings.TrimSpace(ui.quantityEntry.Text) == "" {
		ui.setQuantityValue(holding.Quantity)
		return
	}
	ui.syncQuantityFromEntry(ui.quantityEntry.Text)
}

func (ui *UI) configureQuantitySlider(maxQuantity int) {
	if ui.quantitySlider == nil {
		return
	}
	ui.suppressQuantitySync = true
	ui.quantitySlider.Min = 0
	ui.quantitySlider.Max = float64(maxQuantity)
	if maxQuantity <= 0 {
		ui.quantitySlider.SetValue(0)
		ui.suppressQuantitySync = false
		ui.quantitySlider.Refresh()
		return
	}
	current := ui.quantitySlider.Value
	if current > float64(maxQuantity) {
		current = float64(maxQuantity)
	}
	ui.quantitySlider.SetValue(current)
	ui.suppressQuantitySync = false
	ui.quantitySlider.Refresh()
}

func (ui *UI) syncQuantityFromEntry(value string) {
	if ui.suppressQuantitySync || ui.tradeTypeSelect == nil || ui.tradeTypeSelect.Selected != "平仓" || ui.quantitySlider == nil {
		return
	}
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		ui.suppressQuantitySync = true
		ui.quantitySlider.SetValue(0)
		ui.suppressQuantitySync = false
		return
	}
	quantity, err := strconv.Atoi(trimmed)
	if err != nil {
		return
	}
	maxQuantity := int(ui.quantitySlider.Max)
	if maxQuantity > 0 && quantity > maxQuantity {
		quantity = maxQuantity
		ui.setQuantityValue(quantity)
		return
	}
	if quantity < 0 {
		quantity = 0
		ui.setQuantityValue(quantity)
		return
	}
	ui.suppressQuantitySync = true
	ui.quantitySlider.SetValue(float64(quantity))
	ui.suppressQuantitySync = false
}

func (ui *UI) syncQuantityFromSlider(value float64) {
	if ui.suppressQuantitySync || ui.tradeTypeSelect == nil || ui.tradeTypeSelect.Selected != "平仓" {
		return
	}
	ui.setQuantityValue(int(value + 0.5))
}

func (ui *UI) setQuantityValue(quantity int) {
	if ui.quantityEntry == nil {
		return
	}
	ui.suppressQuantitySync = true
	ui.quantityEntry.SetText(strconv.Itoa(quantity))
	if ui.quantitySlider != nil {
		ui.quantitySlider.SetValue(float64(quantity))
	}
	ui.suppressQuantitySync = false
}

func (ui *UI) lookupNameByCode(code string) string {
	ui.tableMu.RLock()
	defer ui.tableMu.RUnlock()

	for i := 1; i < len(ui.holdingsRows); i++ {
		row := ui.holdingsRows[i]
		if len(row) > 1 && row[0] == code {
			return row[1]
		}
	}
	return ""
}

func (ui *UI) adjustBalance(direction float64) {
	amount, err := parseFloatField(ui.balanceAdjustEntry.Text, "余额调整金额")
	if err != nil {
		dialog.ShowError(err, ui.window)
		return
	}
	if amount <= 0 {
		dialog.ShowError(fmt.Errorf("余额调整金额必须大于 0"), ui.window)
		return
	}
	if err := ui.portfolio.AdjustBalanceByCurrency(amount, ui.balanceCurrencySelect.Selected, direction); err != nil {
		dialog.ShowError(err, ui.window)
		return
	}
	if err := ui.portfolio.SaveToFile(ui.statePath); err != nil {
		dialog.ShowError(err, ui.window)
		return
	}
	ui.balanceAdjustEntry.SetText("")
	ui.balanceCurrencySelect.SetSelected("RMB")
	ui.refreshView()
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

func formatMoneyWithUnit(unit string, value float64) string {
	switch unit {
	case "港币", "HKD":
		return fmt.Sprintf("HK$%.2f", value)
	default:
		return fmt.Sprintf("¥%.2f", value)
	}
}

func formatCurrencyBreakdown(primaryUnit string, primaryValue float64, secondaryUnit string, secondaryValue float64, totalCNY float64) string {
	return fmt.Sprintf("%s  %s\n%s  %s\n折合  %s",
		primaryUnit, formatMoneyWithUnit(primaryUnit, primaryValue),
		secondaryUnit, formatMoneyWithUnit(secondaryUnit, secondaryValue),
		formatMoney(totalCNY),
	)
}

func formatNumber(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}
