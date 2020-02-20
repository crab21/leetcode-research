package base

import (
	"fmt"
	"testing"
)

func TestSliceSprintf(t *testing.T) {
	t.Run("print-warn", func(t *testing.T) {
		fmt.Println(t.Name())

		SliceSprintf()

	})
}
