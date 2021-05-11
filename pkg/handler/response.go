package handler

import (
	"log"
	"net/http"

	"github.com/Askaell/homework/pkg/render"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Print("error ", message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

func writeResponse(c *gin.Context, code int, data interface{}) {
	output, err := renderByFormatType(c, data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(code, "%s", output)
}

func renderByFormatType(c *gin.Context, input interface{}) (interface{}, error) {
	formatType := getFormatFromQuery(c)
	renderer := render.GetRenderer(formatType)

	output, err := renderer.Render(input)
	return output, err
}

func getFormatFromQuery(c *gin.Context) string {
	format := c.Query("format")
	return format
}
