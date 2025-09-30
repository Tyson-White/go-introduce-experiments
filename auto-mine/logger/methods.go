package logger

import "fmt"

func (l *logger) Line(str ...any) {
	if !l.enable {
		return
	}

	fmt.Println(str...)
}
