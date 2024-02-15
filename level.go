// Copyright 2024 Leandro "inkel" López. All rights reserved.

/*
Package slogflag implements command-line flag parsing for slog.Level
variables.

# Usage

Define flags using [slogflag.Level]. This declares a slog.Level flag, -log-level, stored in the pointer level, with type *slog.Level:

	var level = slogflag.Level("log-level", slog.LevelInfo, "help message for flag -log-level")

If you like, you can bind the flag to a variable using the [slogflag.LevelVar] function:

	var levelVar slog.Level
	func init() {
		slogflag.LevelVar(&levelVar, "log-level", slog.LevelInfo, "help message for flag -log-level")
	}

After all the flags are defined, call [flag.Parse] as usual. Then you can set the log level of your [slog.Logger] like this:

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: *level}))
*/
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

// ErrParse is the error returned when it's impossible to parse the
// log level to something slog accepts.
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

// LevelVar defines a [slog.Level] flag with specified name, default
// value, and usage string. The argument p points to a [slog.Level]
// variable in which to store the value of the flag.
func LevelVar(p *slog.Level, name string, value slog.Level, usage string) {
	*p = value
	lv := levelVar(*p)
	flag.CommandLine.Var(&lv, name, usage)
}

// Level defines a [slog.Level] flag with specified name, default
// value, and usage string. The return value is the address of a
// [slog.Level] variable that stores the value of the flag.
func Level(name string, value slog.Level, usage string) *slog.Level {
	var l slog.Level
	LevelVar(&l, name, value, usage)
	return &l
}
