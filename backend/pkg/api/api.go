package api

import (
  "log"
  "net/http"
  "encoding/json"

  "github.com/rs/cors"
  "github.com/gorilla/mux"
  storage "backend/pkg/storage"
)

type APIServer struct {
	listenAddr  string
	store       storage.Storage
}

func NewAPIServer(listenAddr string, store storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	r := router.PathPrefix("/v1/").Subrouter()

	// Handlers::
	r.HandleFunc("/ping", makeHTTPHandlerFunc(s.handlePing))
	r.HandleFunc("/people", makeHTTPHandlerFunc(s.handlePeople))
	
	c := cors.New(cors.Options{
		//AllowedOrigins: []string{webPort},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	log.Print("Server running on port: ", s.listenAddr)
	log.Fatal(http.ListenAndServe(s.listenAddr, handler))
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Context-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func (s *APIServer) handlePing(w http.ResponseWriter, r *http.Request) error {
	type PingData struct {
		Health bool `json:"health"`
	}
	return WriteJSON(w, http.StatusOK, PingData{ Health: true })
}
