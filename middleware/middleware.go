package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/anime454/go_hexagonal_architecture/logs"
	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
	code int
}

type responseBody struct {
	Code    int
	Message string
	Data    interface{}
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// func (w bodyLogWriter) WriteString(s string) (int, error) {
// 	fmt.Println("on writeString")
// 	w.body.WriteString(s)
// 	return w.ResponseWriter.WriteString(s)
// }

//thank you alot https://github.com/gin-gonic/gin/issues/961
func InitAccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		b, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(b))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(b)) //We have to create a new Buffer, because rdr1 will be read.

		buf := new(bytes.Buffer)
		buf.ReadFrom(rdr1)
		RequestBodyStr := buf.String()

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Request.Body = rdr2
		c.Next()

		r := responseBody{}
		err := json.Unmarshal(blw.body.Bytes(), &r)
		if err != nil {
			logs.Error(err)
		}
		blw.code = r.Code

		responseTime := time.Since(now)
		resBody := logs.AccessLog{
			Ip:            c.Request.RemoteAddr,
			RequestId:     c.GetHeader("requestId"),
			Method:        c.Request.Method,
			Url:           c.Request.RequestURI,
			RequestBody:   RequestBodyStr,
			ResponseBody:  blw.body.String(),
			ResponseCode:  blw.code,
			ResponserTime: responseTime.String(),
		}
		logs.RequestLog(resBody)
	}
}
