package designpatterns

import (
	"fmt"
	"sync"
)

var once sync.Once
var instance *MyInstance

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("interation:", i)
		instance = getMyInstance()
	}
	fmt.Println("got an instance:", instance)
}

type MyInstance struct {
	DbHost   string
	User     string
	Password string
	Port     int
}

func getMyInstance() *MyInstance {
	once.Do(func() {
		fmt.Println("getting instance")
		instance = &MyInstance{}
	})
	return instance
}
