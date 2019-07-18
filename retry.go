package main

import (
	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/Rican7/retry/backoff"
	"errors"
	"fmt"
	"time"
)

func main() {
	action := func(attempt uint) error {
		fmt.Println(time.Now())
		return errors.New("basic error")
	}

	const Incremental = time.Second
	incremental := backoff.BinaryExponential(duration, Incremental)
	strategy1 := strategy.Backoff(incremental)

	strategy2 := strategy.Limit(5)

	err := retry.Retry(action, strategy1, strategy2)
	fmt.Println(err)
}