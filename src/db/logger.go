package db

import (
	"github.com/go-xorm/core"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Entry
	level core.LogLevel
}

func (l *Logger) SetLevel(level core.LogLevel) {
	l.level = level
}

func (l *Logger) Level() core.LogLevel {
	return l.level
}

func (l *Logger) ShowSQL(show ...bool) {

}

func (l *Logger) IsShowSQL() bool {
	return false
}
