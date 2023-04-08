package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tentangKode/users"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var APPLICATION_NAME = "yaza barudak"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNATURE_KEY = []byte("opewfjdi3f84f339fu3")

type MyClaims struct {
	Username    string `json:"username"`
	NamaLengkap string `json:"nama_lengkap"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {

	var userInput users.User
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		response, _ := json.Marshal(map[string]string{"message": "failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}
	defer r.Body.Close()

	fmt.Println(userInput.Username, userInput.Password)
	user, err := checkUsernameAndPassword(userInput.Username, userInput.Password)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": "invalid username or password"})
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(response)
		return
	}

	claims := &MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username:    user.Username,
		NamaLengkap: user.NamaLengkap,
	}

	// mendeklarasikan algoritma yang akan digunakan untuk signing
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	response := map[string]string{"token": signedToken}
	json.NewEncoder(w).Encode(response)

}

func checkUsernameAndPassword(username, password string) (*users.User, error) {
	var user users.User

	db, err := users.ConnectDBUsers()
	if err != nil {
		panic(err)
	}

	err = db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func Register(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan dari json
	var userInput users.User

	db, err := users.ConnectDBUsers()
	if err != nil {
		panic(err)
	}

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		response, _ := json.Marshal(map[string]string{"message": "failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}
	defer r.Body.Close()

	// hass pass menggunakan bcrypt
	hassPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hassPassword)

	// insert ke database
	if err := db.Create(&userInput).Error; err != nil {
		response, _ := json.Marshal(map[string]string{"message": "failed to save data"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	response, _ := json.Marshal(map[string]string{"message": "Succes"})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
