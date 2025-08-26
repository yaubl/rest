/*
since gin's context is stupid and cannot be extended
without hardcoding shit, i made this middleware to
"extend" gin context.
*/

package middlewares

import (
	"api/db"

	"context"
	"github.com/gin-gonic/gin"
)

type AppContext struct {
	DB      *db.Queries
	Context context.Context
}

const appContextKey = "appContext"

func WithAppContext(ctx *AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(appContextKey, ctx)
		c.Next()
	}
}

func GetAppContext(c *gin.Context) *AppContext {
	if ctx, exists := c.Get(appContextKey); exists {
		if appCtx, ok := ctx.(*AppContext); ok {
			return appCtx
		}
	}
	return nil
}
