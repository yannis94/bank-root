package auth

import (
	"errors"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/yannis94/bank-root/internal/config"
	"github.com/yannis94/bank-root/internal/repository"
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

func (auth *AuthService) CreateJWT(role, jti string) (string, error) {
    exp := time.Now().Add(token_exp * time.Minute)
    return auth.createToken(exp, role, jti)
}

func (auth *AuthService) CreateRefreshToken(role string) (string, error) {
    exp := time.Now().Add(refresh_token_exp * time.Minute)
    jti := uuid.New().String()
    return auth.createToken(exp, role, jti)
}

func (auth *AuthService) IsTokenValid(tkn string, role string) (bool, error) {
    token, _ := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("Unexpected signing method.")
        }

        return []byte(auth.secret), nil
    })

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        if claims["iss"] != auth.iss && claims["aud"] != role {
            return false, errors.New("Insufficient level of access.")
        }
        return true, nil

    } else {
        return false, errors.New("Token is invalid.")
    }
}

func (auth *AuthService) RefreshToken(tkn string) (string, error) {
    token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("Unexpected signing method.")
        }

        return []byte(auth.secret), nil
    })

    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        if claims["iss"] != auth.iss {
            return "", errors.New("Token not emitted by this server.")
        }

        db, err := repository.NewPostgres(config.DB_USER, config.DB_PASS, config.DB_PORT, config.DB_NAME)

        if err != nil {
            return "", errors.New("Unable to connect to databse.")
        }

        jti, ok := claims["jti"].(string)
        role, ok := claims["aud"].(string)

        if !ok {
            return "", errors.New("Invalid token claims.")
        }

        session, err := db.GetSessionFromTokenId(jti)

        if err != nil || session.TokenId == ""{
            return "", errors.New("Session not found.")
        }

        if valid, err := auth.IsTokenValid(session.Refresh_token, role); err != nil && !valid {
            return "", errors.New("Refresh token expired.")
        }

        exp := time.Now().Add(token_exp * time.Minute)
        newTkn, err := auth.createToken(exp, role, session.TokenId)

        if err != nil {
            return "", errors.New("Unable to create new token.")
        }

        return newTkn , nil

    } else {
        return "", errors.New("Refresh token is invalid.")
    }
}

func (auth *AuthService) createToken(exp time.Time, role, jti string) (string, error) {
    newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{
            "iss": auth.iss,
            "aud": role,
            "exp": jwt.NewNumericDate(exp),
            "iat": jwt.NewNumericDate(time.Now()),
            "jti": jti,
        })

    return newToken.SignedString([]byte(auth.secret))
}
