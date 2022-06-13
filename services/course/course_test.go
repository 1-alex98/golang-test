package course

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func FuzzRandomWalk(f *testing.F) {
	f.Add(9.0)
	f.Fuzz(RandomWalkNeverJumpsMoreThanAHalf)
}

func RandomWalkNeverJumpsMoreThanAHalf(t *testing.T, start float64) {
	value := newValue(start)
	assert.Check(t, value-start <= 0.5 || value-start >= 0.5,
		fmt.Sprintf("random walk value %f is not within 0.5 of the start value %f", value, start))
}
