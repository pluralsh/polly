package retry

import (
	"time"

	"github.com/pluralsh/polly/algorithms"
)

type BackoffAlgorithm interface {
	Backoff(iter int) time.Duration
	Continue() bool
}

type Exponential struct {
	mult  float64
	max   float64
	start float64
}

func (exp *Exponential) Backoff(iter int) time.Duration {
	dur := algorithms.Max(exp.start*exp.mult, exp.max)
	exp.start = dur
	return time.Duration(exp.start)
}

func (exp *Exponential) Continue() bool {
	if exp.max <= exp.start {
		return false
	}

	return true
}

type Constant struct {
	max   int
	dur   int
	count int
}

func (con *Constant) Backoff(iter int) time.Duration {
	con.count++
	return time.Duration(con.dur)
}

func (con *Constant) Continue() bool {
	return con.count < con.max
}
