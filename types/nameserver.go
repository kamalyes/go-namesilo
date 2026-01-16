/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:52:05
 * @FilePath: \go-namesilo\types\nameserver.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package types

import (
	"encoding/xml"
	"fmt"
)

// ============== Change NameServers ==============

// ChangeNameServersRequest 更改域名服务器请求
type ChangeNameServersRequest struct {
	Domain      string   `json:"domain"`      // 域名（逗号分隔，最多 200 个）
	Nameservers []string `json:"nameservers"` // 域名服务器列表（2-13个）
}

// ChangeNameServersResponse 更改域名服务器响应
type ChangeNameServersResponse struct {
	BaseResponse
}

// ============== List Registered NameServers ==============

// ListRegisteredNameServersRequest 列出已注册的域名服务器请求
type ListRegisteredNameServersRequest struct {
	Domain string `json:"domain"` // 要查询的域名
}

// ListRegisteredNameServersResponse 列出已注册的域名服务器响应
type ListRegisteredNameServersResponse struct {
	XMLName xml.Name                       `xml:"namesilo" json:"-"`
	Request RequestSection                 `xml:"request" json:"request"`
	Reply   ListRegisteredNameServersReply `xml:"reply" json:"reply"`
}

// ListRegisteredNameServersReply 列出已注册的域名服务器响应内容
type ListRegisteredNameServersReply struct {
	Code   int                    `xml:"code" json:"code"`
	Detail string                 `xml:"detail" json:"detail"`
	Hosts  []RegisteredNameServer `xml:"hosts" json:"hosts"`
}

// RegisteredNameServer 已注册的域名服务器
type RegisteredNameServer struct {
	Host string   `xml:"host" json:"host"` // 主机名
	IPs  []string `xml:"ip" json:"ips"`    // IP 地址列表
}

// ============== Add Registered NameServer ==============

// AddRegisteredNameServerRequest 添加已注册的域名服务器请求
type AddRegisteredNameServerRequest struct {
	Domain  string   `json:"domain"`   // 域名
	NewHost string   `json:"new_host"` // 新的主机名（不包括域名）
	IPs     []string `json:"ips"`      // IP 地址列表（1-13个）
}

// AddRegisteredNameServerResponse 添加已注册的域名服务器响应
type AddRegisteredNameServerResponse struct {
	BaseResponse
}

// ============== Modify Registered NameServer ==============

// ModifyRegisteredNameServerRequest 修改已注册的域名服务器请求
type ModifyRegisteredNameServerRequest struct {
	Domain      string   `json:"domain"`       // 域名
	CurrentHost string   `json:"current_host"` // 当前主机名（不包括域名）
	NewHost     string   `json:"new_host"`     // 新的主机名（不包括域名）
	IPs         []string `json:"ips"`          // IP 地址列表（1-13个）
}

// ModifyRegisteredNameServerResponse 修改已注册的域名服务器响应
type ModifyRegisteredNameServerResponse struct {
	BaseResponse
}

// ============== Delete Registered NameServer ==============

// DeleteRegisteredNameServerRequest 删除已注册的域名服务器请求
type DeleteRegisteredNameServerRequest struct {
	Domain      string `json:"domain"`       // 域名
	CurrentHost string `json:"current_host"` // 要删除的主机名（不包括域名）
}

// DeleteRegisteredNameServerResponse 删除已注册的域名服务器响应
type DeleteRegisteredNameServerResponse struct {
	BaseResponse
}

// ============== Reply Success/Error Methods ==============

// Success 判断响应是否成功
func (r *ListRegisteredNameServersReply) Success() bool {
	return r.Code == StatusSuccess ||
		r.Code == StatusSuccessWithPartialFailure ||
		r.Code == StatusSuccessWithContactIssue
}

// Error 实现 error 接口
func (r *ListRegisteredNameServersReply) Error() string {
	if r.Success() {
		return ""
	}
	return fmt.Sprintf("NameSilo API error (code %d): %s", r.Code, r.Detail)
}
