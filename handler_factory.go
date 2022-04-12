package metalog

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/luraproject/lura/v2/config"
	"github.com/luraproject/lura/v2/proxy"
	krakendgin "github.com/luraproject/lura/v2/router/gin"
)

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
			var meta = make(map[string]interface{})
			for key, value := range *metaData {
				meta[key] = value
			}
			ctx := c.Request.Context()
			c.Request = c.Request.WithContext(context.WithValue(ctx, Metalog, meta))
			handlerFunc(c)
		}
	}
}
