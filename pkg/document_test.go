package pkg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocument_Separar(t *testing.T) {
	type fields struct {
		ID               string
		DocumentoDeVenda DocumentoDeVenda
		Cliente          Cliente
		Movimentacoes    Movimentacoes
	}
	type args struct {
		estoquista Estoquista
		canal      string
	}

	estq := Estoquista{Nome: "Marcos Vieira Jr", Matricula: 12345, Filial: 1000}
	arg := args{estoquista: estq, canal: "padrao"}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "com Situacao CANCELAR",
			fields:  fields{DocumentoDeVenda: DocumentoDeVenda{Situacao: Cancelar}},
			args:    arg,
			wantErr: assert.Error,
		},
		{
			name:    "com Situacao SEPARAR",
			fields:  fields{DocumentoDeVenda: DocumentoDeVenda{Situacao: Separar}},
			args:    arg,
			wantErr: assert.NoError,
		},
		{
			name:    "com Situacao SEPARADO",
			fields:  fields{DocumentoDeVenda: DocumentoDeVenda{Situacao: Separado}},
			args:    arg,
			wantErr: assert.Error,
		},
		{
			name:    "com Situacao ENTREGUE",
			fields:  fields{DocumentoDeVenda: DocumentoDeVenda{Situacao: Entregue}},
			args:    arg,
			wantErr: assert.Error,
		},
		{
			name:    "com Situacao CANCELADO",
			fields:  fields{DocumentoDeVenda: DocumentoDeVenda{Situacao: Cancelado}},
			args:    arg,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Document{
				ID:               tt.fields.ID,
				DocumentoDeVenda: tt.fields.DocumentoDeVenda,
				Cliente:          tt.fields.Cliente,
				Movimentacoes:    tt.fields.Movimentacoes,
			}

			err := d.Separar(tt.args.estoquista, tt.args.canal)

			tt.wantErr(t, err, fmt.Sprintf("Document.Separar(%#v, %#v)", tt.args.estoquista, tt.args.canal))

			if err == nil {
				assert.Equal(t, Situacao(Separado), d.DocumentoDeVenda.Situacao)
				assert.Equal(t, 3, d.DocumentoDeVenda.OrdenacaoSituacao)
				assert.Equal(t, tt.args.canal, d.DocumentoDeVenda.SeparacaoDeEstoque.Canal)
				assert.Equal(t, tt.args.estoquista, d.DocumentoDeVenda.SeparacaoDeEstoque.Estoquista)
			}
		})
	}
}

func TestDocument_Entregar(t *testing.T) {
	type fields struct {
		ID               string
		DocumentoDeVenda DocumentoDeVenda
		Cliente          Cliente
		Movimentacoes    Movimentacoes
	}
	type args struct {
		estoquista Estoquista
		canal      string
		mov        Movimentacoes
	}

	estq := Estoquista{Nome: "Marcos Vieira Jr", Matricula: 12345, Filial: 1000}
	arg := args{estoquista: estq, canal: "pdv"}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "com Situacao CANCELAR",
			fields:  fields{DocumentoDeVenda: DocumentoDeVenda{Situacao: Cancelar}},
			args:    arg,
			wantErr: true,
		},
		{
			name:    "com Situacao SEPARAR",
			fields:  fields{DocumentoDeVenda: DocumentoDeVenda{Situacao: Separar}},
			args:    arg,
			wantErr: false,
		},
		{
			name:    "com Situacao SEPARADO",
			fields:  fields{DocumentoDeVenda: DocumentoDeVenda{Situacao: Separado}},
			args:    arg,
			wantErr: false,
		},
		{
			name:    "com Situacao ENTREGUE",
			fields:  fields{DocumentoDeVenda: DocumentoDeVenda{Situacao: Entregue}},
			args:    arg,
			wantErr: true,
		},
		{
			name:    "com Situacao CANCELADO",
			fields:  fields{DocumentoDeVenda: DocumentoDeVenda{Situacao: Cancelado}},
			args:    arg,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Document{
				ID:               tt.fields.ID,
				DocumentoDeVenda: tt.fields.DocumentoDeVenda,
				Cliente:          tt.fields.Cliente,
				Movimentacoes:    tt.fields.Movimentacoes,
			}

			var err error
			if err = d.Entregar(tt.args.estoquista, tt.args.canal, tt.args.mov); (err != nil) != tt.wantErr {
				t.Errorf("Document.Entregar() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				assert.Equal(t, Situacao(Entregue), d.DocumentoDeVenda.Situacao)
				assert.Equal(t, 4, d.DocumentoDeVenda.OrdenacaoSituacao)
				assert.Equal(t, tt.args.canal, d.DocumentoDeVenda.EntregaDeEstoque.Canal)
				assert.Equal(t, tt.args.estoquista, d.DocumentoDeVenda.EntregaDeEstoque.Estoquista)
				assert.NotNil(t, d.Movimentacoes)
			}
		})
	}
}

func TestDocument_Cancelar(t *testing.T) {
	type fields struct {
		ID               string
		DocumentoDeVenda DocumentoDeVenda
		Cliente          Cliente
		Movimentacoes    Movimentacoes
	}
	type args struct {
		estoquista Estoquista
		canal      string
	}

	estq := Estoquista{Nome: "Marcos Vieira Jr", Matricula: 12345, Filial: 1000}
	arg := args{estoquista: estq, canal: "padrao"}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "com Situacao CANCELAR",
			fields:  fields{DocumentoDeVenda: DocumentoDeVenda{Situacao: Cancelar}},
			args:    arg,
			wantErr: assert.NoError,
		},
		{
			name:    "com Situacao SEPARAR",
			fields:  fields{DocumentoDeVenda: DocumentoDeVenda{Situacao: Separar}},
			args:    arg,
			wantErr: assert.Error,
		},
		{
			name:    "com Situacao SEPARADO",
			fields:  fields{DocumentoDeVenda: DocumentoDeVenda{Situacao: Separado}},
			args:    arg,
			wantErr: assert.Error,
		},
		{
			name:    "com Situacao ENTREGUE",
			fields:  fields{DocumentoDeVenda: DocumentoDeVenda{Situacao: Entregue}},
			args:    arg,
			wantErr: assert.Error,
		},
		{
			name:    "com Situacao CANCELADO",
			fields:  fields{DocumentoDeVenda: DocumentoDeVenda{Situacao: Cancelado}},
			args:    arg,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Document{
				ID:               tt.fields.ID,
				DocumentoDeVenda: tt.fields.DocumentoDeVenda,
				Cliente:          tt.fields.Cliente,
				Movimentacoes:    tt.fields.Movimentacoes,
			}
			err := d.Cancelar(tt.args.estoquista, tt.args.canal)
			tt.wantErr(t, err, fmt.Sprintf("Cancelar(%v)", tt.args.estoquista))
			if err == nil {
				assert.Equal(t, Situacao(Cancelado), d.DocumentoDeVenda.Situacao)
				assert.Equal(t, 5, d.DocumentoDeVenda.OrdenacaoSituacao)
				assert.Equal(t, "padrao", d.DocumentoDeVenda.CancelamentoDeEstoque.Canal)
				assert.Equal(t, estq, d.DocumentoDeVenda.CancelamentoDeEstoque.Estoquista)
			}
		})
	}
}
