package main

import (
	"fmt"
)

func max(x,y int) (int) {
	if x > y { return x }
	return y
}

/*
题目: 给定一个无序数组 arr, 其中元素可正可负可 0, 给定一个整数 k,
求 arr 所有的子数组中累加和为 k 的最长子数组长度

分析: 如果s(i)表示子数组arr[0..i]中所有元素的累加和, 则子数组arr[j..i]的累加和为s(i)-s(j-1)
1. sum表示arr[0..i]的累加和, 初始化为0;
   len表示累加和为k的最长子数组长度, 初始化为0;
   哈希表m, 其中, key表示arr最左边开始累加的过程中出现的sum, value表示sum最早出现的位置.
2. 从左到右遍历, 当前元素记为arr[i]
3. sum+=arr[i], 此时sum就是s(i), 查看m中是否存在sum-k
  3.1. 如果m中存在sum-k, 说明s(j-1)=k, 此时计算子数组arr[j..i]的长度
  3.2. 如果不存在, 则更新m
*/
func maxLenth01(arr []int, k int) (length int) {
	if len(arr) < 0 {
		return 0
	}
	m := make(map[int]int)
	m[0] = -1
	length, sum := 0, 0
	for index, value := range arr {
		sum += value
		if _, ok := m[sum-k]; ok {
			length = max(index-m[sum-k], length)
		}
		if _, ok := m[sum]; !ok {
			m[sum] = index
		}
	}
	return length
}

func main() {
	arr := []int {1, 2, 1, 1, 1}
	k := 3
	length := maxLenth01(arr, k)
	fmt.Printf("%d", length)
}
