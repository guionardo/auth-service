package data

import (
	"errors"
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
)

func TestCacheRedis_Set(t *testing.T) {
	db, mock := redismock.NewClientMock()

	mock.ExpectGet("unexistent_key").SetErr(errors.New("unexistent"))
	mock.Regexp().ExpectSet("a", true, time.Hour).SetVal("true")
	mock.ExpectGet("a").SetVal("true")
	mock.Regexp().ExpectSet("b", true, time.Hour*time.Duration(-1)).SetVal("true")
	mock.ExpectGet("b").SetErr(errors.New("expired"))

	c := &CacheRedis{}
	c.Setup(db)

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
	if v, e := c.Get("b"); e == nil {
		t.Errorf("Expected getting expired value with errors and nil data: %v - %v", v, e)
	}
}
