package dto


type CreateProductDto struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Quantity uint `json:"quantity"`
	Image_url string `json:"image_url"`
}

type UpdateProductDto struct {
	Name        *string  `json:"name"`
    Description *string  `json:"description"`
    Price       *float64 `json:"price"`
    Quantity    *uint    `json:"quantity"`
    ImageURL    *string  `json:"imageURL"`
}