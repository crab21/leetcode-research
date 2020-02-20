package sortAlgorithm

import (
	"fmt"
	"time"
)

/*
快速排序
*/

/*
leetcode [1 1 2 5 6 7 10] 8
输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Aug() {
	//smp()
	//algorithm.Calc_cut()
	var num = []int{2, 5, 2, 1, 2}

	//fmt.Print(num)
	sum2 := combinationSum2(num, 5)
	fmt.Println(sum2)

}

//思路： 先排序  再回溯查找  一旦发现比target大 跳出循环
/*
快速排序
*/

func SortStartByQuick(num []int, low, high int) {

	if num == nil || low >= high || len(num) <= 1 {
		return
	}
	var partion = partionSortOfIndex(num, low, high)
	SortStartByQuick(num, low, partion-1)
	SortStartByQuick(num, partion+1, high)
}

func partionSortOfIndex(num []int, low, high int) int {
	var flagNumber = num[low]
	for ; low < high; {
		for ; num[high] >= flagNumber && high > low; {
			high--
		}
		num[low] = num[high]

		for ; num[low] <= flagNumber && high > low; {
			low++
		}
		num[high] = num[low]
	}
	num[low] = flagNumber

	return low
}

func combinationSum2(candidates []int, target int) [][]int {
	SortStartByQuick(candidates, 0, len(candidates)-1)
	var ll = time.Now().UnixNano()
	var res [][]int
	op :=backtrack(candidates, 0, target, []int{}, 0,res)
	fmt.Println(time.Now().UnixNano()-ll)
	return op

}



func backtrack(numbers []int, start int, target int, tmp []int, sum int,res [][]int) [][]int{
	var s = make(map[int]interface{})
	for i := start; i < len(numbers); i++ {
		if _, ok := s[numbers[i]]; ok {
			continue
		} else {
			s[numbers[i]] = nil
		}

		tmp = append(tmp, numbers[i])
		sum += numbers[i]

		if sum == target {
			var ssmp []int = make([]int, len(tmp))
			copy(ssmp, tmp)
			res = append(res, ssmp)
		} else if sum < target {
			res = backtrack(numbers, i+1, target, tmp, sum,res)
		} else {
			if len(tmp) > 0 {
				tmp = tmp[:len(tmp)-1]
			}
			sum -= numbers[i]
			break
		}
		if len(tmp) > 0 {
			tmp = tmp[:len(tmp)-1]
		}
		sum -= numbers[i]
	}
	return res
}
