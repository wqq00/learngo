package channel

import (
	"sync"
	"time"
	"fmt"
)

type Accout struct {
	flag sync.Mutex
	money int
}

func (a *Accout) Do_Prepare(){
	time.Sleep(time.Second)
}

func (a *Accout) Get_Gongzi(n int){
	a.money += n
}

func (a *Accout) Give_Wife(n int){
	a.flag.Lock()
	defer a.flag.Unlock()
	if a.money > n{
		a.Do_Prepare()
		a.money -= n
	}else{
		fmt.Println("no money give wife")
	}
}

func (a *Accout) Buy(n int){
	a.flag.Lock()
	defer a.flag.Unlock()
	if a.money > n {
		a.Do_Prepare()
		a.money -= n
	}else{
		fmt.Println("no money buy")
	}
}

func (a *Accout) Left() int{
	return a.money
}

func main(){
	var account Accout
	account.Get_Gongzi(10000)

	var work_info chan string
	work_info = make(chan string, 2)
	defer close(work_info)

	go func(){
		account.Give_Wife(6000)
		work_info <- "I done"
	}()

	go func(){
		account.Buy(5000)
		work_info <- "I have done too!"
	}()

	cnt := 0
	for i := range work_info{
		fmt.Println(i)
		cnt ++
		if cnt >= 2{
			break
		}
	}

	fmt.Printf("您的剩余工资是\033[31;1m%d\033[0m",account.Left())

}



















