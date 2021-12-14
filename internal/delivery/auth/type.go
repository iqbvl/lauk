package auth

type Response struct {
	Data          string `json:"data"`
	ErrorMessages string `json:"error"`
}

type GeneratePassword struct {
	Password string `json:"password"`
}
