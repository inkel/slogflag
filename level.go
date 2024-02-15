package slogflag

import (
	"errors"
	"flag"
	"log/slog"
	"strconv"
	"strings"
)

type levelVar slog.Level

var _ flag.Value = new(levelVar)
var _ slog.Leveler = new(levelVar)

var ErrParse = errors.New("cannot parse log level")

func parse(val string) (slog.Level, error) {
	var l slog.Level
	switch val {
	case "DEBUG":
		l = slog.LevelDebug
	case "INFO":
		l = slog.LevelInfo
	case "WARN":
		l = slog.LevelWarn
	case "ERROR":
		l = slog.LevelError
	default:
		return 0, ErrParse
	}

	return l, nil
}

func (lv *levelVar) Set(val string) error {
	var l slog.Level

	val = strings.ToUpper(val)
	idx := len(val)

	for i, r := range val {
		if r == '+' || r == '-' {
			idx = i
			break
		}
	}

	l, err := parse(val[:idx])
	if err != nil {
		return err
	}

	d, _ := strconv.Atoi(val[idx:])
	l += slog.Level(d)

	*lv = levelVar(l)

	return nil
}

func (lv *levelVar) String() string {
	l := slog.Level(*lv)
	return l.String()
}

func (lv *levelVar) Level() slog.Level { return slog.Level(*lv) }

func LevelVar(p *slog.Level, name string, value slog.Level, usage string) {
	*p = value
	lv := levelVar(*p)
	flag.CommandLine.Var(&lv, name, usage)
}

func Level(name string, value slog.Level, usage string) *slog.Level {
	var l slog.Level
	LevelVar(&l, name, value, usage)
	return &l
}
