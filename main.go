//@author D-S
//Created on 2019/9/18 3:41 下午
package main

import (
	"fmt"
)
func main() {

	slice := []int{0,1,2,3}
	m := make(map[int]*int)

	for key,val := range slice {
		m[key] = &val
	}

	for k,v := range m {
		fmt.Println(k,"->",*v)
	}

}


func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}



//删除  idea命令
// git rm --cached -r .idea
// git rm -r --cached .