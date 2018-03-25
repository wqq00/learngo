package base

import (
	"fmt"
)

func variableZeroValue(){
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue(){
	var a, b int = 3,4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeefuction(){
	var a, b, c, s = 3,4, true, "def"
	fmt.Println(a, b, c, s)
}

func main(){
	fmt.Println("Hello word")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeefuction()
}

