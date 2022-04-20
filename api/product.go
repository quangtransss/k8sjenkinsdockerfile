package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type getProductRequest struct {
	ID int64 `uri:"id" binding:"required,min=0"`
}

func (server *Server) GetProductById(ctx *gin.Context)  {

	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product ,err := server.service.GetProductById(ctx,req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK,product)
}