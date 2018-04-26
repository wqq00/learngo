package main

import "fmt"

//冒泡排序
func maopao(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

//选择
func xuanze(a []int) {
	for i := 0; i < len(a); i++ {
		var min int = i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}

//插入
func charu(a []int){
	for i := 1; i < len(a); i++{
		for j := i; j > 0; j--{
			if a[j] > a[j-1]{
				break
			}
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

//快速排序
func kuaisu(a []int, left, right int){

	if left >= right{
		return
	}

	val := a[left]
	k := left
	//确定val所在的位置
	for i := left + 1; i < right; i++{
		if(a[i] < val){
			a[k] = a[i]
			a[i] = a[k+1]
			k++
		}
	}

	kuaisu(a, left, k-1)
	kuaisu(a, k+1, right)

}

func main() {
	b := [...]int{7, 4, 35, 4, 10, 4, 6, 16, 43}
	//maopao(b[:])
	//charu(b[:])
	kuaisu(b[:], 0, len(b)-1)
	fmt.Println(b)
}
