package apiserver

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/JohnNON/tasker/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *server) response(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.response(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) handleStart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			s.response(w, r, http.StatusBadRequest, nil)
			return
		}

		defer r.Body.Close()
		data, err := io.ReadAll(r.Body)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		req := &model.Request{}
		err = json.Unmarshal(data, req)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if err := req.Validate(); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		task := &model.Task{
			ID:        primitive.NewObjectID(),
			Type:      req.Type,
			State:     0,
			StartedAt: time.Now().UTC(),
		}

		err = s.store.CreateTask(task)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.response(w, r, http.StatusCreated, nil)
	}
}

func (s *server) handleFinish() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			s.response(w, r, http.StatusBadRequest, nil)
			return
		}

		defer r.Body.Close()
		data, err := io.ReadAll(r.Body)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		req := &model.Request{}
		err = json.Unmarshal(data, req)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if err := req.Validate(); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		err = s.store.CompleteTask(req.Type)
		if err != nil {
			status := http.StatusInternalServerError
			if err == mongo.ErrNoDocuments {
				status = http.StatusNotFound
			}

			s.error(w, r, status, err)
			return
		}

		s.response(w, r, http.StatusOK, nil)
	}
}
