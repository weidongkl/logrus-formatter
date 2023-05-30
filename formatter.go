package formatter

import (
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

const (
	defaultLogFormat       = "[%levelname%] %time%  [%filename%:%lineno%]: %extends% %message%\n"
	defaultTimestampFormat = time.RFC3339Nano
)

type Formatter struct {
	TimestampFormat string
	LogFormat       string
}

func (m Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := m.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := m.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	output = strings.Replace(output, "%message%", entry.Message, 1)

	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%levelname%", level, 1)
	if entry.HasCaller() {
		output = strings.Replace(output, "%lineno%", strconv.Itoa(entry.Caller.Line), 1)
		output = strings.Replace(output, "%filename%", entry.Caller.File, 1)
	}
	for k, val := range entry.Data {
		switch v := val.(type) {
		case string:
			output = strings.Replace(output, "%extends%", k+"="+v+" "+"%extends%", 1)
		case int:
			s := strconv.Itoa(v)
			output = strings.Replace(output, "%extends%", k+"="+s+" "+"%extends%", 1)
		case bool:
			s := strconv.FormatBool(v)
			output = strings.Replace(output, "%extends%", k+"="+s+" "+"%extends%", 1)
		case time.Time:
			s := v.String()
			output = strings.Replace(output, "%extends%", k+"="+s+" "+"%extends%", 1)
		}
	}
	output = strings.Replace(output, "%extends%", "", 1)
	return []byte(output), nil
}
