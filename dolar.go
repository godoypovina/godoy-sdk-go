package godoy

// DolarHoy ...
type DolarHoy struct {
	Dolar struct {
		DolarHoy float32 `json:"dolar_hoy"`
	} `json:"dolar"`
}

// GetStock ...
func (g *Godoy) GetDolarHoy() (*DolarHoy, error) {
	dolar := DolarHoy{}

	err := g.get("/dolar/hoy", nil, &dolar)
	if err != nil {
		return nil, err
	}

	return &dolar, nil
}
