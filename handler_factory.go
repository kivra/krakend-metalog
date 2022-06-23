package metalog

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luraproject/lura/v2/config"
	"github.com/luraproject/lura/v2/proxy"
	krakendgin "github.com/luraproject/lura/v2/router/gin"
)

func Add(moreData map[string]interface{}, c *gin.Context) {
	metaData := Get(c.Request)
	for key, value := range moreData {
		metaData[key] = value
	}
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), metalog, metaData))
}

func Get(req *http.Request) map[string]interface{} {
	metaData, ok := req.Context().Value(metalog).(Config)
	if ok {
		return metaData
	}
	return make(map[string]interface{})
}

func HandlerFactory(next krakendgin.HandlerFactory) krakendgin.HandlerFactory {
	// runs when endpoint is registered
	return func(remote *config.EndpointConfig, p proxy.Proxy) gin.HandlerFunc {
		handlerFunc := next(remote, p)

		metaData, ok := ConfigGetter(remote.ExtraConfig)
		if !ok {
			return handlerFunc
		}

		// runs when request is executed
		return func(c *gin.Context) {
			Add(*metaData, c)
			handlerFunc(c)
		}
	}
}
