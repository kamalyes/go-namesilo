# 🌐 Domains - 域名管理模块

域名管理模块提供完整的域名生命周期管理功能,包括注册、续费、转移、查询、锁定等操作。

## 📋 功能列表

- ✅ 检查域名注册可用性
- ✅ 注册新域名
- ✅ 续费域名
- ✅ 域名转移
- ✅ 检查域名转移可用性
- ✅ 域名锁定/解锁
- ✅ 域名转发配置
- ✅ 子域名转发配置
- ✅ 删除子域名转发
- ✅ 域名推送
- ✅ 查询域名信息
- ✅ 列出所有域名
- ✅ WHOIS 查询
- ✅ 自动续费设置
- ✅ 域名 Drop Catch 注册

## 🚀 快速开始

```go
import (
    "context"
    "github.com/kamalyes/go-namesilo/client"
    "github.com/kamalyes/go-namesilo/domains"
)

// 创建客户端
c, _ := client.New("your-api-key")

// 创建域名服务
domainService := domains.NewService(c)
ctx := context.Background()
```

## 📖 使用示例

### 检查域名可用性

```go
req := &domains.CheckRegisterAvailabilityRequest{
    Domains: []string{"example.com", "test.com", "mysite.org"},
}

resp, err := domainService.CheckAvailability(ctx, req)
if err != nil {
    log.Fatal(err)
}

// 可注册的域名
for _, domain := range resp.Available {
    fmt.Printf("✅ %s 可注册 - 价格: $%s\n", domain.Domain, domain.Price)
}

// 不可注册的域名
for _, domain := range resp.Unavailable {
    fmt.Printf("❌ %s 不可注册 - 原因: %s\n", domain.Domain, domain.Reason)
}
```

### 注册域名

```go
req := &domains.RegisterDomainRequest{
    Domain:    "example.com",
    Years:     1,                    // 注册年限 (1-10)
    Private:   1,                    // 启用 WHOIS 隐私保护
    AutoRenew: 1,                    // 启用自动续费
    ContactID: "12345",              // 联系人ID (可选)
    // 或者直接提供联系人信息
    FirstName: "John",
    LastName:  "Doe",
    Address:   "123 Main St",
    City:      "New York",
    State:     "NY",
    Zip:       "10001",
    Country:   "US",
    Email:     "john@example.com",
    Phone:     "+12125551234",
}

resp, err := domainService.RegisterDomain(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("域名注册成功! 订单金额: $%.2f\n", resp.Reply.OrderAmount)
fmt.Printf("域名: %s\n", resp.Reply.Domain)
```

### 续费域名

```go
req := &domains.RenewDomainRequest{
    Domain: "example.com",
    Years:  1,              // 续费年限 (1-10)
}

resp, err := domainService.RenewDomain(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("域名续费成功! 费用: $%.2f\n", resp.Reply.OrderAmount)
```

### 域名锁定/解锁

```go
// 锁定域名 (防止未授权转移)
lockReq := &domains.DomainLockRequest{
    Domain: "example.com",
}
_, err := domainService.Lock(ctx, lockReq)

// 解锁域名
unlockReq := &domains.DomainUnlockRequest{
    Domain: "example.com",
}
_, err = domainService.Unlock(ctx, unlockReq)
```

### 查询域名信息

```go
req := &domains.GetDomainInfoRequest{
    Domain: "example.com",
}

resp, err := domainService.GetDomainInfo(ctx, req)
if err != nil {
    log.Fatal(err)
}

info := resp.Reply
fmt.Printf("域名: %s\n", info.Domain)
fmt.Printf("创建日期: %s\n", info.Created)
fmt.Printf("过期日期: %s\n", info.Expires)
fmt.Printf("状态: %s\n", info.Status)
fmt.Printf("是否锁定: %v\n", info.Locked == "Yes")
fmt.Printf("是否私有: %v\n", info.Private == "Yes")
fmt.Printf("是否自动续费: %v\n", info.AutoRenew == "Yes")
```

### 列出所有域名

```go
req := &domains.ListDomainsRequest{
    Page:     1,
    PageSize: 20,
}

resp, err := domainService.List(ctx, req)
if err != nil {
    log.Fatal(err)
}

for _, domain := range resp.Reply.Domains {
    fmt.Printf("域名: %s (创建: %s, 过期: %s)\n", 
        domain.Name, domain.Created, domain.Expires)
}
```

