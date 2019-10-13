package marker

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

type LogMarkerOption func(*LogMarker)

type LogMarkerRule struct {
	Matcher MatcherFunc
	Color   *color.Color
}

type LogMarker struct {
	rules  []LogMarkerRule
	logger *log.Logger
}

func NewLogMarker(options ...LogMarkerOption) *LogMarker {
	logMarker := &LogMarker{logger: log.New(os.Stdout, "", 0)}
	for _, option := range options {
		option(logMarker)
	}
	return logMarker
}

func SetLogger(logger *log.Logger) LogMarkerOption {
	return func(logMarker *LogMarker) {
		logMarker.logger = logger
	}
}

func (l *LogMarker) AddRule(rule LogMarkerRule) *LogMarker {
	l.rules = append(l.rules, rule)
	return l
}

func (l *LogMarker) Print(v ...interface{}) {
	marked := fmt.Sprint(v...)
	for _, rule := range l.rules {
		marked = Mark(marked, rule.Matcher, rule.Color)
	}
	l.logger.Print(marked)
}
