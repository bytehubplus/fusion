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

package internal

import "github.com/gin-gonic/gin"

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService LoginService
	jwtService   JWTService
}

type Credentials struct {
	VaultID   string `json:"vaultID"`
	Nonce     string `json:"nonce"`
	Signature string `json:"signature"`
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credentials Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	//用户用自己的私钥对VaultID+Nonce签名，如果签名者是VaultID的controller则通过验证，为用户生成Token
	isAuthenticated := controller.loginService.Login(credentials.VaultID, credentials.Nonce, credentials.Signature)
	if isAuthenticated {
		return controller.jwtService.GenerateToken(credentials.VaultID, credentials.Nonce, credentials.Signature)
	}
	return ""
}

func NewLoginController(loginService LoginService, jwtService JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}
