package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type getRoleRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetRoleById(ctx *gin.Context) {

	var req getRoleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	role, err := server.service.GetRoleByIdBiz(ctx,req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, role)
}
