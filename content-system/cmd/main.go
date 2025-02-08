package main

import (
	"content-system/internal/api"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.NewRouters(r)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := r.Run()
	if err != nil {
		fmt.Printf("r run error = %v", err)
		return
	}

}
