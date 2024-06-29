package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jmavila/golang/web-server-4/models"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := models.RegisterUserPayload{
			FirstName: "user",
			LastName:  "person",
			Email:     "",
			Password:  "123123",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should correctly register the user", func(t *testing.T) {
		payload := models.RegisterUserPayload{
			FirstName: "user",
			LastName:  "person",
			Email:     "my@mail.com",
			Password:  "123123",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*models.User, error) {
	return nil, fmt.Errorf("Not found")
}

func (m *mockUserStore) GetUserById(id int) (*models.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user models.User) error {
	return nil
}
