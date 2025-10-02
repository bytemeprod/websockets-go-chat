package login

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bytemeprod/websockets-go-chat/internal/redisstore"
	"github.com/bytemeprod/websockets-go-chat/internal/tokens"
	"github.com/labstack/echo"
)

// Errors
var (
	ErrBadRequest    = "invalid request"
	ErrInternalError = "internal server error"
	ErrUserExist     = "user already exist"
)

type Request struct {
	Username string `json:"username"`
}

func NewHandler(ctx context.Context, storage *redisstore.RedisClient, secret_key string) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request Request
		if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
			fmt.Printf("Failed to decode request: %s", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": ErrBadRequest,
			})
		}
		exists, err := storage.UsernameExist(ctx, request.Username)
		if err != nil {
			fmt.Printf("Failed to check if username exist: %s", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": ErrInternalError,
			})
		}
		if exists {
			return c.JSON(http.StatusConflict, map[string]string{
				"error": ErrUserExist,
			})
		}

		token, err := tokens.GenerateJWT([]byte(secret_key), request.Username)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": ErrInternalError,
			})
		}

		storage.AddClient(ctx, request.Username)

		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}
