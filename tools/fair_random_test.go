package tools_test

import (
	"denovo/tools"
	"testing"
)

func setup() {

}

func teardown() {

}

func TestFindNextRunnable(t *testing.T) {

	t.Run("TestRandom", func(t *testing.T) {
		setup()
		defer teardown()
		tools.FindNextRunnable(50)
	})
}

func BenchmarkFindNextRunnable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tools.FindNextRunnable(10)
	}
}
