package mailbox

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// middleware，集中进行参数校验以及鉴权

func ValidatorAddMessage(c *gin.Context) {
	if checkAddMessageRequest(c, "content") && checkAddMessageRequest(c, "source") {
		c.Next()
	} else {
		return
	}
}

func checkAddMessageRequest(c *gin.Context, arg string) bool {
	value, ok := c.GetQuery(arg)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("arg %s do not exist", arg)})
		return false
	}
	if len(strings.TrimSpace(value)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("the arg of %s is invalid", arg)})
		return false
	}
	return true
}
