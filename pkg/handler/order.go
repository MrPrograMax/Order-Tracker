package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"net/http"
)

func (h *Handler) CreateOrder(m *stan.Msg) {
	h.services.CreateOrder(m)
}

func (h *Handler) getOrderById(c *gin.Context) {
	id := c.Param("id")

	order, err := h.services.GetOrderById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}
