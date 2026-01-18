/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:15:55
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-15 22:05:31
 * @FilePath: \go-namesilo\types\domain.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package types

import (
	"encoding/xml"
	"fmt"
)

// ============== Domain Register ==============

// RegisterDomainRequest 注册域名请求
type RegisterDomainRequest struct {
	Domain string `json:"domain"` // 要注册的域名
	Years  int    `json:"years"`  // 注册年限
	// 以下为可选参数
	PaymentID string   `json:"payment_id,omitempty"` // 支付 ID（可选）
	Private   bool     `json:"private,omitempty"`    // 是否启用隐私保护
	AutoRenew bool     `json:"auto_renew,omitempty"` // 是否自动续费
	Portfolio string   `json:"portfolio,omitempty"`  // 作品集名称（可选）
	NS        []string `json:"ns,omitempty"`         // 域名服务器列表（最多13个）
	Coupon    string   `json:"coupon,omitempty"`     // 优惠券代码
	ContactID string   `json:"contact_id,omitempty"` // 联系人 ID（可选）
	// WHOIS 联系人信息字段（可选）
	FirstName string `json:"fn,omitempty"`  // 名
	LastName  string `json:"ln,omitempty"`  // 姓
	Address   string `json:"ad,omitempty"`  // 地址
	City      string `json:"cy,omitempty"`  // 城市
	State     string `json:"st,omitempty"`  // 州/省
	Zip       string `json:"zp,omitempty"`  // 邮编
	Country   string `json:"ct,omitempty"`  // 国家代码
	Email     string `json:"em,omitempty"`  // 邮箱
	Phone     string `json:"ph,omitempty"`  // 电话
	Nickname  string `json:"nn,omitempty"`  // 昵称
	Company   string `json:"cp,omitempty"`  // 公司
	Address2  string `json:"ad2,omitempty"` // 地址2
	Fax       string `json:"fx,omitempty"`  // 传真
	// .US 域名专用字段
	USNexusCategory string `json:"usnc,omitempty"` // .US Nexus Category
	USAppPurpose    string `json:"usap,omitempty"` // .US Application Purpose
}

// RegisterDomainResponse 注册域名响应
// 使用 CommonReply（包含 Code, Detail, Message, Domain, OrderAmount）
type RegisterDomainResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   CommonReply    `xml:"reply" json:"reply"`
}

// ============== Domain List ==============

// ListDomainsRequest 列出域名请求
type ListDomainsRequest struct {
	Portfolio    string `json:"portfolio,omitempty"`     // 按作品集名称过滤（可选）
	PageSize     int    `json:"page_size,omitempty"`     // 每页显示的域名数量（可选）
	Page         int    `json:"page,omitempty"`          // 页码（可选）
	WithBid      int    `json:"with_bid,omitempty"`      // 显示竞价信息（可选，1=显示）
	SkipExpired  int    `json:"skip_expired,omitempty"`  // 跳过过期域名（可选，1=跳过）
	ExpiredGrace int    `json:"expired_grace,omitempty"` // 仅显示宽限期内的过期域名（可选，1=仅显示）
}

// ListDomainsResponse 列出域名响应
type ListDomainsResponse struct {
	XMLName xml.Name         `xml:"namesilo" json:"-"`
	Request RequestSection   `xml:"request" json:"request"`
	Reply   ListDomainsReply `xml:"reply" json:"reply"`
}

// ListDomainsReply 列出域名响应内容
type ListDomainsReply struct {
	Code    int      `xml:"code" json:"code"`
	Detail  string   `xml:"detail" json:"detail"`
	Domains []Domain `xml:"domains>domain" json:"domains"`
	Pager   *Pager   `xml:"pager" json:"pager,omitempty"` // 分页信息（可选）
}

// Pager 分页信息
type Pager struct {
	Total    int `xml:"total" json:"total"`        // 总数
	PageSize int `xml:"pageSize" json:"page_size"` // 每页大小
	Page     int `xml:"page" json:"page"`          // 当前页
}

// Domain 域名信息
type Domain struct {
	Domain  string  `xml:",chardata" json:"domain"`
	Created string  `xml:"created,attr" json:"created,omitempty"`
	Expires string  `xml:"expires,attr" json:"expires,omitempty"`
	MaxBid  float64 `xml:"maxBid,attr" json:"max_bid,omitempty"` // 最高出价（仅在 withBid=1 时返回）
}

