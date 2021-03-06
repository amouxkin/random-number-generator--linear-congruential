package lcg_test

import (
	"avala/lcg"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func lehmerRunner(seed, sampleSpace int, generated *[]int) {
	generator := lcg.LehmerGenerator((*generated)[0], seed, sampleSpace)

	for {
		random := generator()

		fmt.Printf("%v", random)

		if random == (*generated)[0] {
			break
		}

		*generated = append(*generated, random)
	}

}

func TestLehmerGeneratorValues(t *testing.T) {
	tests := []struct {
		name        string
		seed        int
		sampleSpace int
		generated   []int
	}{
		{"should randomly generate [0...10]", 7, 11, []int{1}},
		{"should randomly generate [0...12]", 7, 13, []int{1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lehmerRunner(test.seed, test.sampleSpace, &test.generated)
			assert.Equal(t, test.sampleSpace-1, len(test.generated))

			for i := 1; i < test.sampleSpace; i++ {
				assert.Contains(t, test.generated, i)
			}
		})
	}
}

func TestLehmerGeneratorGuards(t *testing.T) {
	tests := []struct {
		name         string
		seed         int
		sampleSpace  int
		previous     int
		panicMessage string
	}{
		{"should panic for invalid seed", 8, 11, 12, "seed is not prime"},
		{"should not panic for invalid previous", 7, 4, 12, "invalid previous value"},
		{"should not panic for invalid previous", 7, 6, 12, "sampleSpace is not a prime / power of prime"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.PanicsWithValue(t, test.panicMessage, func() {
				lcg.LehmerGenerator(test.previous, test.seed, test.sampleSpace)
			})
		})
	}
}

func TestFindClosestSampleSpace(t *testing.T) {
	tests := []struct {
		num                int
		closestSampleSpace int
	}{
		{10, 11},
		{0xfff, 4_099},
		{0xfff_fff, 16_777_259},
		{0xf_fff_fff, 268_435_459},
		// Takes a couple of minutes to calculate this. (The value is very close to the hexadecimal number required
		// Just a difference of 16.)
		{0xff_fff_fff, 4_294_967_311},
	}

	for _, test := range tests {
		assert.Equal(t, lcg.FindClosestSampleSpace(test.num), test.closestSampleSpace)
	}
}
