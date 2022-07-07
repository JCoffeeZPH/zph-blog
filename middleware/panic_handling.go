package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	api_error "zph/error"
	"zph/lib/common"
	"zph/logger"
)

var log = logger.Log

func GlobalPanicHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			c.Abort()

			if err := recover(); err != nil {
				switch err.(type) {
				case *api_error.ApiError:
					e := err.(*api_error.ApiError)
					if e.HttpStatusCode > 400 && e.HttpStatusCode < 500 {
						log.Warnf("User is unauthorized, err: %+v, Authorization: %s,  url: %s", e, c.GetHeader("Authorization"), c.Request.URL)
					} else {
						printStack(c, err)
					}
					c.JSON(e.HttpStatusCode, gin.H{"message": e.ErrorCode + e.ErrorMsg, "error_code": e.ErrorCode + e.ErrorMsg, "data": e.Data})
				case common.DBError:
					printStack(c, err)
					c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("DB error")})
				case common.ServiceError:
					printStack(c, err)
					c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Service error")})
				default:
					printStack(c, err)
					log.Errorf("default: %+v", err)
					c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
				}
			}
		}()

		c.Next()
	}
}

func printStack(c *gin.Context, err interface{}) {
	buf := make([]byte, 8192)
	n := runtime.Stack(buf, false)
	stackInfo := fmt.Sprintf("%s", buf[:n])
	log.Errorf("msg: %+v, url: %s panic stack info %s", err, c.Request.URL, stackInfo)
}
