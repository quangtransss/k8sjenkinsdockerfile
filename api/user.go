package api

import (
	"golang/internal/model/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	RoleId	 int64 `json:"roleid" binding:"required"`
	PhoneNumber	int64 `json:"mobile" binding:"required"`

}
type createUserResponse  struct {
	Username string `json:"username"`
	RoleId	 int64 `json:"roleid"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Mobile	 int64	`json:"mobile"`
}
func (server *Server)  CreateUser(ctx *gin.Context)  {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req) ; err != nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

    arg := domain.CreateUserParams{
		Username: req.Username,
		FullName: req.FullName,
		HashedPassword: req.Password,
		Email: req.Email,
		Mobile: req.PhoneNumber,
		Roleid: req.RoleId,
    }
	user ,err := server.service.CreateUser(ctx,arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	userReponse := createUserResponse{
		Username: user.Username,
		RoleId: user.Roleid,
		FullName: user.FullName,
		Email: user.FullName,
		Mobile: user.Mobile,
	}
	ctx.JSON(http.StatusOK, userReponse)
}
type listUserRequest struct {
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}
func (server *Server) ListUser(ctx *gin.Context)  {
	var req listUserRequest

	arg := domain.ListUsersParams{
		Limit: req.Limit,
		Offset: req.Offset,
	}
	users , err := server.service.ListUsers(ctx,arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK,users)
}
