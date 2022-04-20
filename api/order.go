package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetOrderRequest struct {
	Customerid int64 `uri:"id" binding:"required"`
}

func (server *Server) GetOrderByid(ctx *gin.Context) {

	var req GetOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	data, err := server.service.OrderByIdBiz(ctx,req.Customerid)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, data)
}
