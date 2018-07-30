package main

import (
	"fmt"
	"time"
)

//
func ex20(n int) [][]int {
	if n%2 != 1 {
		return [][]int{}
	}
	var ms [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		temp := []int{}
		for j := 0; j < n; j++ {
			temp = append(temp, 0)
		}
		ms[i] = temp
	}
	ms[0][n/2] = 1
	val := 1
	k := 0     //行
	l := n / 2 // 列
	for {
		if k == 0 {
			k = n
		}
		k--
		if l == 0 {
			l = n
		}
		l--
		fmt.Println(k, l)
		time.Sleep(time.Millisecond)
		if ms[k][l] == 0 {
			val ++
			ms[k][l] = val
			if (val == n*n) {
				break
			}
			fmt.Println(ms, val)
		} else {
			l += 2
			if l>=n{
				l -= n
			}
			k += 3
			if k>=n{
				k -= n
			}

		}
	}
	return ms
}

func main() {
	fmt.Println("魔方程序：",ex20(5))
}
