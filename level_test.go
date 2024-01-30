package slogflag_test

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/inkel/slogflag"
)

func TestSet(t *testing.T) {
	tests := []struct {
		in  string
		l   slog.Level
		err error
	}{
		{"DEBUG", slog.LevelDebug, nil},
		{"info", slog.LevelInfo, nil},
		{"Warn", slog.LevelWarn, nil},
		{"eRRor", slog.LevelError, nil},
		{"foo", 0, slogflag.ErrParse},
		{"info+2", slog.LevelInfo + 2, nil},
		{"error-1", slog.LevelError - 1, nil},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			var lv slogflag.LevelValue
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

func BenchmarkSet(b *testing.B) {
	ins := []string{"DEBUG", "INFO", "WARN", "ERROR", "INFO+2", "ERROR-1", "FOO"}

	for _, in := range ins {
		b.Run(in, func(b *testing.B) {
			var lv slogflag.LevelValue

			for i := 0; i < b.N; i++ {
				lv.Set(in)
			}
		})
	}
}
