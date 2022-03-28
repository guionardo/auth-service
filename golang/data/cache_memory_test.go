package data

import (
	"testing"
	"time"
)

func TestCacheMemory_Set(t *testing.T) {
	t.Run("set, get", func(t *testing.T) {
		c := &CacheMemory{}
		c.Setup(nil)
		if _, e := c.Get("unexistent_key"); e == nil {
			t.Errorf("Expected error on unexistent key")
		}
		if e := c.Set("a", true, time.Hour); e != nil {
			t.Errorf("Expected setting value without errors %v", e)
		}
		if v, e := c.Get("a"); e != nil || v == false {
			t.Errorf("Expected getting value without errors and true data: %v - %v", v, e)
		}
		if e := c.Set("b", true, time.Hour*time.Duration(-1)); e != nil {
			t.Errorf("Expected setting value without errors %v", e)
		}
		if v, e := c.Get("b"); e == nil || v != nil {
			t.Errorf("Expected getting expired value with errors and nil data: %v - %v", v, e)
		}

	})
}
