package auth

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/iqbvl/lauk/internal/model"
	"github.com/iqbvl/lauk/internal/platform/middleware"
	"github.com/iqbvl/lauk/internal/platform/util"
	"golang.org/x/net/context"
)

func (d *REST) GeneratePasswordHandler(w http.ResponseWriter, r *http.Request) {
	rsp := GeneratePassword{}
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		responses := &Response{
			Data:          "",
			ErrorMessages: postMethodSupported,
		}
		middleware.WriteResponse(w, responses, http.StatusMethodNotAllowed)
		return
	}

	user := model.User{}
	user, err := util.UserRequestBodyDecoder(r)
	if err != nil {
		responses := &Response{
			Data:          "",
			ErrorMessages: errorConvertRequestBody,
		}
		middleware.WriteResponse(w, responses, http.StatusBadRequest)
		return
	}

	respUser, err := d.AuthUsecase.GetUser(d.Context, user)
	if err != nil {
		response, _ := json.Marshal(rsp)
		w.Write(response)
		return
	}

	if respUser.Name != "" { //means exists
		rsp.Password = respUser.Password
		response, _ := json.Marshal(rsp)
		w.Write(response)
	} else {
		pwd := util.GeneratePassword()
		user.Password = pwd
		err := d.AuthUsecase.SetUser(d.Context, user)
		if err != nil {
			response, _ := json.Marshal(rsp)
			w.Write(response)

			return
		}

		rsp.Password = pwd
		response, _ := json.Marshal(rsp)
		w.Write(response)
	}
}

func (d *REST) GetJWTHandler(w http.ResponseWriter, r *http.Request) {
	rsp := GetJWTResponse{}
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		responses := &Response{
			Data:          "",
			ErrorMessages: postMethodSupported,
		}
		middleware.WriteResponse(w, responses, http.StatusMethodNotAllowed)
		return
	}

	user := model.User{}
	user, err := util.UserGetJWTRequestBodyDecoder(r)
	if err != nil {
		responses := &Response{
			Data:          "",
			ErrorMessages: errorConvertRequestBody,
		}
		middleware.WriteResponse(w, responses, http.StatusBadRequest)
		return
	}

	respUser, err := d.AuthUsecase.GetUser(d.Context, user)
	if err != nil {
		rsp.ErrorMessages = "not valid to get JWT, empty data from cache"
		response, _ := json.Marshal(rsp)
		w.Write(response)
		return
	}

	if respUser.Name == "" {
		response, _ := json.Marshal(rsp)
		w.Write(response)
		return
	}

	expired := util.ParseDuration("P1Y")
	tmstmp := time.Now().Format(`2006-01-02T15:04:05.000-07:00`)
	_, tokenString, _ := d.TokenJWT.Encode(jwt.MapClaims{
		"name":       respUser.Name,
		"phone":      respUser.Phone,
		"role":       respUser.Role,
		"timestamp":  tmstmp,
		"claim_date": tmstmp,
		"expire":     expired,
	})

	ctx := context.WithValue(r.Context(), "TokenAuth", d.TokenJWT)
	r.WithContext(ctx)

	//Success
	rsp.Token = tokenString
	rsp.Claims = JWTClaims{
		Name:      respUser.Name,
		Phone:     respUser.Phone,
		Role:      respUser.Role,
		Timestamp: tmstmp,
		ClaimDate: tmstmp,
		Expire:    expired,
	}
	rsp.ErrorMessages = ""

	response, _ := json.Marshal(rsp)
	w.Write(response)
}

func (d *REST) ValidateJWTHandler(w http.ResponseWriter, r *http.Request) {
	rsp := GetJWTResponse{}
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		responses := &Response{
			Data:          "",
			ErrorMessages: getMethodSupported,
		}
		middleware.WriteResponse(w, responses, http.StatusMethodNotAllowed)
		return
	}

	_, claims, _ := jwtauth.FromContext(r.Context())

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	rsp.Token = reqToken
	rsp.Claims = JWTClaims{
		Name:      claims["name"].(string),
		Phone:     claims["phone"].(string),
		Role:      claims["role"].(string),
		Timestamp: claims["timestamp"].(string),
		ClaimDate: claims["claim_date"].(string),
		Expire:    time.Duration(claims["expire"].(float64)),
	}
	response, _ := json.Marshal(rsp)
	w.Write(response)
}
