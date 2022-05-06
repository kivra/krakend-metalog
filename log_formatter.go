package metalog

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func LogFormatter(param gin.LogFormatterParams) string { // nolint
	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency -= param.Latency % time.Second
	}

	var logMessage = make(map[string]interface{})

	ctx := param.Request.Context()
	metaData, ok := ctx.Value(Metalog).(Config)
	if ok {
		logMessage = metaData
	}

	logMessage["client_ip"] = param.ClientIP
	logMessage["duration_us"] = param.Latency.Microseconds()
	logMessage["headers"] = format(param.Request.Header)
	logMessage["level"] = "INFO"
	logMessage["method"] = param.Method
	logMessage["module"] = "GIN"
	logMessage["path"] = param.Request.URL.Path
	logMessage["query"] = param.Request.URL.Query()
	logMessage["response_body_size"] = param.BodySize
	logMessage["status"] = param.StatusCode
	logMessage["timestamp"] = param.TimeStamp.Format("2006-01-02T15:04:05.000+00:00")

	if param.ErrorMessage != "" {
		logMessage["message"] = strings.TrimSuffix(param.ErrorMessage, "\n")
	}

	jsonLog, err := json.Marshal(logMessage)

	if err != nil {
		return "{\"level\": \"ERROR\", \"module\": \"metalog\", \"message\": \"failed to convert log message to JSON\"}\n"
	}

	return string(jsonLog) + "\n"
}

func format(m map[string][]string) map[string]string {
	formatted := make(map[string]string)
	for key, values := range m {
		formatted[key] = strings.Join(values, ", ")
	}
	return formatted
}
