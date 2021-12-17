package auth

import (
	"encoding/json"
	"net/http"
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
		Name:      respUser.Name,
		Phone:     respUser.Phone,
		Role:      respUser.Role,
		Timestamp: tmstmp,
		ClaimDate: tmstmp,
		Expire:    expired,
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

	rsp.Token = util.GetToken(r)
	rsp.Claims = JWTClaims{
		Name:      claims[Name].(string),
		Phone:     claims[Phone].(string),
		Role:      claims[Role].(string),
		Timestamp: claims[Timestamp].(string),
		ClaimDate: claims[ClaimDate].(string),
		Expire:    time.Duration(claims[Expire].(float64)),
	}
	response, _ := json.Marshal(rsp)
	w.Write(response)
}
