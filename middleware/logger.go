package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	// [REQUEST] [GET /all] request-id:abc123 request-header:%s request-body:%s
	requestFormat = "[REQUEST] [%s %s] request-id:%s request-header:%s request-body:%s\n"

	// [RESPONSE] [GET /all 200] request-id:abc123 response-header:%s response-body:%s response-time:%sms
	responseFormat = "[RESPONSE] [%s %s %d] request-id:%s response-header:%s response-body:%s response-time:%dms\n"
)

type LogWriter struct {
}

func (writer LogWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format(time.RFC3339) + " " + string(bytes))
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// A custom response writer to also store the the written response body during writing
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		// Pre-Request Logging
		startTime := time.Now()
		requestId, _ := uuid.NewRandom()
		method := c.Request.Method
		path := c.Request.URL.Path
		requestHeaderJson, _ := json.Marshal(c.Request.Header)
		requestHeader := string(requestHeaderJson)
		requestBody, _ := io.ReadAll(c.Request.Body) // Consume request body for log

		if json.Valid(requestBody) {
			compactBody := &bytes.Buffer{}
			json.Compact(compactBody, requestBody) // Convert multi-line JSON to one-liner for tidy logs
			log.Printf(requestFormat, method, path, requestId, requestHeader, compactBody)
		} else {
			log.Printf(requestFormat, method, path, requestId, requestHeader, requestBody)
		}

		c.Request.Body = io.NopCloser(bytes.NewReader(requestBody)) // Return request body for controllers

		c.Next()

		// Post-Request Logging
		responseTime := time.Since(startTime).Milliseconds()
		status := c.Writer.Status()
		responseHeaderJson, _ := json.Marshal(c.Writer.Header())
		responseHeader := string(responseHeaderJson)
		responseBody := w.body.Bytes()

		if json.Valid(responseBody) {
			compactBody := &bytes.Buffer{}
			json.Compact(compactBody, responseBody) // Convert multi-line JSON to one-liner for tidy logs
			log.Printf(responseFormat, method, path, status, requestId, responseHeader, compactBody, responseTime)
		} else {
			log.Printf(responseFormat, method, path, status, requestId, responseHeader, responseBody, responseTime)
		}
	}
}
