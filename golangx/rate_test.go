package golangx

import (
	"testing"
	"time"

	"golang.org/x/time/rate"
)

// TestRate 测试限流器, 被限流的事件直接丢弃
func TestRate(t *testing.T) {
	every := time.NewTicker(2 * time.Millisecond)

	// 参数1 r 令牌放入桶的速率, 参数2 b 为令牌桶的大小, 即允许的 burst 突发流量
	limiter := rate.NewLimiter(rate.Every(50*time.Millisecond), 8)

	for i := 0; i < 100; i++ {
		<-every.C
		if limiter.Allow() {
			t.Logf("limiter allow: %v", i)
		}
	}
}

// TestRateReservation 测试限流器的保留功能
func TestRateReservation(t *testing.T) {

	limiter := rate.NewLimiter(rate.Every(100*time.Millisecond), 8)
	rese := limiter.Reserve()
	go func() {
		<-time.After(rese.Delay())
		t.Log("limiter reservation: ok")
	}()
	every := time.NewTicker(2 * time.Millisecond)
	for i := 0; i < 100; i++ {
		<-every.C
		if limiter.Allow() {
			t.Logf("limiter allow: %v", i)
		}
	}
}