// ============== Domain Info ==============

// GetDomainInfoRequest 获取域名信息请求
type GetDomainInfoRequest struct {
	Domain string `json:"domain"` // 要查询的域名
}

// GetDomainInfoResponse 获取域名信息响应
type GetDomainInfoResponse struct {
	XMLName xml.Name        `xml:"namesilo" json:"-"`
	Request RequestSection  `xml:"request" json:"request"`
	Reply   DomainInfoReply `xml:"reply" json:"reply"`
}

// DomainInfoReply 域名信息响应内容
type DomainInfoReply struct {
	Code                      int               `xml:"code" json:"code"`
	Detail                    string            `xml:"detail" json:"detail"`
	Created                   string            `xml:"created" json:"created"`
	Expires                   string            `xml:"expires" json:"expires"`
	Status                    string            `xml:"status" json:"status"`
	Locked                    string            `xml:"locked" json:"locked"`
	Private                   string            `xml:"private" json:"private"`
	AutoRenew                 string            `xml:"auto_renew" json:"auto_renew"`
	TrafficType               string            `xml:"traffic_type" json:"traffic_type"`
	EmailVerificationRequired string            `xml:"email_verification_required" json:"email_verification_required"`
	Portfolio                 string            `xml:"portfolio" json:"portfolio"`
	ForwardURL                string            `xml:"forward_url" json:"forward_url"`
	ForwardType               string            `xml:"forward_type" json:"forward_type"`
	Nameservers               []NameserverEntry `xml:"nameservers>nameserver" json:"nameservers"`
	ContactIDs                ContactIDs        `xml:"contact_ids" json:"contact_ids"`
}

// NameserverEntry 域名服务器条目
type NameserverEntry struct {
	Nameserver string `xml:",chardata" json:"nameserver"`
	Position   int    `xml:"position,attr" json:"position,omitempty"`
}

// ContactIDs 联系人 ID 信息
type ContactIDs struct {
	Registrant     string `xml:"registrant" json:"registrant"`
	Administrative string `xml:"administrative" json:"administrative"`
	Technical      string `xml:"technical" json:"technical"`
	Billing        string `xml:"billing" json:"billing"`
}

// ============== Domain Renew ==============

// RenewDomainRequest 续费域名请求
type RenewDomainRequest struct {
	Domain    string `json:"domain"`               // 要续费的域名
	Years     int    `json:"years"`                // 续费年限
	PaymentID string `json:"payment_id,omitempty"` // 支付 ID（可选）
	Coupon    string `json:"coupon,omitempty"`     // 优惠券代码（可选）
}

// RenewDomainResponse 续费域名响应
// 使用 CommonReply（包含 Code, Detail, Message, Domain, OrderAmount）
type RenewDomainResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   CommonReply    `xml:"reply" json:"reply"`
}

// ============== Domain Transfer ==============

// TransferDomainRequest 转移域名请求
type TransferDomainRequest struct {
	Domain    string   `json:"domain"`               // 要转移的域名
	Auth      string   `json:"auth,omitempty"`       // 授权码
	PaymentID string   `json:"payment_id,omitempty"` // 支付 ID
	Private   bool     `json:"private,omitempty"`    // 是否启用隐私保护
	AutoRenew bool     `json:"auto_renew,omitempty"` // 是否自动续费
	Portfolio string   `json:"portfolio,omitempty"`  // 作品集名称
	NS        []string `json:"ns,omitempty"`         // 域名服务器列表（最多13个）
	Coupon    string   `json:"coupon,omitempty"`     // 优惠券代码
	ContactID string   `json:"contact_id,omitempty"` // 联系人 ID
	// WHOIS 联系人信息字段（可选）
	FirstName string `json:"fn,omitempty"`  // 名
	LastName  string `json:"ln,omitempty"`  // 姓
	Address   string `json:"ad,omitempty"`  // 地址
	City      string `json:"cy,omitempty"`  // 城市
	State     string `json:"st,omitempty"`  // 州/省
	Zip       string `json:"zp,omitempty"`  // 邮编
	Country   string `json:"ct,omitempty"`  // 国家代码
	Email     string `json:"em,omitempty"`  // 邮箱
	Phone     string `json:"ph,omitempty"`  // 电话
	Nickname  string `json:"nn,omitempty"`  // 昵称
	Company   string `json:"cp,omitempty"`  // 公司
	Address2  string `json:"ad2,omitempty"` // 地址2
	Fax       string `json:"fx,omitempty"`  // 传真
	// .US 域名专用字段
	USNexusCategory string `json:"usnc,omitempty"` // .US Nexus Category
	USAppPurpose    string `json:"usap,omitempty"` // .US Application Purpose
}

