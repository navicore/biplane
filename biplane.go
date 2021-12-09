package biplane

import (
	"fmt"
	"log"
	"net/http"

	"biplane.build/server"
	"github.com/gorilla/mux"
)

// Start server with given server config
// This should also init the database client
// and setup any hooks.
func TakeOff(conf server.Config) {
	r := mux.NewRouter()

	for _, s := range conf.Routers {
		if s, ok := s.(server.Configurer); ok {
			s.Configure(conf)
		}

		s.Routes(r)
	}

	p := conf.Port
	if p == 0 {
		p = 8080
	}

	u := fmt.Sprintf("%s:%d", conf.Host, p)
	log.Printf("Server online and listening at %s", u)
	http.ListenAndServe(u, r)
}
