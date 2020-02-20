package leetcode

import "fmt"

/**
寻找中位数
*/

func StartMedianNumber() {
	a := []int{1, 2}
	b := []int{3, 4}
	find := findMedianSortedArrays(a, b)
	fmt.Println(find)
}
func findMedianSortedArrays(a, b []int) float64 {
	m := len(a)
	n := len(b)

	if m > n {
		a, b = b, a
		m, n = n, m
	}

	min, max, half_len := 0, m, (m+n+1)/2
	i, j := 0,0
	for ; min <= max; {
		i = (min + max) / 2
		j = half_len - i

		if i < max && b[j-1] > a[i] {
			min = i + 1
		} else if i > min && a[i-1] > b[j] {
			max = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = b[j-1]
			} else if j == 0 {
				maxLeft = a[i-1]
			} else {
				if a[i-1] > b[j-1] {
					maxLeft = a[i-1]
				} else {
					maxLeft = b[j-1]
				}
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			if i == m {
				maxLeft += b[j]
			} else if j == n {
				maxLeft += a[i]
			} else {
				if a[i] < b[j] {
					maxLeft += a[i]
				} else {
					maxLeft += b[j]
				}
			}
			return float64(maxLeft) / 2.0
		}

	}
	return 0.0
}
