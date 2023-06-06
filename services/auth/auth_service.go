package auth

type UserAuthInterface interface {
	Login(string, string) (bool, error)
	Logout() bool
	Register(UserAuth)
}

type UserAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
