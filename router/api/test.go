package main

import (
	"github.com/bytehubplus/fusion/router/api/vault"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//ssl tls ..
	r.POST("/register", vault.RegistVault)
	r.POST("/unregister", vault.Unregister)
	r.Run(":8080")
}
