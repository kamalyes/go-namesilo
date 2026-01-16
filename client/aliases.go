/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:15:17
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:13:32
 * @FilePath: \go-namesilo\client\aliases.go
 * @Description: Client 类型别名 - 为 types 包中的类型创建别名，便于在 client 层使用
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package client

import (
	namesilo "github.com/kamalyes/go-namesilo"
	"github.com/kamalyes/go-namesilo/types"
)

// ============================================================================
// 错误类型
// ============================================================================

var (
	ErrAPIKeyRequired  = namesilo.ErrAPIKeyRequired
	ErrInvalidAPIKey   = namesilo.ErrInvalidAPIKey
	ErrRequestFailed   = namesilo.ErrRequestFailed
	ErrInvalidResponse = namesilo.ErrInvalidResponse
)

// ============================================================================
// 常量别名
// ============================================================================

const (
	DefaultBaseURL    = types.DefaultBaseURL
	DefaultAPIVersion = types.DefaultAPIVersion
	DefaultType       = types.DefaultType
	DefaultTimeout    = types.DefaultTimeout
)

// ============================================================================
// 配置类型
// ============================================================================

type Config = types.Config

// ============================================================================
// 通用响应类型
// ============================================================================

type (
	BaseResponse   = types.BaseResponse
	RequestSection = types.RequestSection
	CommonReply    = types.CommonReply
)

// ============================================================================
// 状态码常量
// ============================================================================

const (
	// 成功状态码
	StatusSuccess                   = types.StatusSuccess
	StatusSuccessWithPartialFailure = types.StatusSuccessWithPartialFailure
	StatusSuccessWithContactIssue   = types.StatusSuccessWithContactIssue

	// 一般错误
	StatusNoHTTPS             = types.StatusNoHTTPS
	StatusNoVersion           = types.StatusNoVersion
	StatusInvalidAPIVersion   = types.StatusInvalidAPIVersion
	StatusNoType              = types.StatusNoType
	StatusInvalidAPIType      = types.StatusInvalidAPIType
	StatusNoOperation         = types.StatusNoOperation
	StatusInvalidAPIOperation = types.StatusInvalidAPIOperation
	StatusMissingParameter    = types.StatusMissingParameter
	StatusNoAPIKey            = types.StatusNoAPIKey
	StatusInvalidAPIKey       = types.StatusInvalidAPIKey
	StatusInvalidUser         = types.StatusInvalidUser
	StatusAPINotForSubAccount = types.StatusAPINotForSubAccount
	StatusIPNotAllowed        = types.StatusIPNotAllowed
	StatusInvalidDomainSyntax = types.StatusInvalidDomainSyntax
)
