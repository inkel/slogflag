package slogflag

import (
	"errors"
	"log/slog"
	"testing"
)

func TestLevelVarSet(t *testing.T) {
	tests := []struct {
		in  string
		l   slog.Level
		err error
	}{
		{"DEBUG", slog.LevelDebug, nil},
		{"info", slog.LevelInfo, nil},
		{"Warn", slog.LevelWarn, nil},
		{"eRRor", slog.LevelError, nil},
		{"foo", 0, ErrParse},
		{"info+2", slog.LevelInfo + 2, nil},
		{"error-1", slog.LevelError - 1, nil},
		{"info+warn", 0, ErrParse},
		{"0", slog.LevelInfo, nil},
		{"-4", slog.LevelDebug, nil},
		{"4", slog.LevelWarn, nil},
		{"8", slog.LevelError, nil},
		{"-10", slog.Level(-10), nil},
		{"10", slog.Level(10), nil},
		{"+15", slog.Level(15), nil},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			var lv levelVar
			err := lv.Set(tt.in)
			if !errors.Is(err, tt.err) {
				t.Fatalf("expecting error %v, got %v", tt.err, err)
			}
			if got := lv.Level(); got != tt.l {
				t.Errorf("expecting level %v, got %v", tt.l, got)
			} else {
				t.Logf("input: %s - level: %v", tt.in, got)
			}
		})
	}
}

func BenchmarkLevelVarSet(b *testing.B) {
	ins := []string{"DEBUG", "INFO", "WARN", "ERROR", "INFO+2", "ERROR-1", "FOO", "WARN+ERROR", "-4", "0", "4", "8", "-10", "+10", "10"}

	for _, in := range ins {
		b.Run(in, func(b *testing.B) {
			var lv levelVar

			for i := 0; i < b.N; i++ {
				_ = lv.Set(in)
			}
		})
	}
}
