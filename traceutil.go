package mgtrace

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/maczh/mgcache"
	"github.com/maczh/utils"
	"runtime"
	"strconv"
	"time"
)

func PutRequestId(c *gin.Context) {
	requestId := c.GetHeader("X-Request-Id")
	if requestId == "" {
		requestId = utils.GetRandomHexString(16)
	}
	routineId := GetGID()
	mgcache.OnGetCache("RequestId").Add(routineId, requestId, 5*time.Minute)
}

func GetRequestId() string {
	requestId, found := mgcache.OnGetCache("RequestId").Value(GetGID())
	if found {
		return requestId.(string)
	} else {
		return ""
	}
}

//获取当前协程Id
func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
