package base

import "fmt"

func main()  {
	var array1 = [...]int{1,2,3}
	fmt.Println(array1)

	for i := 0; i< len(array1); i++{
		fmt.Println(array1[i])
	}

	for i,v := range array1{
		fmt.Println(array1[i])
		fmt.Println(i)
		fmt.Println(v)

	}

}
