package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("simple shell")
	fmt.Println("-------------")
	fmt.Print("pleas input your name :")

	char,err := reader.ReadString('\n')

	if (err != nil){
		fmt.Println(err)
	}

	fmt.Println("yourname is " + char)
	fmt.Println("apakah akan melanjutkan ke step selanjutnya y/n ?")
	reader2 := bufio.NewReader(os.Stdin)
	
	s,e := reader2.ReadString('\n')
	
	if (e != nil){
		fmt.Println(e)
	}

	if (s == "y") {
		res := getInput()
		fmt.Println(res)
	}
	 if (s == "n") {
		fmt.Println("program selesai")
	 }

	 fmt.Println(s)
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin);
	s,e := reader.ReadString('\n')
	if (e != nil){
		fmt.Println(e)
	}
	return s
}


