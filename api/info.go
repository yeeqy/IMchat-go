package api

import (
	"IM-chat/model"

	"github.com/gin-gonic/gin"
)

// {
// 	fromUserId,
// 	toUserId,
// 	pageNum,
// 	pageSize
// }

func SelectMessageByUserIdAndPage(c *gin.Context) {
	var req = model.Info{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, nil)
		return
	}

	infos, err := model.ListInfo(&req)
	if err != nil {
		c.JSON(200, nil)
		return
	}

	c.JSON(200, infos)
}
