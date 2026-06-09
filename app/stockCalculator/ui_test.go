package main

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

	rows = buildRecordRows(nil, recordSortState{Column: recordColumnFee, Direction: sortDescending})
	if rows[0][recordColumnFee] != "手续费 ↓" {
		t.Fatalf("fee header = %q, want descending marker", rows[0][recordColumnFee])
	}

	rows = buildRecordRows(nil, recordSortState{Column: recordColumnFee, Direction: sortNone})
	if rows[0][recordColumnFee] != "手续费" {
		t.Fatalf("fee header = %q, want no marker", rows[0][recordColumnFee])
	}
}

func TestSortedRecordsByTimeAndFee(t *testing.T) {
	base := time.Date(2026, 6, 9, 10, 0, 0, 0, time.Local)
	ui := &UI{recordSort: recordSortState{Column: recordColumnTime, Direction: sortAscending}}
	records := []TradeSummary{
		{DisplayCode: "US:B", Time: base.Add(2 * time.Hour), Fee: 3},
		{DisplayCode: "US:A", Time: base, Fee: 1},
		{DisplayCode: "US:C", Time: base.Add(time.Hour), Fee: 2},
	}

	got := ui.sortedRecords(records)
	if got[0].DisplayCode != "US:A" || got[2].DisplayCode != "US:B" {
		t.Fatalf("time asc records = %+v", got)
	}

	ui.recordSort = recordSortState{Column: recordColumnFee, Direction: sortDescending}
	got = ui.sortedRecords(records)
	if got[0].DisplayCode != "US:B" || got[2].DisplayCode != "US:A" {
		t.Fatalf("fee desc records = %+v", got)
	}

	ui.recordSort = recordSortState{Column: recordColumnFee, Direction: sortNone}
	got = ui.sortedRecords(records)
	if got[0].DisplayCode != "US:B" || got[1].DisplayCode != "US:A" || got[2].DisplayCode != "US:C" {
		t.Fatalf("unsorted records = %+v", got)
	}
}
