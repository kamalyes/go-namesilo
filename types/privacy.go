/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:00:00
 * @FilePath: \go-namesilo\types\privacy.go
 * @Description: Privacy 隐私保护相关类型定义
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package types

// ============================================================================
// Privacy 请求类型
// ============================================================================

// AddPrivacyRequest 添加隐私保护请求
type AddPrivacyRequest struct {
	Domain string `json:"domain" xml:"domain"` // 域名
}

// RemovePrivacyRequest 移除隐私保护请求
type RemovePrivacyRequest struct {
	Domain string `json:"domain" xml:"domain"` // 域名
}

// ============================================================================
// Privacy 响应类型
// ============================================================================

// AddPrivacyResponse 添加隐私保护响应
type AddPrivacyResponse struct {
	BaseResponse
	Reply AddPrivacyReply `xml:"reply" json:"reply"`
}

// AddPrivacyReply 添加隐私保护响应详情
type AddPrivacyReply struct {
	CommonReply
	OrderAmount float64 `xml:"order_amount" json:"order_amount"` // 订单金额
}

// RemovePrivacyResponse 移除隐私保护响应
type RemovePrivacyResponse struct {
	BaseResponse
	Reply RemovePrivacyReply `xml:"reply" json:"reply"`
}

// RemovePrivacyReply 移除隐私保护响应详情
type RemovePrivacyReply struct {
	CommonReply
}
