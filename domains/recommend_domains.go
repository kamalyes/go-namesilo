/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 13:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:56:25
 * @FilePath: \go-namesilo\domains\recommend_domains.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"
	"fmt"

	"github.com/kamalyes/go-toolbox/pkg/mathx"
	"github.com/kamalyes/go-toolbox/pkg/random"
)

// RecommendDomains 推荐域名（基于关键词生成推荐域名列表）
// Docs: https://www.namesilo.com/api-reference#domains/check-register-availability
//
// 根据关键词生成多种变体并检查可用性，同时检查原始关键词的可用性
// 自动过滤超过最大价格的高级域名
func (s *Service) RecommendDomains(ctx context.Context, req *RecommendDomainsRequest) (*RecommendDomainsResponse, error) {
	// 设置默认值
	tlds := mathx.IfEmpty(req.TLDs, DefaultTLDs)
	includeMatched := req.IncludeMatched // 默认 false，需要显式设置

	// 1. 生成推荐域名（基于关键词变体）
	recommendKeywords := random.NewDomainKeywordBuilder(req.Keyword).Generate()
	recommendDomains := random.JoinDomainsWithTLDs(recommendKeywords, tlds)

	// 2. 生成匹配域名（原始关键词）
	var allDomains []string
	var originalDomains []string

	if includeMatched {
		originalDomains = random.JoinDomainsWithTLDs([]string{req.Keyword}, tlds)
		allDomains = append(recommendDomains, originalDomains...)
	} else {
		allDomains = recommendDomains
	}

	// 3. 创建原始域名集合（用于快速查找）
	originalDomainSet := make(map[string]bool)
	for _, domain := range originalDomains {
		originalDomainSet[domain] = true
	}

	// 4. 检查所有域名可用性（不提前截断，让API返回所有结果）
	checkResp, err := s.CheckAvailability(ctx, &CheckRegisterAvailabilityRequest{
		Domains: allDomains,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to check domain availability: %w", err)
	}

	// 5. 分离并过滤域名
	resp := &RecommendDomainsResponse{
		Recommended: make([]AvailableDomain, 0),
		Matched:     make([]AvailableDomain, 0),
		Unavailable: checkResp.Reply.Unavailable,
	}

	// 处理可用域名：先按价格过滤，再限制数量
	for _, domain := range checkResp.Reply.Available {
		// 跳过超过最大价格的高级域名
		if req.MaxPrice > 0 && domain.Price > req.MaxPrice {
			continue
		}

		availDomain := AvailableDomain{
			Domain: domain.Domain,
			Price:  domain.Price,
		}
		// 根据域名类型分类
		if originalDomainSet[domain.Domain] {
			resp.Matched = append(resp.Matched, availDomain)
		} else {
			resp.Recommended = append(resp.Recommended, availDomain)
		}
	}

	return resp, nil
}
