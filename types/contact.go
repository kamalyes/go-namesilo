/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-16 22:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 22:30:00
 * @FilePath: \go-namesilo\types\contact.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package types

import (
	"encoding/xml"
	"fmt"
)

// ============== Contact List ==============

// ContactListRequest 列出联系人请求
type ContactListRequest struct {
	ContactID string `json:"contact_id,omitempty"` // 可选，特定联系人ID
	Offset    int    `json:"offset,omitempty"`     // 可选，分页偏移量
}

// ContactListResponse 列出联系人响应
type ContactListResponse struct {
	XMLName xml.Name         `xml:"namesilo" json:"-"`
	Request RequestSection   `xml:"request" json:"request"`
	Reply   ContactListReply `xml:"reply" json:"reply"`
}

// ContactListReply 列出联系人响应内容
type ContactListReply struct {
	Code     int       `xml:"code" json:"code"`
	Detail   string    `xml:"detail" json:"detail"`
	Contacts []Contact `xml:"contact" json:"contacts"`
}

// Contact 联系人信息
type Contact struct {
	ContactID      string `xml:"contact_id" json:"contact_id"`
	DefaultProfile int    `xml:"default_profile" json:"default_profile"` // 1=默认联系人，0=非默认
	Nickname       string `xml:"nickname" json:"nickname"`
	Company        string `xml:"company" json:"company,omitempty"`
	FirstName      string `xml:"first_name" json:"first_name"`
	LastName       string `xml:"last_name" json:"last_name"`
	Address        string `xml:"address" json:"address"`
	Address2       string `xml:"address2" json:"address2,omitempty"`
	City           string `xml:"city" json:"city"`
	State          string `xml:"state" json:"state"`
	Zip            string `xml:"zip" json:"zip"`
	Country        string `xml:"country" json:"country"`
	Email          string `xml:"email" json:"email"`
	Phone          string `xml:"phone" json:"phone"`
	Fax            string `xml:"fax" json:"fax,omitempty"`
	// .US 专用字段
	USNC string `xml:"usnc" json:"usnc,omitempty"` // .US Nexus Category
	USAP string `xml:"usap" json:"usap,omitempty"` // .US Application Purpose
	// .CA 专用字段
	CALF string `xml:"calf" json:"calf,omitempty"` // CIRA Legal Form
	CALN string `xml:"caln" json:"caln,omitempty"` // CIRA Language
	CAAG string `xml:"caag" json:"caag,omitempty"` // CIRA Agreement Version
	CAWD string `xml:"cawd" json:"cawd,omitempty"` // CIRA WHOIS Display
	// .EU 专用字段
	EUCS string `xml:"eucs" json:"eucs,omitempty"` // Citizenship Country
}

// ============== Contact Add ==============

// ContactAddRequest 添加联系人请求
type ContactAddRequest struct {
	// 必填字段
	FirstName string `json:"fn"` // 名（最多32字符）
	LastName  string `json:"ln"` // 姓（最多32字符）
	Address   string `json:"ad"` // 地址（最多128字符）
	City      string `json:"cy"` // 城市（最多64字符）
	State     string `json:"st"` // 州/省（最多64字符）
	Zip       string `json:"zp"` // 邮编（最多16字符）
	Country   string `json:"ct"` // 国家代码（最多4字符）
	Email     string `json:"em"` // 邮箱（最多128字符）
	Phone     string `json:"ph"` // 电话（最多32字符）
	// 可选字段
	Nickname string `json:"nn,omitempty"`  // 昵称（最多24字符）
	Company  string `json:"cp,omitempty"`  // 公司（最多64字符）
	Address2 string `json:"ad2,omitempty"` // 地址2（最多128字符）
	Fax      string `json:"fx,omitempty"`  // 传真（最多32字符）
	// .US 专用字段
	USNC string `json:"usnc,omitempty"` // .US Nexus Category（最多3字符）
	USAP string `json:"usap,omitempty"` // .US Application Purpose（最多2字符）
	// .CA 专用字段
	CALF string `json:"calf,omitempty"` // CIRA Legal Form
	CALN string `json:"caln,omitempty"` // CIRA Language
	CAAG string `json:"caag,omitempty"` // CIRA Agreement Version
	CAWD string `json:"cawd,omitempty"` // CIRA WHOIS Display
	// .EU 专用字段
	EUCS string `json:"eucs,omitempty"` // Citizenship Country Abbreviation
}

