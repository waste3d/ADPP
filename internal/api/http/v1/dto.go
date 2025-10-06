package v1

type CreateJobRequest struct {
	Input int `json:"input"`
}

type CreateJobResponse struct {
	ID uint `json:"id"`
}
