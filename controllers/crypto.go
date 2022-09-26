package controllers

import (
	"errors"
	"net/http"
	"strconv"

	datastore "github.com/caiomp87/crypto-votes-datastore/postgres"
	"github.com/caiomp87/crypto-votes-entities/models"
	"github.com/gin-gonic/gin"
)

func ListCryptos(ctx *gin.Context) {
	cryptos, err := datastore.CryptoDatastore.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, cryptos)
}

func GetCrypto(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("id field is required"),
		})
		return
	}

	convertedID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	crypto, err := datastore.CryptoDatastore.FindByID(convertedID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, crypto)
}

func CreateCrypto(ctx *gin.Context) {
	var crypto *models.Crypto
	if err := ctx.BindJSON(&crypto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := datastore.CryptoDatastore.Create(crypto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "crypto created successfully",
	})
}

func UpdateCrypto(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("id field is required"),
		})
		return
	}

	convertedID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var crypto *models.Crypto
	if err := ctx.BindJSON(&crypto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := datastore.CryptoDatastore.UpdateByID(convertedID, crypto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "crypto updated successfully",
	})
}

func DeleteCrypto(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("id field is required"),
		})
		return
	}

	convertedID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := datastore.CryptoDatastore.DeleteByID(convertedID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "crypto deleted successfully",
	})
}
