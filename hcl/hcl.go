// hcl implements compatibility with github.com/hashicorp/go-hclog
package hcl

import (
	"bytes"
	"log"
	"regexp"
	"strings"

	"github.com/digitalrebar/logger"
	hclog "github.com/hashicorp/go-hclog"
)

type L struct {
	logger.Logger
	defaultLvl logger.Level
}

func (l *L) Named(s string) *L {
	return &L{l.Logger.Switch(s), l.defaultLvl}
}

func (l *L) ResetNamed(s string) *L {
	return l.Named(s)
}

type adp struct {
	*L
}

var lvlRE = regexp.MustCompile(`^\[([^/]]+)\]`)

func (a *adp) parse(s string) (logger.Level, string) {
	matches := lvlRE.FindStringSubmatch(s)
	if matches == nil || len(matches) != 2 {
		return a.defaultLvl, s
	}
	matchLvl, err := logger.ParseLevel(matches[1])
	if err != nil {
		return a.defaultLvl, s
	}
	return matchLvl, strings.TrimPrefix(s, matches[0])
}

func (a *adp) Write(buf []byte) (int, error) {
	str := string(bytes.TrimRight(buf, " \t\n"))
	lvl, msg := a.parse(str)
	a.Logf(lvl, "%s", msg)
	return len(buf), nil
}

func (l *L) StandardLogger(*hclog.StandardLoggerOptions) *log.Logger {
	a := &adp{l}
	return log.New(a, "", 0)
}

func HCL(log logger.Logger, defaultLvl logger.Level) *L {
	return &L{log, defaultLvl}
}
