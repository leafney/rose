package memc

import (
	"testing"
	"time"
)

func TestMemc_BasicOperations(t *testing.T) {
	m := NewMemc()

	// 测试Set和Get
	err := m.Set("key1", []byte("value1"))
	if err != nil {
		t.Errorf("Set failed: %v", err)
	}

	value, err := m.Get("key1")
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if string(value) != "value1" {
		t.Errorf("Expected 'value1', got '%s'", string(value))
	}

	// 测试SetS和GetS
	err = m.SetS("key2", "value2")
	if err != nil {
		t.Errorf("SetS failed: %v", err)
	}

	strValue, err := m.GetS("key2")
	if err != nil {
		t.Errorf("GetS failed: %v", err)
	}
	if strValue != "value2" {
		t.Errorf("Expected 'value2', got '%s'", strValue)
	}

	// 测试Exists
	if !m.Exists("key1") {
		t.Error("key1 should exist")
	}
	if m.Exists("nonexistent") {
		t.Error("nonexistent key should not exist")
	}

	// 测试Del
	err = m.Del("key1")
	if err != nil {
		t.Errorf("Del failed: %v", err)
	}
	if m.Exists("key1") {
		t.Error("key1 should not exist after deletion")
	}
}

func TestMemc_ExpirationOperations(t *testing.T) {
	m := NewMemc()

	// 测试SetEx
	err := m.SetEx("expkey", []byte("expvalue"), 100*time.Millisecond)
	if err != nil {
		t.Errorf("SetEx failed: %v", err)
	}

	// 立即检查值存在
	value, err := m.Get("expkey")
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if string(value) != "expvalue" {
		t.Errorf("Expected 'expvalue', got '%s'", string(value))
	}

	// 等待过期
	time.Sleep(150 * time.Millisecond)
	_, err = m.Get("expkey")
	if err == nil {
		t.Error("Expected key to be expired")
	}

	// 测试SetExSec
	err = m.SetExSec("expkey2", []byte("expvalue2"), 1)
	if err != nil {
		t.Errorf("SetExSec failed: %v", err)
	}

	// 测试SetExSecS
	err = m.SetExSecS("expkey3", "expvalue3", 1)
	if err != nil {
		t.Errorf("SetExSecS failed: %v", err)
	}
}

func TestMemc_TTLOperations(t *testing.T) {
	m := NewMemc()

	// 测试永不过期的key
	m.Set("persistkey", []byte("value"))
	ttl, err := m.TTL("persistkey")
	if err != nil {
		t.Errorf("TTL failed: %v", err)
	}
	if ttl != -1 {
		t.Errorf("Expected TTL -1 for persistent key, got %d", ttl)
	}

	// 测试带过期时间的key
	m.SetEx("tempkey", []byte("value"), 5*time.Second)
	ttl, err = m.TTL("tempkey")
	if err != nil {
		t.Errorf("TTL failed: %v", err)
	}
	if ttl <= 0 || ttl > 5 {
		t.Errorf("Expected TTL between 1-5 seconds, got %d", ttl)
	}

	// 测试不存在的key
	ttl, err = m.TTL("nonexistent")
	if err == nil {
		t.Error("Expected error for nonexistent key")
	}
	if ttl != -2 {
		t.Errorf("Expected TTL -2 for nonexistent key, got %d", ttl)
	}
}

func TestMemc_ExpireOperations(t *testing.T) {
	m := NewMemc()

	// 设置一个key
	m.Set("key", []byte("value"))

	// 测试Expire
	err := m.Expire("key", 100*time.Millisecond)
	if err != nil {
		t.Errorf("Expire failed: %v", err)
	}

	// 检查TTL
	ttl, _ := m.TTL("key")
	if ttl <= 0 {
		t.Error("Key should have positive TTL after Expire")
	}

	// 测试ExpireSec
	m.Set("key2", []byte("value2"))
	err = m.ExpireSec("key2", 1)
	if err != nil {
		t.Errorf("ExpireSec failed: %v", err)
	}

	// 测试ExpireAt
	m.Set("key3", []byte("value3"))
	futureTime := time.Now().Add(1 * time.Second)
	err = m.ExpireAt("key3", futureTime)
	if err != nil {
		t.Errorf("ExpireAt failed: %v", err)
	}

	// 测试对不存在key的操作
	err = m.Expire("nonexistent", time.Second)
	if err == nil {
		t.Error("Expected error when setting expiration on nonexistent key")
	}
}

func TestMemc_IncrDecrOperations(t *testing.T) {
	m := NewMemc()

	// 测试Incr（key不存在时）
	value, err := m.Incr("counter")
	if err != nil {
		t.Errorf("Incr failed: %v", err)
	}
	if value != 1 {
		t.Errorf("Expected 1, got %d", value)
	}

	// 测试IncrBy
	value, err = m.IncrBy("counter", 5)
	if err != nil {
		t.Errorf("IncrBy failed: %v", err)
	}
	if value != 6 {
		t.Errorf("Expected 6, got %d", value)
	}

	// 测试Decr
	value, err = m.Decr("counter")
	if err != nil {
		t.Errorf("Decr failed: %v", err)
	}
	if value != 5 {
		t.Errorf("Expected 5, got %d", value)
	}

	// 测试DecrBy
	value, err = m.DecrBy("counter", 3)
	if err != nil {
		t.Errorf("DecrBy failed: %v", err)
	}
	if value != 2 {
		t.Errorf("Expected 2, got %d", value)
	}

	// 测试对非数字值的操作
	m.Set("nonnum", []byte("notanumber"))
	_, err = m.Incr("nonnum")
	if err == nil {
		t.Error("Expected error when incrementing non-numeric value")
	}
}