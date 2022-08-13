package godoy

import (
	"net/url"
)

// Articles ...
type Articles struct {
	Articles []struct {
		Code          string  `json:"code"`
		Price         float32 `json:"price"`
		PriceBaseDlr  float32 `json:"price_base_dlr"`
		PriceCostDlr  float32 `json:"price_cost_dlr"`
		PriceSaleDlr  float32 `json:"price_sale_dlr"`
		Name          string  `json:"name"`
		BrandName     string  `json:"brand_name"`
		InnerDiameter float32 `json:"inner_diameter"`
		OuterDiameter float32 `json:"outer_diameter"`
		Width         float32 `json:"width"`
		Category      string  `json:"category"`
		Equivalence   uint    `json:"equivalence"`
		ListID        uint    `json:"list_id"`
		Line          string  `json:"line"`
		LineID        string  `json:"line_id"`
		LineGroup     string  `json:"line_group"`
		LineGroupID   string  `json:"line_group_id"`
	} `json:"articles"`
	Total int `json:"total"`
}

// Article ...
type Article struct {
	Article struct {
		Code          string  `json:"code"`
		Price         float32 `json:"price"`
		PriceBaseDlr  float32 `json:"price_base_dlr"`
		PriceCostDlr  float32 `json:"price_cost_dlr"`
		PriceSaleDlr  float32 `json:"price_sale_dlr"`
		Name          string  `json:"name"`
		BrandName     string  `json:"brand_name"`
		InnerDiameter float32 `json:"inner_diameter"`
		OuterDiameter float32 `json:"outer_diameter"`
		Width         float32 `json:"width"`
		Category      string  `json:"category"`
		Equivalence   uint    `json:"equivalence"`
		ListID        uint    `json:"list_id"`
		Line          string  `json:"line"`
		LineID        string  `json:"line_id"`
		LineGroup     string  `json:"line_group"`
		LineGroupID   string  `json:"line_group_id"`
		ProviderID    uint    `json:"provider_id"`
	} `json:"article"`
}

// ArticlePack ...
type ArticlePack struct {
	Articles []struct {
		Code         string  `json:"code"`
		Price        float32 `json:"price"`
		Pack         string  `json:"pack"`
		Nivel        string  `json:"nivel"`
		PriceCostDlr float32 `json:"price_cost_dlr"`
		PriceSaleDlr float32 `json:"price_sale_dlr"`
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

// GetAllArticlesCatalogo returns list of articles
func (g *Godoy) GetAllArticlesCatalogo(query url.Values) (*Articles, error) {
	articles := Articles{}

	err := g.get("/articles/catalogo", query, &articles)
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

// GetAllArticlesPacks returns list of articles
func (g *Godoy) GetAllArticlesPacks(query url.Values) (*ArticlePack, error) {
	articles := ArticlePack{}

	err := g.get("/articles/packs", query, &articles)
	if err != nil {
		return nil, err
	}

	return &articles, nil
}
