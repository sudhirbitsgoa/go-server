package mylib

import (
	"fmt"
	"net/http"

	wsroutes "go-server/wsroutes"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// fs := http.StripPrefix("/", http.FileServer(http.Dir("../client/")))
	// http.Handle("/", fs)
	// http.Handle("/ws", websocket.Handler(sy.EchoWS))
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	router.PathPrefix("/ws").Handler(websocket.Handler(wsroutes.RoutesHandler))
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("client/"))))

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	// http.StripPrefix("/", http.FileServer(http.Dir("client/")))
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/v2/",
		Index,
	},

	Route{
		"DeleteBackTest",
		"DELETE",
		"/api/v2/backtest/{id}",
		DeleteBackTest,
	},

	Route{
		"GetBackTest",
		"GET",
		"/api/v2/backtest/{id}",
		GetBackTest,
	},

	Route{
		"GetBackTests",
		"GET",
		"/api/v2/strategy/{id}/backtests",
		GetBackTests,
	},

	Route{
		"GetCommunityBackTest",
		"GET",
		"/api/v2/backtest/share",
		GetCommunityBackTest,
	},

	Route{
		"ShareBackTest",
		"POST",
		"/api/v2/backtest/share/{id}",
		ShareBackTest,
	},

	Route{
		"CreateStrategy",
		"POST",
		"/api/v2/strategy",
		CreateStrategy,
	},

	Route{
		"DeleteStrategy",
		"DELETE",
		"/api/v2/strategy/{id}",
		DeleteStrategy,
	},

	Route{
		"ExecStrategy",
		"POST",
		"/api/v2/strategy/{id}/exec",
		ExecStrategy,
	},

	Route{
		"GetStrategy",
		"GET",
		"/api/v2/strategy/{id}",
		GetStrategy,
	},

	Route{
		"GetStrategys",
		"GET",
		"/api/v2/strategy",
		GetStrategys,
	},

	Route{
		"UpdateStrategy",
		"PUT",
		"/api/v2/strategy/{id}",
		UpdateStrategy,
	},

	Route{
		"UploadStrategyCode",
		"POST",
		"/api/v2/strategy/code/{id}",
		UploadStrategyCode,
	},

	Route{
		"CreateThread",
		"POST",
		"/api/v2/thread",
		CreateThread,
	},

	Route{
		"FollowThread",
		"POST",
		"/api/v2/thread/{threadId}/follow",
		FollowThread,
	},

	Route{
		"GetThread",
		"GET",
		"/api/v2/thread/{threadId}",
		GetThread,
	},

	Route{
		"GetThreads",
		"GET",
		"/api/v2/thread",
		GetThreads,
	},

	Route{
		"LikeThread",
		"POST",
		"/api/v2/thread/{threadId}/like",
		LikeThread,
	},

	Route{
		"ReplyToThread",
		"POST",
		"/api/v2/thread/{threadId}",
		ReplyToThread,
	},

	Route{
		"ViewThread",
		"POST",
		"/api/v2/thread/{threadId}/view",
		ViewThread,
	},

	Route{
		"ActivateUser",
		"GET",
		"/api/v2/user/activate",
		ActivateUser,
	},

	Route{
		"ForgotPassword",
		"GET",
		"/api/v2/user/forgotpassword",
		ForgotPassword,
	},

	Route{
		"GetProfile",
		"GET",
		"/api/v2/me",
		GetProfile,
	},

	Route{
		"Regiteruser",
		"POST",
		"/api/v2/user",
		Regiteruser,
	},

	Route{
		"ResetPassword",
		"POST",
		"/api/v2/user/setpassword",
		ResetPassword,
	},

	Route{
		"Userlogin",
		"POST",
		"/api/v2/user/login",
		Userlogin,
	},
}
