package pdf

import (
	"fmt"
	"html"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Greet(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		name = "Gopher"
	}
	c.HTML(http.StatusOK, "pdf.html", gin.H{
		"message": fmt.Sprintf("hello, %s!", html.EscapeString(name)),
	})
}
