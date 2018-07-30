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
	//product := 1
	//for i := 1; i <= m; i++ {
	//	//不要使用 product *=  (n - i + 1) / i 由于(n - i + 1) / i会有小数 导致数据不对
	//	product = product * (n - i + 1) / i
	//}
	//return product

	if n == m || m == 0 {
		return 1
	}
	var ll [][]int = make([][]int, 0)
	ll = append(ll, []int{m, n})
	sum := 0
	for {
		if len(ll) == 0 {
			break
		}

		temp := ll[len(ll)-1]
		ll = ll[:len(ll)-1]
		m = temp[0]
		n = temp[1]

		//fmt.Println(ll, "m:", m, "n:", n)
		if n > 1 {
			n--
			if m!=n {
				ll = append(ll, []int{m, n})
			}else {
				sum++
			}
			if m > 1 {
				ll = append(ll, []int{m - 1, n})
			} else {
				sum++
			}
		} else {
			sum++
		}

		//fmt.Println(ll, sum)
		//time.Sleep(time.Second)
	}

	return sum
}

func a9_recursive(n, m int) int {

	if n == m || m == 0 {
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
	fmt.Println("a9: iteration:", a9_iteration(10, 3), " recursive:", a9_recursive(10, 3))

	//a10 ================
	fmt.Println("a10 iteration:", a10_iteration(3, 2), " recursive:", a10_recursive(3, 2))

	//a11 ================
	fmt.Println("a11 iteration:", a11_iteration(10), " recursive:", a11_recursive(10))

	//a12 ================
	fmt.Println("a12  recursive:", a12_recursive([]byte{'a', 'b', 'c'}))
}

func a10_iteration(m, n int) int {
	var ll [][]int = make([][]int, 0)
	ll = append(ll, []int{m, n})

	for {
		if len(ll) == 0 {
			break
		}
		for m > 0 {
			for n > 0 {
				//(m,n-1)->(m,0)
				n--
				ll = append(ll, []int{m, n})
			}
			//(m,0) (m-1,1)
			ll = ll[:len(ll)-1]
			n++
			m--
			ll = append(ll, []int{m, n})
		}

		temp := ll[len(ll)-1]
		m = temp[0]
		n = temp[1]
		n++
		ll = ll[:len(ll)-1]
		//fmt.Println(ll)
		if len(ll) != 0 {
			temp = ll[len(ll)-1]
			ll = ll[:len(ll)-1]
			m = temp[0]
			m--
			ll = append(ll, []int{m, n})
		}
		//fmt.Println(ll)
		//time.Sleep(time.Second)
	}
	return n
}

func a10_recursive(m, n int) int {
	if m == 0 {
		return n + 1
	} else if n == 0 {
		return a10_recursive(m-1, 1)
	} else {
		return a10_recursive(m-1, a10_recursive(m, n-1))
	}
}

func a11_iteration(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f = 2*f + 1
	}
	return f
}

func a11_recursive(n int) int {
	if n == 1 {
		return 1
	}
	return a11_recursive(n-1)*2 + 1
}

func a12_recursive(ll []byte) [][]byte {
	length := len(ll)
	if length == 1 {
		return [][]byte{[]byte{}, ll}
	}
	res := a12_recursive(ll[1:])
	for _, v := range res {
		res = append(res, append(v, ll[0]))
	}
	return res
}
