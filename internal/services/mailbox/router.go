package mailbox

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yang20150702/base-microservice-web-project-layout/framework"
	"github.com/yang20150702/base-microservice-web-project-layout/internal/services/mailbox/storage"
)

var log = framework.GetLogger()

func init() {
	g := framework.GlobalEngine.Group("/mailbox")

	g.GET("/add", ValidatorAddMessage, addMessage)

}

func addMessage(c *gin.Context) {
	content, _ := c.GetQuery("content")

	m, err := storage.GetBackendImpl().CreateMessage(content)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.String(http.StatusOK, "the content of message is %s", m.Content)
}
