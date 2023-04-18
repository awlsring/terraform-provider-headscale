package utils

import (
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
)

func IsExpired(t string) (bool, error) {
	if t == "0001-01-01T00:00:00.000Z" {
		return false, nil
	}

	expireTime, err := time.Parse(time.RFC3339Nano, t)
	if err != nil {
		return false, err
	}

	if expireTime.Before(time.Now()) {
		return true, nil
	}

	return false, nil
}

func MakeExpiraryTime(e string) (*strfmt.DateTime, error) {
	expireDuration, err := ParseDurationPlus(e)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	then := now.Add(expireDuration)

	time := strfmt.DateTime(then)
	return &time, nil
}

func ParseDurationPlus(input string) (time.Duration, error) {
	var duration time.Duration
	for len(input) > 0 {
		var (
			num  string
			unit string
		)
		i := 0
		for i < len(input) && input[i] >= '0' && input[i] <= '9' {
			num += string(input[i])
			i++
		}
		if i == 0 {
			return 0, &time.ParseError{Layout: "", Value: input, LayoutElem: ""}
		}
		if i >= len(input) {
			return 0, &time.ParseError{Layout: "", Value: input, LayoutElem: ""}
		}
		unit = string(input[i])
		if unit != "ns" && unit != "us" && unit != "µs" && unit != "ms" && unit != "s" && unit != "m" && unit != "h" && unit != "d" && unit != "w" && unit != "M" && unit != "y" {
			return 0, &time.ParseError{Layout: "", Value: input, LayoutElem: unit}
		}
		i++
		if i < len(input) && input[i] >= 'a' && input[i] <= 'z' {
			unit += string(input[i])
			i++
		}
		input = input[i:]

		numVal, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			return 0, err
		}

		switch unit {
		case "ns":
			duration += time.Duration(numVal) * time.Nanosecond
		case "us", "µs":
			duration += time.Duration(numVal) * time.Microsecond
		case "ms":
			duration += time.Duration(numVal) * time.Millisecond
		case "s":
			duration += time.Duration(numVal) * time.Second
		case "m":
			duration += time.Duration(numVal) * time.Minute
		case "h":
			duration += time.Duration(numVal) * time.Hour
		case "d":
			duration += time.Duration(numVal) * time.Hour * 24
		case "w":
			duration += time.Duration(numVal) * time.Hour * 24 * 7
		case "M":
			duration += time.Duration(float64(numVal) * 30.44 * 24 * float64(time.Hour))
		case "y":
			duration += time.Duration(numVal) * time.Hour * 24 * 365
		}
	}

	return duration, nil
}
