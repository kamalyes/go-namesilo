# 📮 Forwarding 转发管理

转发管理模块,提供域名转发和邮件转发功能。

## 📦 安装

```bash
go get -u github.com/kamalyes/go-namesilo
```

## 🚀 快速开始

### 初始化服务

```go
import (
    "context"
    "github.com/kamalyes/go-namesilo/client"
    "github.com/kamalyes/go-namesilo/forwarding"
)

// 创建客户端
c, err := client.New("your-api-key")
if err != nil {
    log.Fatal(err)
}

// 创建转发服务
fwdService := forwarding.NewService(c)
ctx := context.Background()
```

## 📚 功能列表

### 🌐 域名转发

#### 1. 配置域名转发

将域名转发到目标 URL。

```go
req := &forwarding.ForwardDomainRequest{
    Domain:   "example.com",
    Protocol: "https",        // http 或 https
    Address:  "target.com",
    Method:   "301",          // 301 永久重定向 或 302 临时重定向
}

resp, err := fwdService.ForwardDomain(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("✅ 域名转发配置成功: %s -> %s://%s\n", 
    req.Domain, req.Protocol, req.Address)
```

#### 2. 配置子域名转发

将子域名转发到目标 URL。

```go
req := &forwarding.ForwardSubDomainRequest{
    Domain:    "example.com",
    SubDomain: "blog",        // 子域名前缀
    Protocol:  "https",
    Address:   "blog.target.com",
    Method:    "301",
}

resp, err := fwdService.ForwardSubDomain(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("✅ 子域名转发配置成功: %s.%s -> %s://%s\n", 
    req.SubDomain, req.Domain, req.Protocol, req.Address)
```

#### 3. 删除转发

删除域名或子域名的转发配置。

```go
req := &forwarding.DeleteForwardRequest{
    Domain: "example.com",
    // SubDomain: "blog", // 可选,删除子域名转发时指定
}

resp, err := fwdService.DeleteForward(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("✅ 转发配置已删除\n")
```

### 📧 邮件转发

#### 4. 列出邮件转发

查询域名的所有邮件转发配置。

```go
req := &forwarding.ListEmailForwardsRequest{
    Domain: "example.com",
}

resp, err := fwdService.ListEmailForwards(ctx, req)
if err != nil {
    log.Fatal(err)
}

for _, forward := range resp.Addresses {
    fmt.Printf("📧 %s -> %s\n", forward.Email, forward.ForwardsTo)
}
```

#### 5. 配置邮件转发

添加或更新邮件转发规则。

```go
req := &forwarding.ConfigureEmailForwardRequest{
    Domain:     "example.com",
    Email:      "info",           // 邮箱前缀
    ForwardsTo: "real@email.com", // 转发目标
}

resp, err := fwdService.ConfigureEmailForward(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("✅ 邮件转发配置成功: %s@%s -> %s\n", 
    req.Email, req.Domain, req.ForwardsTo)
```

#### 6. 删除邮件转发

删除指定的邮件转发规则。

```go
req := &forwarding.DeleteEmailForwardRequest{
    Domain: "example.com",
    Email:  "info", // 要删除的邮箱前缀
}

resp, err := fwdService.DeleteEmailForward(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("✅ 邮件转发已删除: %s@%s\n", req.Email, req.Domain)
```

## 💡 实用示例

### 批量配置邮件转发

```go
emails := map[string]string{
    "info":    "contact@company.com",
    "support": "help@company.com",
    "sales":   "sales@company.com",
}

for email, target := range emails {
    req := &forwarding.ConfigureEmailForwardRequest{
        Domain:     "example.com",
        Email:      email,
        ForwardsTo: target,
    }
    
    if _, err := fwdService.ConfigureEmailForward(ctx, req); err != nil {
        log.Printf("❌ 配置 %s 失败: %v\n", email, err)
        continue
    }
    fmt.Printf("✅ %s@example.com -> %s\n", email, target)
}
```

### 域名转发到 HTTPS

```go
// 将旧域名重定向到新域名
req := &forwarding.ForwardDomainRequest{
    Domain:   "old-domain.com",
    Protocol: "https",
    Address:  "new-domain.com",
    Method:   "301", // 永久重定向,SEO 友好
}

resp, err := fwdService.ForwardDomain(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Println("✅ 域名永久重定向配置成功")
```

## 🎯 错误处理

```go
import namesilo "github.com/kamalyes/go-namesilo"

resp, err := fwdService.ForwardDomain(ctx, req)
if err != nil {
    if namesilo.IsError(err, namesilo.ErrDomainRequired) {
        fmt.Println("域名参数缺失")
    } else if namesilo.IsError(err, namesilo.ErrInvalidProtocol) {
        fmt.Println("无效的协议(仅支持 http/https)")
    } else {
        fmt.Printf("其他错误: %v\n", err)
    }
    return
}
```

## 📖 相关资源

- [NameSilo Forwarding API 文档](https://www.namesilo.com/api-reference#forwarding)
- [返回主文档](../)

## ⚠️ 注意事项

1. **协议选择**: 仅支持 `http` 和 `https` 协议
2. **重定向方法**: 
   - `301` - 永久重定向,推荐用于 SEO
   - `302` - 临时重定向
3. **邮件转发**: 
   - 源邮箱仅需提供前缀(不含 @domain.com)
   - 目标邮箱必须是完整的邮箱地址
4. **生效时间**: 配置后立即生效,但 DNS 传播可能需要时间
5. **数量限制**: 每个域名的邮件转发数量可能有限制
6. **子域名转发**: 配置子域名转发前,确保该子域名存在对应的 DNS 记录
