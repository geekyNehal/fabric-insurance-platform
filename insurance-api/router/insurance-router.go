package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {

	engine.POST("api/addUser", func(c *gin.Context) {

		c.JSON(http.StatusOK, "Received Thanks!")
	})

	err := engine.Run(":3003")
	if err != nil {
		fmt.Println(err)
	}

}
