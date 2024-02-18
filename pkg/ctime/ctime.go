package ctime

import (
	"time"

	"github.com/uptrace/bun"
)

func NewJST() (time.Time, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Time{}, err
	}
	nowJST := time.Now().In(jst)

	return nowJST, nil
}

func NewExpirationTimeInJST(d time.Duration) (time.Time, error) {
	nowJST, err := NewJST()
	if err != nil {
		return time.Time{}, err
	}

	return nowJST.Add(d), nil
}

func NullTimeToPtrJST(t bun.NullTime) *time.Time {
	if !t.IsZero() {
		return &t.Time
	}

	return nil
}
