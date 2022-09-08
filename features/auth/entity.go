package auth

type Core struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UsecaseInterface interface {
	Login(data Core) (string, error)
}

type DataInterface interface {
	InsertData(data Core) (string, error)
}
