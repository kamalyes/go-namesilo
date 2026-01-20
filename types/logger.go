/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-30 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 19:23:00
 * @FilePath: \go-namesilo\types\common.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package types

import "context"

// Logger 日志接口
type Logger interface {
	DebugContext(ctx context.Context, msg string, keysAndValues ...interface{})
	InfoContext(ctx context.Context, msg string, keysAndValues ...interface{})
	WarnContext(ctx context.Context, msg string, keysAndValues ...interface{})
	ErrorContext(ctx context.Context, msg string, keysAndValues ...interface{})
}

// EmptyLogger 空日志记录器实现,不执行任何操作
type EmptyLogger struct{}

// DebugContext 实现 Logger 接口
func (e *EmptyLogger) DebugContext(ctx context.Context, msg string, keysAndValues ...interface{}) {}

// InfoContext 实现 Logger 接口
func (e *EmptyLogger) InfoContext(ctx context.Context, msg string, keysAndValues ...interface{}) {}

// WarnContext 实现 Logger 接口
func (e *EmptyLogger) WarnContext(ctx context.Context, msg string, keysAndValues ...interface{}) {}

// ErrorContext 实现 Logger 接口
func (e *EmptyLogger) ErrorContext(ctx context.Context, msg string, keysAndValues ...interface{}) {}

// noLogger 全局空日志记录器实例（单例）
var noLogger Logger = &EmptyLogger{}

// NewEmptyLogger 返回一个空日志记录器（单例模式）
func NewEmptyLogger() Logger {
	return noLogger
}
