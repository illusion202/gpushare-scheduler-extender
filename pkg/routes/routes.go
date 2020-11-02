package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	versionPath       = "/version"
	apiPrefix         = "/gpushare-scheduler"
	bindPrefix        = apiPrefix + "/bind"
	predicatesPrefix  = apiPrefix + "/filter"
	inspectPrefix     = apiPrefix + "/inspect/:nodename"
	inspectListPrefix = apiPrefix + "/inspect"
)

var (
	version = "0.1.0"
	// mu      sync.RWMutex
)

func checkBody(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
}

func OnSubmitRoute()  {
	
}

//func InspectRoute(inspect *scheduler.Inspect) httprouter.Handle {
//	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//		result := inspect.Handler(ps.ByName("nodename"))
//
//		if resultBody, err := json.Marshal(result); err != nil {
//			// panic(err)
//			log.Printf("warn: Failed due to %v", err)
//			w.Header().Set("Content-Type", "application/json")
//			w.WriteHeader(http.StatusInternalServerError)
//			errMsg := fmt.Sprintf("{'error':'%s'}", err.Error())
//			w.Write([]byte(errMsg))
//		} else {
//			w.Header().Set("Content-Type", "application/json")
//			w.WriteHeader(http.StatusOK)
//			w.Write(resultBody)
//		}
//	}
//}

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

//func AddInspect(router *httprouter.Router, inspect *scheduler.Inspect) {
//	router.GET(inspectPrefix, DebugLogging(InspectRoute(inspect), inspectPrefix))
//	router.GET(inspectListPrefix, DebugLogging(InspectRoute(inspect), inspectListPrefix))
//}
