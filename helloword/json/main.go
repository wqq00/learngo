package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name string
	ProductID int64
	Number int
	Price float64
	IsOnSale bool
}

func main(){
	p := Product{}
	p.Name = "Xiao mi 66"
	p.ProductID = 1
	p.Number = 666
	p.Price = 66
	p.IsOnSale = true
	data, _ := json.Marshal(&p)
	fmt.Println(string(data))
}
