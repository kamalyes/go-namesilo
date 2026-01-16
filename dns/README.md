# 📝 DNS - DNS 记录管理模块

DNS 记录管理模块提供完整的 DNS 资源记录管理功能,支持多种记录类型和 DNSSEC。

## 📋 功能列表

- ✅ 列出所有 DNS 记录
- ✅ 添加 DNS 记录 (A/AAAA/CNAME/MX/TXT/SRV/CAA)
- ✅ 更新 DNS 记录
- ✅ 删除 DNS 记录
- ✅ DNSSEC 管理
  - 列出 DNSSEC 记录
  - 添加 DNSSEC 记录
  - 删除 DNSSEC 记录

## 🚀 快速开始

```go
import (
    "context"
    "github.com/kamalyes/go-namesilo/client"
    "github.com/kamalyes/go-namesilo/dns"
)

// 创建客户端
c, _ := client.New("your-api-key")

// 创建 DNS 服务
dnsService := dns.NewService(c)
ctx := context.Background()
```

## 📖 使用示例

### 列出 DNS 记录

```go
req := &dns.DNSListRecordsRequest{
    Domain: "example.com",
}

resp, err := dnsService.ListRecords(ctx, req)
if err != nil {
    log.Fatal(err)
}

for _, record := range resp.Reply.Records {
    fmt.Printf("ID: %s | Type: %s | Host: %s | Value: %s | TTL: %s\n",
        record.RecordID, record.Type, record.Host, record.Value, record.TTL)
}
```

### 添加 A 记录

```go
req := &dns.DNSAddRecordRequest{
    Domain:   "example.com",
    Type:     dns.RecordTypeA,     // "A"
    Host:     "www",                // www.example.com
    Value:    "192.168.1.1",        // IP 地址
    TTL:      3600,                 // 生存时间 (秒)
    Distance: 0,                    // MX/SRV 记录的优先级
}

resp, err := dnsService.AddRecord(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("记录已添加! ID: %s\n", resp.Reply.RecordID)
```

### 添加 MX 记录

```go
req := &dns.DNSAddRecordRequest{
    Domain:   "example.com",
    Type:     dns.RecordTypeMX,
    Host:     "",                   // 空表示根域名
    Value:    "mail.example.com",   // 邮件服务器
    TTL:      3600,
    Distance: 10,                   // MX 优先级
}

_, err := dnsService.AddRecord(ctx, req)
```

### 添加 CNAME 记录

```go
req := &dns.DNSAddRecordRequest{
    Domain: "example.com",
    Type:   dns.RecordTypeCNAME,
    Host:   "blog",                 // blog.example.com
    Value:  "bloghost.example.com", // 目标域名
    TTL:    7200,
}

_, err := dnsService.AddRecord(ctx, req)
```

### 添加 TXT 记录

```go
req := &dns.DNSAddRecordRequest{
    Domain: "example.com",
    Type:   dns.RecordTypeTXT,
    Host:   "",
    Value:  "v=spf1 include:_spf.example.com ~all", // SPF 记录
    TTL:    3600,
}

_, err := dnsService.AddRecord(ctx, req)
```

### 更新 DNS 记录

```go
req := &dns.DNSUpdateRecordRequest{
    Domain:   "example.com",
    RecordID: "12345678",           // 记录 ID
    Host:     "www",
    Value:    "192.168.1.2",        // 新的 IP 地址
    TTL:      7200,                 // 新的 TTL
    Distance: 0,
}

_, err := dnsService.UpdateRecord(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Println("记录已更新!")
```

### 删除 DNS 记录

```go
req := &dns.DNSDeleteRecordRequest{
    Domain: "example.com",
    RRID:   "12345678",             // 记录 ID
}

_, err := dnsService.DeleteRecord(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Println("记录已删除!")
```

## 🔒 DNSSEC 管理

DNSSEC 为 DNS 提供额外的安全层,防止 DNS 劫持和缓存投毒攻击。

### 列出 DNSSEC 记录

```go
req := &dns.DNSSecListRecordsRequest{
    Domain: "example.com",
}

resp, err := dnsService.ListDNSSecRecords(ctx, req)
if err != nil {
    log.Fatal(err)
}

for _, record := range resp.Reply.Records {
    fmt.Printf("Digest: %s | Algorithm: %s | Type: %s\n",
        record.Digest, record.DigestType, record.Algorithm)
}
```

### 添加 DNSSEC 记录

```go
req := &dns.DNSSecAddRecordRequest{
    Domain:     "example.com",
    Digest:     "ABC123...",         // DS 记录的摘要
    KeyTag:     "12345",             // 密钥标签
    DigestType: "2",                 // 摘要类型 (1=SHA-1, 2=SHA-256)
    Algorithm:  "8",                 // 算法 (8=RSA/SHA-256)
}

_, err := dnsService.AddDNSSecRecord(ctx, req)
if err != nil {
    log.Fatal(err)
}

fmt.Println("DNSSEC 记录已添加!")
```

