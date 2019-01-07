package log_output

func Debug(logRecord *LogRecord) {
	Logger.Log(DEBUG, logRecord, LogCallFlag(3))
}

func Info(logRecord *LogRecord) {
	Logger.Log(INFO, logRecord, LogCallFlag(3))
}

func Warning(logRecord *LogRecord) {
	Logger.Log(WARNING, logRecord, LogCallFlag(3))
}

func Error(logRecord *LogRecord) {
	Logger.Log(ERROR, logRecord, LogCallFlag(3))
}

func Critical(logRecord *LogRecord) {
	Logger.Log(CRITICAL, logRecord, LogCallFlag(3))
}

func Fixed(logRecord *LogRecord) {
	Logger.Log(FIXED, logRecord, LogCallFlag(3))
}
