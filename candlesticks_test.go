package shingo

import (
	"testing"
	"time"
)

func TestCandlesticksAppend(t *testing.T) {
	cs, err := NewCandlesticks(IntervalOneDay, 5)
	if err != nil {
		t.Fatalf("Error creating new candlesticks %+v", err)
	}
	c, err := NewCandlestick(45.5, 46, 46.5, 42.2, time.Now(), 2302)
	if err != nil {
		t.Fatalf("Error creating new candlestick %+v", err)
	}
	if err := cs.AppendCandlestick(c); err != nil {
		t.Fatalf("Error appending candlestick %+v", err)
	}

	if cs.Interval() != IntervalOneDay {
		t.Errorf("Expected interval to be %+v but got %+v", IntervalOneDay, cs.Interval())
	}

	if cs.Total() != 1 {
		t.Errorf("Expected total to be 1 but got %d", cs.Total())
	}

	if cs.ItemAtIndex(0) != c {
		t.Errorf("Expected candlestick to be what was added but got %+v", cs.ItemAtIndex(0))
	}
	if cs.ItemAtIndex(1) != nil {
		t.Errorf("Expected candlestick to be nil at 1 but got %+v", cs.ItemAtIndex(1))
	}
}

func TestCandlesticksMaxCap(t *testing.T) {
	c1, _ := NewCandlestick(45.5, 46, 46.5, 42.2, time.Now(), 2302)
	c2, _ := NewCandlestick(47, 46, 46.5, 42.2, time.Now().Add(time.Minute*1), 3000)
	c3, _ := NewCandlestick(47.2, 46, 46.5, 42.2, time.Now().Add(time.Minute*2), 4111)
	c4, _ := NewCandlestick(43.1, 46, 46.5, 42.2, time.Now().Add(time.Minute*3), 1022)
	max := 3
	testCases := []struct {
		candles  []*Candlestick
		expected []*Candlestick
	}{
		{
			candles:  []*Candlestick{c1},
			expected: []*Candlestick{c1},
		},
		{
			candles:  []*Candlestick{c1, c2, c3},
			expected: []*Candlestick{c1, c2, c3},
		},
		{
			candles:  []*Candlestick{c1, c2, c3, c4},
			expected: []*Candlestick{c2, c3, c4},
		},
	}
	for _, tc := range testCases {
		cs, err := NewCandlesticks(IntervalOneDay, max)
		if err != nil {
			t.Fatalf("Error creating new candlesticks %+v", err)
		}
		for _, c := range tc.candles {
			cs.AppendCandlestick(c)
		}
		for i, e := range tc.expected {
			if len(tc.expected) != cs.Total() && cs.Total() != max {
				t.Fatalf("Expected max candles to %d but got %d and %d", max, cs.Total(), len(tc.expected))
			}
			item := cs.ItemAtIndex(i)
			if item.OpenTime.Equal(e.OpenTime) == false {
				t.Errorf("Expected candle to be the same but it was not")
			}
		}
	}
}

func TestCandlesticksGetLast(t *testing.T) {
	c1, _ := NewCandlestick(45.5, 46, 46.5, 42.2, time.Now(), 2302)
	c2, _ := NewCandlestick(47, 46, 46.5, 42.2, time.Now().Add(time.Minute*1), 3000)
	max := 3
	testCases := []struct {
		candles  []*Candlestick
		expected *Candlestick
	}{
		{
			candles:  []*Candlestick{},
			expected: nil,
		},
		{
			candles:  []*Candlestick{c1},
			expected: c1,
		},
		{
			candles:  []*Candlestick{c1, c2},
			expected: c2,
		},
	}
	for _, tc := range testCases {
		cs, err := NewCandlesticks(IntervalOneDay, max)
		if err != nil {
			t.Fatalf("Error creating new candlesticks %+v", err)
		}
		for _, c := range tc.candles {
			cs.AppendCandlestick(c)
		}
		c := cs.GetLastItem()
		if c != tc.expected {
			t.Errorf("Expected %+v but got %+v", tc.expected, c)
		}
	}
}
