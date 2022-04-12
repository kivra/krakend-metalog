package metalog

import (
	"bytes"
	"encoding/json"

	"github.com/luraproject/lura/v2/config"
)

type Config = map[string]interface{}

type Key struct {
	namespace string
}

const Namespace = "kivra/metalog"

var Metalog = Key{namespace: Namespace}

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
