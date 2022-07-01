package apiserver

import (
	"fmt"
	"net/http"

	"github.com/JohnNON/tasker/internal/config"
	"github.com/JohnNON/tasker/internal/store"
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
