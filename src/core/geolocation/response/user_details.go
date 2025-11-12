package response

type UserDetails struct {
	UserId     uint   `json:"userId"`
	Nickname   string `json:"nickname"`
	ProfilePic string `json:"profilePic"`
}
