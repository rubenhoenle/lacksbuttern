package main

import (
	"context"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

const (
	AddressCTX = "server_address"
	TypeCTX    = "server_type"
)

type Server struct {
	ctx context.Context
}

func NewServer(ctx context.Context) *Server {
	return &Server{ctx}
}

func (server *Server) Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/lazy", server.LazyHandler)
	r.HandleFunc("/", server.RootHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         server.ctx.Value(AddressCTX).(string),
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
	}
	go func() {
		log.Ctx(server.ctx).Info().Msg("starting server")
		if err := srv.ListenAndServe(); err != nil {
			log.Ctx(server.ctx).Fatal().Err(err).Msg("listen and serve crashed")
		}
	}()
	for {
		select {
		case <-server.ctx.Done():
			log.Ctx(server.ctx).Debug().Msg("starting grace period on shutdown")
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			srv.Shutdown(ctx)
			log.Ctx(server.ctx).Debug().Msg("HTTP server shutdown complete")
			return
		}
	}
}

func (server *Server) RootHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Category: %v\n", vars["category"])
	var tmplFile = "rsc/index.html.tpl"
	tmpl, err := template.New("index.html.tpl").ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	inputs := map[string]string{
		"timestamp": time.Now().String(),
	}
	err = tmpl.Execute(w, inputs)
	if err != nil {
		panic(err)
	}
}

func (server *Server) LazyHandler(w http.ResponseWriter, r *http.Request) {
	println("triggered lazy stuff")
	fmt.Fprint(w, "<p>hello world</p>")
}
