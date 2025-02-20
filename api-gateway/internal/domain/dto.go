package domain

type Product struct {
	Id       string `json:"_id"`
	Image    string `json:"image"`
	Title    string `json:"title"`
	Price    int    `json:"price"`
	Desc     string `json:"desc"`
	Category string `json:"category"`
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseErr struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
