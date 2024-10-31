package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	types "backend/pkg/types"
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

func (s *APIServer) handleCreateMessage(w http.ResponseWriter, r *http.Request) error {	
	if r.Method == "POST" {
		username := getID(r, "username")
		
	    req := new(types.CreateMessageRequest)
	    if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			return err
	    }

	    message, err := types.NewMessage(req.Text, username)
	    if err != nil {
	      return err
	    }
	
		id, err := s.store.CreateMessage(message)
		if err != nil {
	      return err
	    }

	    message.ID = id
	    //message.CreatedAt = Date(..)??

	    return WriteJSON(w, http.StatusOK, message)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}
