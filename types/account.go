/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:15:57
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 11:15:12
 * @FilePath: \go-namesilo\types\account.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package types

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

// ============== Get Prices ==============

// GetPricesRequest 获取价格列表请求
type GetPricesRequest struct {
	RetailPrices        bool `json:"retail_prices,omitempty"`        // 是否返回零售价格
	RegistrationDomains int  `json:"registration_domains,omitempty"` // 基于注册域名数量的价格
}

// GetPricesResponse 获取价格列表响应
type GetPricesResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   PricesReply    `xml:"reply" json:"reply"`
}

// PricesReply 价格信息响应
type PricesReply struct {
	Code    int                    `xml:"code" json:"code"`
	Detail  string                 `xml:"detail" json:"detail"`
	Message string                 `xml:"message" json:"message"`
	Prices  map[string]PriceDetail `xml:"-" json:"prices"` // TLD 价格信息
	Extra   map[string]interface{} `json:"-" xml:"-"`      // 额外字段
}

// PriceDetail TLD 价格详情
type PriceDetail struct {
	Registration float64 `json:"registration"` // 注册价格
	Transfer     float64 `json:"transfer"`     // 转移价格
	Renew        float64 `json:"renew"`        // 续费价格
	Restore      float64 `json:"restore"`      // 恢复价格
}

// UnmarshalJSON 自定义 JSON 解析
func (p *PricesReply) UnmarshalJSON(data []byte) error {
	type alias struct {
		Code    json.RawMessage `json:"code"`
		Detail  string          `json:"detail"`
		Message string          `json:"message"`
	}

	var aux alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// 解析 code 字段
	if len(aux.Code) > 0 {
		var codeInt int
		if err := json.Unmarshal(aux.Code, &codeInt); err == nil {
			p.Code = codeInt
		} else {
			var codeStr string
			if err := json.Unmarshal(aux.Code, &codeStr); err != nil {
				return fmt.Errorf("failed to parse code field: %w", err)
			}
			codeStr = strings.TrimSpace(codeStr)
			if codeStr == "" {
				return fmt.Errorf("code field is empty")
			}
			code, err := strconv.Atoi(codeStr)
			if err != nil {
				return fmt.Errorf("failed to parse code field: %w", err)
			}
			p.Code = code
		}
	}

	p.Detail = aux.Detail
	p.Message = aux.Message

	// 解析价格信息
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"code": {}, "detail": {}, "message": {},
	}

	p.Prices = make(map[string]PriceDetail)
	p.Extra = make(map[string]interface{})

	for key, value := range raw {
		if _, ok := knownKeys[key]; ok {
			continue
		}

		// 尝试解析为价格详情
		var nestedObj map[string]interface{}
		if err := json.Unmarshal(value, &nestedObj); err == nil {
			if hasPriceFields(nestedObj) {
				priceDetail := parsePriceDetail(nestedObj)
				if priceDetail.Registration > 0 || priceDetail.Renew > 0 || priceDetail.Transfer > 0 {
					p.Prices[key] = priceDetail
					continue
				}
			}
		}

		// 否则作为额外字段存储
		var v interface{}
		if err := json.Unmarshal(value, &v); err != nil {
			p.Extra[key] = string(value)
		} else {
			p.Extra[key] = v
		}
	}

	return nil
}

// hasPriceFields 检查 map 是否包含价格相关字段
func hasPriceFields(m map[string]interface{}) bool {
	priceKeys := []string{"registration", "transfer", "renew", "restore"}
	for _, key := range priceKeys {
		if _, ok := m[key]; ok {
			return true
		}
	}
	return false
}

// parsePriceDetail 从 map 中解析价格详情
func parsePriceDetail(m map[string]interface{}) PriceDetail {
	var detail PriceDetail

	parsePrice := func(val interface{}) float64 {
		switch v := val.(type) {
		case float64:
			return v
		case string:
			if f, err := strconv.ParseFloat(v, 64); err == nil {
				return f
			}
		case json.Number:
			if f, err := v.Float64(); err == nil {
				return f
			}
		}
		return 0
	}

	if val, ok := m["registration"]; ok {
		detail.Registration = parsePrice(val)
	}
	if val, ok := m["transfer"]; ok {
		detail.Transfer = parsePrice(val)
	}
	if val, ok := m["renew"]; ok {
		detail.Renew = parsePrice(val)
	}
	if val, ok := m["restore"]; ok {
		detail.Restore = parsePrice(val)
	}

	return detail
}

// ============== Get Account Balance ==============

// GetAccountBalanceRequest 获取账户余额请求
type GetAccountBalanceRequest struct {
	// 无需参数
}

// GetAccountBalanceResponse 获取账户余额响应
type GetAccountBalanceResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   struct {
		Code    int    `xml:"code" json:"code"`
		Detail  string `xml:"detail" json:"detail"`
		Balance string `xml:"balance" json:"balance"` // 账户余额
	} `xml:"reply" json:"reply"`
}

