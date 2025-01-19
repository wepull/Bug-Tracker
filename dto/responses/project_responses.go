package responses

type ProjectResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      *uint  `json:"user_id"`
	TeamID      *uint  `json:"team_id"`
}
