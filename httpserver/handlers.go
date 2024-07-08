package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outagelab/outagelab/models"
	"github.com/outagelab/outagelab/services/authservice"
	"github.com/outagelab/outagelab/services/datapageservice"
)

func (hs *HttpServer) handleLogin(ctx *gin.Context) {
	var loginRequest authservice.LoginRequest
	if err := ctx.BindJSON(&loginRequest); err != nil {
		ctx.AbortWithError(400, err)
	}

	loginResponse, err := hs.authService.Login(ctx, &loginRequest)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	if loginResponse == nil {
		ctx.AbortWithStatus(401)
		return
	}

	ctx.JSON(http.StatusOK, loginResponse)
}

func (hs *HttpServer) handleGetAccount(ctx *gin.Context) {
	claims := getAuthClaims(ctx)
	if claims == nil {
		ctx.AbortWithStatus(401)
		return
	}

	// need account ID

	account, err := hs.accountService.GetAccount(ctx, claims.AccountId)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (hs *HttpServer) handleUpdateAccount(ctx *gin.Context) {
	claims := getAuthClaims(ctx)
	if claims == nil {
		ctx.AbortWithStatus(401)
		return
	}

	account := &models.Account{}
	if err := ctx.BindJSON(account); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	if claims.AccountId != account.ID {
		ctx.AbortWithStatus(403)
		return
	}

	account, err := hs.accountService.UpdateAccount(ctx, claims.AccountId, account)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (hs *HttpServer) handleGenerateDatapageDeprecated(ctx *gin.Context) {
	apiKey := ctx.GetHeader("x-api-key")
	if apiKey == "" {
		ctx.AbortWithStatus(401)
		return
	}

	account, err := hs.datapageService.GenerateDatapageDeprecated(ctx, apiKey, &datapageservice.GenerateDatapageRequest{})
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	if account == nil {
		ctx.AbortWithError(401, err)
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (hs *HttpServer) handleGenerateDatapage(ctx *gin.Context) {
	apiKey := ctx.GetHeader("x-api-key")
	if apiKey == "" {
		ctx.AbortWithStatus(401)
		return
	}

	request := &datapageservice.GenerateDatapageRequest{}
	if err := ctx.BindJSON(request); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	account, err := hs.datapageService.GenerateDatapage(ctx, apiKey, request)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	if account == nil {
		ctx.AbortWithError(401, err)
		return
	}

	ctx.JSON(http.StatusOK, account)
}
