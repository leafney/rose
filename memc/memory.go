package memc

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

// Memc 内存缓存结构体
type Memc struct {
	mu      sync.RWMutex
	data    map[string][]byte
	expires map[string]time.Time
}

// NewMemc 创建新的内存缓存实例
func NewMemc() *Memc {
	return &Memc{
		data:    make(map[string][]byte),
		expires: make(map[string]time.Time),
	}
}

// isExpired 检查key是否过期
func (m *Memc) isExpired(key string) bool {
	if expireTime, exists := m.expires[key]; exists {
		return time.Now().After(expireTime)
	}
	return false
}

// cleanExpired 清理过期的key
func (m *Memc) cleanExpired(key string) {
	if m.isExpired(key) {
		delete(m.data, key)
		delete(m.expires, key)
	}
}

// Set 设置key-value对（[]byte类型）
func (m *Memc) Set(key string, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.data[key] = value
	// 清除过期时间
	delete(m.expires, key)
	return nil
}

// SetS 设置key-value对（string类型）
func (m *Memc) SetS(key string, value string) error {
	return m.Set(key, []byte(value))
}

// Get 获取key对应的值（[]byte类型）
func (m *Memc) Get(key string) ([]byte, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	// 检查并清理过期key
	m.cleanExpired(key)
	
	value, exists := m.data[key]
	if !exists {
		return nil, errors.New("key not found")
	}
	return value, nil
}

// GetS 获取key对应的值（string类型）
func (m *Memc) GetS(key string) (string, error) {
	value, err := m.Get(key)
	if err != nil {
		return "", err
	}
	return string(value), nil
}

// Del 删除指定key
func (m *Memc) Del(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	delete(m.data, key)
	delete(m.expires, key)
	return nil
}

// Exists 检查key是否存在
func (m *Memc) Exists(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	// 检查并清理过期key
	m.cleanExpired(key)
	
	_, exists := m.data[key]
	return exists
}

// SetEx 设置带过期时间的key-value对
func (m *Memc) SetEx(key string, value []byte, expires time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.data[key] = value
	m.expires[key] = time.Now().Add(expires)
	return nil
}

// SetExSec 设置带过期时间的key-value对（秒为单位）
func (m *Memc) SetExSec(key string, value []byte, seconds int64) error {
	return m.SetEx(key, value, time.Duration(seconds)*time.Second)
}

// SetExSecS 设置带过期时间的key-value对（string类型，秒为单位）
func (m *Memc) SetExSecS(key string, value string, seconds int64) error {
	return m.SetExSec(key, []byte(value), seconds)
}

// TTL 获取key的剩余生存时间（秒）
func (m *Memc) TTL(key string) (int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	// 检查并清理过期key
	m.cleanExpired(key)
	
	// 检查key是否存在
	if _, exists := m.data[key]; !exists {
		return -2, errors.New("key not found")
	}
	
	// 检查是否设置了过期时间
	expireTime, hasExpire := m.expires[key]
	if !hasExpire {
		return -1, nil // -1表示永不过期
	}
	
	// 计算剩余时间
	remaining := expireTime.Sub(time.Now())
	if remaining <= 0 {
		return 0, nil // 已过期
	}
	
	// 返回剩余秒数，至少返回1秒如果还有剩余时间
	seconds := int64(remaining.Seconds())
	if seconds == 0 && remaining > 0 {
		seconds = 1
	}
	return seconds, nil
}

// ExpireSec 设置key的过期时间（秒为单位）
func (m *Memc) ExpireSec(key string, seconds int64) error {
	return m.Expire(key, time.Duration(seconds)*time.Second)
}

// Expire 设置key的过期时间
func (m *Memc) Expire(key string, expires time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	// 检查key是否存在
	if _, exists := m.data[key]; !exists {
		return errors.New("key not found")
	}
	
	m.expires[key] = time.Now().Add(expires)
	return nil
}

// ExpireAt 设置key在指定时间过期
func (m *Memc) ExpireAt(key string, tm time.Time) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	// 检查key是否存在
	if _, exists := m.data[key]; !exists {
		return errors.New("key not found")
	}
	
	m.expires[key] = tm
	return nil
}

// IncrBy 将key对应的数字值增加指定数量
func (m *Memc) IncrBy(key string, increment int64) (int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	// 检查并清理过期key
	m.cleanExpired(key)
	
	var currentValue int64 = 0
	if value, exists := m.data[key]; exists {
		// 尝试将现有值转换为int64
		var err error
		currentValue, err = strconv.ParseInt(string(value), 10, 64)
		if err != nil {
			return 0, errors.New("value is not a valid integer")
		}
	}
	
	newValue := currentValue + increment
	m.data[key] = []byte(strconv.FormatInt(newValue, 10))
	return newValue, nil
}

// Incr 将key对应的数字值增加1
func (m *Memc) Incr(key string) (int64, error) {
	return m.IncrBy(key, 1)
}

// DecrBy 将key对应的数字值减少指定数量
func (m *Memc) DecrBy(key string, decrement int64) (int64, error) {
	return m.IncrBy(key, -decrement)
}

// Decr 将key对应的数字值减少1
func (m *Memc) Decr(key string) (int64, error) {
	return m.IncrBy(key, -1)
}