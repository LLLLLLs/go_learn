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
