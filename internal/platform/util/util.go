package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/iqbvl/lauk/internal/model"
)

var (
	loc, _ = time.LoadLocation("Asia/Jakarta")
)

func GeneratePassword() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 4
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

func ParseDuration(str string) time.Duration {
	durationRegex := regexp.MustCompile(`P(?P<years>\d+Y)?(?P<months>\d+M)?(?P<days>\d+D)?T?(?P<hours>\d+H)?(?P<minutes>\d+M)?(?P<seconds>\d+S)?`)
	matches := durationRegex.FindStringSubmatch(str)

	years := ParseInt64(matches[1])
	months := ParseInt64(matches[2])
	days := ParseInt64(matches[3])
	hours := ParseInt64(matches[4])
	minutes := ParseInt64(matches[5])
	seconds := ParseInt64(matches[6])

	hour := int64(time.Hour)
	minute := int64(time.Minute)
	second := int64(time.Second)
	return time.Duration(years*24*365*hour + months*30*24*hour + days*24*hour + hours*hour + minutes*minute + seconds*second)
}

func ParseInt64(value string) int64 {
	if len(value) == 0 {
		return 0
	}
	parsed, err := strconv.Atoi(value[:len(value)-1])
	if err != nil {
		return 0
	}
	return int64(parsed)
}

func GenerateKey(in model.User) string {
	return fmt.Sprintf("%s", in.Phone)
}

func UserRequestBodyDecoder(req *http.Request) (model.User, error) {
	var t model.User = model.User{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return t, err
	}

	if t.Name == "" || t.Phone == "" || t.Role == "" {
		return model.User{}, errors.New("invalid request")
	}
	defer req.Body.Close()
	return t, nil
}

func UserGetJWTRequestBodyDecoder(req *http.Request) (model.User, error) {
	var t model.User = model.User{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return t, err
	}

	defer req.Body.Close()
	return t, nil
}

func GetToken(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")

	return splitToken[1]
}

//GetRatesExpiryTime return time duration of remaining time from request happen until end of day
func GetRatesExpiryTime() time.Duration {
	now := time.Now().In(loc)
	eod := now

	nxt := eod.AddDate(0, 0, 1)
	nxt = time.Date(nxt.Year(), nxt.Month(), nxt.Day(), 0, 0, 0, 0, loc)
	return nxt.Sub(now)
}

func ParseDate(in string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, in)
	if err != nil {
		return time.Time{}, err
	}

	result := t.In(loc)
	return result, nil
}

func FindMinAndMax(a []int) (min int, max int) {

	mapW := make(map[int]int)
	for _, v := range a {
		mapW[v] = mapW[v] + 1
	} 

	min = 0
	max = 0

	for _, value := range mapW {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	if len(a) == 1 {
		min = max
	}

	return min, max
}

func FindAvg(a []int) int {

	var sum int
	mapW := make(map[int]int)
	for _, v := range a {
		mapW[v] = mapW[v] + 1
	}

	for _, v := range mapW {
		sum = sum + v
	}

	return sum / len(a)
}

func FindMedian(a []int) int {
	mapW := make(map[int]int)
	for _, v := range a {
		mapW[v] = mapW[v] + 1
	}

	var tA []int
	for _, v := range mapW {
		tA = append(tA, v)
	}
	medianNum := len(tA) / 2
	sort.Ints(tA)
	if len(tA)%2 == 0 {
		return (tA[medianNum-1] + tA[medianNum]) / 2
	}

	return tA[medianNum]
}

func FindFinalData(provinceData map[int]model.StorageAgg) map[int]model.StorageAgg {
	res := make(map[int]model.StorageAgg)
	for k, v := range provinceData {
		min, max := FindMinAndMax(v.TxnInAWeek)
		avg := FindAvg(v.TxnInAWeek)
		median := FindMedian(v.TxnInAWeek)

		d := model.StorageAgg{
			Avg:        avg,
			Median:     median,
			Min:        min,
			Max:        max,
			TxnInAWeek: v.TxnInAWeek,
		}
		res[k] = d
	}

	return res
}
