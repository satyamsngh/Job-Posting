package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type key string

const TraceIdKey key = "1"

func (m *Mid) Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := uuid.NewString()
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, TraceIdKey, traceId)
		req := c.Request.WithContext(ctx)
		c.Request = req

		log.Info().Str("Trace Id", traceId).Str("Method", c.Request.Method).
			Str("URL Path", c.Request.URL.Path).Msg("request started")

		defer log.Info().Str("Trace Id", traceId).Str("Method", c.Request.Method).
			Str("URL Path", c.Request.URL.Path).
			Int("status Code", c.Writer.Status()).Msg("Request processing completed")

		c.Next()
	}
}
