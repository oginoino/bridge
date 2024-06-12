package models

type PixRes struct {
	Calendario         int          `json:"calendario"`
	Txid               string       `json:"txid"`
	Revisao            int          `json:"revisao"`
	Loc                LocModel     `json:"loc"`
	Location           string       `json:"location"`
	Status             string       `json:"status"`
	Devedor            DevedorModel `json:"devedor"`
	Valor              float64      `json:"valor"`
	Chave              string       `json:"chave"`
	SolicitacaoPagador string       `json:"solicitacaoPagador"`
}

type LocModel struct {
	Id       int    `json:"id"`
	Location string `json:"location"`
	TipoCob  string `json:"tipoCob"`
}
