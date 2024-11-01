package api

import (
	"encoding/json"
	"fmt"
	"bytes"
	"io"
	"log"
	"strings"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	jwt "github.com/golang-jwt/jwt/v4"
	types "backend/pkg/types"
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

	// User handlers:
	r.HandleFunc("/log-in", makeHTTPHandlerFunc(s.handleLogin))
	r.HandleFunc("/sign-up", makeHTTPHandlerFunc(s.handleSignUp))

	/*r.HandleFunc("/user", makeHTTPHandlerFunc(s.handleUser))
	r.HandleFunc("/user/id/{id}", makeHTTPHandlerFunc(s.handleUserByID))
	r.HandleFunc("/user/username/{username}", makeHTTPHandlerFunc(s.handleUserByUsername))*/
	r.HandleFunc("/delete-account/{username}", withJWTAuth(makeHTTPHandlerFunc(s.handleDeleteAccount)))


	// Message handlers:
	r.HandleFunc("/message/new/{username}", withJWTAuth(makeHTTPHandlerFunc(s.handleCreateMessage)))
	//---->>!!!add auth to every except just get message route
	r.HandleFunc("/message/{id}", makeHTTPHandlerFunc(s.handleMessage))
	r.HandleFunc("/messages", makeHTTPHandlerFunc(s.handleMessages))

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

func createJWT(user *types.User) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt":     15000,
		"username":      user.Username,
	}

	secret := os.Getenv("BACKEND_JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func permissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
}

func withJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("calling JWT auth middleware")//, r.Header.Get("Authorization"))

		tokenString := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
		token, err := validateJWT(tokenString)
		if err != nil {
			permissionDenied(w)
			return
		}
		if !token.Valid {
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

	    type Request struct {
			Username      string    `json:"username"`
	    }
	    if r.Body != nil {
			body, req, err := GetBodyData[Request](r)
			if err != nil {
				permissionDenied(w)
				return
			}

			if req.Username != claims["username"]	{
				permissionDenied(w)
				return
			}

			r.Body = io.NopCloser(bytes.NewBuffer(body))
	    } else if getID(r, "username") != claims["username"] {
			permissionDenied(w)
			return
	    }

		if err != nil {
			WriteJSON(w, http.StatusForbidden, ApiError{Error: "invalid token"})
			return
		}
		
	    handlerFunc(w, r)

	    defer r.Body.Close()
	}
}

func GetBodyData[T any](r *http.Request) ([]byte, *T, error) {
  if r.Body == nil {
    return nil, nil, fmt.Errorf("body missing data")
  }
  req := new(T)
  body, err := io.ReadAll(r.Body)
  if err != nil {
    return nil, nil, err
  }
  
  rdr1 := body
  rdr2 := body
  
  err = json.Unmarshal(rdr1, &req)
  if err != nil {
    return nil, nil, err
  }
  return rdr2, req, nil
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("BACKEND_JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
}

func getID(r *http.Request, name string) string {
	return mux.Vars(r)[name]
}
