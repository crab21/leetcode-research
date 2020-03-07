package sortAlgorithm

import "fmt"

/**
堆排序

*/

func HeapSort(arr []int) {
	//构建大顶堆

	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, len(arr))
	}

	fmt.Println("=======", arr)
	//调整堆结构+交换首尾元素
	for j := len(arr) - 1; j > 0; j-- {
		arr[0], arr[j] = arr[j], arr[0]
		adjustHeap(arr, 0, j)
	}

}

func adjustHeapCopy(arr []int, i, length int) {
	tmp := arr[i]

	for k := 2*i + 1; k < length; k = k*2 + 1 {
		if k+1 < length && arr[k] > arr[k+1] {
			k++
		}
		if arr[k]< tmp {
			arr[i] = arr[k]
			 i = k
		}else {
			break
		}

	}
	arr[i] = tmp

}

func adjustHeap(arr []int, i int, length int) {
	tmp := arr[i]

	for k := 2*i + 1; k < length; k = k*2 + 1 {
		if k+1 < length && arr[k] > arr[k+1] {
			k++
		}
		if arr[k] < tmp {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}
	arr[i] = tmp
}
