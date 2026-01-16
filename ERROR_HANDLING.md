# 错误处理指南

本项目采用统一的错误处理机制，所有错误定义都集中在 `errors.go` 文件中，并通过各个包的 `aliases.go` 文件导出错误别名。

## 错误组织结构

### 三层架构

1. **中央定义层** (`errors.go`)：所有错误的源头定义
2. **包别名层** (各包的 `aliases.go`)：为每个包导入需要的错误别名
3. **使用层** (业务代码)：直接使用本地的错误别名

### Error 结构体

```go
type Error struct {
    Code    string // 错误代码
    Message string // 错误消息
    Err     error  // 原始错误
}
```

### 错误代码常量

- `ErrCodeInvalidParam`: 无效参数
- `ErrCodeMissingParam`: 缺少必需参数
- `ErrCodeParamOutOfRange`: 参数超出范围
- `ErrCodeAPIRequest`: API 请求失败
- `ErrCodeAPIResponse`: API 响应错误
- `ErrCodeOperationFailed`: 操作失败

## 使用方法

### 1. 在包内使用错误（推荐）

在各个包内部，直接使用本地的错误别名（通过 `aliases.go` 导入）：

```go
package contact

// 在 contact 包中，直接使用 ErrDomainRequired
func (s *Service) SomeFunction(domain string) error {
    if domain == "" {
        return ErrDomainRequired  // 不需要 namesilo. 前缀
    }
    // ...
}
```

### 2. 在包外使用错误

如果需要在包外部引用错误，使用包前缀：

```go
import "github.com/kamalyes/go-namesilo"

func ExternalFunction() error {
    err := someOperation()
    if namesilo.IsError(err, namesilo.ErrDomainRequired) {
        // 处理域名缺失错误
    }
    return err
}
```

### 3. API 操作错误

当 API 调用失败时，使用 `NewAPIError` 创建错误：

```go
if !resp.Reply.Success() {
    return nil, namesilo.NewAPIError("add contact", resp.Reply.Error())
}
```

### 4. 请求错误

HTTP 请求失败时：

```go
resp, err := c.httpClient.Get(fullURL).Do(ctx)
if err != nil {
    return nil, namesilo.NewRequestError(err)
}
```

### 5. 解析错误

响应解析失败时：

```go
if err := xml.Unmarshal(data, v); err != nil {
    return namesilo.NewParseError(err)
}
```

### 6. 错误检查

检查是否为特定错误：

```go
err := someFunction()
if namesilo.IsError(err, namesilo.ErrDomainRequired) {
    // 处理域名缺失错误
}
```

或使用标准库的 `errors.Is`:

```go
import "errors"

if errors.As(err, &namesilo.Error{}) {
    // 处理自定义错误
}
```

## 预定义错误列表

### 通用字段验证

- `ErrDomainRequired`: 域名必需
- `ErrContactIDRequired`: 联系人 ID 必需
- `ErrRecordIDRequired`: 记录 ID 必需
- `ErrOrderNumberRequired`: 订单号必需

### Contact 字段验证

- `ErrFirstNameRequired`: 名字必需
- `ErrLastNameRequired`: 姓氏必需
- `ErrAddressRequired`: 地址必需
- `ErrCityRequired`: 城市必需
- `ErrStateRequired`: 州/省必需
- `ErrZipRequired`: 邮编必需
- `ErrCountryRequired`: 国家必需
- `ErrEmailRequired`: 邮箱必需
- `ErrPhoneRequired`: 电话必需

### Contact 业务错误

- `ErrContactNotFound`: 联系人未找到
- `ErrContactInUse`: 联系人正在使用中
- `ErrInvalidContactData`: 无效的联系人数据

### Domain 相关验证

- `ErrSubDomainRequired`: 子域名必需
- `ErrProtocolRequired`: 协议必需
- `ErrAddressForwardRequired`: 转发地址必需
- `ErrMethodRequired`: 方法必需
- `ErrDomainsRequired`: 域名列表必需

### Domain 业务错误

- `ErrDomainNotFound`: 域名未找到
- `ErrDomainAlreadyExists`: 域名已存在
- `ErrDomainLocked`: 域名已锁定
- `ErrDomainNotAvailable`: 域名不可用
- `ErrInvalidDomain`: 无效的域名
- `ErrDomainTransferDenied`: 域名转移被拒绝
- `ErrInsufficientBalance`: 余额不足

### DNS 相关验证

- `ErrRRTypeRequired`: DNS 记录类型必需
- `ErrRRHostRequired`: DNS 主机必需
- `ErrRRValueRequired`: DNS 值必需
- `ErrDigestRequired`: Digest 必需
- `ErrRecordTypeRequired`: 记录类型必需
- `ErrRecordValueRequired`: 记录值必需

### DNS 业务错误

- `ErrInvalidRecordType`: 无效的记录类型
- `ErrRecordNotFound`: DNS 记录未找到
- `ErrInvalidTTL`: 无效的 TTL 值
- `ErrInvalidPriority`: 无效的优先级值

### Forwarding 相关错误

- `ErrURLRequired`: URL 必需
- `ErrSubdomainRequired`: 子域名必需
- `ErrForwardNotFound`: 转发未找到
- `ErrInvalidURL`: 无效的 URL 格式
- `ErrInvalidEmail`: 无效的邮箱格式
- `ErrForwardListExceedLimit`: 转发列表不能超过 5 个地址

### Nameserver 相关错误

