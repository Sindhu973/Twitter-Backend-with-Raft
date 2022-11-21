package auth

import (
	"errors"
	"net/http"
	"time"

	"garuda.com/m/web/cmd/auth/storage"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Register(username, password string) error {
	// TODO: check if password is hashed successfully
	hash_password, _ := HashPassword(password)
	return storage.OurStore.AddUser(username, hash_password)

}

func Login(username, password string) (string, time.Time, error) {
	user, ok := storage.OurStore.GetUser(username)
	if ok != nil {
		return "", time.Time{}, errors.New("User does not exist")
	}
	if CheckPasswordHash(password, user.Hash_password) {
		expirationTime := time.Now().Add(5 * time.Minute)
		claims := &Claims{
			Username: username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// TODO: take the value from environment variable instead of using config file
		tokenString, err := token.SignedString([]byte(AuthConfig.SecretKey))
		if err != nil {
			return "", time.Time{}, err
		}
		return tokenString, expirationTime, nil
	}
	return "", time.Time{}, errors.New("Wrong password")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request, claims *Claims)) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tokenObject, err := request.Cookie("token")
		claims := &Claims{}
		if err == nil {
			tkn, err := jwt.ParseWithClaims(tokenObject.Value, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(AuthConfig.SecretKey), nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					writer.WriteHeader(http.StatusUnauthorized)
					return
				} else {
					writer.Write([]byte(err.Error()))
				}
			}
			if !tkn.Valid {
				writer.WriteHeader(http.StatusUnauthorized)
				return
			}
			// #============================================================================================
			// refreshing the token here
			// #============================================================================================
			expirationTime := time.Now().Add(5 * time.Minute)
			claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString([]byte(AuthConfig.SecretKey))
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			http.SetCookie(writer, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})
			// #================================================================================================
			endpointHandler(writer, request, claims)
			return
		} else {
			if err == http.ErrNoCookie {
				writer.WriteHeader(http.StatusUnauthorized)
				return
			}
			writer.WriteHeader(http.StatusBadRequest)
			return

		}
	})
}

func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		return "", nil
	})
	if err != nil {
		return false, err
	}
	if token.Valid {
		return true, nil
	}
	return false, errors.New("Invalid token")
}
