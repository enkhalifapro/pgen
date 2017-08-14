package models

type SimilarTechTechnology struct {
	ID         int `json:"id"`
	Name       string `json:"name"`
	Pages      []string `json:"pages"`
	Coverage   float64     `json:"coverage"`
	Categories []string `json:"categories"`
	Paying     string `json:"paying"`
}

type SimilarTechResponse struct {
	Site         string `json:"site"`
	Found        bool        `json:"found"`
	Technologies []SimilarTechTechnology `json:"technologies"`
}
