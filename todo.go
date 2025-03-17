package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	instruction := os.Args[1]
	var (
		taskId int
		arg    string
	)

	if len(os.Args) >= 3 {
		if getTask, err := strconv.Atoi(os.Args[2]); err == nil {
			taskId = getTask
		}

		// argList is used as temp memory to store only part with arguments
		argList := os.Args[2:]

		for i := 0; i < len(argList); i++ {
			arg += argList[i] + " "
		}
	}

	switch instruction {
	case "show":
		show()
	case "add":
		add(arg)
	case "delete":
		del(taskId)
	default:
		help()
	}
}

func show() {
	var result []string
	lastCut := 0

	f, err := os.ReadFile("todo.txt")

	if err != nil {
		fmt.Println("You must add some tasks first.")
	}

	// Loop inside of tasks stored in todo.txt file, separate them and store in array
	for i := 0; i < len(f); i++ {
		// Check for \n stuff
		if f[i] == 10 {
			task := f[lastCut:i]
			result = append(result, string(task))
			lastCut = i + 1
		}
	}

	// Loop inside array with separated tasks.
	// Two loops in one function? I would consider
	// this unefficient but I don't think that there's any better option
	for i := 0; i < len(result); i++ {
		fmt.Printf("%v. %s\n", i+1, result[i])
	}

}

func add(input string) {
	f, err := os.OpenFile("todo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	// To make it more redable
	input += "\n"

	f.WriteString(input)
}

func del(id int) {
	var tasks []string
	lastCut := 0

	f, err := os.ReadFile("todo.txt")
	if err != nil {
		fmt.Println("Something went wrong... ")
	}

	// Loop inside of tasks stored in todo.txt file, separate them and store in array like in the show function
	for i := 0; i < len(f); i++ {
		// Check for \n stuff
		if f[i] == 10 {
			task := f[lastCut:i]
			tasks = append(tasks, string(task))
			lastCut = i + 1
		}
	}

	// To make it coherent with displayed list
	id -= 1
	// And remove the element with given id
	tasks = append(tasks[:id], tasks[id+1:]...)

	// Now apply changes to todo.txt file
	os.Remove("todo.txt")

	for i := 0; i < len(tasks); i++ {
		add(tasks[i])
	}
}

func help() {
	fmt.Println("+--------------------------------------------------------+")
	fmt.Println("|                          Help                          |")
	fmt.Println("+--------------------------------------------------------+")
	fmt.Println("|  Param  |        Usage         |       Definition      |")
	fmt.Println("|---------+----------------------+-----------------------|")
	fmt.Println("|  show   |      todo show       |     Shows the list.   |")
	fmt.Println("|---------+----------------------+-----------------------|")
	fmt.Println("|   add   |   todo add <task>    |  Adds the given task. |")
	fmt.Println("|---------+----------------------+-----------------------|")
	fmt.Println("| delete  | todo delete <taskId> | Deletes the task with |")
	fmt.Println("|         |                      | given id.             |")
	fmt.Println("|---------+----------------------+-----------------------|")
	fmt.Println("|  help   |      todo help       |  Shows this window.   |")
	fmt.Println("+--------------------------------------------------------+")
}
