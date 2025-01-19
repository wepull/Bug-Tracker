package requests

type CreateInviteRequest struct {
	InviteeID uint `json:"invitee_id" binding:"required"`
}

type RespondInviteRequest struct {
	Accept bool `json:"accept"` // or "status" string
}
