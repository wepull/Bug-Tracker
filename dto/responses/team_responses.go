package responses

type TeamResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   uint   `json:"created_by"`
}
