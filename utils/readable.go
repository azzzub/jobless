package utils

import "github.com/icza/gox/fmtx"

const (
	currency  = "Rp"
	separator = '.'
	ending    = ",-"
)

func ReadablePrice(price int) string {
	return currency + fmtx.FormatInt(int64(price), 3, separator) + ending
}
