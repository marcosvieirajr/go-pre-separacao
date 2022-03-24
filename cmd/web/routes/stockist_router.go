package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viavarejo-internal/pre-separacao-api/cmd/web/routes/handlers"
)

func MakeStockistRoutes(r *gin.Engine, h *handlers.Stockist, m ...gin.HandlerFunc) {
	v1 := r.Group("/v1", m...)
	{
		v1.GET("/estoquistas", h.ListAll())
	}
}

// import (
// 	"github.com/gin-gonic/gin"
// )

// type DocumentHandler interface {
// 	ListAll() func(ctx *gin.Context)
// 	// ListOne() http.HandlerFunc
// 	// ChangeSituation() http.HandlerFunc
// 	// ChangeSituationPDV() http.HandlerFunc
// }

// // func MakeDocumentRoutes(r *mux.Router, h DocumentHandler) {
// // 	listAll := r.Methods(http.MethodGet).Subrouter()
// // 	listAll.HandleFunc("/v1/pre-separacoes", h.ListAll())

// // 	// listOne := r.Methods(http.MethodGet).Subrouter()
// // 	// listOne.HandleFunc("/v1/pre-separacoes/{filial-pedido}/{documento-venda}", h.ListOne())

// // 	// changeSituation := r.Methods(http.MethodPut).Subrouter()
// // 	// changeSituation.HandleFunc("/v1/pre-separacoes/alterar-situacao", h.ChangeSituation())

// // 	// changeSituationPDV := r.Methods(http.MethodPut).Subrouter()
// // 	// changeSituationPDV.HandleFunc("/v1/pre-separacoes/{filial-pedido}/{documento-venda}/{situacao}", h.ChangeSituationPDV())
// // }

// // type StockistHandler interface {
// // 	ListAll() http.HandlerFunc
// // }

// // func MakeStockistRoutes(r *mux.Router, h StockistHandler) {
// // 	listAll := r.Methods(http.MethodGet).Subrouter()
// // 	listAll.HandleFunc("/v1/estoquistas", h.ListAll())
// // }

// func MakeDocumentRoutesGin(r *gin.Engine, h DocumentHandler) { //func(c *gin.Context)
// 	r.GET("/v1/pre-separacoes", h.ListAll())
// }
