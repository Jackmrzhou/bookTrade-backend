package middlewares

import (
	error2 "bookTrade-backend/app/error"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type CommonResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type bodyCacheWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyCacheWriter) Write(b []byte) (int, error) {
	if strings.Contains(w.ResponseWriter.Header().Get("Content-Type"), "application/json") {
		w.body.Write(b)
		return len(b), nil
	} else {
		return w.ResponseWriter.Write(b)
	}
}

func WrapResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		bcw := &bodyCacheWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bcw
		c.Next()
		// only trigger on json response or error cases
		logrus.Info(c.Writer.Header().Get("Content-Type"))
		logrus.Info(strings.Contains(c.Writer.Header().Get("Content-Type"), "application/json"))
		if strings.Contains(c.Writer.Header().Get("Content-Type"), "application/json") || c.GetInt(error2.CodeKey) != 0 {
			// actually no difference whether using 200 or 500 in frontend backend split project
			// using 200 for making debugging using postman easier
			c.Writer.WriteHeader(http.StatusOK)
			resp := CommonResp{}
			code := c.GetInt(error2.CodeKey)
			resp.Code = code
			resp.Msg = error2.Translate(code)
			if code == error2.OK {
				c.Writer.WriteHeader(http.StatusOK)
				resp.Data = json.RawMessage(bcw.body.String())
			}
			data, _ := json.Marshal(resp)
			bcw.ResponseWriter.Write(data)
		}
	}
}
