package main

import (
	"fmt"
)

//ID not export
type ID int
//Priority not export
type Priority int
//Task not export
type Task struct {
	Num int
	Detail string
	done bool
}

func main() {
	fmt.Println("-Process-")
	Process()
	fmt.Println("-tasks-")
	tasks()
	
	fmt.Println("-NewTask-")
	//task := NewTask(1, "buy the milk")
    // &{ID:1 Detail:buy the milk done:false}
	//fmt.Printf("%+v", task)
	task := NewTask(1, "buy the milk")
	fmt.Printf("%s\n", task) // 1) buy the milk
	fmt.Println(task.Num, task.Detail, task.done)
}

//Process not export
func Process() {
	var id ID = 3
	var priority Priority = 5
	id, _ = ProcessTask(id, priority)
	fmt.Println(id)
}

//ProcessTask not export
func ProcessTask(id ID, priority Priority) (ID, Priority) {
	return id, priority
}

func tasks() {
	var task *Task = new(Task)

	var task1 Task = Task{
		Num: 1,
		Detail: "buy the milk",
		done: true,
	}
	var task2 Task = Task{
		2, "by the tomato", false,
	}
	fmt.Println(task.Num)
	fmt.Println(task.Detail)
	fmt.Println(task.done)
	fmt.Println("---")
	fmt.Println(task1.Num)
	fmt.Println(task1.Detail)
	fmt.Println(task1.done)
	fmt.Println("---")
	fmt.Println(task2.Num)
	fmt.Println(task2.Detail)
	fmt.Println(task2.done)
}

func NewTask(id int, detail string) *Task {
    task := &Task{
        Num: id,
        Detail: detail,
        done: false,
    }
    return task
}

func (task Task) String() string {
    str := fmt.Sprintf("%d) %s", task.Num, task.Detail)
    return str
}