package godoy

import "net/url"

// Customers ...
type Customers struct {
	Customers []struct {
		ID                         uint    `json:"id"`
		Name                       string  `json:"name"`
		ProtheusCode               string  `json:"protheus_code"`
		Direccion                  string  `json:"address"`
		City                       string  `json:"city"`
		ZipCode                    string  `json:"zip_code"`
		State                      string  `json:"state"`
		Cuit                       string  `json:"cuit"`
		Code                       string  `json:"code"`
		SellerName                 string  `json:"seller_name"`
		Nivel1                     string  `json:"nivel1"`
		Nivel2                     string  `json:"nivel2"`
		Nivel3                     string  `json:"nivel3"`
		PaymentCondition           string  `json:"payment_condition"`
		PaymentConditionPercentage float32 `json:"payment_condition_percentage"`
		PaymentConditionDays       string  `json:"payment_condition_days"`
		Store                      string  `json:"store"`
		FleteMontoMinimo           float32 `json:"flete_monto_minimo"`
	} `json:"customers"`
	Total int `json:"total"`
}

type CustomersCode struct {
	CustomersCodes []struct {
		ErpClienteId               string `json:"erp_cliente_id"`
		ErpClienteCode             string `json:"erp_cliente_code"`
		ErpClienteTienda           string `json:"erp_cliente_tienda"`
		ProductoProtheusCodigo     string `json:"producto_protheus_codigo"`
		ProductoClienteCodigo      string `json:"producto_cliente_codigo"`
		ProductoClienteDescripcion string `json:"producto_cliente_descripcion"`
		Homologado                 string `json:"homologado"`
		DescripcionCliente         string `json:"descripcion_cliente"`
	} `json:"clients_codes"`
	Total int `json:"total"`
}

// GetAllCustomers returns list of customers
func (g *Godoy) GetAllCustomers(query url.Values) (*Customers, error) {
	customers := Customers{}

	err := g.get("/customers", query, &customers)
	if err != nil {
		return nil, err
	}

	return &customers, nil
}

// GetAllCustomers returns list of customersCodes
func (g *Godoy) GetAllCustomerCodes(query url.Values) (*CustomersCode, error) {
	customersCodes := CustomersCode{}

	err := g.get("/clientsCodes", query, &customersCodes)
	if err != nil {
		return nil, err
	}

	return &customersCodes, nil
}
