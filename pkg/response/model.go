package response

// 1 model for whole Create user flow
type CreateResponse struct {
	Uuid        string
	Description string `json:"description"`
	PostId      string `json:"post_id"`
	Subject     string `json:"subject"`
}

type Response struct {
	Uuid        string
	PostId      string `json:"post_id"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	Poster      string `json:"subject'`
}
