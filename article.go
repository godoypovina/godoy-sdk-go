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
		BrandName        string  `json:"brand_name"`
		InnerDiameter    float32 `json:"inner_diameter"`
		OuterDiameter    float32 `json:"outer_diameter"`
		Width            float32 `json:"width"`
		Category         string  `json:"category"`
		Equivalence      uint    `json:"equivalence"`
		ListID           uint    `json:"list_id"`
		Line             string  `json:"line"`
		LineGroup        string  `json:"line_group"`
		StockReconquista int     `json:"stock_reconquista"`
		StockResistencia int     `json:"stock_resistencia"`
		StockPosadas     int     `json:"stock_posadas"`
		StockTucuman     int     `json:"stock_tucuman"`
		StockBuenosAires int     `json:"stock_buenosaires"`
	} `json:"articles"`
	Total int `json:"total"`
}

// Article ...
type Article struct {
	Article struct {
		Code             string  `json:"code"`
		Price            float32 `json:"price"`
		Name             string  `json:"name"`
		BrandName        string  `json:"brand_name"`
		InnerDiameter    float32 `json:"inner_diameter"`
		OuterDiameter    float32 `json:"outer_diameter"`
		Width            float32 `json:"width"`
		Category         string  `json:"category"`
		Equivalence      uint    `json:"equivalence"`
		ListID           uint    `json:"list_id"`
		Line             string  `json:"line"`
		LineGroup        string  `json:"line_group"`
		ProviderID       uint    `json:"provider_id"`
		StockReconquista int     `json:"stock_reconquista"`
		StockResistencia int     `json:"stock_resistencia"`
		StockPosadas     int     `json:"stock_posadas"`
		StockTucuman     int     `json:"stock_tucuman"`
		StockBuenosAires int     `json:"stock_buenosaires"`
	} `json:"article"`
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

// GetArticle returns an article
func (g *Godoy) GetArticle(code string) (*Article, error) {
	article := Article{}

	err := g.get(`/articles/`+code, nil, &article)
	if err != nil {
		return nil, err
	}

	return &article, nil
}
