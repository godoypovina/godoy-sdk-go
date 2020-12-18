package godoy

import "net/url"

// ArticlesComparative ...
type ArticlesComparative struct {
	Articles []struct {
		Code           string  `json:"code"`
		PriceCostDlrT2 float32 `json:"price_cost_dlr_t2"`
		PriceCostDlrT3 float32 `json:"price_cost_dlr_t3"`
		PriceSaleDlr   float32 `json:"price_sale_dlr"`
	} `json:"articles"`
	Total int `json:"total"`
}

// GetAllArticlesComparative returns an article
func (g *Godoy) GetAllArticlesComparative(query url.Values) (*ArticlesComparative, error) {
	articles := ArticlesComparative{}

	err := g.get("/articles-comparative", nil, &articles)
	if err != nil {
		return nil, err
	}

	return &articles, nil
}
