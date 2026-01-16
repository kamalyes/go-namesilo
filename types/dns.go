/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:55:29
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 11:55:55
 * @FilePath: \go-namesilo\types\dns.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package types

import "encoding/xml"

// DNS 记录类型常量
const (
	RecordTypeA     = "A"
	RecordTypeAAAA  = "AAAA"
	RecordTypeCNAME = "CNAME"
	RecordTypeMX    = "MX"
	RecordTypeTXT   = "TXT"
	RecordTypeSRV   = "SRV"
	RecordTypeCAA   = "CAA"
)

// ============== DNS List Records ==============

// DNSListRecordsRequest 列出 DNS 记录请求
type DNSListRecordsRequest struct {
	Domain string `json:"domain"` // 要查询的域名
}

// DNSListRecordsResponse 列出 DNS 记录响应
type DNSListRecordsResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   struct {
		Code            int         `xml:"code" json:"code"`
		Detail          string      `xml:"detail" json:"detail"`
		ResourceRecords []DNSRecord `xml:"resource_record" json:"resource_records"`
	} `xml:"reply" json:"reply"`
}

// DNSRecord DNS 记录
type DNSRecord struct {
	RecordID string `xml:"record_id" json:"record_id"` // 记录 ID
	Type     string `xml:"type" json:"type"`           // 记录类型（A, AAAA, CNAME, MX, TXT）
	Host     string `xml:"host" json:"host"`           // 主机名
	Value    string `xml:"value" json:"value"`         // 记录值
	TTL      int    `xml:"ttl" json:"ttl"`             // TTL
	Distance int    `xml:"distance" json:"distance"`   // MX 优先级（仅 MX 记录）
}

// ============== DNS Add Record ==============

// DNSAddRecordRequest 添加 DNS 记录请求
type DNSAddRecordRequest struct {
	Domain   string `json:"domain"`               // 要更新的域名
	Type     string `json:"rrtype"`               // 记录类型（A, AAAA, CNAME, MX, TXT, SRV, CAA）
	Host     string `json:"rrhost"`               // 主机名（无需包含域名）
	Value    string `json:"rrvalue"`              // 记录值
	Distance int    `json:"rrdistance,omitempty"` // MX 优先级（仅 MX 记录，默认 10）
	TTL      int    `json:"rrttl,omitempty"`      // TTL（默认 7207）
}

// DNSAddRecordResponse 添加 DNS 记录响应
type DNSAddRecordResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   struct {
		Code     int    `xml:"code" json:"code"`
		Detail   string `xml:"detail" json:"detail"`
		RecordID string `xml:"record_id" json:"record_id"` // 新创建的记录 ID
	} `xml:"reply" json:"reply"`
}

// ============== DNS Update Record ==============

// DNSUpdateRecordRequest 更新 DNS 记录请求
type DNSUpdateRecordRequest struct {
	Domain   string `json:"domain"`               // 域名
	RecordID string `json:"rrid"`                 // 要修改的记录 ID
	Host     string `json:"rrhost"`               // 主机名（无需包含域名）
	Value    string `json:"rrvalue"`              // 记录值
	Distance int    `json:"rrdistance,omitempty"` // MX 优先级（仅 MX 记录，默认 10）
	TTL      int    `json:"rrttl,omitempty"`      // TTL（默认 7207）
}

// DNSUpdateRecordResponse 更新 DNS 记录响应
type DNSUpdateRecordResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   struct {
		Code     int    `xml:"code" json:"code"`
		Detail   string `xml:"detail" json:"detail"`
		RecordID string `xml:"record_id" json:"record_id"` // 更新后的新记录 ID
	} `xml:"reply" json:"reply"`
}

// ============== DNS Delete Record ==============

// DNSDeleteRecordRequest 删除 DNS 记录请求
type DNSDeleteRecordRequest struct {
	Domain string `json:"domain"` // 域名
	RRID   string `json:"rrid"`   // 记录 ID
}

// DNSDeleteRecordResponse 删除 DNS 记录响应
type DNSDeleteRecordResponse struct {
	BaseResponse
}

// ============== DNSSEC List Records ==============

// DNSSecListRecordsRequest 列出 DNSSEC 记录请求
type DNSSecListRecordsRequest struct {
	Domain string `json:"domain"` // 要查询的域名
}

// DNSSecListRecordsResponse 列出 DNSSEC 记录响应
type DNSSecListRecordsResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   struct {
		Code      int              `xml:"code" json:"code"`
		Detail    string           `xml:"detail" json:"detail"`
		DSRecords []DNSSecDSRecord `xml:"ds_record" json:"ds_records"`
	} `xml:"reply" json:"reply"`
}

// DNSSecDSRecord DNSSEC DS 记录
type DNSSecDSRecord struct {
	Digest     string `xml:"digest" json:"digest"`           // 摘要
	DigestType int    `xml:"digest_type" json:"digest_type"` // 摘要类型
	Algorithm  int    `xml:"algorithm" json:"algorithm"`     // 算法
	KeyTag     int    `xml:"key_tag" json:"key_tag"`         // 密钥标签
}

// ============== DNSSEC Add Record ==============

// DNSSecAddRecordRequest 添加 DNSSEC 记录请求
type DNSSecAddRecordRequest struct {
	Domain     string `json:"domain"`     // 域名
	Digest     string `json:"digest"`     // 摘要
	KeyTag     int    `json:"keyTag"`     // 密钥标签
	DigestType int    `json:"digestType"` // 摘要类型
	Algorithm  int    `json:"alg"`        // 算法
}

// DNSSecAddRecordResponse 添加 DNSSEC 记录响应
type DNSSecAddRecordResponse struct {
	BaseResponse
}

// ============== DNSSEC Delete Record ==============

// DNSSecDeleteRecordRequest 删除 DNSSEC 记录请求
type DNSSecDeleteRecordRequest struct {
	Domain     string `json:"domain"`     // 域名
	Digest     string `json:"digest"`     // 摘要
	KeyTag     int    `json:"keyTag"`     // 密钥标签
	DigestType int    `json:"digestType"` // 摘要类型
	Algorithm  int    `json:"alg"`        // 算法
}

// DNSSecDeleteRecordResponse 删除 DNSSEC 记录响应
type DNSSecDeleteRecordResponse struct {
	BaseResponse
}
