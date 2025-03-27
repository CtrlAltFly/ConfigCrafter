package ConfigCrafter

import (
	"ConfigCrafter/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/convert", handlers.ConvertHandler)

	r.Run(":8080") // Start server on port 8080
}
