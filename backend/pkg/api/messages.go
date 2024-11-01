package api

import (
	"fmt"
	"net/http"

	types "backend/pkg/types"
)

func (s *APIServer) handleGetMessageById(w http.ResponseWriter, r *http.Request) error {
	id := getID(r, "id")
	
	message, err := s.store.GetMessage(id)
	if err != nil {
		return err
	}
	
	return WriteJSON(w, http.StatusOK, message)
}

func (s *APIServer) handleMessage(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleUpdateMessage(w, r)
	} else if r.Method == "DELETE" {
		return s.handleDeleteMessage(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleUpdateMessage(w http.ResponseWriter, r *http.Request) error {	
	_, req, err := GetBodyData[types.UpdateMessageRequest](r)
	if err != nil {
		return err
	}

	if err = s.store.UpdateMessage(req); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, req)
}

func (s *APIServer) handleDeleteMessage(w http.ResponseWriter, r *http.Request) error {
	_, req, err := GetBodyData[types.DeleteMessageRequest](r)
	if err != nil {
		return err
	}
	
	if err := s.store.DeleteMessage(req.ID); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]string{ "deleted": req.ID })
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
		_, req, err := GetBodyData[types.CreateMessageRequest](r)
		if err != nil {
			return err
		}

	    message, err := types.NewMessage(req.Text, req.Username)
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
