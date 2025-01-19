package responses

type InviteResponse struct {
	ID        uint   `json:"id"`
	TeamID    uint   `json:"team_id"`
	InviterID uint   `json:"inviter_id"`
	InviteeID uint   `json:"invitee_id"`
	Status    string `json:"status"`
}
