package main

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

/**
 *                    .::::.
 *                  .::::::::.
 *                 :::::::::::  
 *             ..:::::::::::'
 *           '::::::::::::'
 *             .::::::::::
 *        '::::::::::::::..
 *             ..::::::::::::.
 *           ``::::::::::::::::
 *            ::::``:::::::::'        .:::.
 *           ::::'   ':::::'       .::::::::.
 *         .::::'      ::::     .:::::::'::::.
 *        .:::'       :::::  .:::::::::' ':::::.
 *       .::'        :::::.:::::::::'      ':::::.
 *      .::'         ::::::::::::::'         ``::::.
 *  ...:::           ::::::::::::'              ``::.
 * ```` ':.          ':::::::::'                  ::::..
 *                    '.:::::'                    ':'````..
 */

type matrix struct {
	col int //列
	row int //行
	val int
}

func create(col int, row int) []matrix {
	var m matrix = matrix{col, row, 0}
	return []matrix{m}
}

func addEle(m []matrix, col, row, val int) ([]matrix, error) {

	if col > m[0].col || row > m[0].row {
		return m, errors.New("参数错误")
	}
	m = append(m, matrix{col, row, val})
	m[0].val++
	return m, nil
}

func transpose(a []matrix) []matrix {
	b := create(a[0].row, a[0].col)
	b[0].val = a[0].val
	for i := 0; i < b[0].row; i++ {
		for j := 1; j < b[0].val+1; j++ {
			if a[j].col == i {
				b = append(b, matrix{a[j].row, a[j].col, a[j].val})
			}
		}

	}
	return b
}

func fast_transpose(a []matrix) []matrix {
	b := make([]matrix, a[0].val+1)
	b[0].col = a[0].row
	b[0].row = a[0].col
	b[0].val = a[0].val
	aCol := make([]int, a[0].col)
	bRow := make([]int, a[0].col)
	//获取原矩阵 每列数目 为转置矩阵 每行的数目
	for i := 1; i < a[0].val+1; i++ {
		aCol[a[i].col]++
	}
	// 获取新矩阵每行 的开始索引
	bRow[0] = 1
	for i := 1; i < a[0].col; i++ {
		bRow[i] = bRow[i-1] + aCol[i-1]
	}
	//
	for i := 1; i < a[0].val+1; i++ {
		b[bRow[a[i].col]] = matrix{a[i].row, a[i].col, a[i].val}
		bRow[a[i].col]++
	}
	return b
}
func product(a, b []matrix) []matrix {
	c := create(b[0].col, a[0].row)
	for i := 0; i < c[0].row; i++ {
		for j := 0; j < c[0].col; j++ {
			temp := 0
			for m := 1; m < a[0].val+1; m++ {
				if a[m].row == i {
					for n := 1; n < b[0].val+1; n++ {
						if a[m].col == b[n].row && b[n].col == j {
							temp += a[m].val * b[n].val
						}
					}
				}
			}
			if temp != 0 {
				c = append(c, matrix{j, i, temp})
				c[0].val ++
			}
		}
	}
	return c
}

//转置 b 获取b每一行的个数 a的i行 ✖️ b的j行 即为c[i][j]
func fast_product(a, b []matrix) []matrix {
	b = fast_transpose(b)
	aRow := make([]int, a[0].row+1)
	aRow[0] = 1
	aRow[a[0].row] = a[0].val + 1
	start := 1
	for i := 1; i < a[0].val+1; i++ {
		if a[0].row == a[i].row+1 {
			break
		}
		start++
		aRow[a[i].row+1] = start
	}

	bRow := make([]int, b[0].row+1)
	bRow[0] = 1
	bRow[a[0].row] = b[0].val + 1
	start = 1
	for i := 1; i < a[0].val+1; i++ {
		if b[0].row == b[i].row+1 {
			break
		}
		start++
		bRow[a[i].row+1] = start
	}
	fmt.Println(aRow, bRow)

	c := create(b[0].row, a[0].row)
	for i := 0; i < a[0].row; i++ {
		for j := 0; j < b[0].row; j++ {
			temp := 0
			for m := aRow[i]; m < aRow[i+1]; m++ {
				for n := bRow[j]; n < bRow[j+1]; n++ {
					if a[m].col == b[n].col {
						temp += a[m].val * b[n].val
					}
				}
			}
			if temp != 0 {
				c = append(c, matrix{j, i, temp})
				c[0].val ++
			}
		}
	}

	return c
}
func main() {
	var m []matrix = create(8, 5)
	fmt.Println(m)
	for i := 0; i < 5; i++ {
		for j := 0; j < 8; j++ {
			rand.Seed(int64(time.Now().Nanosecond()))
			if rand.Intn(2) == 1 {
				m, _ = addEle(m, j, i, rand.Intn(1000))
			}
		}
	}
	fmt.Println(m)
	fmt.Println(transpose(m))
	fmt.Println(fast_transpose(m))
	fmt.Println(product(m, fast_transpose(m)))
	fmt.Println(fast_product(m, fast_transpose(m)))
}
