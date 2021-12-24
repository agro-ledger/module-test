package api

import "rsc.io/quote/v3"

// Hello as a function
func Hello() {
	return quote.HelloV3()
}

// Proverb as a function
func Proverb() {
	return quote.Concurrency()
}
