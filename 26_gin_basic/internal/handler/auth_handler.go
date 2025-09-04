package handler

import (
	"gin-basic/internal/service"
	"gin-basic/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: &service}
}

func (h *AuthHandler) LoaginHandler(ctx *gin.Context) {

	type HelloWorldHeader struct {
		Auth string `header:"Authorization"`
	}

	type HelloWorldPath struct {
		P1 string `uri:"userId" binding:"omitempty,len=4"`
	}

	type HelloWorldQuery struct {
		Q1 string `form:"q1" binding:"omitempty,max=4"`
		Q2 int    `form:"q2" binding:"omitempty,len=2"`
	}

	type HelloWorldRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Passowrd string `json:"password" binding:"required,min=8,max=12"`
		//
		Num1  int    `json:"num1" binding:"omitempty,gte=18,lte=21"`
		Str1  string `json:"str1" binding:"omitempty,len=4"`
		Bool1 bool   `json:"bool1" binding:"omitempty"`
	}

	var header HelloWorldHeader
	var path HelloWorldPath
	var query HelloWorldQuery
	var request HelloWorldRequest

	err := ctx.ShouldBindHeader(&header)
	if err != nil {
		ctx.AbortWithStatusJSON(response.Error(http.StatusBadRequest, err))
		return
	}

	err = ctx.ShouldBindUri(&path)
	if err != nil {
		ctx.AbortWithStatusJSON(response.Error(http.StatusBadRequest, err))
		return
	}

	err = ctx.ShouldBindQuery(&query)
	if err != nil {
		ctx.AbortWithStatusJSON(response.Error(http.StatusBadRequest, err))
		return
	}

	err = ctx.ShouldBindBodyWithJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(response.Error(http.StatusBadRequest, err))
		return
	}

	data, err := h.service.LoginService(request.Email, request.Passowrd)

	if err != nil {
		ctx.AbortWithStatusJSON(response.Error(http.StatusUnauthorized, err))
		return
	}
	ctx.JSON(response.Success(http.StatusOK, map[string]any{
		"header":  header,
		"path":    path,
		"query":   query,
		"request": request,
		"result":  data,
	}))
}
