/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 13:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:52:15
 * @FilePath: \go-namesilo\domains\check_availability_filter.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"
	"strings"
)

// CheckAvailabilityFilter 检查域名可用性（支持多条件过滤）
// Docs: https://www.namesilo.com/api-reference#domains/check-register-availability
//
// 支持多种过滤条件组合：价格、TLD、长度、Premium 等
// 所有条件采用 AND 逻辑（必须同时满足）
func (s *Service) CheckAvailabilityFilter(ctx context.Context, req *CheckAvailabilityFilterRequest) (*CheckAvailabilityFilterResponse, error) {
	// 检查域名可用性
	checkResp, err := s.CheckAvailability(ctx, &CheckRegisterAvailabilityRequest{
		Domains: req.Domains,
	})
	if err != nil {
		return nil, err
	}

	resp := &CheckAvailabilityFilterResponse{
		Available:   make([]AvailableDomain, 0),
		Unavailable: checkResp.Reply.Unavailable,
		Filtered:    make([]AvailableDomain, 0),
	}

	// 如果没有过滤条件，直接返回所有可用域名
	if req.Filter == nil {
		resp.Available = checkResp.Reply.Available
		return resp, nil
	}

	// 应用过滤条件
	for _, domain := range checkResp.Reply.Available {
		availDomain := AvailableDomain{
			Domain: domain.Domain,
			Price:  domain.Price,
		}

		if matchesFilter(domain, req.Filter) {
			resp.Available = append(resp.Available, availDomain)
		} else {
			resp.Filtered = append(resp.Filtered, availDomain)
		}
	}

	return resp, nil
}

// matchesFilter 检查域名是否匹配所有过滤条件（AND 逻辑）
func matchesFilter(domain AvailableDomain, filter *DomainFilter) bool {
	// 1. 价格过滤
	if filter.MaxPrice != nil && domain.Price > *filter.MaxPrice {
		return false
	}
	if filter.MinPrice != nil && domain.Price < *filter.MinPrice {
		return false
	}

	// 2. Premium 过滤
	if filter.ExcludePremium && domain.Premium == 1 {
		return false
	}

	// 3. TLD 过滤
	tld := extractTLD(domain.Domain)

	// 包含 TLD 列表（白名单）
	if len(filter.IncludeTLDs) > 0 {
		found := false
		for _, allowedTLD := range filter.IncludeTLDs {
			if strings.EqualFold(tld, allowedTLD) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	// 排除 TLD 列表（黑名单）
	if len(filter.ExcludeTLDs) > 0 {
		for _, excludedTLD := range filter.ExcludeTLDs {
			if strings.EqualFold(tld, excludedTLD) {
				return false
			}
		}
	}

	// 4. 长度过滤（不含 TLD 的域名主体长度）
	domainName := extractDomainName(domain.Domain)
	nameLength := len(domainName)

	if filter.MaxLength != nil && nameLength > *filter.MaxLength {
		return false
	}
	if filter.MinLength != nil && nameLength < *filter.MinLength {
		return false
	}

	return true
}

// extractTLD 提取域名的 TLD 部分
func extractTLD(domain string) string {
	parts := strings.Split(domain, ".")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}

// extractDomainName 提取域名主体（不含 TLD）
func extractDomainName(domain string) string {
	parts := strings.Split(domain, ".")
	if len(parts) > 1 {
		return strings.Join(parts[:len(parts)-1], ".")
	}
	return domain
}