// TransferDomainResponse 转移域名响应
// 使用 CommonReply（包含 Code, Detail, Message, Domain, OrderAmount）
type TransferDomainResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   CommonReply    `xml:"reply" json:"reply"`
}

// ============== Domain Register Drop ==============

// RegisterDomainDropRequest 使用 Drop-Catching 注册域名请求
type RegisterDomainDropRequest struct {
	Domain    string `json:"domain"`               // 要注册的域名
	Years     int    `json:"years"`                // 注册年限
	Private   bool   `json:"private,omitempty"`    // 是否启用隐私保护
	AutoRenew bool   `json:"auto_renew,omitempty"` // 是否自动续费
}

// RegisterDomainDropResponse 使用 Drop-Catching 注册域名响应
// 使用 CommonReply（包含 Code, Detail, Message, Domain, OrderAmount）
type RegisterDomainDropResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   CommonReply    `xml:"reply" json:"reply"`
}

// ============== Domain Register Claims ==============

// TrademarkClaim 商标声明信息
type TrademarkClaim struct {
	ClaimCode    string             `xml:"claim_code" json:"claim_code"`       // 声明代码
	NoticeID     string             `xml:"notice_id" json:"notice_id"`         // 通知 ID
	NotAfter     string             `xml:"not_after" json:"not_after"`         // 有效期
	AcceptedDate string             `xml:"accepted_date" json:"accepted_date"` // 接受日期
	Info         TrademarkClaimInfo `xml:"info" json:"info"`                   // 商标信息
}

// TrademarkClaimInfo 商标信息详情
type TrademarkClaimInfo struct {
	MarkName         string           `xml:"markName" json:"mark_name"`                  // 商标名称
	Holder           TrademarkContact `xml:"holder" json:"holder"`                       // 持有人信息
	Contact          TrademarkContact `xml:"contact" json:"contact"`                     // 联系人信息
	JurDesc          string           `xml:"jurDesc" json:"jur_desc"`                    // 司法管辖描述
	Classes          string           `xml:"classes" json:"classes"`                     // 类别
	GoodsAndServices string           `xml:"goodsAndServices" json:"goods_and_services"` // 商品和服务
	Jurisdiction     string           `xml:"jurisdiction" json:"jurisdiction"`           // 司法管辖区
}

// TrademarkContact 商标联系人信息
type TrademarkContact struct {
	Name    string `xml:"name" json:"name"`       // 名称
	Org     string `xml:"org" json:"org"`         // 组织
	Street  string `xml:"street" json:"street"`   // 街道
	City    string `xml:"city" json:"city"`       // 城市
	State   string `xml:"state" json:"state"`     // 州/省
	Zip     string `xml:"zip" json:"zip"`         // 邮编
	Country string `xml:"country" json:"country"` // 国家
	Phone   string `xml:"phone" json:"phone"`     // 电话
	Fax     string `xml:"fax" json:"fax"`         // 传真
	Email   string `xml:"email" json:"email"`     // 邮箱
}

// RegisterDomainClaimsResponse 带商标声明的域名注册响应（中间响应）
type RegisterDomainClaimsResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   struct {
		Code   int              `xml:"code" json:"code"`
		Detail string           `xml:"detail" json:"detail"`
		Claims []TrademarkClaim `xml:"claims>claim" json:"claims"`
	} `xml:"reply" json:"reply"`
}

// ============== Domain Push ==============

// DomainPushRequest 域名推送请求
type DomainPushRequest struct {
	RecipientLogin string   `json:"recipient_login"` // 接收方登录名
	Domains        []string `json:"domains"`         // 要推送的域名列表
}

// DomainPushResponse 域名推送响应
type DomainPushResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   struct {
		Code   int              `xml:"code" json:"code"`
		Detail string           `xml:"detail" json:"detail"`
		Body   DomainPushResult `xml:"body" json:"body"`
	} `xml:"reply" json:"reply"`
}