// ContactAddResponse 添加联系人响应
type ContactAddResponse struct {
	XMLName xml.Name        `xml:"namesilo" json:"-"`
	Request RequestSection  `xml:"request" json:"request"`
	Reply   ContactAddReply `xml:"reply" json:"reply"`
}

// ContactAddReply 添加联系人响应内容
type ContactAddReply struct {
	Code      int    `xml:"code" json:"code"`
	Detail    string `xml:"detail" json:"detail"`
	ContactID string `xml:"contact_id" json:"contact_id"` // 新创建的联系人ID
}

// ============== Contact Update ==============

// ContactUpdateRequest 更新联系人请求
type ContactUpdateRequest struct {
	// 必填字段
	ContactID string `json:"contact_id"` // 要更新的联系人ID
	FirstName string `json:"fn"`         // 名（最多32字符）
	LastName  string `json:"ln"`         // 姓（最多32字符）
	Address   string `json:"ad"`         // 地址（最多128字符）
	City      string `json:"cy"`         // 城市（最多64字符）
	State     string `json:"st"`         // 州/省（最多64字符）
	Zip       string `json:"zp"`         // 邮编（最多16字符）
	Country   string `json:"ct"`         // 国家代码（最多4字符）
	Email     string `json:"em"`         // 邮箱（最多128字符）
	Phone     string `json:"ph"`         // 电话（最多32字符）
	// 可选字段
	Nickname string `json:"nn,omitempty"`  // 昵称（最多24字符）
	Company  string `json:"cp,omitempty"`  // 公司（最多64字符）
	Address2 string `json:"ad2,omitempty"` // 地址2（最多128字符）
	Fax      string `json:"fx,omitempty"`  // 传真（最多32字符）
	// .US 专用字段
	USNC string `json:"usnc,omitempty"` // .US Nexus Category（最多3字符）
	USAP string `json:"usap,omitempty"` // .US Application Purpose（最多2字符）
	// .CA 专用字段
	CALF string `json:"calf,omitempty"` // CIRA Legal Form
	CALN string `json:"caln,omitempty"` // CIRA Language
	CAAG string `json:"caag,omitempty"` // CIRA Agreement Version
	CAWD string `json:"cawd,omitempty"` // CIRA WHOIS Display
	// .EU 专用字段
	EUCS string `json:"eucs,omitempty"` // Citizenship Country Abbreviation
}

// ContactUpdateResponse 更新联系人响应
type ContactUpdateResponse struct {
	BaseResponse
}

// ============== Contact Delete ==============

// ContactDeleteRequest 删除联系人请求
type ContactDeleteRequest struct {
	ContactID string `json:"contact_id"` // 要删除的联系人ID
}

// ContactDeleteResponse 删除联系人响应
type ContactDeleteResponse struct {
	BaseResponse
}

// ============== Contact Domain Associate ==============

// ContactDomainAssociateRequest 关联联系人到域名请求
type ContactDomainAssociateRequest struct {
	Domain         string `json:"domain"`                   // 要更新的域名
	Registrant     string `json:"registrant,omitempty"`     // 注册人联系人ID
	Administrative string `json:"administrative,omitempty"` // 管理联系人ID
	Billing        string `json:"billing,omitempty"`        // 账单联系人ID
	Technical      string `json:"technical,omitempty"`      // 技术联系人ID
}

// ContactDomainAssociateResponse 关联联系人到域名响应
type ContactDomainAssociateResponse struct {
	BaseResponse
}

// ============== Reply Success/Error Methods ==============

// Success 判断响应是否成功
func (r *ContactListReply) Success() bool {
	return r.Code == StatusSuccess ||
		r.Code == StatusSuccessWithPartialFailure ||
		r.Code == StatusSuccessWithContactIssue
}

// Error 实现 error 接口
func (r *ContactListReply) Error() string {
	if r.Success() {
		return ""
	}
	return fmt.Sprintf("NameSilo API error (code %d): %s", r.Code, r.Detail)
}

// Success 判断响应是否成功
func (r *ContactAddReply) Success() bool {
	return r.Code == StatusSuccess ||
		r.Code == StatusSuccessWithPartialFailure ||
		r.Code == StatusSuccessWithContactIssue
}

// Error 实现 error 接口
func (r *ContactAddReply) Error() string {
	if r.Success() {
		return ""
	}
	return fmt.Sprintf("NameSilo API error (code %d): %s", r.Code, r.Detail)
}
