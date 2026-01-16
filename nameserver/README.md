# 🔧 Nameserver 域名服务器管理

域名服务器（Name Server）管理模块,提供域名 NS 服务器的增删改查功能。

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
    "github.com/kamalyes/go-namesilo/nameserver"
)

// 创建客户端
c, err := client.New("your-api-key")
if err != nil {
    log.Fatal(err)
}

// 创建域名服务器服务
nsService := nameserver.NewService(c)
ctx := context.Background()
```

## 📚 功能列表

### 1. 修改域名服务器

更改域名使用的 NS 服务器。

```go
req := &nameserver.ChangeNameServersRequest{
    Domain: "example.com",
    NS1:    "ns1.example.com",
    NS2:    "ns2.example.com",
    NS3:    "ns3.example.com", // 可选
    NS4:    "ns4.example.com", // 可选
    NS5:    "ns5.example.com", // 可选
}

resp, err := nsService.ChangeNameServers(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("✅ 域名服务器修改成功\n")
```

### 2. 列出已注册的域名服务器

查询在 NameSilo 注册的自定义 NS 服务器。

```go
req := &nameserver.ListRegisteredNameServersRequest{
    Domain: "example.com",
}

resp, err := nsService.ListRegisteredNameServers(ctx, req)
if err != nil {
    log.Fatal(err)
}

for _, ns := range resp.Hosts {
    fmt.Printf("域名服务器: %s\n", ns.Host)
    for _, ip := range ns.IPs {
        fmt.Printf("  IP: %s\n", ip)
    }
}
```

### 3. 添加域名服务器

注册新的自定义域名服务器。

```go
req := &nameserver.AddRegisteredNameServerRequest{
    Domain: "example.com",
    NewHost: "ns1.example.com",
    IP1:     "192.0.2.1",
    IP2:     "192.0.2.2", // 可选
}

resp, err := nsService.AddRegisteredNameServer(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("✅ 域名服务器 %s 添加成功\n", req.NewHost)
```

### 4. 修改域名服务器 IP

更新已注册域名服务器的 IP 地址。

```go
req := &nameserver.ModifyRegisteredNameServerRequest{
    Domain:      "example.com",
    CurrentHost: "ns1.example.com",
    NewIP1:      "192.0.2.10",
    NewIP2:      "192.0.2.20", // 可选
}

resp, err := nsService.ModifyRegisteredNameServer(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("✅ 域名服务器 IP 修改成功\n")
```

### 5. 删除域名服务器

删除已注册的自定义域名服务器。

```go
req := &nameserver.DeleteRegisteredNameServerRequest{
    Domain:      "example.com",
    CurrentHost: "ns1.example.com",
}

resp, err := nsService.DeleteRegisteredNameServer(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("✅ 域名服务器 %s 已删除\n", req.CurrentHost)
```

## 🎯 错误处理

```go
import namesilo "github.com/kamalyes/go-namesilo"

resp, err := nsService.ChangeNameServers(ctx, req)
if err != nil {
    if namesilo.IsError(err, namesilo.ErrDomainRequired) {
        fmt.Println("域名参数缺失")
    } else if namesilo.IsError(err, namesilo.ErrInvalidDomain) {
        fmt.Println("无效的域名")
    } else {
        fmt.Printf("其他错误: %v\n", err)
    }
    return
}
```

## 📖 相关资源

- [NameSilo NS API 文档](https://www.namesilo.com/api-reference#dns/nameserver-update)
- [返回主文档](../)

## ⚠️ 注意事项

1. **NS 服务器数量**: 至少需要 2 个 NS 服务器,最多支持 5 个
2. **自定义 NS**: 注册自定义 NS 前,确保该域名已在 NameSilo 管理
3. **IP 地址**: 添加/修改 NS 时,至少需要 1 个 IP 地址
4. **生效时间**: NS 修改可能需要 24-48 小时全球生效
5. **删除限制**: 删除 NS 前,确保没有域名正在使用该 NS
