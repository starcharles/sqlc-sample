package response

import "time"

type User struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserResponse struct {
	User User `json:"user"`
}
type GetUsersResponse struct {
	Users []*User `json:"users"`
}
