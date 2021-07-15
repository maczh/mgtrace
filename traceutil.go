package utils

import (
	"github.com/gin-gonic/gin"
	"time"
)

func PutRequestId(c *gin.Context) {
	requestId := c.GetHeader("X-Request-Id")
	requestId = If(requestId == "", GetRandomHexString(16), requestId).(string)
	routineId := GetGID()
	OnGetCache("RequestId").Add(routineId, requestId, 5*time.Minute)
}

func GetRequestId() string {
	requestId, found := OnGetCache("RequestId").Value(GetGID())
	if found {
		return requestId.(string)
	} else {
		return ""
	}
}
