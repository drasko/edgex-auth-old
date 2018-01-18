package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"go.uber.org/zap"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateUser creates new user account based on provided username and password.
// The account is assigned with one master key - a key with all permissions on
// all owned resources regardless of their type. Provided password in encrypted
// using bcrypt algorithm.
func CreateUser(username, password string) (User, error) {
	u, err := uuid.NewV4()
	if err != nil {
		logger.Error("Failed to generate uuid", zap.Error(err))
		return User{}, err
	}

	user := User{
		Id:       u.String(),
		Username: username,
	}

	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return user, &AuthError{Code: http.StatusInternalServerError}
	}

	user.Password = string(p)

	return user, nil
}

// CheckPassword tries to determine whether or not the submitted password
// matches the one stored (and hashed) during registration. An error will be
// used to indicate an invalid password.
func CheckPassword(plain, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &User{}
	if err = json.Unmarshal(body, user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &User{}
	if err = json.Unmarshal(body, user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: query mongoDB

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func authorize(w http.ResponseWriter, r *http.Request) {
	// TODO
}
