package formatter

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
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

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	format := f.LogFormat
	if format == "" {
		format = defaultLogFormat
	}
	timeFormat := f.TimestampFormat
	if timeFormat == "" {
		timeFormat = defaultTimestampFormat
	}

	var b bytes.Buffer
	formatted := format

	// 使用 strings.NewReplacer 预构建替换器提高效率
	replacer := []string{
		"%levelname%", strings.ToUpper(entry.Level.String()),
		"%time%", entry.Time.Format(timeFormat),
		"%message%", entry.Message,
	}

	if entry.HasCaller() {
		replacer = append(replacer,
			"%filename%", filepath.Base(entry.Caller.File),
			"%lineno%", strconv.Itoa(entry.Caller.Line),
		)
	} else {
		// 移除 caller 信息块
		formatted = removeCallerInfo(formatted)
	}

	r := strings.NewReplacer(replacer...)
	formatted = r.Replace(formatted)

	// 扩展字段
	if strings.Contains(formatted, "%extends%") {
		var extends strings.Builder
		for k, v := range entry.Data {
			extends.WriteString(fmt.Sprintf("%s=%v ", k, v))
		}
		formatted = strings.ReplaceAll(formatted, "%extends%", extends.String())
	}

	b.WriteString(formatted)
	return b.Bytes(), nil
}

func removeCallerInfo(s string) string {
	// 简单移除形如 [%filename%:%lineno%] 的字段
	start := strings.Index(s, "[%filename%:%lineno%]")
	if start != -1 {
		return strings.Replace(s, "[%filename%:%lineno%]", "", 1)
	}
	return s
}
