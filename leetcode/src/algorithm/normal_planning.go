package algorithm

import "fmt"

/*
正太规划问题
 */

//1.钢筋切割过程中利益最大化问题.

//递归算法
func cut(p []int, n int) int {
	if n == 0 {
		return 0
	}

	var q = 0
	for i := 1; i <= n; i += 1 {
		cp := p[i-1] + cut(p, n-i)
		if q < cp {
			q = cp
		}
	}
	return q
}

//备忘录策略

func cutDemo(p []int) {
	r := make([]int,len(p)+1)
	for i := 0; i <= len(p); i += 1 {
		r[i] = -1
	}
	cutBWL(p, len(p), r)
}

// 1:判断备忘录中是否有值，有值则直接返回  2:没有值，则进行计算
func cutBWL(p []int, n int, r []int) int {
	q := -1
	if r[n] > 0 {
		return r[n]
	}
	if n == 0 {
		q = 0
	} else {
		for i := 1; i <= n; i += 1 {
			ct := cutBWL(p, n-i, r) + p[i-1]
			if q < ct {
				q = ct
			}
		}
	}
	r[n] = q
	return q
}

//自顶向下   自底向上
func b_u_cut(p []int) int{
	r := make([]int,len(p)+1)
	for i := 1; i <= len(p); i += 1 {
		q := -1
		for j := 1; j <= i; j += 1 {
			ct:=p[j-1]+r[i-j]
			if q< ct {
				q = ct
			}
		}
		r[i]=q
	}
	return r[len(p)]
}


func Calc_cut() {
	var p =[]int{1,5,8,9,10,17,17,20,24,30}
	u_cut := cut(p,9)
	fmt.Print(u_cut)
}