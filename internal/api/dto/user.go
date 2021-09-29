package user_dto

type UserDto struct {
	Id int `json:"id"`
	Name string `json:"name"`
	UserName string `json:"username"`
	Email string `json:"email"`
}

type UserListDto struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type CreateUserDto struct {
	Name string `json:"name"`
	UserName string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}