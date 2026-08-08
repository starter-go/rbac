package users

import "context"

// User_Service 是针对 UserDTO 的服务
type UserService interface {
	Insert(c context.Context, o *UserDTO) (*UserDTO, error)
	Update(c context.Context, id UserID, o *UserDTO) (*UserDTO, error)
	Delete(c context.Context, id UserID) error

	Find(c context.Context, id UserID) (*UserDTO, error)
	List(c context.Context, q *UserQuery) ([]*UserDTO, error)
}
