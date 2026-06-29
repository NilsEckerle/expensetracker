package currency

import (
	"sync"
	"time"
)

type cacheEntry struct {
	rate      float64
	fetchedAt time.Time
}

type CachingExchangeRateAPI struct {
	inner ICurrencyConvertionAPI
	ttl   time.Duration

	mu    sync.RWMutex
	cache map[string]cacheEntry
}

func NewCachingExchangeRateAPI(inner ICurrencyConvertionAPI, ttl time.Duration) *CachingExchangeRateAPI {
	return &CachingExchangeRateAPI{
		inner: inner,
		ttl:   ttl,
		cache: make(map[string]cacheEntry),
	}
}

func (c *CachingExchangeRateAPI) GetExchangeRate(from, to string) (float64, error) {
	key := from + ":" + to

	c.mu.RLock()
	entry, ok := c.cache[key]
	c.mu.RUnlock()

	if ok && time.Since(entry.fetchedAt) < c.ttl {
		return entry.rate, nil
	}

	// Cache miss or expired — fetch fresh.
	rate, err := c.inner.GetExchangeRate(from, to)
	if err != nil {
		// Optional fallback: serve stale data if the network fails.
		if ok {
			return entry.rate, nil
		}
		return 0, err
	}

	c.mu.Lock()
	c.cache[key] = cacheEntry{rate: rate, fetchedAt: time.Now()}
	c.mu.Unlock()

	return rate, nil
}
