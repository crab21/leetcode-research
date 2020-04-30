package sortAlgorithm

//选择
func SelectSort(a []int) []int {
	n := len(a)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
	return a
}

//插入
func InsertSort(a []int) []int {
	if a == nil || len(a) < 2 {
		return nil
	}
	n := len(a)
	for i := 1; i < n; i++ {
		tmp := a[i]
		k := i - 1
		for ; k >= 0 && a[k] > tmp; {
			k--
		}
		for j := i; j > k+1; j-- {
			a[j] = a[j-1]
		}
		a[k+1] = tmp
	}
	return a
}

//冒泡
func BubbleSort(a []int) []int {
	if a == nil || len(a) < 2 {
		return nil
	}
	n := len(a)
	for i := 0; i < n; i++ {
		flag := true
		for j := 0; j < n-1-i; j++ {
			if a[j+1] < a[j] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return a
}

func MergeSort(a []int, left, right int) []int {
	if left < right {
		mid := (right + left) / 2
		MergeSort(a, left, mid)
		MergeSort(a, mid+1, right)
		MergeToSlice(a, left, mid, right)
	}
	return a
}

func MergeToSlice(a []int, left int, mid int, right int) {
	sa := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	for ; i <= mid && j <= right; {
		if a[i] < a[j] {
			sa[k] = a[i]
			k++
			i++
		} else {
			sa[k] = a[j]
			k++
			j++
		}
	}
	for ; i <= mid; {
		sa[k] = a[i]
		k++
		i++
	}
	for ; j <= right; {
		sa[k] = a[j]
		k++
		j++
	}
	for is := 0; is < len(sa); is++ {
		a[left] = sa[is]
		left++
	}
}

func MergeSortNoRecure(a []int) {
	n := len(a)
	for i := 1; i < n; i += i {
		left := 0
		mid := left + i - 1
		right := mid + i
		for ; right < n; {
			MergeToSlice(a, left, mid, right)
			left = right + 1
			mid = left + i - 1
			right = mid + 1
		}
		if left < n && mid < n {
			MergeToSlice(a, left, mid, n-1)
		}
	}
}

func QuickSort(a []int, left, right int) []int {
	if left < right {
		mid := partition(a, left, right)
		QuickSort(a, left, mid-1)
		QuickSort(a, mid+1, right)
	}
	return a

}

func partition(a []int, left int, right int) int {
	pivot := a[left]
	for ; left < right; {
		for ; left < right && a[right] >= pivot; {
			right--
		}
		a[left] = a[right]
		for ; left < right && a[left] <= pivot; {
			left++
		}
		a[right] = a[left]
	}
	a[left] = pivot
	return left
}

func HeapSort(num []int) {
	length := len(num)
	first := length/2 - 1
	//构建大顶堆
	for i := first; i >= 0; i-- {
		adjustHeap(num, i, length)
	}

	//交换首尾，调整结构
	for j := len(num) - 1; j > 0; j-- {
		num[0], num[j] = num[j], num[0]
		adjustHeap(num, 0, j)

	}

}

func adjustHeap(num []int, i int, length int) {
	tmp := num[i]

	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && num[k] < num[k+1] {
			k++
		}

		if num[k] > tmp {
			num[i] = num[k]
			i = k
		} else {
			break
		}
	}

	num[i] = tmp
}



