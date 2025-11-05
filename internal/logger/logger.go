package logger

import "os"

type scnorionLogger struct {
	LogFile *os.File
}

func (l *scnorionLogger) Close() {
	l.LogFile.Close()
}
