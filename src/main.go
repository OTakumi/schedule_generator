package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Name  string
	Phase string
	Term  int
}

func main() {

}

func numberofTasks() int {
	s := bufio.NewScanner(os.Stdin)
	fmt.Println("How many tasks do you want to enter?")
	s.Scan() // Get tasks info from standard input

	// Converts the input value to a numerical value.
	n, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}

	fmt.Println(n)

	return n
}

func inputTasks(numTasks int) []Task {
	s := bufio.NewScanner(os.Stdin)
	num := numberofTasks()
	task := Task{}
	taskList := []Task{}

	// Get task information
	for i := 0; i < num; i++ {
		// 標準入力から情報を入力
		s.Scan()
		items := strings.Split(s.Text(), " ")

		// Partsにデータを代入
		task.Name = items[0]
		task.Phase = items[1]
		task.Term, _ = strconv.Atoi(items[2])

		// partsListに追加
		taskList = append(taskList, task)
		fmt.Printf("%d) Task name: %s phase: %s term: %d\n", i, task.Name, task.Phase, task.Term)
	}
	fmt.Println(taskList)

	return taskList
}
