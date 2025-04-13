package auth

import (
	"fmt"
	"movie-catalog/config"
	"movie-catalog/internal/core/domain/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT interface {
	GenerateToken(data *entity.JwtData) (string, int64, error)
	VerifyAccessToken(token string) (*entity.JwtData, error)
}

type Options struct {
	SigningKey string
	issuer     string
}

func (o *Options) GenerateToken(data *entity.JwtData) (string, int64, error) {
	now := time.Now().Local()

	expireAt := now.Add(time.Hour * 24)

	data.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expireAt)
	data.RegisteredClaims.Issuer = o.issuer
	data.RegisteredClaims.NotBefore = jwt.NewNumericDate(now)

	acToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	accessToken, err := acToken.SignedString([]byte(o.SigningKey))

	if err != nil {
		return "", 0, err
	}

	return accessToken, expireAt.Unix(), nil
}

func (o *Options) VerifyAccessToken(token string) (*entity.JwtData, error) {
	parseToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(o.SigningKey), nil
	})
	if err != nil {
		return nil, err
	}

	if parseToken.Valid {
		claim, ok := parseToken.Claims.(jwt.MapClaims)
		if !ok || !parseToken.Valid {
			return nil, err
		}
		jwtData := &entity.JwtData{
			UserID: claim["UserID"].(float64),
		}

		return jwtData, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func NewJwt(cfg *config.Config) JWT {
	opt := new(Options)
	opt.SigningKey = cfg.App.JwtSecretKey
	opt.issuer = cfg.App.JwtIssuer
	return opt
}
