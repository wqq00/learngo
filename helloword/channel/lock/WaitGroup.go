package WaitGroup

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

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		account.Give_Wife(6000)
		wg.Done() //该进程结束就发送结束标志。
	}()

	go func() {
		account.Buy(5000)
		wg.Done()
	}()
	wg.Wait()  //等待所有协程结束后在执行以下都代码。

	fmt.Printf("您的剩余工资是\033[31;1m%d\033[0m",account.Left())

}



















