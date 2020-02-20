package sortAlgorithm

/**
归并排序算法
思想：类似快速排序    分治思想
*/

func MergeForSort(arr []int, left int, right int) {
	if left < right {
		middle := (left + right) / 2
		MergeForSort(arr, left, middle)
		MergeForSort(arr, middle+1, right)
		merge(arr, left, middle, right)
	}
}

func merge(arr []int, left int, middle int, right int) {
	arrSize := right - left + 1
	var tmp = make([]int, 0, arrSize)
	leftModer, rightModer := left, middle+1
	k := 0
	for ; leftModer <= middle && rightModer <= right; {
		if arr[leftModer] <= arr[rightModer] {
			tmp = append(tmp, arr[leftModer])
			leftModer++
		} else {
			tmp = append(tmp, arr[rightModer])
			rightModer++
		}
		k++
	}

	for ; leftModer <= middle; {
		tmp = append(tmp, arr[leftModer])
		leftModer++
		k++
	}
	for ; rightModer <= right; {
		tmp = append(tmp, arr[rightModer])
		rightModer++
		k++
	}
	for k2 := 0; k2 < len(tmp); k2++ {
		arr[k2+left] = tmp[k2]
	}
}
