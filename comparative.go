package godoy

import "net/url"

// ArticleComparative ...
type ArticleComparative struct {
	Article struct {
		Code         string  `json:"code"`
		Name         string  `json:"name"`
		BrandName    string  `json:"brand_name"`
		Equivalence  uint    `json:"equivalence"`
		PriceCostDlr float32 `json:"price_cost_dlr"`
		PriceSaleDlr float32 `json:"price_sale_dlr"`
		ProviderID   float32 `json:"provider_id"`
	} `json:"article"`
}

// GetArticleComparative returns an article
func (g *Godoy) GetArticleComparative(query url.Values) (*ArticleComparative, error) {
	article := ArticleComparative{}

	err := g.get("/comparative", query, &article)
	if err != nil {
		return nil, err
	}

	return &article, nil
}
