package userVdr

type GetReq struct {
	Id int64 `form:"id" json:"id" binding:"required,min=1"`
}
type GetReplyUser struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
