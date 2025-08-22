package libs

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func GetUid(c *gin.Context) (id int) {
	uid, ok := c.Get("UID")
	if !ok {
		return 0
	}
	id, ok = uid.(int)
	if !ok {
		return 0
	}
	return
}

func GetRole(c *gin.Context) (role int) {
	r, ok := c.Get("ROLE")
	if !ok {
		return 0
	}
	role, ok = r.(int)
	if !ok {
		return 0
	}
	return
}

func GetTimeStamp() int64 {
	return time.Now().Unix()
}

func GetLimitByPage(page, pageSize int) (limit string) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	limit = strconv.Itoa((page-1)*pageSize) + "," + strconv.Itoa(pageSize)
	return
}

func Addslashes(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}