- `ErrNameserverRequired`: 域名服务器必需
- `ErrNameserverNotFound`: 域名服务器未找到
- `ErrInvalidNameserver`: 无效的域名服务器
- `ErrNameserverInUse`: 域名服务器正在使用中
- `ErrInvalidIPAddress`: 无效的 IP 地址
- `ErrIPAddressExceedLimit`: 最多只能添加 13 个 IP 地址

### 范围验证

- `ErrYearsOutOfRange`: 年份必须在 1-10 之间
- `ErrDomainsExceedLimit`: 域名数量不能超过 200
- `ErrDaysCountInvalid`: 天数必须大于 0
- `ErrAmountInvalid`: 金额必须大于 0

### 复杂验证

- `ErrContactRoleRequired`: 至少需要一个联系人角色（registrant、administrative、billing 或 technical）
- `ErrRecipientLoginRequired`: 接收者登录名必需

### API 客户端错误

- `ErrAPIKeyRequired`: API Key 必需
- `ErrInvalidAPIKey`: 无效的 API Key
- `ErrRequestFailed`: API 请求失败
- `ErrInvalidResponse`: 无效的 API 响应

### Account 相关错误

- `ErrInvalidAmount`: 无效的金额
- `ErrPaymentFailed`: 支付失败

## 包级别错误别名

每个包在其 `aliases.go` 文件中导入所需的错误别名。例如：

### contact/aliases.go

```go
package contact

import (
    namesilo "github.com/kamalyes/go-namesilo"
    "github.com/kamalyes/go-namesilo/types"
)

var (
    // 通用错误
    ErrAPIKeyRequired = namesilo.ErrAPIKeyRequired
    ErrDomainRequired = namesilo.ErrDomainRequired

    // Contact 相关错误
    ErrContactIDRequired   = namesilo.ErrContactIDRequired
    ErrFirstNameRequired   = namesilo.ErrFirstNameRequired
    // ... 其他错误
)
```

### dns/aliases.go

```go
package dns

import (
    namesilo "github.com/kamalyes/go-namesilo"
    "github.com/kamalyes/go-namesilo/types"
)

var (
    // 通用错误
    ErrDomainRequired = namesilo.ErrDomainRequired

    // DNS 相关错误
    ErrRecordIDRequired    = namesilo.ErrRecordIDRequired
    ErrRecordTypeRequired  = namesilo.ErrRecordTypeRequired
    // ... 其他错误
)
```

## 使用示例

### 包内部使用（推荐）

```go
package contact

// 通过 aliases.go 导入的错误别名，直接使用
func (s *Service) AddContact(req *ContactAddRequest) error {
    if req.Domain == "" {
        return ErrDomainRequired  // 直接使用，无需前缀
    }
    if req.FirstName == "" {
        return ErrFirstNameRequired
    }
    // ...
}
```

### 包外部使用

```go
package main

import namesilo "github.com/kamalyes/go-namesilo"

func main() {
    err := someOperation()
    if namesilo.IsError(err, namesilo.ErrDomainRequired) {
        // 处理错误
    }
}
```

### API 错误处理

```go
if !resp.Reply.Success() {
    return nil, namesilo.NewAPIError("add contact", resp.Reply.Error())
}
```

## 最佳实践

1. **使用预定义错误**：优先使用 `errors.go` 中定义的错误常量
2. **包内直接使用**：在包内部直接使用错误别名，无需 `namesilo.` 前缀
3. **包外使用前缀**：在包外部引用错误时使用 `namesilo.` 前缀
4. **保持一致性**：所有参数验证使用相同的错误处理模式
5. **包含上下文**：使用 `NewAPIError` 时，提供清晰的操作描述
6. **错误包装**：使用 `WrapError` 保留原始错误信息
7. **错误检查**：使用 `IsError` 或 `errors.Is` 进行类型安全的错误检查
8. **避免硬编码**：不要在业务代码中直接使用 `fmt.Errorf` 创建验证错误

## 添加新错误

如果需要添加新的错误定义，按以下步骤操作：

### 1. 在 errors.go 中添加错误定义

```go
// 在 errors.go 的预定义错误部分添加
var (
    // ... 现有错误
    ErrNewFieldRequired = NewError(ErrCodeMissingParam, "new_field is required")
)
```

### 2. 在相关包的 aliases.go 中添加别名

```go
// 在 contact/aliases.go 中添加
var (
    // ... 现有别名
    ErrNewFieldRequired = namesilo.ErrNewFieldRequired
)
```

### 3. 在业务代码中使用

```go
// 在 contact 包的业务代码中
func (s *Service) SomeFunction(req *SomeRequest) error {
    if req.NewField == "" {
        return ErrNewFieldRequired  // 直接使用
    }
    // ...
}
```

### 4. 更新此文档

在"预定义错误列表"部分添加新错误的说明。

## 错误别名文件清单

各个包的 `aliases.go` 文件中都包含了错误别名：

- `client/aliases.go` - 客户端相关错误
- `contact/aliases.go` - 联系人相关错误
- `dns/aliases.go` - DNS 相关错误
- `domains/aliases.go` - 域名相关错误
- `forwarding/aliases.go` - 转发相关错误
- `nameserver/aliases.go` - 域名服务器相关错误
- `account/aliases.go` - 账户相关错误

## 注意事项

- **所有错误都应该通过 `errors.go` 集中管理**
- **在包内部使用错误时，通过 `aliases.go` 导入，直接使用别名**
- **在包外部引用错误时，使用 `namesilo.` 前缀**
- **不要在业务代码中直接使用 `fmt.Errorf` 创建验证错误**
- **保持错误消息简洁明了，使用英文**
- **错误代码应该具有描述性**
- **新增错误时，同时更新相关包的 `aliases.go`**
