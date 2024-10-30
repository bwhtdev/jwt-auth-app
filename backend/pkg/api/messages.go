package api

import (
	"net/http"
)

func (s *APIServer) handleMessages(w http.ResponseWriter, r *http.Request) error {
	message, err := s.store.GetMessages()
	if err != nil {
		return err
	}
	
	return WriteJSON(w, http.StatusOK, message)
}
