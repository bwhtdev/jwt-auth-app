package api

import (
	"net/http"
)

func (s *APIServer) handlePeople(w http.ResponseWriter, r *http.Request) error {
	people, err := s.store.GetPeople()
	if err != nil {
		return err
	}
	
	return WriteJSON(w, http.StatusOK, people)
}
