package base

import "fmt"

func main()  {
	arr := [...]int{0,1,2,3,4,5,6,7}
	//fmt.Println(arr[2:6])

	s1 := arr[2:6]
	s2 := s1[1]

	fmt.Println(s1)
	fmt.Println(s2)

}
