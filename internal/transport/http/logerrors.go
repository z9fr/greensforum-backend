package http

import (
	"net/http"

	"github.com/z9fr/greensforum-backend/internal/utils"
)

func LogWarningsWithRequestInfo(r *http.Request, message interface{}) {
	loginfo := struct {
		Error      interface{}
		RemoteAddr string
		UserAgent  string
	}{
		Error:      message,
		RemoteAddr: r.RemoteAddr,
		UserAgent:  r.UserAgent(),
	}
	utils.LogWarn(loginfo)

}
