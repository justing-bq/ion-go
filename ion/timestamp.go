package ion

import (
	"fmt"
	"time"
)

// TimestampPrecision is for tracking the precision of a timestamp
type TimestampPrecision uint8

// Possible TimestampPrecision values
const (
	NoPrecision TimestampPrecision = iota
	Year
	Month
	Day
	Minute
	Second
	Nanosecond
)

func (tp TimestampPrecision) String() string {
	switch tp {
	case NoPrecision:
		return "<no precision>"
	case Year:
		return "Year"
	case Month:
		return "Month"
	case Day:
		return "Day"
	case Minute:
		return "Minute"
	case Second:
		return "Second"
	case Nanosecond:
		return "Nanosecond"
	default:
		return fmt.Sprintf("<unknown precision %v>", uint8(tp))
	}
}

// Timestamp struct
type Timestamp struct {
	DateTime  time.Time
	precision TimestampPrecision
	offset    bool
}

// NewTimestamp constructor
func NewTimestamp(dateTime time.Time, precision TimestampPrecision) Timestamp {
	return Timestamp{dateTime, precision, false}
}

// NewTimestampWithOffset constructor
func NewTimestampWithOffset(dateTime time.Time, precision TimestampPrecision, offset bool) Timestamp {
	if precision <= Day {
		// offset does not apply to Timestamps with Year, Month, or Day precision
		return Timestamp{dateTime, precision, false}
	}
	return Timestamp{dateTime, precision, offset}
}

func emptyTimestamp() Timestamp {
	return Timestamp{time.Time{}, NoPrecision, false}
}

// Format returns a formatted Timestamp string.
func (ts Timestamp) Format() string {
	var dateFormat string

	switch ts.precision {
	case NoPrecision:
		dateFormat = ""
	case Year:
		dateFormat = "2006T"
	case Month:
		dateFormat = "2006-01T"
	case Day:
		dateFormat = "2006-01-02T"
	case Minute:
		if ts.offset {
			dateFormat = "2006-01-02T15:04Z07:00"
		} else {
			dateFormat = "2006-01-02T15:04Z"
		}
	case Second:
		if ts.offset {
			dateFormat = "2006-01-02T15:04:05Z07:00"
		} else {
			dateFormat = "2006-01-02T15:04:05Z"
		}
	case Nanosecond:
		if ts.offset {
			dateFormat = "2006-01-02T15:04:05.999999999Z07:00"
		} else {
			dateFormat = "2006-01-02T15:04:05.999999999Z"
		}
	default:
		dateFormat = time.RFC3339Nano
	}

	return ts.DateTime.Format(dateFormat)
}