### 域名转发

```go
req := &domains.DomainForwardRequest{
    Domain:   "example.com",
    Protocol: "https",                    // http 或 https
    Address:  "https://newsite.com",      // 目标地址
    Method:   "301",                      // 301 或 302
}

_, err := domainService.ForwardDomain(ctx, req)
if err != nil {
    log.Fatal(err)
}
```

### 子域名转发

```go
req := &domains.DomainForwardSubDomainRequest{
    Domain:    "example.com",
    SubDomain: "blog",                    // blog.example.com
    Protocol:  "https",
    Address:   "https://blog.newsite.com",
    Method:    "301",
}

_, err := domainService.ForwardSubDomain(ctx, req)
```

### 域名转移

```go
req := &domains.TransferDomainRequest{
    Domain:    "example.com",
    Auth:      "transfer-auth-code",      // EPP 授权码
    Private:   true,                      // 启用隐私保护
    AutoRenew: true,                      // 启用自动续费
    ContactID: "12345",                   // 可选
}

resp, err := domainService.Transfer(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("转移已提交! 订单金额: $%.2f\n", resp.Reply.OrderAmount)
```

### 检查域名转移可用性

```go
req := &domains.CheckTransferAvailabilityRequest{
    Domains: []string{"example.com", "test.com"},
}

resp, err := domainService.CheckTransferAvailability(ctx, req)
if err != nil {
    log.Fatal(err)
}

// 可转移的域名
for _, domain := range resp.Reply.Available {
    fmt.Printf("✅ %s 可转移 - 价格: $%s\n", domain.Domain, domain.Price)
}

// 不可转移的域名
for _, domain := range resp.Reply.Unavailable {
    fmt.Printf("❌ %s 不可转移 - 原因: %s\n", domain.Domain, domain.Reason)
}
```

### WHOIS 查询

```go
req := &domains.WhoisRequest{
    Domain: "example.com",
}

resp, err := domainService.Whois(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("WHOIS 信息:\n%s\n", resp.Reply.Whois)
```

### 自动续费设置

```go
// 启用自动续费
addReq := &domains.AddAutoRenewalRequest{
    Domain: "example.com",
}
_, err := domainService.AddAutoRenewal(ctx, addReq)

// 移除自动续费
removeReq := &domains.RemoveAutoRenewalRequest{
    Domain: "example.com",
}
_, err = domainService.RemoveAutoRenewal(ctx, removeReq)
```

### 域名推送

```go
req := &domains.DomainPushRequest{
    Domain:   "example.com",
    Recipient: "recipient@example.com",   // 接收方邮箱
}

resp, err := domainService.Push(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("推送状态: %s\n", resp.Reply.Push.Status)
```

## 🔧 高级功能

### Drop Catch 注册

Drop Catch 允许您在域名释放时自动注册。

```go
req := &domains.RegisterDomainDropRequest{
    Domain:    "expired-domain.com",
    Years:     1,
    Private:   1,
    AutoRenew: 1,
    ContactID: "12345",
}

resp, err := domainService.RegisterDomainDrop(ctx, req)
```

## ⚠️ 注意事项

1. **年限限制**: 域名注册/续费年限为 1-10 年
2. **转移授权码**: 域名转移需要从原注册商获取 EPP 授权码
3. **锁定状态**: 锁定的域名无法转移,需先解锁
4. **联系人信息**: 注册域名时必须提供有效的联系人信息或 ContactID
5. **自动续费**: 建议为重要域名启用自动续费,避免过期
6. **WHOIS 隐私**: 启用隐私保护可隐藏 WHOIS 中的个人信息
7. **资金要求**: 注册和转移操作需要账户有足够余额或已验证信用卡

## 📚 API 文档

详细的 API 文档请参考:
- [NameSilo 官方 API 文档](https://www.namesilo.com/api-reference#domains)
- [GoDoc 文档](https://pkg.go.dev/github.com/kamalyes/go-namesilo/domains)

## 🔗 相关模块

- [Privacy 隐私保护](../privacy/) - 管理域名 WHOIS 隐私
- [Transfer 转移管理](../transfer/) - 管理域名转移流程
- [Contact 联系人管理](../contact/) - 管理域名联系人
- [DNS 记录管理](../dns/) - 管理域名 DNS 记录
