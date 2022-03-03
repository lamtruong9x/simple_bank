package api

import (
	"database/sql"
	"net/http"
	db "simple_bank/db/sqlc"
	"simple_bank/db/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username       string `json:"username" binding:"required,alphanum"`
	Password 	   string `json:"password" binding:"required,min=6"`
	FullName       string `json:"full_name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(u *db.User) createUserResponse {
	return createUserResponse{
		Username: u.Username,
		FullName: u.FullName,
		Email: u.Email,
		PasswordChangedAt: u.PasswordChangedAt,
		CreatedAt: u.CreatedAt,
	}
} 

func (s *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err!=nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username: req.Username,
		HashedPassword: hashedPassword,
		FullName: req.FullName,
		Email: req.Email,
	}
	user, err := s.store.CreateUser(ctx ,arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			//log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(pqErr))
				return
			}
			
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, newUserResponse(&user))
}

type getUserRequest struct {
	Username  string `uri:"username" binding:"required"`
}

func (s *Server) getUser(ctx *gin.Context) {
	var req getUserRequest 
	
	if err := ctx.ShouldBindUri(&req); err != nil  {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	user, err := s.store.GetUser(ctx, req.Username)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, newUserResponse(&user))
}
