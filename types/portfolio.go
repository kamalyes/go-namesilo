/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 01:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 01:00:00
 * @FilePath: \go-namesilo\types\portfolio.go
 * @Description: Portfolio 域名组合相关类型定义
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package types

// ============================================================================
// Portfolio 请求类型
// ============================================================================

// PortfolioListRequest 列出组合请求
type PortfolioListRequest struct {
	// 无参数
}

// PortfolioAddRequest 添加组合请求
type PortfolioAddRequest struct {
	Portfolio string `json:"portfolio" xml:"portfolio"` // 组合名称
}

// PortfolioDeleteRequest 删除组合请求
type PortfolioDeleteRequest struct {
	Portfolio string `json:"portfolio" xml:"portfolio"` // 组合名称
}

// PortfolioDomainAssociateRequest 关联域名到组合请求
type PortfolioDomainAssociateRequest struct {
	Domains   []string `json:"domains" xml:"domains"`     // 域名列表
	Portfolio string   `json:"portfolio" xml:"portfolio"` // 组合名称
}

// ============================================================================
// Portfolio 响应类型
// ============================================================================

// PortfolioListResponse 列出组合响应
type PortfolioListResponse struct {
	BaseResponse
	Reply PortfolioListReply `xml:"reply" json:"reply"`
}

// PortfolioListReply 列出组合响应详情
type PortfolioListReply struct {
	CommonReply
	Portfolios []Portfolio `xml:"portfolios>portfolio" json:"portfolios"` // 组合列表
}

// Portfolio 组合信息
type Portfolio struct {
	Name        string `xml:",chardata" json:"name"`                 // 组合名称
	DomainCount int    `xml:"domain_count,attr" json:"domain_count"` // 域名数量
}

// PortfolioAddResponse 添加组合响应
type PortfolioAddResponse struct {
	BaseResponse
	Reply PortfolioAddReply `xml:"reply" json:"reply"`
}

// PortfolioAddReply 添加组合响应详情
type PortfolioAddReply struct {
	CommonReply
	Portfolio string `xml:"portfolio" json:"portfolio"` // 组合名称
}

// PortfolioDeleteResponse 删除组合响应
type PortfolioDeleteResponse struct {
	BaseResponse
	Reply PortfolioDeleteReply `xml:"reply" json:"reply"`
}

// PortfolioDeleteReply 删除组合响应详情
type PortfolioDeleteReply struct {
	CommonReply
}

// PortfolioDomainAssociateResponse 关联域名到组合响应
type PortfolioDomainAssociateResponse struct {
	BaseResponse
	Reply PortfolioDomainAssociateReply `xml:"reply" json:"reply"`
}

// PortfolioDomainAssociateReply 关联域名到组合响应详情
type PortfolioDomainAssociateReply struct {
	CommonReply
}
