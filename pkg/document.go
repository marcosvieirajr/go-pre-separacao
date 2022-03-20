package pkg

import (
	"fmt"
	"time"
)

type Document struct {
	ID               string
	DocumentoDeVenda DocumentoDeVenda
	Cliente          Cliente
	Movimentacoes    Movimentacoes
}
type Filial struct {
	Codigo  int
	Empresa int
}

type Estoquista struct {
	Matricula int
	Nome      string
	Filial    int
}

type MovimentacaoDeEstoque struct {
	Estoquista Estoquista
	Canal      string
	Data       time.Time
}

type Produtos struct {
	Sku        string
	Ean        string
	Dun        string
	Tipo       string
	Quantidade int
	Descricao  string
	Estado     string
	Valor      string
}
type DocumentoDeVenda struct {
	CodigoDocumento       string
	PedidoInternet        string
	DataPedido            time.Time
	Tipo                  string
	TipoEntrega           string
	LocalEstoque          string
	Situacao              Situacao
	FilialPedido          Filial
	FilialRetirada        Filial
	SeparacaoDeEstoque    MovimentacaoDeEstoque
	EntregaDeEstoque      MovimentacaoDeEstoque
	CancelamentoDeEstoque MovimentacaoDeEstoque
	Produtos              []Produtos
	OrdenacaoSituacao     uint8
	OrdenacaoLocalEstoque uint8
}
type Cliente struct {
	Codigo    int
	Nome      string
	Documento string
}
type Movimentacoes struct {
	Pdv                                int
	FilialVenda                        int
	FilialEmissao                      int
	FilialEntrega                      int
	ConsultaRetiraRapido               interface{}
	TipoConsultaRetiraRapido           interface{}
	TipoValidacaoRetiraRapido          interface{}
	InicioValidacaoRetiraRapido        interface{}
	FimValidacaoRetiraRapido           interface{}
	InicioImpressaoComprovanteRetirada interface{}
	FimImpressaoComprovanteRetirada    interface{}
	ConsultaRetiraLoja                 time.Time
	MatriculaEstoquista                int
	InicioEntrega                      time.Time
	InicioDoubleCheck                  interface{}
	FimDoubleCheck                     time.Time
	InicioConfirmacaoVistoria          time.Time
	FimConfirmacaoVistoria             time.Time
	InicioProcessamentoDocumento       time.Time
	FimProcessamentoDocumento          time.Time
	LiberacaoGerencial                 bool
}

type Situacao string

const (
	Separar   Situacao = "SEPARAR"
	Separado  Situacao = "SEPARADO"
	Entregue  Situacao = "ENTREGUE"
	Cancelar  Situacao = "CANCELAR"
	Cancelado Situacao = "CANCELADO"
)

func (s Situacao) IsValid() error {
	switch s {
	case Separar, Separado, Entregue, Cancelar, Cancelado:
		return nil
	}
	return fmt.Errorf("situação inválida")
}

type ordens map[Situacao]uint8

var ordem = ordens{
	Separar:   1,
	Separado:  3,
	Entregue:  4,
	Cancelar:  2,
	Cancelado: 5,
}

func (d *DocumentoDeVenda) Movimentar(situacao string, estoquista Estoquista, canal string) error {
	sit := Situacao(situacao)

	switch sit {
	case Separado:
		if d.Situacao != Separar {
			return fmt.Errorf("não é possível %s um documento com a siturção '%s'", "Separar", d.Situacao)
		}

	case Entregue:
		if d.Situacao != Separar && d.Situacao != Separado {
			return fmt.Errorf("não é possível %s um documento com a siturção '%s'", "Entregar", d.Situacao)
		}

	case Cancelado:
		if d.Situacao != Cancelar {
			return fmt.Errorf("não é possível %s um documento com a siturção '%s'", "Cancelar", d.Situacao)
		}

	case Separar, Cancelar:
		return fmt.Errorf("não é possível alterar a Situação para '%v'", situacao)

	default:
		return fmt.Errorf("situação inválida")
	}

	d.movimentar(sit, estoquista, canal)
	return nil
}

func (d *DocumentoDeVenda) movimentar(sit Situacao, estoquista Estoquista, canal string) {
	d.Situacao = sit
	d.OrdenacaoSituacao = ordem[sit]

	movEstq := MovimentacaoDeEstoque{
		Estoquista: estoquista,
		Canal:      canal,
		Data:       time.Now(),
	}

	if sit == Separado {
		d.SeparacaoDeEstoque = movEstq
	} else if sit == Entregue {
		d.EntregaDeEstoque = movEstq
	} else if sit == Cancelado {
		d.CancelamentoDeEstoque = movEstq
	}
}
