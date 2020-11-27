package router

import (
	"net/url"
	"strconv"
)

type Query struct {
	path   string
	values url.Values
}

func (p Query) Path() string {
	return p.path
}

func (p Query) Str(key string) string {
	return p.values.Get(key)
}

func (p Query) Int(key string) int {
	i, _ := strconv.ParseInt(p.Str(key), 10, 64)
	return int(i)
}

func (p Query) Bool(key string) bool {
	b, _ := strconv.ParseBool(p.Str(key))
	return b
}

func (p Query) Float64(key string) float64 {
	f, _ := strconv.ParseFloat(p.Str(key), 64)
	return f
}
