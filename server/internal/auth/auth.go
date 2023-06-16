package auth

import (
	"errors"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService struct {
    iss string
    secret string
}

const token_exp = 3
const refresh_token_exp = 60

func NewAuthService(secret string) *AuthService {
    return &AuthService{ iss: "bank-root-server", secret: secret }
}

func (auth *AuthService) CreateJWT(role string) (string, error) {
    exp := time.Now().Add(token_exp * time.Minute).String()
    jti := uuid.New().String()
    return auth.createToken(exp, role, jti)
}

func (auth *AuthService) CreateRefreshToken(role string) (string, error) {
    exp := time.Now().Add(refresh_token_exp * time.Minute).String()
    jti := uuid.New().String()
    return auth.createToken(exp, role, jti)
}

func (auth *AuthService) IsTokenValid(tkn string, role string) (bool, error) {
    token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("Unexpected signing method.")
        }

        return []byte(auth.secret), nil
    })

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        if claims["iss"] != auth.iss && claims["aud"] != role {
            log.Println(err)
            return false, errors.New("Insufficient level of access.")
        }
        return true, nil

    } else {
        log.Println(err)
        return false, errors.New("Token is invalid.")
    }
}

func (auth *AuthService) RefreshToken(refresh_tkn, tkn_id string) (string, error) {
    /*
    token, err := jwt.Parse(refresh_tkn, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
            return nil, errors.New("Unexpected signing method.")
        }

        return []byte(auth.secret), nil
    })

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        if claims["iss"] != auth.iss {
            return "", errors.New("Token not emitted by this server.")
        } else if claims["jti"] != tkn_id {
            return "", errors.New("Token id invalid.")
        }
        role, _ := claims["aud"]
        exp := time.Now().Add(token_exp * time.Minute)
        newTkn := auth.createToken(exp, role, tkn_id)
        return "", nil

    } else {
        log.Println(err)
        return "", errors.New("Refresh token is invalid.")
    }
    */
    return "", nil
}

func (auth *AuthService) createToken(exp, role, jti string) (string, error) {
    log.Println("Secret", auth.secret)
    newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{
            "iss": auth.iss,
            "aud": role,
            "exp": exp,
            "jti": jti,
        })

    return newToken.SignedString([]byte(auth.secret))
}