// DomainPushResult 域名推送结果
type DomainPushResult struct {
	RecipientLogin      string             `xml:"recipientLogin" json:"recipient_login"`
	DomainsPushStatuses []DomainPushStatus `xml:"domainsPushStatuses>entry" json:"domains_push_statuses"`
}

// DomainPushStatus 域名推送状态
type DomainPushStatus struct {
	Domain  string `xml:"domain" json:"domain"`   // 域名
	Success bool   `xml:"success" json:"success"` // 是否成功
	Error   string `xml:"error" json:"error"`     // 错误信息
}

// ============== Domain Forward ==============

// DomainForwardRequest 域名转发请求
type DomainForwardRequest struct {
	Domain          string `json:"domain"`                     // 域名
	Protocol        string `json:"protocol"`                   // URL 协议 (http 或 https)
	Address         string `json:"address"`                    // 要转发到的网站地址
	Method          string `json:"method"`                     // 转发方法 (301, 302 或 cloaked)
	MetaTitle       string `json:"meta_title,omitempty"`       // META title (仅用于 cloaked)
	MetaDescription string `json:"meta_description,omitempty"` // META description (仅用于 cloaked)
	MetaKeywords    string `json:"meta_keywords,omitempty"`    // META keywords (仅用于 cloaked)
}

// DomainForwardResponse 域名转发响应
// 使用 CommonReply（包含 Code, Detail）
type DomainForwardResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   CommonReply    `xml:"reply" json:"reply"`
}

// ============== Domain Forward SubDomain ==============

// DomainForwardSubDomainRequest 子域名转发请求
type DomainForwardSubDomainRequest struct {
	Domain          string `json:"domain"`                     // 域名
	SubDomain       string `json:"sub_domain"`                 // 子域名
	Protocol        string `json:"protocol"`                   // URL 协议 (http 或 https)
	Address         string `json:"address"`                    // 要转发到的网站地址
	Method          string `json:"method"`                     // 转发方法 (301, 302 或 cloaked)
	MetaTitle       string `json:"meta_title,omitempty"`       // META title (仅用于 cloaked)
	MetaDescription string `json:"meta_description,omitempty"` // META description (仅用于 cloaked)
	MetaKeywords    string `json:"meta_keywords,omitempty"`    // META keywords (仅用于 cloaked)
}

// DomainForwardSubDomainResponse 子域名转发响应
// 使用 CommonReply（包含 Code, Detail, Message）
type DomainForwardSubDomainResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   CommonReply    `xml:"reply" json:"reply"`
}

// ============== Delete Domain Forward SubDomain ==============

// DeleteDomainForwardSubDomainRequest 删除子域名转发请求
type DeleteDomainForwardSubDomainRequest struct {
	Domain    string `json:"domain"`     // 主域名
	SubDomain string `json:"sub_domain"` // 要删除的子域名
}

// DeleteDomainForwardSubDomainResponse 删除子域名转发响应
type DeleteDomainForwardSubDomainResponse struct {
	BaseResponse
}

// ============== Check Transfer Availability ==============

// CheckTransferAvailabilityRequest 检查转移可用性请求
type CheckTransferAvailabilityRequest struct {
	Domains []string `json:"domains"` // 要检查的域名列表
}

// CheckTransferAvailabilityResponse 检查转移可用性响应
type CheckTransferAvailabilityResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   struct {
		Code        int                         `xml:"code" json:"code"`
		Detail      string                      `xml:"detail" json:"detail"`
		Available   []TransferAvailableDomain   `xml:"available>domain" json:"available"`
		Unavailable []TransferUnavailableDomain `xml:"unavailable>domain" json:"unavailable"`
	} `xml:"reply" json:"reply"`
}

// TransferAvailableDomain 可转移的域名信息
type TransferAvailableDomain struct {
	Domain  string  `xml:",chardata" json:"domain"`
	Price   float64 `xml:"price,attr" json:"price,omitempty"`
	Premium int     `xml:"premium,attr" json:"premium,omitempty"`
}

// TransferUnavailableDomain 不可转移的域名信息
type TransferUnavailableDomain struct {
	Domain string `xml:",chardata" json:"domain"`
	Reason string `xml:"reason,attr" json:"reason,omitempty"`
}

