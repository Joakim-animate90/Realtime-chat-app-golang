package v1

import (
	"net/http"
	"realtime-chat-app/internal/service"
	"realtime-chat-app/pkg/common/request"
	"realtime-chat-app/pkg/common/response"
	"realtime-chat-app/pkg/global/log"

	"github.com/gin-gonic/gin"
)

// Get the message list
func GetMessage(c *gin.Context) {
	log.Logger.Info(c.Query("uuid"))
	var messageRequest request.MessageRequest
	err := c.BindQuery(&messageRequest)
	if nil != err {
		log.Logger.Error("bindQueryError", log.Any("bindQueryError", err))
	}
	log.Logger.Info("messageRequest params: ", log.Any("messageRequest", messageRequest))

	messages, err := service.MessageService.GetMessages(messageRequest)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessMsg(messages))
}
