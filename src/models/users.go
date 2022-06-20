package emulator

type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var Users = map[string]string{
	"superuser1": "qwerty123",
	"admin":      "password",
}
