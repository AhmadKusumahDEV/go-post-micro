package domain

type PostProduct struct {
	Image    string `validate:"required" json:"image"`
	Title    string `validate:"required" json:"title"`
	Price    int    `validate:"required" json:"price"`
	Desc     string `validate:"required" json:"desc"`
	Category string `validate:"required" json:"category"`
}

type UpdateProduct struct {
	Id       string `json:"_id"`
	Image    string `json:"image"`
	Title    string `json:"title"`
	Price    int    `json:"price"`
	Desc     string `json:"desc"`
	Category string `json:"category"`
}

type DeleteProduct struct {
	Id string `validate:"required" json:"_id"`
}
