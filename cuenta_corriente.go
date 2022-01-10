package godoy

import "net/url"

// CuentaCorriente ...
type CuentaCorriente struct {
	CuentaCorriente []struct {
		Emision     string `json:"emision"`
		Hora        string `json:"hora"`
		Vencimiento string `json:"vencimiento"`
		Tipo        string `json:"tipo"`
		Naturaleza  string `json:"naturaleza"`
		Dias        string `json:"dias"`
		Numero      string `json:"numero"`
		Debe        string `json:"debe"`
		Haber       string `json:"haber"`
		Saldo       string `json:"saldo"`
	} `json:"cuenta_corriente"`
	Total int `json:"total"`
}

// GetCuentaCorriente lista de cuenta corriente.
func (g *Godoy) GetCuentaCorriente(id string, query url.Values) (*CuentaCorriente, error) {
	cuentaCorriente := CuentaCorriente{}

	err := g.get("/cuentacorriente/"+id, query, &cuentaCorriente)
	if err != nil {
		return nil, err
	}

	return &cuentaCorriente, nil
}
