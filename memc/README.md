# Memc - 内存缓存组件

Memc 是一个基于 Go map 实现的高性能内存缓存组件，提供了类似 Redis 的 API 接口，支持过期时间管理、TTL 查询、数字递增递减等功能。

## 特性

- **并发安全**: 使用读写锁保证并发访问安全
- **过期管理**: 支持设置 key 的过期时间，自动清理过期数据
- **多种数据类型**: 支持 `[]byte` 和 `string` 两种数据类型
- **数字操作**: 支持对数字值进行递增和递减操作
- **TTL 查询**: 可以查询 key 的剩余生存时间
- **灵活的过期设置**: 支持多种方式设置过期时间

## 安装

```go
import "github.com/leafney/rose/memc"
```

## 快速开始

```go
package main

import (
    "fmt"
    "time"
    "github.com/leafney/rose/memc"
)

func main() {
    // 创建缓存实例
    cache := memc.NewMemc()
    
    // 基础操作
    cache.Set("key1", []byte("value1"))
    cache.SetS("key2", "value2")
    
    // 获取数据
    value, err := cache.Get("key1")
    if err == nil {
        fmt.Printf("key1: %s\n", string(value))
    }
    
    strValue, err := cache.GetS("key2")
    if err == nil {
        fmt.Printf("key2: %s\n", strValue)
    }
    
    // 设置带过期时间的数据
    cache.SetEx("temp_key", []byte("temp_value"), 5*time.Second)
    cache.SetExSec("temp_key2", []byte("temp_value2"), 10)
    
    // 数字操作
    cache.SetS("counter", "0")
    newValue, _ := cache.Incr("counter")  // 返回 1
    newValue, _ = cache.IncrBy("counter", 5)  // 返回 6
    newValue, _ = cache.Decr("counter")  // 返回 5
}
```

## API 文档

### 基础操作

#### Set(key string, value []byte) error
设置 key-value 对（[]byte 类型）

#### SetS(key string, value string) error
设置 key-value 对（string 类型）

#### Get(key string) ([]byte, error)
获取 key 对应的值（[]byte 类型）

#### GetS(key string) (string, error)
获取 key 对应的值（string 类型）

#### Del(key string) error
删除指定的 key

#### Exists(key string) bool
检查 key 是否存在

### 过期时间操作

#### SetEx(key string, value []byte, expires time.Duration) error
设置带过期时间的 key-value 对

#### SetExSec(key string, value []byte, seconds int64) error
设置带过期时间的 key-value 对（秒为单位）

#### SetExSecS(key string, value string, seconds int64) error
设置带过期时间的 key-value 对（string 类型，秒为单位）

### TTL 和过期管理

#### TTL(key string) (int64, error)
获取 key 的剩余生存时间（秒）
- 返回 -1：key 存在但没有设置过期时间
- 返回 -2：key 不存在
- 返回 0：key 已过期
- 返回正数：剩余秒数

#### Expire(key string, expires time.Duration) error
设置 key 的过期时间

#### ExpireSec(key string, seconds int64) error
设置 key 的过期时间（秒为单位）

#### ExpireAt(key string, tm time.Time) error
设置 key 在指定时间过期

### 数字操作

#### Incr(key string) (int64, error)
将 key 对应的数字值增加 1

#### IncrBy(key string, increment int64) (int64, error)
将 key 对应的数字值增加指定数量

#### Decr(key string) (int64, error)
将 key 对应的数字值减少 1

#### DecrBy(key string, decrement int64) (int64, error)
将 key 对应的数字值减少指定数量

## 注意事项

1. **并发安全**: 所有操作都是并发安全的，可以在多个 goroutine 中安全使用
2. **过期清理**: 过期的 key 会在访问时自动清理，不需要手动清理
3. **数字操作**: 数字操作要求 key 对应的值必须是有效的整数字符串
4. **内存管理**: 这是一个纯内存缓存，重启后数据会丢失

## 性能特点

- 基于 Go 原生 map 实现，读写性能优秀
- 使用读写锁优化并发读取性能
- 惰性删除过期 key，减少后台清理开销
- 内存占用低，适合作为应用内缓存使用

## 使用场景

- 应用内缓存
- 会话存储
- 临时数据存储
- 计数器实现
- 限流器状态存储