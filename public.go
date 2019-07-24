package log

import "log"

var (
	// loggers stores the logger configuration used in project.
	logger = newSeveroLog()
)

// Print uses formatting and arguments to print text to logger output.
func Print(formatter string, args ...interface{}) {
	logger.Print(formatter, args)
}

// Panic prints and then call panic.
//noinspection GoUnusedExportedFunction
func Panic(formatter string, args ...interface{}) {
	logger.Panic(formatter, args)
}

// Fatal prints and then os.Exit(1).
func Fatal(formatter string, args ...interface{}) {
	logger.Fatal(formatter, args)
}

// SetLogContext choose between different contexts and setups log to
// limit information depending on context.
func SetLogContext(lc Context) {
	logger.SetLogContext(lc)
}

// Logger return the modified logger used package.
func Logger() (l *log.Logger) {
	l = logger.Logger()
	return
}
