package metalog

import (
	"bytes"
	"encoding/json"
	"strings"
	"time"

	"github.com/luraproject/lura/v2/config"
)

type Config = map[string]interface{}

type Key struct {
	namespace string
}

const Namespace = "kivra/metalog"

var metalog = Key{namespace: Namespace}

// TimeFormat used to format time in logs
var TimeFormat = time.RFC3339

// TimeKey that holds the log time in logs
var TimeKey = "timestamp"

// LevelFormatter to format log levels
var LevelFormatter = strings.ToUpper

func ConfigGetter(e config.ExtraConfig) (*Config, bool) { // nolint
	cfg := new(Config)

	tmp, ok := e[Namespace]
	if !ok {
		return cfg, false
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(tmp); err != nil {
		return cfg, false
	}
	if err := json.NewDecoder(buf).Decode(cfg); err != nil {
		return cfg, false
	}

	return cfg, true
}
