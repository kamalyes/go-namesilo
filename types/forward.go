/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 15:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 22:06:15
 * @FilePath: \go-namesilo\types\forward.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package types

import "encoding/xml"

// Forward 方法类型常量
const (
	ForwardMethod301     = "301"     // 301 永久重定向
	ForwardMethod302     = "302"     // 302 临时重定向
	ForwardMethodCloaked = "cloaked" // 伪装转发
)

// Forward 协议类型常量
const (
	ForwardProtocolHTTP  = "http"
	ForwardProtocolHTTPS = "https"
)

// ============== Sub-Domain Forward ==============

// SubDomainForwardRequest 子域名转发请求
type SubDomainForwardRequest struct {
	Domain          string `json:"domain"`                     // 主域名
	SubDomain       string `json:"sub_domain"`                 // 要转发的子域名
	Protocol        string `json:"protocol"`                   // URL 协议（http 或 https）
	Address         string `json:"address"`                    // 转发目标地址
	Method          string `json:"method"`                     // 转发方式（301, 302 或 cloaked）
	MetaTitle       string `json:"meta_title,omitempty"`       // META 标题（仅 cloaked 方式使用，需 URL 编码）
	MetaDescription string `json:"meta_description,omitempty"` // META 描述（仅 cloaked 方式使用，需 URL 编码）
	MetaKeywords    string `json:"meta_keywords,omitempty"`    // META 关键词（仅 cloaked 方式使用，需 URL 编码）
}

// SubDomainForwardResponse 子域名转发响应
type SubDomainForwardResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   CommonReply    `xml:"reply" json:"reply"`
}

// ============== Delete Sub-Domain Forward ==============

// DeleteSubDomainForwardRequest 删除子域名转发请求
type DeleteSubDomainForwardRequest struct {
	Domain    string `json:"domain"`     // 主域名
	SubDomain string `json:"sub_domain"` // 要删除转发的子域名
}

// DeleteSubDomainForwardResponse 删除子域名转发响应
type DeleteSubDomainForwardResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   CommonReply    `xml:"reply" json:"reply"`
}
