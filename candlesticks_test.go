package shingo

import (
	"testing"
	"time"
)

func TestCandlesticksAppend(t *testing.T) {
	cs, err := NewCandlesticks(IntervalOneDay)
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
