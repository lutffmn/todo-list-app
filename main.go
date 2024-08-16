package main

import "fmt"

type Task struct {
	description string
	status      bool
}

func todoList(list []Task) {
	fmt.Printf("No.\tTask\t\tStatus\n")
	for i, v := range list {
		fmt.Printf("%d\t%s\t%v\n", (i + 1), v.description, v.status)
	}
}

func deleteTodo(number int, todo []Task) {

}

func main() {
	todo := []Task{
		{description: "wash the dish", status: false},
		{description: "walk the dog", status: false},
	}
	todoList(todo)
	for {
		choice := ""
		fmt.Println("\n1.Add todo\n2.Complete todo\n3.Delete todo")
		fmt.Println("Input your choice.")
		fmt.Scanln(&choice)

		if choice == "q" {
			break
		}
	}
}
