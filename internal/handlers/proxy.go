package handlers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Proxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := http.NewRequest(c.Request.Method, target+c.Request.URL.Path, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		req.Header = c.Request.Header

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		for k, v := range resp.Header {
			for _, vv := range v {
				c.Writer.Header().Add(k, vv)
			}
		}
		c.Status(resp.StatusCode)
		io.Copy(c.Writer, resp.Body)
	}
}
