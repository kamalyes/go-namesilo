/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:15:17
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 21:37:18
 * @FilePath: \go-namesilo\client\request.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package client

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"
	"strings"

	namesilo "github.com/kamalyes/go-namesilo"
	"github.com/kamalyes/go-toolbox/pkg/errorx"
	"github.com/kamalyes/go-toolbox/pkg/osx"
)

// DoRequest 执行 API 请求（公开给子模块使用）
func (c *Client) DoRequest(ctx context.Context, operation string, params map[string]string) ([]byte, error) {
	// 生成 traceId
	traceId := osx.HashUnixMicroCipherText()

	// 构建请求参数
	queryParams := url.Values{}

	// 添加业务参数（先添加，以便系统参数可以覆盖）
	for key, value := range params {
		if value != "" {
			queryParams.Set(key, value)
		}
	}

	// 强制设置系统参数（这些参数不应被 params 覆盖）
	queryParams.Set("version", c.config.APIVersion)
	queryParams.Set("type", c.config.ResponseType.String())
	queryParams.Set("key", c.config.APIKey)

	// 构建完整 URL（NameSilo API 使用 GET 方法）
	baseURL, err := osx.JoinURL(c.config.BaseURL, operation)
	if err != nil {
		return nil, namesilo.WrapError(ErrCodeAPIRequest, "failed to build request URL", err)
	}
	fullURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())

	// 用于记录响应信息
	var respStatus string
	var respHeaders interface{}
	var respBody []byte
	var apiErr error
	var recoverErr error

	// defer 记录完整的请求和响应日志
	defer func() {
		c.logRequestResponse(ctx, traceId, operation, fullURL, queryParams.Encode(),
			respStatus, respHeaders, respBody, apiErr, recoverErr)
	}()

	// 使用 errorx.Recover 捕获 panic
	recoverErr = errorx.Recover(func() {
		// 使用 GET 请求（NameSilo API 标准方法）
		// httpx 已修复参数覆盖问题，现在可以直接传完整 URL
		resp, err := c.httpClient.Get(fullURL).WithContext(ctx).Send()
		if err != nil {
			apiErr = err
			return
		}
		defer resp.Close()

		// 记录响应状态和头
		respStatus = resp.Status
		respHeaders = resp.Header

		// 读取响应体
		body, err := resp.Bytes()
		if err != nil {
			apiErr = err
			return
		}

		respBody = body
	})

	// 如果发生 panic，返回错误
	if recoverErr != nil {
		return nil, namesilo.WrapError(ErrCodeAPIRequest, fmt.Sprintf("panic occurred during request to operation: %s", operation), recoverErr)
	}

	// 如果有 API 错误，返回错误
	if apiErr != nil {
		return nil, namesilo.WrapError(ErrCodeAPIRequest, fmt.Sprintf("failed to send request to operation: %s", operation), apiErr)
	}

	// 检查响应是否为空
	if len(respBody) == 0 {
		apiErr = fmt.Errorf("empty response body")
		return nil, namesilo.NewError(ErrCodeAPIResponse, fmt.Sprintf("received empty response from NameSilo API for operation: %s", operation))
	}

	return respBody, nil
}

// ParseResponse 解析响应（公开给子模块使用）
func (c *Client) ParseResponse(data []byte, v interface{}) error {
	// 检查响应是否为空
	if len(data) == 0 {
		return namesilo.WrapError(ErrCodeAPIResponse, "received empty response from NameSilo API", fmt.Errorf("empty response body"))
	}

	// 根据配置的响应类型进行解析
	var err error
	switch c.config.ResponseType {
	case ResponseTypeJSON:
		err = json.Unmarshal(data, v)
	default:
		// 默认使用 XML 解析
		err = xml.Unmarshal(data, v)
	}

	if err != nil {
		return namesilo.WrapError(ErrCodeAPIResponse,
			fmt.Sprintf("failed to unmarshal response as %s, response preview: %s", c.config.ResponseType.String(), string(data)), err)
	}

	return nil
}

// logRequestResponse 记录请求和响应日志
func (c *Client) logRequestResponse(ctx context.Context, traceId, operation, fullURL, queryParams string,
	respStatus string, respHeaders interface{}, respBody []byte, apiErr, recoverErr error) {

	if recoverErr != nil {
		// panic 情况记录错误日志
		c.logPanicRequest(ctx, traceId, operation, fullURL, queryParams, recoverErr)
		return
	}

	// 正常情况记录调试日志
	c.logNormalRequestResponse(ctx, traceId, operation, fullURL, queryParams, respStatus, respHeaders, respBody, apiErr)
}

// logPanicRequest 记录 panic 请求日志
func (c *Client) logPanicRequest(ctx context.Context, traceId, operation, fullURL, queryParams string, panicErr error) {
	logMsg := c.formatRequestLog("🚨 NameSilo API Request PANIC", traceId, operation, fullURL, queryParams) +
		fmt.Sprintf("❌ Panic Error : %v\n", panicErr) +
		c.getLogSeparator()
	c.logger.ErrorContext(ctx, logMsg)
}

// logNormalRequestResponse 记录正常请求响应日志
func (c *Client) logNormalRequestResponse(ctx context.Context, traceId, operation, fullURL, queryParams string,
	respStatus string, respHeaders interface{}, respBody []byte, apiErr error) {

	statusIcon := "✅"
	if apiErr != nil {
		statusIcon = "❌"
	}

	logMsg := c.formatRequestLog(fmt.Sprintf("%s NameSilo API Request/Response", statusIcon), traceId, operation, fullURL, queryParams) +
		c.formatResponseLog(respStatus, respHeaders, respBody, apiErr)
	c.logger.DebugContext(ctx, logMsg)
}

// formatRequestLog 格式化请求日志
func (c *Client) formatRequestLog(title, traceId, operation, fullURL, queryParams string) string {
	return fmt.Sprintf(`
%s
%s
%s
📝 Trace ID    : %s
🎯 Operation   : %s
🔗 Request URL : %s
📋 Query Params: %s
%s
`,
		c.getLogSeparator(), title, c.getLogSeparator(),
		traceId, operation, fullURL, queryParams,
		c.getLogSeparator())
}

// formatResponseLog 格式化响应日志
func (c *Client) formatResponseLog(respStatus string, respHeaders interface{}, respBody []byte, apiErr error) string {
	return fmt.Sprintf(`📊 Status      : %s
📦 Headers     : %v
📏 Body Length : %d bytes
📄 Body        : %s
⚠️  Error       : %v
%s`,
		respStatus, respHeaders, len(respBody), string(respBody), apiErr,
		c.getLogSeparator())
}

// getLogSeparator 获取日志分隔线
func (c *Client) getLogSeparator() string {
	return strings.Repeat("━", 70) + "\n"
}
