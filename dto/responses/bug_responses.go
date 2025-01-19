package responses

type BugResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
	Status      string `json:"status"`
	ProjectID   uint   `json:"project_id"`
	CreatedBy   uint   `json:"created_by"`
	AssignedTo  uint   `json:"assigned_to"`
}
