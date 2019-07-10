package algoutil

import (
	"encoding/json"
	mathrand "math/rand"
	"net/http"
	"time"
)

//生成随机字符串
func RandStr(strlen int) string {
	mathrand.Seed(time.Now().Unix())
	data := make([]byte, strlen)
	var num int
	for i := 0; i < strlen; i++ {
		num = mathrand.Intn(57) + 65
		for {
			if num > 90 && num < 97 {
				num = mathrand.Intn(57) + 65
			} else {
				break
			}
		}
		data[i] = byte(num)
	}
	return string(data)
}

func AccessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		h.ServeHTTP(w, r)
	})
}

func OptionControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			json.NewEncoder(w).Encode(`{"code":0,"data":"success"}`)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func TimeRange(start, end int64) (int64, int64) {
	if start < 0 {
		start = 0
	}
	if end < 0 || end > time.Now().Unix() {
		end = time.Now().Unix()
	}
	if start > end {
		start, end = end, start
	}
	return start, end
}
