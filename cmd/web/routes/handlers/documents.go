package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// type DocumentUC interface {
// 	ListAll() []interface{}
// 	ListOne() interface{}
// 	ChangeSituation() interface{}
// }

type Document struct {
	log *logrus.Entry
	// uc  DocumentUC
}

func NewDocument(l *logrus.Entry) *Document {
	return &Document{
		log: l,
	}
}

// func (d *Documents) ListAll() http.HandlerFunc {
// 	return func(rw http.ResponseWriter, r *http.Request) {
// 		d.log.Debug("Get all documents")
// 		rw.Write([]byte(`{"status":"Get all documents!"}`))
// 	}
// }

func (d *Document) ListAll() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		d.log.Debug("Get all documents")
		ctx.JSON(http.StatusOK, gin.H{"status": "Get all documents!"})
	}
}

func (d *Document) ListOne() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		d.log.Debug("Get one document")
		ctx.JSON(http.StatusOK, gin.H{"status": "get one document"})
	}
}
func (d *Document) ChangeSituation() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		d.log.Debug("Change documents situation")
		ctx.JSON(http.StatusOK, gin.H{"status": "Change documents situation"})
	}
}

func (d *Document) ChangeSituationPDV() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		d.log.Debug("Change document situation")
		ctx.JSON(http.StatusOK, gin.H{"status": "Change document situation"})
	}
}
