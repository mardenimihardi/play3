package types

type JSONResponse struct {
	Header JSONResponseHeader `json:"header"`
	Data   interface{}        `json:"data,omitempty"`
}
type JSONResponseHeader struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

type Dishes struct {
	Id         int    `json:"id,omitempty", db:"id"`
	Dish       string `json:"dish", db:"name"`
	Restaurant string `json:"restaurant"`
	Event      string `json:"event"`
	Venue      string `json:"venue"`
	Price      int    `json:"price"`
	Count      int    `json:"count"`
}
