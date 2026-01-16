/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:15:17
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 18:55:29
 * @FilePath: \go-namesilo\client\options.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package client

import (
	"time"
)

// Option 客户端选项函数
type Option func(*Config)

// WithID 设置客户端 ID
func WithID(id uint64) Option {
	return func(c *Config) {
		c.ID = id
	}
}

// WithAPIKey 设置 API Key（会覆盖 New 函数参数传入的 apiKey）
func WithAPIKey(apiKey string) Option {
	return func(c *Config) {
		c.APIKey = apiKey
	}
}

// WithBaseURL 设置 API 基础 URL
func WithBaseURL(baseURL string) Option {
	return func(c *Config) {
		c.BaseURL = baseURL
	}
}

// WithAPIVersion 设置 API 版本
func WithAPIVersion(apiVersion string) Option {
	return func(c *Config) {
		c.APIVersion = apiVersion
	}
}

// WithPublicURL 设置公共访问 URL
func WithPublicURL(publicURL string) Option {
	return func(c *Config) {
		c.PublicURL = publicURL
	}
}

// WithResponseType 设置响应类型（json 或 xml）
func WithResponseType(responseType ResponseType) Option {
	return func(c *Config) {
		c.ResponseType = responseType
	}
}

// WithTimeout 设置请求超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// WithDebug 开启调试模式
func WithDebug(debug bool) Option {
	return func(c *Config) {
		c.Debug = debug
	}
}

// WithLogger 设置日志记录器
func WithLogger(logger Logger) Option {
	return func(c *Config) {
		c.Logger = logger
	}
}
