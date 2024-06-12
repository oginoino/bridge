package models

type PixReq struct {
	Calendario         int                   `json:"calendario"`
	Devedor            DevedorModel          `json:"devedor"`
	Valor              float64               `json:"valor"`
	Chave              string                `json:"chave"`
	SolicitacaoPagador string                `json:"solicitacao_pagador"`
	InfoAdicionais     []InfoAdicionaisModel `json:"additional_info"`
}
