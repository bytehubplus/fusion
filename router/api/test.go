package main

import (
	entry "github.com/bytehubplus/fusion/router/api/Entry"
	"github.com/bytehubplus/fusion/router/api/controller"
	"github.com/bytehubplus/fusion/router/api/vault"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//ssl tls ..
	r.POST("/register", vault.RegistVault)
	r.POST("/unregister", vault.Unregister)
	r.POST("/append", controller.Append)
	r.POST("/pubkey_delete", controller.Delete)
	r.POST("/list", controller.List)
	r.POST("/put", entry.Put)
	r.POST("/get", entry.Get)
	r.POST("/entry_delete", entry.Delete)
	r.Run(":8080")
}
