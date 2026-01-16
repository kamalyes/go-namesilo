/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 00:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:30:00
 * @FilePath: \go-namesilo\types\transfer.go
 * @Description: Transfer 域名转移相关类型定义
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package types

// ============================================================================
// Transfer 请求类型
// ============================================================================

// RetrieveAuthCodeRequest 获取授权码请求
type RetrieveAuthCodeRequest struct {
	Domain string `json:"domain" xml:"domain"` // 域名
}

// CheckTransferStatusRequest 检查转移状态请求
type CheckTransferStatusRequest struct {
	Domain string `json:"domain" xml:"domain"` // 域名
}

// TransferUpdateResubmitRequest 重新提交转移请求
type TransferUpdateResubmitRequest struct {
	Domain string `json:"domain" xml:"domain"` // 域名
}

// TransferUpdateResendEmailRequest 重新发送管理员邮件请求
type TransferUpdateResendEmailRequest struct {
	Domain string `json:"domain" xml:"domain"` // 域名
}

// TransferUpdateChangeEPPCodeRequest 更改 EPP 授权码请求
type TransferUpdateChangeEPPCodeRequest struct {
	Domain  string `json:"domain" xml:"domain"`     // 域名
	EPPCode string `json:"epp_code" xml:"epp_code"` // 新的 EPP 授权码
}

// ============================================================================
// Transfer 响应类型
// ============================================================================

// RetrieveAuthCodeResponse 获取授权码响应
type RetrieveAuthCodeResponse struct {
	BaseResponse
	Reply RetrieveAuthCodeReply `xml:"reply" json:"reply"`
}

// RetrieveAuthCodeReply 获取授权码响应详情
type RetrieveAuthCodeReply struct {
	CommonReply
	AuthCode string `xml:"auth" json:"auth"` // 授权码(EPP Code)
}

// CheckTransferStatusResponse 检查转移状态响应
type CheckTransferStatusResponse struct {
	BaseResponse
	Reply CheckTransferStatusReply `xml:"reply" json:"reply"`
}

// CheckTransferStatusReply 检查转移状态响应详情
type CheckTransferStatusReply struct {
	CommonReply
	Transfer TransferStatus `xml:"transfer" json:"transfer"` // 转移状态信息
}

// TransferStatus 转移状态信息
type TransferStatus struct {
	Date    string `xml:"date" json:"date"`       // 转移日期
	Status  string `xml:"status" json:"status"`   // 转移状态
	Message string `xml:"message" json:"message"` // 状态消息
}

// TransferUpdateResubmitResponse 重新提交转移响应
type TransferUpdateResubmitResponse struct {
	BaseResponse
	Reply TransferUpdateResubmitReply `xml:"reply" json:"reply"`
}

// TransferUpdateResubmitReply 重新提交转移响应详情
type TransferUpdateResubmitReply struct {
	CommonReply
}

// TransferUpdateResendEmailResponse 重新发送管理员邮件响应
type TransferUpdateResendEmailResponse struct {
	BaseResponse
	Reply TransferUpdateResendEmailReply `xml:"reply" json:"reply"`
}

// TransferUpdateResendEmailReply 重新发送管理员邮件响应详情
type TransferUpdateResendEmailReply struct {
	CommonReply
}

// TransferUpdateChangeEPPCodeResponse 更改 EPP 授权码响应
type TransferUpdateChangeEPPCodeResponse struct {
	BaseResponse
	Reply TransferUpdateChangeEPPCodeReply `xml:"reply" json:"reply"`
}

// TransferUpdateChangeEPPCodeReply 更改 EPP 授权码响应详情
type TransferUpdateChangeEPPCodeReply struct {
	CommonReply
}
