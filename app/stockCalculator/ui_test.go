package stockcalculator

import (
	"testing"
	"time"
)

func TestFormatNextRefreshCountdown(t *testing.T) {
	now := time.Date(2026, 6, 3, 12, 0, 0, 0, time.Local)
	tests := []struct {
		name string
		next time.Time
		want string
	}{
		{name: "zero", next: time.Time{}, want: "下次更新: --"},
		{name: "seconds", next: now.Add(25 * time.Second), want: "下次更新: 25秒"},
		{name: "minutes", next: now.Add(90 * time.Second), want: "下次更新: 1分30秒"},
		{name: "expired", next: now.Add(-time.Second), want: "下次更新: 0秒"},
	}

	for _, tt := range tests {
		if got := formatNextRefreshCountdown(tt.next, now); got != tt.want {
			t.Fatalf("%s: formatNextRefreshCountdown = %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestFormatRealizedPnLWithFeesUsesMarketNativeCurrency(t *testing.T) {
	got := formatRealizedPnLWithFees(120, []FeeMarketTotal{
		{Market: "美股", Currency: "USD", Amount: 3.3},
		{Market: "港股", Currency: "HKD", Amount: 29.27},
	})
	want := "¥120.00\n手续费 美股: $3.30 / 港股: HK$29.27"
	if got != want {
		t.Fatalf("formatRealizedPnLWithFees = %q, want %q", got, want)
	}
}

func TestBuildRecordRowsAppliesSortHeader(t *testing.T) {
	rows := buildRecordRows(nil, recordSortState{Column: recordColumnTime, Direction: sortAscending})
	if rows[0][recordColumnTime] != "时间 ↑" {
		t.Fatalf("time header = %q, want ascending marker", rows[0][recordColumnTime])
	}
	if recordColumnRealizedPnL >= recordColumnFee {
		t.Fatalf("column order: realized pnl column=%d fee column=%d, want realized before fee", recordColumnRealizedPnL, recordColumnFee)
	}

	rows = buildRecordRows(nil, recordSortState{Column: recordColumnFee, Direction: sortDescending})
	if rows[0][recordColumnFee] != "手续费 ↓" {
		t.Fatalf("fee header = %q, want descending marker", rows[0][recordColumnFee])
	}

	rows = buildRecordRows(nil, recordSortState{Column: recordColumnRealizedPnL, Direction: sortDescending})
	if rows[0][recordColumnRealizedPnL] != "盈亏(CNY) ↓" {
		t.Fatalf("realized pnl header = %q, want descending marker", rows[0][recordColumnRealizedPnL])
	}

	rows = buildRecordRows(nil, recordSortState{Column: recordColumnFee, Direction: sortNone})
	if rows[0][recordColumnFee] != "手续费" {
		t.Fatalf("fee header = %q, want no marker", rows[0][recordColumnFee])
	}
}

func TestSortedRecordsByTimeAndFee(t *testing.T) {
	base := time.Date(2026, 6, 9, 10, 0, 0, 0, time.Local)
	ui := &UI{recordSort: recordSortState{Column: recordColumnTime, Direction: sortDescending}}
	records := []TradeSummary{
		{DisplayCode: "US:B", Time: base.Add(2 * time.Hour), Fee: 3, RealizedPnL: 20},
		{DisplayCode: "US:A", Time: base, Fee: 1, RealizedPnL: -5},
		{DisplayCode: "US:C", Time: base.Add(time.Hour), Fee: 2, RealizedPnL: 10},
	}

	got := ui.sortedRecords(records)
	if got[0].DisplayCode != "US:B" || got[2].DisplayCode != "US:A" {
		t.Fatalf("time desc records = %+v", got)
	}

	ui.recordSort = recordSortState{Column: recordColumnFee, Direction: sortDescending}
	got = ui.sortedRecords(records)
	if got[0].DisplayCode != "US:B" || got[2].DisplayCode != "US:A" {
		t.Fatalf("fee desc records = %+v", got)
	}

	ui.recordSort = recordSortState{Column: recordColumnRealizedPnL, Direction: sortAscending}
	got = ui.sortedRecords(records)
	if got[0].DisplayCode != "US:A" || got[2].DisplayCode != "US:B" {
		t.Fatalf("realized pnl asc records = %+v", got)
	}

	ui.recordSort = recordSortState{Column: recordColumnFee, Direction: sortNone}
	got = ui.sortedRecords(records)
	if got[0].DisplayCode != "US:B" || got[1].DisplayCode != "US:A" || got[2].DisplayCode != "US:C" {
		t.Fatalf("unsorted records = %+v", got)
	}
}

func TestBuildRecordRowsShowsRealizedPnLForSellOnly(t *testing.T) {
	rows := buildRecordRows([]TradeSummary{
		{TypeLabel: "买入", Time: time.Date(2026, 6, 12, 9, 0, 0, 0, time.Local), DisplayCode: "US:AAPL", Market: "美股", Currency: "USD", Quantity: 1, Price: 100, FXRate: 7.2, AmountBase: 720, Fee: 1.5},
		{TypeLabel: "平仓", Time: time.Date(2026, 6, 12, 10, 0, 0, 0, time.Local), DisplayCode: "US:AAPL", Market: "美股", Currency: "USD", Quantity: 1, Price: 120, FXRate: 7.2, AmountBase: 864, Fee: 1.8, RealizedPnL: 144},
	}, recordSortState{})

	if rows[1][recordColumnRealizedPnL] != "" {
		t.Fatalf("buy realized pnl = %q, want empty", rows[1][recordColumnRealizedPnL])
	}
	if rows[2][recordColumnRealizedPnL] != "¥144.00" {
		t.Fatalf("sell realized pnl = %q, want ¥144.00", rows[2][recordColumnRealizedPnL])
	}
}

func TestToggleRecordSortCyclesDescendingAscendingNone(t *testing.T) {
	state := recordSortState{Column: recordColumnTime, Direction: sortDescending}

	state = nextRecordSortState(state, recordColumnFee)
	if state.Column != recordColumnFee || state.Direction != sortDescending {
		t.Fatalf("first toggle = %+v, want fee desc", state)
	}

	state = nextRecordSortState(state, recordColumnFee)
	if state.Column != recordColumnFee || state.Direction != sortAscending {
		t.Fatalf("second toggle = %+v, want fee asc", state)
	}

	state = nextRecordSortState(state, recordColumnFee)
	if state.Column != recordColumnFee || state.Direction != sortNone {
		t.Fatalf("third toggle = %+v, want fee none", state)
	}
}
