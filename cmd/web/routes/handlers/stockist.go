package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Stockist struct {
	log *logrus.Entry
	// uc  DocumentUC
}

func NewStockist(l *logrus.Entry) *Stockist {
	return &Stockist{
		log: l,
	}
}

func (d *Stockist) ListAll() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		d.log.Debug("Get all stockists")
		ctx.JSON(http.StatusOK, gin.H{"status": "Get all stockists"})
	}
}
