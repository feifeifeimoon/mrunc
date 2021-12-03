package util

import (
	"fmt"
	"testing"
	"time"
)

func TestReadSpecFromBundle(t *testing.T) {
	start := time.Now()

	for i := 0; i < 1000; i++ {
		_, err := ReadSpecFromBundle("../../bin")
		if err != nil {
			t.Errorf("read spec from bundle err: %v", err)
		}
	}

	spend := time.Since(start)
	fmt.Println("time ", spend)
}
