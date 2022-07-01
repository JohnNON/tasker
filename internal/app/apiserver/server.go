package apiserver

import (
	"fmt"
	"net/http"

	"ex.ex/ex/internal/config"
	"ex.ex/ex/internal/store"
)

type server struct {
	endPoint string
	store    store.Store
}

func newServer(config *config.Config, st store.Store) *server {
	s := &server{
		endPoint: config.EndPoint,
		store:    st,
	}

	return s
}

func (s *server) configureRouter() {
	http.HandleFunc(fmt.Sprintf("%s/start", s.endPoint), s.handleStart())
	http.HandleFunc(fmt.Sprintf("%s/finish", s.endPoint), s.handleFinish())
}
