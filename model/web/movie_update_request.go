package web

type MovieUpdateRequest struct {
	Id          int     `validate:"required" json:"id"`
	Title       string  `validate:"required,min=1,max=255" json:"title"`
	Description string  `json:"description"`
	Rating      float32 `validate:"number,min=0,max=10" json:"rating"`
	Image       string  `json:"image"`
}