// ============== Domain Whois ==============

// WhoisRequest Whois 查询请求
type WhoisRequest struct {
	Domain string `json:"domain"` // 要查询的域名
}

// WhoisResponse Whois 查询响应
type WhoisResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   WhoisReply     `xml:"reply" json:"reply"`
}

// WhoisReply Whois 查询响应内容
type WhoisReply struct {
	Code       int    `xml:"code" json:"code"`
	Detail     string `xml:"detail" json:"detail"`
	Domain     string `xml:"domain" json:"domain"`
	Registered string `xml:"registered" json:"registered"` // "yes" or "no"
	Changed    string `xml:"changed" json:"changed,omitempty"`
	Created    string `xml:"created" json:"created,omitempty"`
	Expires    string `xml:"expires" json:"expires,omitempty"`
	Registrar  string `xml:"registrar" json:"registrar,omitempty"`
}

// ============== Check Register Availability ==============

// CheckRegisterAvailabilityRequest 检查域名注册可用性请求
type CheckRegisterAvailabilityRequest struct {
	Domains []string `json:"domains"` // 要检查的域名列表（逗号分隔）
}

// CheckRegisterAvailabilityResponse 检查域名注册可用性响应
type CheckRegisterAvailabilityResponse struct {
	XMLName xml.Name                       `xml:"namesilo" json:"-"`
	Request RequestSection                 `xml:"request" json:"request"`
	Reply   CheckRegisterAvailabilityReply `xml:"reply" json:"reply"`
}

// CheckRegisterAvailabilityReply 检查域名注册可用性响应内容
type CheckRegisterAvailabilityReply struct {
	Code        int               `xml:"code" json:"code"`
	Detail      string            `xml:"detail" json:"detail"`
	Available   []AvailableDomain `xml:"available>domain" json:"available"`
	Unavailable []string          `xml:"unavailable>domain" json:"unavailable"`
	Invalid     []string          `xml:"invalid>domain" json:"invalid,omitempty"` // 无效的域名列表
}

// AvailableDomain 可用域名信息
type AvailableDomain struct {
	Domain   string  `xml:",chardata" json:"domain"`
	Price    float64 `xml:"price,attr" json:"price,omitempty"`
	Premium  int     `xml:"premium,attr" json:"premium,omitempty"`
	Duration int     `xml:"duration,attr" json:"duration,omitempty"` // 注册时长（年）
}

// ============== Domain Lock/Unlock ==============

// DomainLockRequest 域名锁定请求
type DomainLockRequest struct {
	Domain string `json:"domain"` // 要锁定的域名
}

// DomainLockResponse 域名锁定响应
type DomainLockResponse struct {
	BaseResponse
}

// DomainUnlockRequest 域名解锁请求
type DomainUnlockRequest struct {
	Domain string `json:"domain"` // 要解锁的域名
}

// DomainUnlockResponse 域名解锁响应
type DomainUnlockResponse struct {
	BaseResponse
}

// ============== Auto Renewal ==============

// AddAutoRenewalRequest 添加自动续费请求
type AddAutoRenewalRequest struct {
	Domain string `json:"domain"` // 要设置自动续费的域名
}

// AddAutoRenewalResponse 添加自动续费响应
type AddAutoRenewalResponse struct {
	BaseResponse
}

// RemoveAutoRenewalRequest 移除自动续费请求
type RemoveAutoRenewalRequest struct {
	Domain string `json:"domain"` // 要移除自动续费的域名
}

// RemoveAutoRenewalResponse 移除自动续费响应
type RemoveAutoRenewalResponse struct {
	BaseResponse
}

// ============== Reply Success/Error Methods ==============

// Success 判断响应是否成功
func (r *ListDomainsReply) Success() bool {
	return r.Code == StatusSuccess ||
		r.Code == StatusSuccessWithPartialFailure ||
		r.Code == StatusSuccessWithContactIssue
}

// Error 实现 error 接口
func (r *ListDomainsReply) Error() string {
	if r.Success() {
		return ""
	}
	return fmt.Sprintf("NameSilo API error (code %d): %s", r.Code, r.Detail)
}

