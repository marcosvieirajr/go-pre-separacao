package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viavarejo-internal/pre-separacao-api/cmd/web/routes/handlers"
)

func MakeDocumentRoutes(r *gin.Engine, h *handlers.Document, m ...gin.HandlerFunc) {
	v1Auth := r.Group("/v1", m...)
	{
		v1Auth.GET("/pre-separacoes", h.ListAll())
		v1Auth.PUT("/pre-separacoes/alterar-situacao", h.ChangeSituation())
	}

	v1 := r.Group("/v1")
	{
		v1.GET("/pre-separacoes/:filial/:document", h.ListOne())
		v1.PUT("/pre-separacoes/:filial/:document/:situation", h.ChangeSituationPDV())
	}
}
