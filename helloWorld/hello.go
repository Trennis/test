package main

import "fmt"

func main() {
	//n, err:= fmt.Println("Hello class, let's hear something from foo")
	n, _ := fmt.Println("Hello class, let's hear something from foo")
	fmt.Println(n)
	//fmt.Println(err)
	foo()
	fmt.Println("something more")
<<<<<<< HEAD
	fmt.Println("blablabla")
=======
	fmt.Println("new windows line")
>>>>>>> 971de377dce5aba09ef03725cd434ade881e93ff
	fmt.Println("Some new thing")

	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
	bar()
	test()
}

func foo() {
	fmt.Println("Hello, I am in FOO")
}

func bar() {
	fmt.Println("and then we exited...")
}

func test() {
	fmt.Println("This is not the end")
}
