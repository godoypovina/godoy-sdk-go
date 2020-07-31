package godoy

import (
	"net/url"
)

// Articles ...
type Articles struct {
	Articles []struct {
		Code             string  `json:"code"`
		Price            float32 `json:"price"`
		Name             string  `json:"name"`
		BrandName        string  `json:"brand"`
		InnerDiameter    float32 `json:"inner_diameter"`
		OuterDiameter    float32 `json:"outer_diameter"`
		Width            float32 `json:"width"`
		RubroDescripcion string  `json:"category"`
		Equivalence      uint32  `json:"equivalence"`
		Line             string  `json:"line"`
		LineGroup        string  `json:"line_group"`
		ProviderID       uint    `json:"provider_id"`
	} `json:"articles"`
	Total int `json:"total"`
}

// GetAllArticles returns list of articles
func (g *Godoy) GetAllArticles(query url.Values) (*Articles, error) {
	articles := Articles{}

	err := g.get("/articles", query, &articles)
	if err != nil {
		return nil, err
	}

	return &articles, nil
}