// Success 判断响应是否成功
func (r *CheckRegisterAvailabilityReply) Success() bool {
	return r.Code == StatusSuccess ||
		r.Code == StatusSuccessWithPartialFailure ||
		r.Code == StatusSuccessWithContactIssue
}

// Error 实现 error 接口
func (r *CheckRegisterAvailabilityReply) Error() string {
	if r.Success() {
		return ""
	}
	return fmt.Sprintf("NameSilo API error (code %d): %s", r.Code, r.Detail)
}

// Success 判断响应是否成功
func (r *DomainInfoReply) Success() bool {
	return r.Code == StatusSuccess ||
		r.Code == StatusSuccessWithPartialFailure ||
		r.Code == StatusSuccessWithContactIssue
}

// Error 实现 error 接口
func (r *DomainInfoReply) Error() string {
	if r.Success() {
		return ""
	}
	return fmt.Sprintf("NameSilo API error (code %d): %s", r.Code, r.Detail)
}

// Success 判断响应是否成功
func (r *WhoisReply) Success() bool {
	return r.Code == StatusSuccess ||
		r.Code == StatusSuccessWithPartialFailure ||
		r.Code == StatusSuccessWithContactIssue
}

// Error 实现 error 接口
func (r *WhoisReply) Error() string {
	if r.Success() {
		return ""
	}
	return fmt.Sprintf("NameSilo API error (code %d): %s", r.Code, r.Detail)
}

// ============== Domain Recommend ==============

// RecommendDomainsRequest 推荐域名请求
type RecommendDomainsRequest struct {
	Keyword        string   `json:"keyword"`                   // 关键词（必填）
	TLDs           []string `json:"tlds,omitempty"`            // 顶级域名列表（可选）
	MaxPrice       float64  `json:"max_price,omitempty"`       // 最大价格（可选）
	MaxDomains     int      `json:"max_domains,omitempty"`     // 最大域名数量（可选）
	IncludeMatched bool     `json:"include_matched,omitempty"` // 是否包含精确匹配的域名（可选）
}

// RecommendDomainsResponse 推荐域名响应
type RecommendDomainsResponse struct {
	Recommended []AvailableDomain `json:"recommended"` // 推荐的域名列表（基于关键词变体）
	Matched     []AvailableDomain `json:"matched"`     // 匹配的域名列表（基于原始关键词）
	Unavailable []string          `json:"unavailable"` // 不可用的域名列表
}

// FilterType 过滤器类型
type FilterType string

const (
	FilterByPrice  FilterType = "price"  // 按价格过滤
	FilterByTLD    FilterType = "tld"    // 按顶级域名过滤
	FilterByLength FilterType = "length" // 按域名长度过滤
)

// DomainFilter 域名过滤条件（支持多条件组合）
type DomainFilter struct {
	// 价格过滤
	MaxPrice *float64 `json:"max_price,omitempty"` // 最大价格（nil 表示不过滤）
	MinPrice *float64 `json:"min_price,omitempty"` // 最小价格（nil 表示不过滤）

	// TLD 过滤
	IncludeTLDs []string `json:"include_tlds,omitempty"` // 包含的 TLD 列表（空表示不过滤）
	ExcludeTLDs []string `json:"exclude_tlds,omitempty"` // 排除的 TLD 列表（空表示不过滤）

	// 长度过滤
	MaxLength *int `json:"max_length,omitempty"` // 最大域名长度（nil 表示不过滤）
	MinLength *int `json:"min_length,omitempty"` // 最小域名长度（nil 表示不过滤）

	// Premium 过滤
	ExcludePremium bool `json:"exclude_premium,omitempty"` // 是否排除高级域名
}

// CheckAvailabilityFilterRequest 带过滤条件的域名可用性检查请求
type CheckAvailabilityFilterRequest struct {
	Domains []string      `json:"domains"`          // 要检查的域名列表（必填）
	Filter  *DomainFilter `json:"filter,omitempty"` // 过滤条件（可选，支持多条件组合）
}

// CheckAvailabilityFilterResponse 带过滤条件的域名可用性检查响应
type CheckAvailabilityFilterResponse struct {
	Available   []AvailableDomain `json:"available"`   // 符合过滤条件的可用域名
	Unavailable []string          `json:"unavailable"` // 不可用的域名
	Filtered    []AvailableDomain `json:"filtered"`    // 可用但不符合过滤条件的域名
}
