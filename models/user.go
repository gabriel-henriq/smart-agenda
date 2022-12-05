package models

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/google/uuid"
	"math"
	"time"
)

type CreateUserRequest struct {
	Name     string `json:"name" binding:"alpha"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type DeleteUserRequest struct {
	ID int32 `uri:"id" uri:"id" binding:"required,min=1"`
}

type GetUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserRequest struct {
	ID       int32  `json:"id" binding:"required,numeric"`
	Name     string `json:"ame" binding:"alpha"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type ListUserResponse struct {
	Users []UserResponse `json:"users"`
	PaginationResponse
}

type LoginUserResponse struct {
	SessionID             uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  UserResponse `json:"user"`
}

func UserToJSON(SQLUser sqlc.User) UserResponse {
	return UserResponse{
		ID:        SQLUser.ID,
		Name:      SQLUser.Name,
		Email:     SQLUser.Email,
		CreatedAt: SQLUser.CreatedAt.Unix(),
		UpdatedAt: SQLUser.UpdatedAt.Unix(),
	}
}

func UsersToJSONList(SQLUsers []sqlc.ListUsersRow, pageID, pageSize int32) ListUserResponse {
	var users []UserResponse

	for _, user := range SQLUsers {
		users = append(users, UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Unix(),
			UpdatedAt: user.UpdatedAt.Unix(),
		})
	}

	totalPages := int32(math.Ceil(float64(SQLUsers[0].TotalItems) / float64(pageSize)))

	return ListUserResponse{
		Users: users,
		PaginationResponse: PaginationResponse{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLUsers[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
