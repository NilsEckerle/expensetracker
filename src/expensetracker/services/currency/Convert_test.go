package currency

import (
	"errors"
	"testing"
)

// fakeAPI is a stub implementation of ICurrencyConvertionAPI for testing.
// It returns a fixed rate (or error) and records how many times it was called.
type fakeAPI struct {
	rate  float64
	err   error
	calls int
}

func (f *fakeAPI) GetExchangeRate(from, to string) (float64, error) {
	f.calls++
	if f.err != nil {
		return 0, f.err
	}
	return f.rate, nil
}

func TestConvertTo_SameCurrency(t *testing.T) {
	// When from == to, no API call should happen and the amount is unchanged.
	fake := &fakeAPI{rate: 999} // rate should be ignored
	c := NewConverter(fake)

	got, err := c.ConvertTo("EUR", 12345, "EUR")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 12345 {
		t.Errorf("got %d, want 12345", got)
	}
	if fake.calls != 0 {
		t.Errorf("API was called %d times, expected 0 for same currency", fake.calls)
	}
}

func TestConvertTo_DifferentCurrency(t *testing.T) {
	// 10000 cents (100.00) at rate 0.9 = 9000 cents (90.00).
	fake := &fakeAPI{rate: 0.9}
	c := NewConverter(fake)

	got, err := c.ConvertTo("EUR", 10000, "USD")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 9000 {
		t.Errorf("got %d, want 9000", got)
	}
}

func TestConvertTo_Rounding(t *testing.T) {
	// 100 cents at rate 0.855 = 85.5 cents, should round to 86.
	fake := &fakeAPI{rate: 0.855}
	c := NewConverter(fake)

	got, err := c.ConvertTo("EUR", 100, "USD")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 86 {
		t.Errorf("got %d, want 86 (0.855 should round up)", got)
	}
}

func TestConvertTo_APIError(t *testing.T) {
	// When the API errors, ConvertTo should propagate it and return 0.
	wantErr := errors.New("network down")
	fake := &fakeAPI{err: wantErr}
	c := NewConverter(fake)

	got, err := c.ConvertTo("EUR", 10000, "USD")
	if !errors.Is(err, wantErr) {
		t.Errorf("got error %v, want %v", err, wantErr)
	}
	if got != 0 {
		t.Errorf("got %d, want 0 on error", got)
	}
}
