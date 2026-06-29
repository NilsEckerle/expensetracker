package currency

import (
	"errors"
	"testing"
	"time"
)

func TestCache_HitAvoidsSecondCall(t *testing.T) {
	// Two lookups of the same pair should call the inner API only once.
	fake := &fakeAPI{rate: 1.1}
	c := NewCachingExchangeRateAPI(fake, time.Hour)

	r1, _ := c.GetExchangeRate("EUR", "USD")
	r2, _ := c.GetExchangeRate("EUR", "USD")

	if r1 != 1.1 || r2 != 1.1 {
		t.Errorf("got %v, %v; want 1.1 both", r1, r2)
	}
	if fake.calls != 1 {
		t.Errorf("inner API called %d times, want 1 (second should hit cache)", fake.calls)
	}
}

func TestCache_DifferentPairsCallSeparately(t *testing.T) {
	// Different currency pairs are cached under different keys.
	fake := &fakeAPI{rate: 1.0}
	c := NewCachingExchangeRateAPI(fake, time.Hour)

	c.GetExchangeRate("EUR", "USD")
	c.GetExchangeRate("EUR", "GBP")

	if fake.calls != 2 {
		t.Errorf("inner API called %d times, want 2 for two distinct pairs", fake.calls)
	}
}

func TestCache_ExpiryRefetches(t *testing.T) {
	// With a zero TTL, every entry is immediately stale, so each call refetches.
	fake := &fakeAPI{rate: 1.0}
	c := NewCachingExchangeRateAPI(fake, 0)

	c.GetExchangeRate("EUR", "USD")
	c.GetExchangeRate("EUR", "USD")

	if fake.calls != 2 {
		t.Errorf("inner API called %d times, want 2 with zero TTL", fake.calls)
	}
}

func TestCache_StaleFallbackOnError(t *testing.T) {
	// First call succeeds and caches. Then the inner API starts failing and
	// the entry is expired. The cache should serve the stale value rather
	// than propagate the error.
	fake := &fakeAPI{rate: 1.23}
	c := NewCachingExchangeRateAPI(fake, time.Nanosecond)

	first, err := c.GetExchangeRate("EUR", "USD")
	if err != nil || first != 1.23 {
		t.Fatalf("first call: got %v, %v; want 1.23, nil", first, err)
	}

	// Make the entry stale and the API fail.
	time.Sleep(2 * time.Nanosecond)
	fake.err = errors.New("network down")

	second, err := c.GetExchangeRate("EUR", "USD")
	if err != nil {
		t.Errorf("expected stale fallback, got error %v", err)
	}
	if second != 1.23 {
		t.Errorf("got %v, want stale 1.23", second)
	}
}

func TestCache_ErrorWithNoCacheReturnsError(t *testing.T) {
	// If the very first call fails (nothing cached), the error propagates.
	fake := &fakeAPI{err: errors.New("network down")}
	c := NewCachingExchangeRateAPI(fake, time.Hour)

	_, err := c.GetExchangeRate("EUR", "USD")
	if err == nil {
		t.Error("expected error when no cached value exists, got nil")
	}
}
