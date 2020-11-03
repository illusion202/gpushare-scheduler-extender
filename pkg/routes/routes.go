package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/KuaishouContainerService/quota-order-webhook/pkg/channel"
	"github.com/julienschmidt/httprouter"
)

const (
	versionPath         = "/version"
	onSubmitPrefix      = "/onsubmit"
	onGetNextActsPrefix = "/ongetnextacts"
)

var (
	version = "0.1.0"
)

func checkBody(w http.ResponseWriter, r *http.Request) bool {
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return false
	}
	return true
}

func checkToken(w http.ResponseWriter, r *http.Request) bool {
	if r.Header == nil || r.Header.Get("Authorization") == "" {
		http.Error(w, "Token is required", http.StatusUnauthorized)
		return false
	}
	return true
}

func OnSubmitRoute(onSubmit *channel.OnSubmit) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if !checkBody(w, r) || !checkToken(w, r) {
			return
		}

		var buf bytes.Buffer
		body := io.TeeReader(r.Body, &buf)

		var postBody channel.PostBody
		var onSubmitResp *channel.OnSubmitResp
		failed := false

		if err := json.NewDecoder(body).Decode(&postBody); err != nil {
			onSubmitResp = &channel.OnSubmitResp{
				ErrorMsg: err.Error(),
			}
			failed = true
		} else {
			log.Printf("debug: OnSubmitRoute PostBody =%v", postBody)
			onSubmitResp = onSubmit.Handler(&postBody)
		}

		if len(onSubmitResp.ErrorMsg) > 0 {
			failed = true
		}

		if resultBody, err := json.Marshal(onSubmitResp); err != nil {
			log.Printf("warn: Failed due to %v", err)
			// panic(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			errMsg := fmt.Sprintf("{'error':'%s'}", err.Error())
			w.Write([]byte(errMsg))
		} else {
			log.Print("info: OnSubmitResponse = ", string(resultBody))
			w.Header().Set("Content-Type", "application/json")
			if failed {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
			}

			w.Write(resultBody)
		}
	}
}

func OnGetNextActsRoute(onGetNextActs *channel.OnGetNextActs) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if !checkBody(w, r) || !checkToken(w, r) {
			return
		}

		var buf bytes.Buffer
		body := io.TeeReader(r.Body, &buf)

		var postBody channel.PostBody
		var onGetNextActsResp *channel.OnGetNextActsResp
		var err error
		failed := false

		if err = json.NewDecoder(body).Decode(&postBody); err != nil {
			failed = true
		} else {
			log.Printf("debug: OnSubmitRoute PostBody =%v", postBody)
			onGetNextActsResp, err = onGetNextActs.Handler(&postBody)
		}

		if err != nil {
			failed = true
		}

		if resultBody, err := json.Marshal(onGetNextActsResp); err != nil {
			log.Printf("warn: Failed due to %v", err)
			// panic(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			errMsg := fmt.Sprintf("{'error':'%s'}", err.Error())
			w.Write([]byte(errMsg))
		} else {
			log.Print("info: OnSubmitResponse = ", string(resultBody))
			w.Header().Set("Content-Type", "application/json")
			if failed {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
			}

			w.Write(resultBody)
		}
	}
}

func VersionRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, fmt.Sprint(version))
}

func AddVersion(router *httprouter.Router) {
	router.GET(versionPath, DebugLogging(VersionRoute, versionPath))
}

func DebugLogging(h httprouter.Handle, path string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.Print("debug: ", path, " request body = ", r.Body)
		h(w, r, p)
		log.Print("debug: ", path, " response=", w)
	}
}

func AddOnSubmit(router *httprouter.Router, onSubmit *channel.OnSubmit) {
	router.POST(onSubmitPrefix, DebugLogging(OnSubmitRoute(onSubmit), onSubmitPrefix))
}

func AddOnGetNextActs(router *httprouter.Router, onGetNextActs *channel.OnGetNextActs) {
	router.POST(onGetNextActsPrefix, DebugLogging(OnGetNextActsRoute(onGetNextActs), onGetNextActsPrefix))
}