### 删除 DNSSEC 记录

```go
req := &dns.DNSSecDeleteRecordRequest{
    Domain: "example.com",
    Digest: "ABC123...",             // 要删除的记录摘要
}

_, err := dnsService.DeleteDNSSecRecord(ctx, req)
```

## 📝 DNS 记录类型

### 支持的记录类型

| 类型 | 常量 | 说明 | 示例 |
|------|------|------|------|
| A | `RecordTypeA` | IPv4 地址记录 | `192.168.1.1` |
| AAAA | `RecordTypeAAAA` | IPv6 地址记录 | `2001:0db8::1` |
| CNAME | `RecordTypeCNAME` | 别名记录 | `target.example.com` |
| MX | `RecordTypeMX` | 邮件交换记录 | `mail.example.com` |
| TXT | `RecordTypeTXT` | 文本记录 | `v=spf1 ...` |
| SRV | `RecordTypeSRV` | 服务记录 | `10 5 5060 sip.example.com` |
| CAA | `RecordTypeCAA` | 证书颁发机构授权 | `0 issue "letsencrypt.org"` |

### 记录类型使用场景

**A 记录**: 将域名指向 IPv4 地址
```go
Type: dns.RecordTypeA
Host: "www"
Value: "192.168.1.1"
```

**AAAA 记录**: 将域名指向 IPv6 地址
```go
Type: dns.RecordTypeAAAA
Host: "www"
Value: "2001:0db8::1"
```

**CNAME 记录**: 创建域名别名
```go
Type: dns.RecordTypeCNAME
Host: "blog"
Value: "hosting.example.com"
```

**MX 记录**: 配置邮件服务器
```go
Type: dns.RecordTypeMX
Host: ""
Value: "mail.example.com"
Distance: 10  // 优先级
```

**TXT 记录**: 存储文本信息 (SPF, DKIM, 域名验证等)
```go
Type: dns.RecordTypeTXT
Host: ""
Value: "v=spf1 include:_spf.google.com ~all"
```

**SRV 记录**: 指定服务位置
```go
Type: dns.RecordTypeSRV
Host: "_sip._tcp"
Value: "10 60 5060 sipserver.example.com"
```

**CAA 记录**: 指定可签发证书的 CA
```go
Type: dns.RecordTypeCAA
Host: ""
Value: "0 issue \"letsencrypt.org\""
```

## 🔧 高级功能

### 批量操作示例

```go
// 批量添加记录
records := []struct {
    Host  string
    Value string
}{
    {"www", "192.168.1.1"},
    {"mail", "192.168.1.2"},
    {"ftp", "192.168.1.3"},
}

for _, r := range records {
    req := &dns.DNSAddRecordRequest{
        Domain: "example.com",
        Type:   dns.RecordTypeA,
        Host:   r.Host,
        Value:  r.Value,
        TTL:    3600,
    }
    
    _, err := dnsService.AddRecord(ctx, req)
    if err != nil {
        log.Printf("添加 %s 失败: %v\n", r.Host, err)
        continue
    }
    
    log.Printf("已添加: %s.example.com -> %s\n", r.Host, r.Value)
}
```

### TTL 设置建议

| TTL 值 | 说明 | 适用场景 |
|--------|------|----------|
| 300 (5分钟) | 短 TTL | 频繁更改的记录、迁移期间 |
| 1800 (30分钟) | 中等 TTL | 一般用途 |
| 3600 (1小时) | 标准 TTL | 推荐默认值 |
| 7200 (2小时) | 较长 TTL | 稳定的生产环境 |
| 86400 (24小时) | 长 TTL | 很少更改的记录 |

## ⚠️ 注意事项

1. **DNS 传播**: DNS 记录更改可能需要几分钟到几小时才能在全球传播
2. **TTL 影响**: 较长的 TTL 可以减少 DNS 查询,但更改生效会更慢
3. **CNAME 限制**: CNAME 记录不能与其他记录类型共存于同一主机名
4. **根域名**: 主机名为空字符串 `""` 表示根域名 `@`
5. **MX 优先级**: MX 记录的 Distance 值越小优先级越高
6. **DNSSEC**: 启用 DNSSEC 需要在域名注册商处配置 DS 记录
7. **记录 ID**: 更新和删除操作需要记录 ID,可通过 ListRecords 获取

## 📚 API 文档

详细的 API 文档请参考:
- [NameSilo 官方 API 文档](https://www.namesilo.com/api-reference#dns)
- [GoDoc 文档](https://pkg.go.dev/github.com/kamalyes/go-namesilo/dns)

## 🔗 相关模块

- [Domains 域名管理](../domains/) - 管理域名生命周期
- [Nameserver 域名服务器](../nameserver/) - 管理域名服务器
