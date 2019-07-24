package log

import (
	"log"

	"github.com/comail/wincolog"
	"github.com/ddsgok/colog"
)

const (
	// DefaultContext it's used as default, prints only warning, errors
	// and alerts.
	DefaultContext = Context(iota)
	// VerboseContext prints more information, logged after some stages
	// in program are being executed.
	VerboseContext = Context(iota)
	// DebuggingContext prints information containing variables values,
	// for debugging purposes.
	DebuggingContext = Context(iota)
)

// Context it's an enum type to choose log contexts for project.
type Context byte

// severoLog it's a setup to use CoLog to configure options, and then
// use a logger created with this configuration.
type severoLog struct {
	*colog.CoLog
	logger     *log.Logger
	currentCtx Context
}

// ResetLogger creates a new logger using configuration chosen on
// CoLog object.
func (sl *severoLog) ResetLogger() {
	sl.logger = sl.CoLog.NewLogger()
}

// Print uses formatting and arguments to print text to logger output.
func (sl *severoLog) Print(formatter string, args []interface{}) {
	sl.logger.Printf(formatter, args...)
}

// Panic prints and then call panic.
func (sl *severoLog) Panic(formatter string, args []interface{}) {
	sl.logger.Panicf(formatter, args...)
}

// Fatal prints and then os.Exit(1).
func (sl *severoLog) Fatal(formatter string, args []interface{}) {
	sl.logger.Fatalf(formatter, args...)
}

// SetLogContext choose between different contexts and setups log to
// limit information depending on context.
func (sl *severoLog) SetLogContext(lc Context) {
	sl.CoLog.SetDefaultLevel(colog.LInfo)
	sl.currentCtx = lc

	switch lc {
	case DefaultContext:
		sl.CoLog.SetOmitHeaders(true)
		sl.CoLog.SetFlags(log.Ltime | log.Ldate)
		sl.CoLog.SetMinLevel(colog.LMessage)
	case VerboseContext:
		sl.CoLog.SetOmitHeaders(false)
		sl.CoLog.SetFlags(log.Ltime | log.Ldate)
		sl.CoLog.SetMinLevel(colog.LInfo)
	case DebuggingContext:
		sl.CoLog.SetOmitHeaders(false)
		sl.CoLog.SetFlags(log.Ltime | log.Llongfile)
		sl.CoLog.SetMinLevel(colog.LDebug)
	}

	sl.ResetLogger()
}

// Logger return the modified logger used in severoLog object.
func (sl *severoLog) Logger() (l *log.Logger) {
	l = sl.logger
	return
}

// newSeveroLog creates a new logger with default configuration for the
// logger on this project. It will use my version of comail/colog, with
// colored output forced, fields colored but non-adjusted, and call
// depth adjusted by two calls: the log.Print() and logger.Print().
func newSeveroLog() (sl *severoLog) {
	sl = &severoLog{
		CoLog: colog.New(wincolog.Stdout(), "", log.LstdFlags),
	}

	sl.CoLog.SetForceColorOutput(true)
	sl.CoLog.SetAdjustFieldsToRight(false)
	sl.CoLog.SetParseFields(true)
	sl.CoLog.AdjustCallDepth(2)

	sl.SetLogContext(DefaultContext)
	sl.ResetLogger()

	return
}
