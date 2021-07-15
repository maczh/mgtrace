package mgtrace

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/maczh/mgcache"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

func PutRequestId(c *gin.Context) {
	requestId := c.GetHeader("X-Request-Id")
	if requestId == "" {
		requestId = getRandomHexString(16)
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

func generateRandString(source string, l int) string {
	bytes := []byte(source)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func getRandomHexString(l int) string {
	str := "0123456789abcdef"
	return generateRandString(str, l)
}
