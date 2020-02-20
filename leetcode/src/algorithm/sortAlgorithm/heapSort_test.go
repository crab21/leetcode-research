package sortAlgorithm

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{9,80,7,6,5,4,3,2,1}
	HeapSort(arr)
	fmt.Println(arr)
}
