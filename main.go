package main

import (
	_ "github.com/RaymondCode/simple-demo/dal"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)

	r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
