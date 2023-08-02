package request

type GetUserRequest struct {
	Id int `json:"id" validate:"required" param:"id"`
}
