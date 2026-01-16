/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:37:28
 * @FilePath: \go-namesilo\types\forwarding.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package types

import "encoding/xml"

// 转发类型常量
const (
	ForwardType301 = "301" // 永久重定向
	ForwardType302 = "302" // 临时重定向
)

// ============== Forward Domain ==============

// ForwardDomainRequest 设置域名转发请求
type ForwardDomainRequest struct {
	Domain      string `json:"domain"`                 // 要设置转发的域名
	Protocol    string `json:"protocol"`               // 协议: http 或 https
	Address     string `json:"address"`                // 转发到的目标地址
	Method      string `json:"method,omitempty"`       // 转发方法: 301 或 302 (默认 301)
	IncludePath string `json:"include_path,omitempty"` // 是否包含路径: Yes 或 No (默认 Yes)
	Wildcard    string `json:"wildcard,omitempty"`     // 是否通配符: Yes 或 No (默认 Yes)
}

// ForwardDomainResponse 设置域名转发响应
type ForwardDomainResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   CommonReply    `xml:"reply" json:"reply"`
}

// ============== Forward Subdomain ==============

// ForwardSubdomainRequest 设置子域名转发请求
type ForwardSubdomainRequest struct {
	Domain      string `json:"domain"`                 // 域名
	SubDomain   string `json:"sub_domain"`             // 子域名（不包含域名部分）
	Protocol    string `json:"protocol"`               // 协议: http 或 https
	Address     string `json:"address"`                // 转发到的目标地址
	Method      string `json:"method,omitempty"`       // 转发方法: 301 或 302 (默认 301)
	IncludePath string `json:"include_path,omitempty"` // 是否包含路径: Yes 或 No (默认 Yes)
}

// ForwardSubdomainResponse 设置子域名转发响应
type ForwardSubdomainResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   CommonReply    `xml:"reply" json:"reply"`
}

// ============== Delete Forward ==============

// DeleteForwardRequest 删除域名/子域名转发请求
type DeleteForwardRequest struct {
	Domain    string `json:"domain"`               // 域名
	SubDomain string `json:"sub_domain,omitempty"` // 子域名（为空则删除主域名转发）
}

// DeleteForwardResponse 删除转发响应
type DeleteForwardResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   CommonReply    `xml:"reply" json:"reply"`
}

// ============== List Email Forwards ==============

// ListEmailForwardsRequest 列出邮件转发请求
type ListEmailForwardsRequest struct {
	Domain string `json:"domain"` // 要查询的域名
}

// ListEmailForwardsResponse 列出邮件转发响应
type ListEmailForwardsResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   struct {
		Code      int            `xml:"code" json:"code"`
		Detail    string         `xml:"detail" json:"detail"`
		Addresses []EmailForward `xml:"address" json:"addresses"` // 邮件转发列表
	} `xml:"reply" json:"reply"`
}

// EmailForward 邮件转发信息
type EmailForward struct {
	Email           string   `xml:"email" json:"email"`                       // 源邮件地址
	ForwardsTo      []string `xml:"forwards_to>email" json:"forwards_to"`     // 转发目标列表
	EmailsForwarded int      `xml:"emails_forwarded" json:"emails_forwarded"` // 已转发邮件数量
}

// ============== Configure Email Forward ==============

// ConfigureEmailForwardRequest 配置邮件转发请求
type ConfigureEmailForwardRequest struct {
	Domain  string   `json:"domain"`  // 域名
	Email   string   `json:"email"`   // 源邮件地址（不包含@domain部分）
	Forward []string `json:"forward"` // 转发目标邮件地址列表（最多 5 个）
}

// ConfigureEmailForwardResponse 配置邮件转发响应
type ConfigureEmailForwardResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   CommonReply    `xml:"reply" json:"reply"`
}

// ============== Delete Email Forward ==============

// DeleteEmailForwardRequest 删除邮件转发请求
type DeleteEmailForwardRequest struct {
	Domain string `json:"domain"` // 域名
	Email  string `json:"email"`  // 要删除的源邮件地址
}

// DeleteEmailForwardResponse 删除邮件转发响应
type DeleteEmailForwardResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   CommonReply    `xml:"reply" json:"reply"`
}
