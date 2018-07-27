package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f2(n int) int {
	return n
}

func a2_iteration(n int, a int) int {
	sum := 0

	for i := n; i >= 0; i-- {
		if i == 0 {
			sum += a
		} else {
			sum = (sum + f2(i)) * a
		}
	}
	return sum
}

func a2_recursive(n int, a int, sum int) int {
	if n == 0 {
		return sum + a
	}
	return a2_recursive(n-1, a, (sum+f2(n))*a)
}

func a3(n int, index int, s []byte) []string {
	if index == n-1 {
		r1 := string(append(s, 't'))
		r2 := string(append(s, 'f'))
		return []string{r1, r2}
	}
	r3 := a3(n, index+1, append(s, 't'))
	r4 := a3(n, index+1, append(s, 'f'))
	var res []string = make([]string, len(r3)+len(r4))
	copy(res, r3)
	copy(res[len(r3):], r4)
	return res
}

func a4(l []int) {
	length := len(l)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if l[i] < l[j] {
				l[i], l[j] = l[j], l[i]
			}
		}
	}
}

func f() int {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(100)
}

func a5(n int) {
	var m map[int]int = make(map[int]int)
	for i := 1; i <= n; i++ {
		t := f()
		if m[t] != 0 {
			//fmt.Println("a5: 输入",m[t],"和",i,"值相同，都为",t)
		}
		m[t] = i
	}
}

func a6(n int) bool {
	var s []int = make([]int, 0)
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			s = append(s, i)
		}
	}

	sum := 0
	for _, v := range s {
		sum += v
	}
	if sum == n {
		return true
	} else {
		return false
	}
}

func a7_iteration(n int) int {
	var product int = 1
	for i := 1; i <= n; i++ {
		product *= i
	}
	return product
}

func a7_recursive(n int) int {
	if n == 1 {
		return 1
	}
	return n * a7_recursive(n-1)
}

func a8_iteration(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		t := b
		b = a + b
		a = t
	}
	return b
}

func a8_recursive(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return a8_recursive(n-1) + a8_recursive(n-2)
}


//求出完整杨辉三角 再取值
func a9_iteration(n, m int) int {
	var index [][]int = make([][]int,0)
	for i:=0;i<=n;i++{
		index = append(index,[]int{})
		for j:=0;j<=i;j++ {
			var t int = 0
			if j==i||j==0{
				t =1
			}else {

				t = index[i-1][j-1]+index[i-1][j]
			}
			index[i] = append(index[i],t)
		}
	}
	//fmt.Println(index)
	return index[n][m]
}

func a9_recursive(n, m int) int {
	if n == m||m==0 {
		return 1
	}
	res := a9_recursive(n-1, m) + a9_recursive(n-1, m-1)
	return res
}

func main() {

	//a2 =================
	fmt.Println("a2: iteration: ", a2_iteration(2, 1), "recursive:", a2_recursive(2, 1, 0))

	//a3 =================
	res := a3(4, 0, []byte{})
	fmt.Println("a3: ", res)

	//a4 =================
	l := []int{1, 2, 3}
	a4(l)
	fmt.Println("a4: ", l)

	//a5 =================
	a5(100)

	//a6 =================
	fmt.Println("a6: 10 是不是它的所有因子之和： ", a6(10))

	//a7 =================
	fmt.Println("a7: iteration:", a7_iteration(10), " recursive:", a7_recursive(10))

	//a8 =================
	fmt.Println("a8: iteration:", a8_iteration(10), " recursive:", a8_recursive(10))

	//a9 =================
	fmt.Println("a9: iteration:", a9_iteration(10, 2), " recursive:", a9_recursive(10, 2))

}

