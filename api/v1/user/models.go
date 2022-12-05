package user

import (
	v1 "github.com/gabriel-henriq/smart-agenda/api/v1"
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/google/uuid"
	"math"
	"time"
)

type createRequest struct {
	Name     string `json:"name" binding:"alpha"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type deleteRequest struct {
	ID int32 `uri:"id" uri:"id" binding:"required,min=1"`
}

type getRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type updateRequest struct {
	ID       int32  `json:"id" binding:"required,numeric"`
	Name     string `json:"ame" binding:"alpha"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
}

type response struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type listResponse struct {
	Users []response `json:"users"`
	v1.PaginationResponse
}

type loginUserResponse struct {
	SessionID             uuid.UUID `json:"session_id"`
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	User                  response  `json:"user"`
}

func toJSON(SQLUser sqlc.User) response {
	return response{
		ID:        SQLUser.ID,
		Name:      SQLUser.Name,
		Email:     SQLUser.Email,
		CreatedAt: SQLUser.CreatedAt.Unix(),
		UpdatedAt: SQLUser.UpdatedAt.Unix(),
	}
}

func toJSONList(SQLUsers []sqlc.ListUsersRow, pageID, pageSize int32) listResponse {
	var users []response

	for _, user := range SQLUsers {
		users = append(users, response{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Unix(),
			UpdatedAt: user.UpdatedAt.Unix(),
		})
	}

	totalPages := int32(math.Ceil(float64(SQLUsers[0].TotalItems) / float64(pageSize)))

	return listResponse{
		Users: users,
		PaginationResponse: v1.PaginationResponse{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLUsers[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
