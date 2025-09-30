package logger

import "myminegame/pkg"

type logger struct {
	enable bool
}

func NewLogger() *logger {
	return &logger{
		enable: pkg.ReadSettings().Logs,
	}
}
