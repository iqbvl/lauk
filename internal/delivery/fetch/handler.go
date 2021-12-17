package fetch

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/iqbvl/lauk/internal/model"
	"github.com/iqbvl/lauk/internal/platform/middleware"
	"github.com/iqbvl/lauk/internal/platform/util"
)

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

func (d *REST) GetStorageHandler(w http.ResponseWriter, r *http.Request) {
	var rsp []model.Storage
	// var rspAdmin []model.StorageAdmin

	var rates float64

	rates, err := d.FetchUsecase.GetRates(d.Context)
	if err != nil {
		responses := &Response{
			Data:          "",
			ErrorMessages: err.Error(),
		}
		middleware.WriteResponse(w, responses, http.StatusInternalServerError)
		return
	}

	_, claims, _ := jwtauth.FromContext(r.Context())
	role := claims["role"].(string)

	rsp, err = d.FetchUsecase.GetStorages(d.Context, model.GetStoragesRequest{Rates: rates, Role: role})
	if err != nil {
		responses := &Response{
			Data:          "",
			ErrorMessages: err.Error(),
		}
		middleware.WriteResponse(w, responses, http.StatusInternalServerError)
		return
	}

	if role == model.Admin { 
		mapProvince := make(map[string]string)
		//mapProvinceWeek := make(map[string][]int) 
		//mapProvinceWeekDays := make(map[int][]int) 

		log.Printf("Total data : %d \n", len(rsp))
		//var days []int
		for _, v := range rsp {  
			_, err := util.ParseDate(v.TglParsed)
			if err != nil {
				responses := &Response{
					Data:          "",
					ErrorMessages: err.Error(),
				}
				middleware.WriteResponse(w, responses, http.StatusInternalServerError)
				return
			}

			if _, ok := mapProvince[v.AreaProvinsi]; !ok {
				mapProvince[v.AreaProvinsi] = v.AreaProvinsi
			} 
		}
 

	}

	response, _ := json.Marshal(rsp)
	w.Write(response)
}
