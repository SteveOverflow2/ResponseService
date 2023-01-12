package response

// 1 model for whole Create user flow
type CreateResponse struct {
	Uuid    string
	Title   string `json:"title"`
	Body    string `json:"body"`
	Subject string `json:"subject"`
}

type Response struct {
	Id          string `json:"_id"`
	PostId      string `json:"post_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	Views       int    `json:"views"`
	Answers     int    `json:"answers"`
	Votes       int    `json:"votes"`
	Responseer  string `json:"responseer"`
}
