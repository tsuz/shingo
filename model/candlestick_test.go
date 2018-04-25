package model

import (
	"testing"
	"time"
)

func TestCandlestick(t *testing.T) {
	now := time.Now()
	c, err := NewCandlestick(45.5, 46, 46.5, 42.2, now, 2302)
	if err != nil {
		t.Fatalf("Error creating new candlestick %+v", err)
	}

	if c.Open != 45.5 {
		t.Errorf("Expected open to be 45.5 but got %f", c.Open)
	}
	if c.Close != 46 {
		t.Errorf("Expected close to be 45.5 but got %f", c.Close)
	}
	if c.High != 46.5 {
		t.Errorf("Expected high to be 46.5 but got %f", c.High)
	}
	if c.Low != 42.2 {
		t.Errorf("Expectd low to be 42.2 but got %f", c.Low)
	}
	if c.OpenTime != now {
		t.Errorf("Expected time to be %+v but got %+v", now, c.OpenTime)
	}
	if c.Volume != 2302 {
		t.Errorf("Expected volume to be 2302 but got %f", c.Volume)
	}
}
