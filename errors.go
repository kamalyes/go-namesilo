/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-16 22:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 22:41:33
 * @FilePath: \go-namesilo\errors.go
 * @Description: 统一错误定义
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package namesilo

import (
	"errors"
	"fmt"
)

// Error 自定义错误类型
type Error struct {
	Code    string // 错误代码
	Message string // 错误消息
	Err     error  // 原始错误
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Err
}

// NewError 创建新的错误
func NewError(code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// WrapError 包装错误
func WrapError(code, message string, err error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// 错误代码常量
const (
	// 参数验证错误
	ErrCodeInvalidParam    = "INVALID_PARAM"
	ErrCodeMissingParam    = "MISSING_PARAM"
	ErrCodeParamOutOfRange = "PARAM_OUT_OF_RANGE"

	// API 请求错误
	ErrCodeAPIRequest  = "API_REQUEST_FAILED"
	ErrCodeAPIResponse = "API_RESPONSE_ERROR"

	// 业务逻辑错误
	ErrCodeOperationFailed = "OPERATION_FAILED"
)

// 预定义的参数验证错误
var (
	// 通用字段验证错误
	ErrDomainRequired      = NewError(ErrCodeMissingParam, "domain is required")
	ErrContactIDRequired   = NewError(ErrCodeMissingParam, "contact_id is required")
	ErrRecordIDRequired    = NewError(ErrCodeMissingParam, "record id (rrid) is required")
	ErrOrderNumberRequired = NewError(ErrCodeMissingParam, "order_number is required")

	// Contact 字段验证错误
	ErrFirstNameRequired = NewError(ErrCodeMissingParam, "first name is required")
	ErrLastNameRequired  = NewError(ErrCodeMissingParam, "last name is required")
	ErrAddressRequired   = NewError(ErrCodeMissingParam, "address is required")
	ErrCityRequired      = NewError(ErrCodeMissingParam, "city is required")
	ErrStateRequired     = NewError(ErrCodeMissingParam, "state is required")
	ErrZipRequired       = NewError(ErrCodeMissingParam, "zip is required")
	ErrCountryRequired   = NewError(ErrCodeMissingParam, "country is required")
	ErrEmailRequired     = NewError(ErrCodeMissingParam, "email is required")
	ErrPhoneRequired     = NewError(ErrCodeMissingParam, "phone is required")

	// Domain 相关验证错误
	ErrSubDomainRequired      = NewError(ErrCodeMissingParam, "sub_domain is required")
	ErrProtocolRequired       = NewError(ErrCodeMissingParam, "protocol is required")
	ErrAddressForwardRequired = NewError(ErrCodeMissingParam, "address is required")
	ErrMethodRequired         = NewError(ErrCodeMissingParam, "method is required")
	ErrDomainsRequired        = NewError(ErrCodeMissingParam, "domains is required")

	// DNS 相关验证错误
	ErrRRTypeRequired      = NewError(ErrCodeMissingParam, "rrtype is required")
	ErrRRHostRequired      = NewError(ErrCodeMissingParam, "rrhost is required")
	ErrRRValueRequired     = NewError(ErrCodeMissingParam, "rrvalue is required")
	ErrDigestRequired      = NewError(ErrCodeMissingParam, "digest is required")
	ErrRecordTypeRequired  = NewError(ErrCodeMissingParam, "record type is required")
	ErrRecordValueRequired = NewError(ErrCodeMissingParam, "record value is required")
	ErrInvalidRecordType   = NewError(ErrCodeInvalidParam, "invalid record type")
	ErrRecordNotFound      = NewError(ErrCodeOperationFailed, "dns record not found")
	ErrInvalidTTL          = NewError(ErrCodeInvalidParam, "invalid TTL value")
	ErrInvalidPriority     = NewError(ErrCodeInvalidParam, "invalid priority value")

	// Transfer/Register 相关错误
	ErrRecipientLoginRequired = NewError(ErrCodeMissingParam, "recipient_login is required")

	// 范围验证错误
	ErrYearsOutOfRange        = NewError(ErrCodeParamOutOfRange, "years must be between 1 and 10")
	ErrDomainsExceedLimit     = NewError(ErrCodeParamOutOfRange, "domains cannot exceed 200")
	ErrForwardListExceedLimit = NewError(ErrCodeParamOutOfRange, "forward list cannot exceed 5 addresses")
	ErrIPAddressExceedLimit   = NewError(ErrCodeParamOutOfRange, "A maximum of 13 IP addresses can be added.")
	ErrDaysCountInvalid       = NewError(ErrCodeParamOutOfRange, "days_count must be greater than 0")
	ErrAmountInvalid          = NewError(ErrCodeParamOutOfRange, "amount must be greater than 0")

	// 复杂验证错误
	ErrContactRoleRequired = NewError(ErrCodeMissingParam, "at least one contact role is required (registrant, administrative, billing, or technical)")

	// Contact 业务错误
	ErrContactNotFound    = NewError(ErrCodeOperationFailed, "contact not found")
	ErrContactInUse       = NewError(ErrCodeOperationFailed, "contact is in use and cannot be deleted")
	ErrInvalidContactData = NewError(ErrCodeInvalidParam, "invalid contact data")

	// Domain 业务错误
	ErrDomainNotFound       = NewError(ErrCodeOperationFailed, "domain not found")
	ErrDomainAlreadyExists  = NewError(ErrCodeOperationFailed, "domain already exists")
	ErrDomainLocked         = NewError(ErrCodeOperationFailed, "domain is locked")
	ErrDomainNotAvailable   = NewError(ErrCodeOperationFailed, "domain is not available for registration")
	ErrInvalidDomain        = NewError(ErrCodeInvalidParam, "invalid domain name")
	ErrDomainTransferDenied = NewError(ErrCodeOperationFailed, "domain transfer denied")
	ErrInsufficientBalance  = NewError(ErrCodeOperationFailed, "insufficient account balance")

	// Transfer 相关错误
	ErrTransferNotFound       = NewError(ErrCodeOperationFailed, "transfer not found")
	ErrAuthCodeRequired       = NewError(ErrCodeMissingParam, "auth code (EPP code) is required")
	ErrInvalidAuthCode        = NewError(ErrCodeInvalidParam, "invalid auth code (EPP code)")
	ErrTransferIDRequired     = NewError(ErrCodeMissingParam, "transfer_id is required")
	ErrTransferAlreadyPending = NewError(ErrCodeOperationFailed, "transfer is already pending")

	// Forwarding 相关错误
	ErrURLRequired       = NewError(ErrCodeMissingParam, "URL is required")
	ErrSubdomainRequired = NewError(ErrCodeMissingParam, "subdomain is required")
	ErrForwardNotFound   = NewError(ErrCodeOperationFailed, "forward not found")
	ErrInvalidURL        = NewError(ErrCodeInvalidParam, "invalid URL format")
	ErrInvalidEmail      = NewError(ErrCodeInvalidParam, "invalid email format")

	// Nameserver 相关错误
	ErrNameserverRequired = NewError(ErrCodeMissingParam, "nameserver is required")
	ErrNameserverNotFound = NewError(ErrCodeOperationFailed, "nameserver not found")
	ErrInvalidNameserver  = NewError(ErrCodeInvalidParam, "invalid nameserver format")
	ErrNameserverInUse    = NewError(ErrCodeOperationFailed, "nameserver is in use")
	ErrInvalidIPAddress   = NewError(ErrCodeInvalidParam, "invalid IP address")

	// Privacy 相关错误
	ErrPrivacyAlreadyEnabled  = NewError(ErrCodeOperationFailed, "privacy is already enabled for this domain")
	ErrPrivacyAlreadyDisabled = NewError(ErrCodeOperationFailed, "privacy is already disabled for this domain")

	// Account 相关错误
	ErrInvalidAmount = NewError(ErrCodeInvalidParam, "invalid amount")
	ErrPaymentFailed = NewError(ErrCodeOperationFailed, "payment failed")

	// API 客户端错误
	ErrAPIKeyRequired  = NewError(ErrCodeMissingParam, "apiKey is required")
	ErrInvalidAPIKey   = NewError(ErrCodeInvalidParam, "invalid API key")
	ErrRequestFailed   = NewError(ErrCodeAPIRequest, "API request failed")
	ErrInvalidResponse = NewError(ErrCodeAPIResponse, "invalid API response")
)

// NewAPIError 创建 API 错误
func NewAPIError(operation string, detail string) *Error {
	return NewError(ErrCodeOperationFailed, fmt.Sprintf("%s failed: %s", operation, detail))
}

// NewRequestError 创建请求错误
func NewRequestError(err error) *Error {
	return WrapError(ErrCodeAPIRequest, "request failed", err)
}

// NewParseError 创建解析错误
func NewParseError(err error) *Error {
	return WrapError(ErrCodeAPIResponse, "failed to unmarshal response (tried JSON and XML)", err)
}

// IsError 检查错误是否为指定的自定义错误
func IsError(err error, target *Error) bool {
	var e *Error
	if errors.As(err, &e) {
		return e.Code == target.Code && e.Message == target.Message
	}
	return false
}
