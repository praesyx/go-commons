package commons

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	loggerBuilder *LoggerBuilder
	once          sync.Once
)

type LoggerBuilder struct {
	serviceCode uint8
}

func NewLoggerBuilder(serviceCode uint8) *LoggerBuilder {
	once.Do(func() {
		loggerBuilder = &LoggerBuilder{serviceCode: serviceCode}
	})
	return loggerBuilder
}

func GetLoggerBuilder() *LoggerBuilder {
	return loggerBuilder
}

func (l *LoggerBuilder) Code(errorCode uint32) int {
	n, _ := strconv.Atoi(fmt.Sprintf("%s%s", strconv.Itoa(int(l.serviceCode)), strconv.Itoa(int(errorCode))))
	return n
}
