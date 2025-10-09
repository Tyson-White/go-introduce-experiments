package dto

type CreateCategory struct {
	Name string `json:"name"`
}

type DeleteCategoryBody struct {
	Category int `json:"category"`
}
