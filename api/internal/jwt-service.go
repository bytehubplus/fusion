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

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	// Token is based on vault ID, nonce and signature, signature must by one of its controller
	GenerateToken(vaultId string, nonce string, signature string) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	VaultID   string
	Nonce     string
	Signature string
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func (service *jwtService) GenerateToken(vaultId string, nonce string, signature string) string {
	claims := &jwtCustomClaims{
		VaultID:   vaultId,
		Nonce:     nonce,
		Signature: signature,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: "richzhao",
		issuer:    "bytehubplus.com",
	}
}
