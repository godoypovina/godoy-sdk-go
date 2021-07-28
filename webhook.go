package godoy

// SendEventCatalogo ...
func (g *Godoy) SendEventCatalogo(data interface{}) (interface{}, error) {
	response := struct{}{}

	err := g.post("/webhook/catalogo", data, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
