package auth

import "time"

type Response struct {
	Data          string `json:"data"`
	ErrorMessages string `json:"error"`
}

type GetJWTResponse struct {
	Token         string    `json:"token"`
	Claims        JWTClaims `json:"claims"`
	ErrorMessages string    `json:"error"`
}

type JWTClaims struct {
	Name      string        `json:"name"`
	Phone     string        `json:"phone"`
	Role      string        `json:"role"`
	Timestamp string        `json:"timestamp"`
	ClaimDate string        `json:"claim_date"`
	Expire    time.Duration `json:"expire"`
}

type GeneratePassword struct {
	Password string `json:"password"`
}
