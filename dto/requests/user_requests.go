package requests

// Update user profile
type UpdateUserRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Username  *string `json:"username"`
	Email     *string `json:"email"`
	Password  *string `json:"password"` // if user wants to change password
}
