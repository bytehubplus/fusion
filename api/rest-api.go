/* The AGPLv3 License (AGPLv3)

Copyright (c) 2022 Zhao Zhenhua <zhao.zhenhua@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>. */

package api

import (
	"net/http"

	"github.com/bytehubplus/fusion/api/internal"
	"github.com/gin-gonic/gin"
)

var (
	loginService LoginService = internal.NewLoginService()
)

func StartService() {
	server := gin.New()
	server.Use(gin.Recovery(),
		gin.Logger(),
	)

	server.POST("/login", func(ctx *gin.Context) {
		token := internal.NewLoginService().Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	server.RunTLS(":8000")
}
