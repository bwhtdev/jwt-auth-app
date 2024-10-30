package api

import (
	"net/http"
)

func (s *APIServer) handleMessage(w http.ResponseWriter, r *http.Request) error {
	id := getID(r, "id")
	
	message, err := s.store.GetMessage(id)
	if err != nil {
		return err
	}
	
	return WriteJSON(w, http.StatusOK, message)
}

func (s *APIServer) handleMessages(w http.ResponseWriter, r *http.Request) error {
	messages, err := s.store.GetMessages()
	if err != nil {
		return err
	}
	
	return WriteJSON(w, http.StatusOK, messages)
}
