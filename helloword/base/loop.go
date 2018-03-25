package base

import (
	"fmt"
	"strconv"
)

func convertToBin(n int) string{
	result := ""
	for ; n>0; n/=2{
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main()  {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)
}
