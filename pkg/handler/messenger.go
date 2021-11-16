package handler

import (
	"chat"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) getAllDialogs(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	dialogs := h.services.IMessenger.GetAllDialogs(userId)

	c.JSON(http.StatusOK, map[string]interface{}{
		"dialogs": dialogs,
	})
}

func (h *Handler) getDialog(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	messages := h.services.IMessenger.GetDialog(userId, id)

	c.JSON(http.StatusOK, map[string]interface{}{
		"messages": messages,
	})
}

func (h *Handler) getUserStatus(c *gin.Context) { //long poll
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	statusOnline := h.services.IMessenger.GetUserStatus(userId)

	c.JSON(http.StatusOK, map[string]interface{}{
		"status_online": statusOnline,
	})
}

func (h *Handler) sendMessage(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	/*id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}*/

	var input chat.Message
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	input.SenderId = userId
	//input.ReceiverId = id
	input.TimeSent = time.Now().Unix()

	err = h.services.IMessenger.Send(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid list id param")
		return
	}
}

func (h *Handler) getMessage(c *gin.Context) { //long poll
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	m, err := h.services.IMessenger.GetMessage(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, m)
}
