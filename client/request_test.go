/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-16 21:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 21:00:00
 * @FilePath: \go-namesilo\client\request_test.go
 * @Description: 测试请求和日志打印
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package client

import (
	"context"
	"fmt"
	"strings"
	"testing"

	logger "github.com/kamalyes/go-logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestLogger 测试用的日志记录器，嵌入 *logger.EmptyLogger 自动实现 ILogger 全部方法
// TestLogger 记录各级别日志，供断言使用
//
// 覆写的必须是 ...ContextKV 这一组：客户端打日志走的是 InfoContextKV / DebugContextKV，
// 而不是 InfoContext / DebugContext。先前这里覆写的是不带 KV 的那组，于是客户端的调用
// 直接落到内嵌的 EmptyLogger 上被丢弃，四个集合永远是空的，断言全部失败
// 这个失败被旧 CI 的 continue-on-error 盖了很久
//
// [EN] The overrides must be the ...ContextKV methods, because that is what the client calls.
// They previously overrode the variants without KV, so every client call fell through to the
// embedded EmptyLogger and was discarded, leaving all four slices empty
type TestLogger struct {
	*logger.EmptyLogger
	DebugLogs []string
	InfoLogs  []string
	WarnLogs  []string
	ErrorLogs []string
}

func (l *TestLogger) DebugContextKV(ctx context.Context, msg string, keysAndValues ...interface{}) {
	logStr := fmt.Sprintf("%s", msg)
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			logStr += fmt.Sprintf(" | %v=%v", keysAndValues[i], keysAndValues[i+1])
		}
	}
	l.DebugLogs = append(l.DebugLogs, logStr)
	fmt.Println("[DEBUG]", logStr)
}

func (l *TestLogger) InfoContextKV(ctx context.Context, msg string, keysAndValues ...interface{}) {
	logStr := fmt.Sprintf("%s", msg)
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			logStr += fmt.Sprintf(" | %v=%v", keysAndValues[i], keysAndValues[i+1])
		}
	}
	l.InfoLogs = append(l.InfoLogs, logStr)
	fmt.Println("[INFO]", logStr)
}

func (l *TestLogger) WarnContextKV(ctx context.Context, msg string, keysAndValues ...interface{}) {
	logStr := fmt.Sprintf("%s", msg)
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			logStr += fmt.Sprintf(" | %v=%v", keysAndValues[i], keysAndValues[i+1])
		}
	}
	l.WarnLogs = append(l.WarnLogs, logStr)
	fmt.Println("[WARN]", logStr)
}

func (l *TestLogger) ErrorContextKV(ctx context.Context, msg string, keysAndValues ...interface{}) {
	logStr := fmt.Sprintf("%s", msg)
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			logStr += fmt.Sprintf(" | %v=%v", keysAndValues[i], keysAndValues[i+1])
		}
	}
	l.ErrorLogs = append(l.ErrorLogs, logStr)
	fmt.Println("[ERROR]", logStr)
}

// TestDoRequestWithLogger 测试带日志的请求
func TestDoRequestWithLogger(t *testing.T) {
	// 创建测试日志记录器
	testLogger := &TestLogger{
		EmptyLogger: logger.NewEmptyLogger(),
		DebugLogs:   make([]string, 0),
		InfoLogs:    make([]string, 0),
		WarnLogs:    make([]string, 0),
		ErrorLogs:   make([]string, 0),
	}

	// 创建客户端
	c, err := New("test-api-key",
		WithDebug(true),
		WithLogger(testLogger),
	)
	assert.NoError(t, err, "创建客户端应该成功")
	assert.NotNil(t, c, "客户端不应该为 nil")

	// 执行请求（这会失败，但我们主要是测试日志）
	ctx := context.Background()
	params := map[string]string{
		"domain": "example.com",
	}

	_, _ = c.DoRequest(ctx, "getDomainInfo", params)

	// 验证日志
	//
	// logRequestResponse 有四条互斥路径：panic / timeout / error / 成功，
	// 四条都记 operation 与 url，并且消息里都含 "NameSilo API"
	// 这里用假 API key 打真实接口，走哪条取决于网络，所以断言四条路径的共同保证，
	// 而不是赌一定走成功路径
	//
	// [EN] logRequestResponse has four mutually exclusive paths and all of them record the
	// operation and the url, so the assertions target what every path guarantees rather than
	// assuming the request succeeded
	t.Run("验证请求日志", func(t *testing.T) {
		recorded := make([]string, 0, 8)
		recorded = append(recorded, testLogger.InfoLogs...)
		recorded = append(recorded, testLogger.WarnLogs...)
		recorded = append(recorded, testLogger.ErrorLogs...)
		require.NotEmpty(t, recorded, "请求结束后应该至少记下一条日志")

		requestLog := ""
		for _, log := range recorded {
			if strings.Contains(log, "NameSilo API") {
				requestLog = log
				break
			}
		}
		require.NotEmpty(t, requestLog, "应该找到一条 NameSilo API 日志")

		assert.Contains(t, requestLog, "operation=getDomainInfo", "日志应该带上操作名")
		assert.Contains(t, requestLog, "url=", "日志应该带上请求 URL")

		t.Logf("请求日志: %s", requestLog)
	})

	t.Run("验证响应日志", func(t *testing.T) {
		// 响应体日志只在拿到非空 body 时才打，所以这条是条件断言：
		// 没有就跳过，有就校验它确实带了 operation 与 body
		// [EN] The response-body log is only emitted for a non-empty body, so this assertion is
		// conditional: skip when absent, and when present check it carries the operation and body
		responseLog := ""
		for _, log := range testLogger.DebugLogs {
			if strings.Contains(log, "NameSilo API response body") {
				responseLog = log
				break
			}
		}
		if responseLog == "" {
			t.Skip("本次请求没有响应体，跳过响应日志校验")
		}

		assert.Contains(t, responseLog, "operation=getDomainInfo", "响应日志应该带上操作名")
		assert.Contains(t, responseLog, "body=", "响应日志应该带上响应体")

		t.Logf("响应日志: %s", responseLog)
	})
}
