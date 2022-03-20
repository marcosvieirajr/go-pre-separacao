package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentoDeVenda_Movimentar(t *testing.T) {
	type fields struct {
		Situacao Situacao
	}
	type args struct {
		mov        string
		estoquista Estoquista
		canal      string
	}

	estq := Estoquista{Nome: "Marcos Vieira Jr", Matricula: 12345, Filial: 1000}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Separar com Situacao SEPARAR",
			fields:  fields{Situacao: Separar},
			args:    args{mov: "SEPARADO", estoquista: estq, canal: "padrao"},
			wantErr: false,
		},
		{
			name:    "Separar com Situacao CANCELAR",
			fields:  fields{Situacao: Cancelar},
			args:    args{mov: "SEPARADO", estoquista: estq, canal: "padrao"},
			wantErr: true,
		},
		{
			name:    "Separar com Situacao SEPARADO",
			fields:  fields{Situacao: Separado},
			args:    args{mov: "SEPARADO", estoquista: estq, canal: "padrao"},
			wantErr: true,
		},
		{
			name:    "Separar com Situacao ENTREGUE",
			fields:  fields{Situacao: Entregue},
			args:    args{mov: "SEPARADO", estoquista: estq, canal: "padrao"},
			wantErr: true,
		},
		{
			name:    "Separar com Situacao CANCELADO",
			fields:  fields{Situacao: Cancelado},
			args:    args{mov: "SEPARADO", estoquista: estq, canal: "padrao"},
			wantErr: true,
		},
		{
			name:    "Entregar com Situacao SEPARAR",
			fields:  fields{Situacao: Separar},
			args:    args{mov: "ENTREGUE", estoquista: estq, canal: "pdv"},
			wantErr: false,
		},
		{
			name:    "Entregar com Situacao CANCELAR",
			fields:  fields{Situacao: Cancelar},
			args:    args{mov: "ENTREGUE", estoquista: estq, canal: "pdv"},
			wantErr: true,
		},
		{
			name:    "Entregar com Situacao SEPARADO",
			fields:  fields{Situacao: Separado},
			args:    args{mov: "ENTREGUE", estoquista: estq, canal: "pdv"},
			wantErr: false,
		},
		{
			name:    "Entregar com Situacao ENTREGUE",
			fields:  fields{Situacao: Entregue},
			args:    args{mov: "ENTREGUE", estoquista: estq, canal: "pdv"},
			wantErr: true,
		},
		{
			name:    "Entregar com Situacao CANCELADO",
			fields:  fields{Situacao: Cancelado},
			args:    args{mov: "ENTREGUE", estoquista: estq, canal: "pdv"},
			wantErr: true,
		},
		{
			name:    "Cancelar com Situacao CANCELAR",
			fields:  fields{Situacao: Cancelar},
			args:    args{mov: "CANCELADO", estoquista: estq, canal: "padrao"},
			wantErr: false,
		},
		{
			name:    "Cancelar com Situacao SEPARAR",
			fields:  fields{Situacao: Separar},
			args:    args{mov: "CANCELADO", estoquista: estq, canal: "padrao"},
			wantErr: true,
		},
		{
			name:    "Cancelar com Situacao SEPARADO",
			fields:  fields{Situacao: Separado},
			args:    args{mov: "CANCELADO", estoquista: estq, canal: "padrao"},
			wantErr: true,
		},
		{
			name:    "Cancelar com Situacao ENTREGUE",
			fields:  fields{Situacao: Entregue},
			args:    args{mov: "CANCELADO", estoquista: estq, canal: "padrao"},
			wantErr: true,
		},
		{
			name:    "Cancelar com Situacao CANCELADO",
			fields:  fields{Situacao: Cancelado},
			args:    args{mov: "CANCELADO", estoquista: estq, canal: "padrao"},
			wantErr: true,
		},
		{
			name:    "Movimentar para SEPARAR",
			fields:  fields{Situacao: Separar},
			args:    args{mov: "SEPARAR"},
			wantErr: true,
		},
		{
			name:    "Movimentar para CANCELAR",
			fields:  fields{Situacao: Cancelar},
			args:    args{mov: "CANCELAR"},
			wantErr: true,
		},
		{
			name:    "Movimentação inválida",
			fields:  fields{Situacao: Separar},
			args:    args{mov: "undefined"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DocumentoDeVenda{
				Situacao: tt.fields.Situacao,
			}
			var err error
			if err = d.Movimentar(tt.args.mov, tt.args.estoquista, tt.args.canal); (err != nil) != tt.wantErr {
				t.Errorf("DocumentoDeVenda.Movimentar() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				s := Situacao(tt.args.mov)
				assert.Equal(t, s, d.Situacao)
				assert.Equal(t, ordem[s], d.OrdenacaoSituacao)
				if s == Separado {
					assert.Equal(t, tt.args.canal, d.SeparacaoDeEstoque.Canal)
					assert.Equal(t, tt.args.estoquista, d.SeparacaoDeEstoque.Estoquista)
				} else if s == Entregue {
					assert.Equal(t, tt.args.canal, d.EntregaDeEstoque.Canal)
					assert.Equal(t, tt.args.estoquista, d.EntregaDeEstoque.Estoquista)
				} else if s == Cancelado {
					assert.Equal(t, tt.args.canal, d.CancelamentoDeEstoque.Canal)
					assert.Equal(t, tt.args.estoquista, d.CancelamentoDeEstoque.Estoquista)
				}
			}
		})
	}
}
