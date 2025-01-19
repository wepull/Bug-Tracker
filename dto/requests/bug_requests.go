package requests

type CreateBugRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Severity    string `json:"severity" binding:"required"`
	AssignedTo  *uint  `json:"assigned_to"`
}

type UpdateBugRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Severity    *string `json:"severity"`
	Status      *string `json:"status"`
	AssignedTo  *uint   `json:"assigned_to"`
}