// ============== Add Account Funds ==============

// AddAccountFundsRequest 添加账户资金请求
type AddAccountFundsRequest struct {
	Amount    float64 `json:"amount"`               // 要添加的金额
	PaymentID string  `json:"payment_id,omitempty"` // 支付 ID（可选）
}

// AddAccountFundsResponse 添加账户资金响应
type AddAccountFundsResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   struct {
		Code       int     `xml:"code" json:"code"`
		Detail     string  `xml:"detail" json:"detail"`
		NewBalance float64 `xml:"new_balance" json:"new_balance"` // 新的账户余额
	} `xml:"reply" json:"reply"`
}

// ============== List Orders ==============

// ListOrdersRequest 列出订单请求
type ListOrdersRequest struct {
	DateFrom string `json:"date_from,omitempty"` // 起始日期（可选）格式：YYYY-MM-DD HH:MM:SS
	DateTo   string `json:"date_to,omitempty"`   // 结束日期（可选）格式：YYYY-MM-DD HH:MM:SS
}

// ListOrdersResponse 列出订单响应
type ListOrdersResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   struct {
		Code   int     `xml:"code" json:"code"`
		Detail string  `xml:"detail" json:"detail"`
		Orders []Order `xml:"order" json:"orders"`
	} `xml:"reply" json:"reply"`
}

// Order 订单信息
type Order struct {
	OrderNumber string  `xml:"order_number" json:"order_number"` // 订单号
	OrderDate   string  `xml:"order_date" json:"order_date"`     // 订单日期
	Method      string  `xml:"method" json:"method"`             // 支付方式
	Total       float64 `xml:"total" json:"total"`               // 总金额
}

// ============== Order Details ==============

// OrderDetailsRequest 获取订单详情请求
type OrderDetailsRequest struct {
	OrderNumber string `json:"order_number"` // 订单号
}

// OrderDetailsResponse 获取订单详情响应
type OrderDetailsResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   struct {
		Code        int           `xml:"code" json:"code"`
		Detail      string        `xml:"detail" json:"detail"`
		OrderNumber string        `xml:"order_number" json:"order_number"` // 订单号
		OrderDate   string        `xml:"order_date" json:"order_date"`     // 订单日期
		Method      string        `xml:"method" json:"method"`             // 支付方式
		Total       float64       `xml:"total" json:"total"`               // 总金额
		OrderDetail []OrderDetail `xml:"order_details" json:"order_details"`
	} `xml:"reply" json:"reply"`
}

// OrderDetail 订单详情项
type OrderDetail struct {
	Description    string  `xml:"description" json:"description"`         // 描述
	YearsQty       int     `xml:"years_qty" json:"years_qty"`             // 年数
	Price          float64 `xml:"price" json:"price"`                     // 单价
	Subtotal       float64 `xml:"subtotal" json:"subtotal"`               // 小计
	Status         string  `xml:"status" json:"status"`                   // 状态
	CreditedDate   string  `xml:"credited_date" json:"credited_date"`     // 退款日期（可选）
	CreditedAmount float64 `xml:"credited_amount" json:"credited_amount"` // 退款金额（可选）
}

// ============== List Expiring Domains ==============

// ListExpiringDomainsRequest 列出即将过期的域名请求
type ListExpiringDomainsRequest struct {
	DaysCount int `json:"days_count"`          // 天数
	Page      int `json:"page,omitempty"`      // 页码（可选）
	PageSize  int `json:"page_size,omitempty"` // 每页数量（可选）
}

// ListExpiringDomainsResponse 列出即将过期的域名响应
type ListExpiringDomainsResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   struct {
		Code   int              `xml:"code" json:"code"`
		Detail string           `xml:"detail" json:"detail"`
		Body   []ExpiringDomain `xml:"body>entry" json:"body"`
	} `xml:"reply" json:"reply"`
}

// ExpiringDomain 即将过期的域名信息
type ExpiringDomain struct {
	ID        int    `xml:"id" json:"id"`                // ID
	Domain    string `xml:"domain" json:"domain"`        // 域名
	CreatedOn string `xml:"createdOn" json:"created_on"` // 创建日期
	ExpiresOn string `xml:"expiresOn" json:"expires_on"` // 过期日期
	Status    int    `xml:"status" json:"status"`        // 状态
	NSServers string `xml:"nsServers" json:"ns_servers"` // 域名服务器
}

// ============== Count Expiring Domains ==============

// CountExpiringDomainsRequest 统计即将过期的域名请求
type CountExpiringDomainsRequest struct {
	DaysCount int `json:"days_count"` // 天数
}

// CountExpiringDomainsResponse 统计即将过期的域名响应
type CountExpiringDomainsResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   struct {
		Code   int    `xml:"code" json:"code"`
		Detail string `xml:"detail" json:"detail"`
		Body   int    `xml:"body" json:"body"` // 数量
	} `xml:"reply" json:"reply"`
}
