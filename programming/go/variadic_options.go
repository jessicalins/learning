package main

import "fmt"

const (
	Low = verbosity(iota)
	Medium
	High
)

type verbosity int
type Logger struct {
	verbosity
	prefix string
}

// Declare a new option setter func type.
type option func(*Logger)

// SetOptions changes the options of Logger.
func (lo *Logger) SetOptions(opts ...option) {
	for _, applyOpt := range opts {
		applyOpt(lo)
	}
}

func LowVerbosity() option {
	return func(lo *Logger) {
		lo.verbosity = Low
	}
}

func HighVerbosity() option {
	return func(lo *Logger) {
		lo.verbosity = High
	}
}

func Prefix(s string) option {
	return func(lo *Logger) {
		lo.prefix = s
	}
}

func (lo *Logger) Info(s string) {
	if lo.verbosity > Medium {
		lo.print("INFO", s)
	}
}

func (lo *Logger) Critical(s string) {
	lo.print("CRITICAL", s)
}

func (lo *Logger) print(level string, msg string) {
	pre := ""
	if lo.prefix != "" {
		pre = "[" + lo.prefix + "]"
	}
	fmt.Printf("%s%-10s: %s\n", pre, level, msg)
}

func main() {
	// Create a new Logger. Default verbosity is Low.
	// Because, the verbosity field has a zero-value of 0, then,
	// it will set the verbosity to Low automatically.
	logger := &Logger{}
	logger.SetOptions(HighVerbosity(), Prefix("ZOMBIE CONTROL"))
	logger.Critical("zombie outbreak!")
	logger.Info("1 sec passed")
}
