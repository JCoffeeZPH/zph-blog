package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zph/constants"
)

type BaseController struct {
}

func (b *BaseController) NoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

func (b *BaseController) Error(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{"error": err})
}

func (b *BaseController) Success(c *gin.Context, body interface{}) {
	c.JSON(http.StatusOK, body)
}

func (b *BaseController) InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Internal server error %s", message)})
}

func (b *BaseController) SuccessWithPageInfo(c *gin.Context, perPage, total uint64, body interface{}) {
	c.Header("Per-Page", strconv.Itoa(int(perPage)))
	c.Header("Total", strconv.Itoa(int(total)))
	c.JSON(http.StatusOK, body)
}

func (b *BaseController) GetUserId(c *gin.Context) uint64 {
	 userId, _ := c.Get(constants.VisitorBlogOwnerId)
	 return userId.(uint64)
}