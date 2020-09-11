package helper

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// FormatTime returns a string value of the time.
func FormatTime(name string, t time.Time) (string, error) {
	t = t.UTC()

	switch name {
	case RFC822TimeFormatName:
		return t.Format(RFC822OutputTimeFormat), nil
	case ISO8601TimeFormatName:
		return t.Format(ISO8601OutputTimeFormat), nil
	case UnixTimeFormatName:
		return strconv.FormatInt(t.Unix(), 10), nil
	default:
		return "", fmt.Errorf("unknown timestamp format name, " + name)
	}
}

// ParseTime attempts to parse the time given the format. Returns
// the time if it was able to be parsed, and fails otherwise.
func ParseTime(formatName, value string) (time.Time, error) {
	switch formatName {
	case RFC822TimeFormatName:
		return time.Parse(RFC822TimeFormat, value)
	case ISO8601TimeFormatName:
		return time.Parse(ISO8601TimeFormat, value)
	case UnixTimeFormatName:
		v, err := strconv.ParseFloat(value, 64)
		_, dec := math.Modf(v)
		dec = math.Round(dec*1e3) / 1e3 //Rounds 0.1229999 to 0.123
		if err != nil {
			return time.Time{}, err
		}
		t := time.Unix(int64(v), int64(dec*(1e9)))
		return t.UTC(), nil

	default:
		panic("unknown timestamp format name, " + formatName)
	}
}
