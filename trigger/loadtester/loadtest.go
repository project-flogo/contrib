package loadtester

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/project-flogo/core/trigger"
)

// Based off github.com/tsliwowicz/go-wrk

// RequesterStats used for collection aggregate statistics
type RequesterStats struct {
	TotDuration    time.Duration
	MinRequestTime time.Duration
	MaxRequestTime time.Duration
	NumRequests    int
	NumErrs        int
}

type LoadTest struct {
	duration         int //seconds
	concurrencyLevel int
	statsAggregator  chan *RequesterStats
	interrupted      int32
}

func NewLoadTest(duration, concurrencyLevel int) *LoadTest {

	lt := &LoadTest{duration: duration, concurrencyLevel: concurrencyLevel}
	lt.statsAggregator = make(chan *RequesterStats, concurrencyLevel)

	return lt
}

func (lt *LoadTest) Run(handler trigger.Handler, data interface{}) error {

	fmt.Printf("Running %vs test\n  %v goroutine(s) running concurrently\n", lt.duration, lt.concurrencyLevel)

	lt.statsAggregator = make(chan *RequesterStats, lt.concurrencyLevel)
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt)

	session := NewLoadSession(lt.duration, lt.statsAggregator)

	for i := 0; i < lt.concurrencyLevel; i++ {
		go session.Run(handler, data)
	}

	responders := 0
	aggStats := RequesterStats{MinRequestTime: time.Minute}

	for responders < lt.concurrencyLevel {
		select {
		case <-sigChan:
			session.Stop()
		case stats := <-lt.statsAggregator:
			aggStats.NumErrs += stats.NumErrs
			aggStats.NumRequests += stats.NumRequests
			aggStats.TotDuration += stats.TotDuration
			aggStats.MaxRequestTime = MaxDuration(aggStats.MaxRequestTime, stats.MaxRequestTime)
			aggStats.MinRequestTime = MinDuration(aggStats.MinRequestTime, stats.MinRequestTime)
			responders++
		}
	}

	if aggStats.NumRequests == 0 {
		fmt.Println("Error: No statistics collected")
		return nil
	}

	avgThreadDur := aggStats.TotDuration / time.Duration(responders) //need to average the aggregated duration

	reqRate := float64(aggStats.NumRequests) / avgThreadDur.Seconds()
	avgReqTime := aggStats.TotDuration / time.Duration(aggStats.NumRequests)
	fmt.Printf("%v requests in %v\n", aggStats.NumRequests, avgThreadDur)
	fmt.Printf("Requests/sec:\t\t%.2f\n", reqRate)
	fmt.Printf("Avg Req Time:\t\t%v\n", avgReqTime)
	fmt.Printf("Fastest Request:\t%v\n", aggStats.MinRequestTime)
	fmt.Printf("Slowest Request:\t%v\n", aggStats.MaxRequestTime)
	fmt.Printf("Number of Errors:\t%v\n", aggStats.NumErrs)

	return nil
}

type LoadSession struct {
	duration        int //seconds
	statsAggregator chan *RequesterStats
	interrupted     int32
}

func NewLoadSession(duration int, statsAggregator chan *RequesterStats) *LoadSession {
	return &LoadSession{duration: duration, statsAggregator: statsAggregator}
}

func (lt *LoadSession) Run(handler trigger.Handler, data interface{}) {
	stats := &RequesterStats{MinRequestTime: time.Minute}
	sessionStart := time.Now()

	for time.Since(sessionStart).Seconds() <= float64(lt.duration) && atomic.LoadInt32(&lt.interrupted) == 0 {

		var duration time.Duration
		duration = -1

		start := time.Now()
		_, err := handler.Handle(context.Background(), data)
		duration = time.Since(start)

		if err == nil {
			stats.TotDuration += duration
			stats.MaxRequestTime = MaxDuration(duration, stats.MaxRequestTime)
			stats.MinRequestTime = MinDuration(duration, stats.MinRequestTime)
			stats.NumRequests++
		} else {
			stats.NumErrs++
		}
	}
	lt.statsAggregator <- stats
}

func (lt *LoadSession) Stop() {
	atomic.StoreInt32(&lt.interrupted, 1)
}

func MaxDuration(d1 time.Duration, d2 time.Duration) time.Duration {
	if d1 > d2 {
		return d1
	} else {
		return d2
	}
}

func MinDuration(d1 time.Duration, d2 time.Duration) time.Duration {
	if d1 < d2 {
		return d1
	} else {
		return d2
	}
}
