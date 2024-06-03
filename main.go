package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Task структура для задачи со свойствами Name и Done
type Task struct {
	Name string
	Done bool
}

// ToDoList структура для списка задач
type ToDoList struct {
	Tasks []Task
}

// AddTask метод для добавления новой задачи в список
func (l *ToDoList) AddTask(taskName string) {
	l.Tasks = append(l.Tasks, Task{Name: taskName})
}

// PrintTasks метод для печати списка задач
func (l *ToDoList) PrintTasks() {
	for i, task := range l.Tasks {
		doneSymbol := " "
		if task.Done {
			doneSymbol = "x"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, doneSymbol, task.Name)
	}
}

// MarkTaskDone метод для отметки задачи выполненной
func (l *ToDoList) MarkTaskDone(index int) {
	if index >= 0 && index < len(l.Tasks) {
		l.Tasks[index].Done = true
	}
}

func main() {
	list := ToDoList{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nВведите задачу (или 'выход' для завершения):")
		scanner.Scan()
		text := scanner.Text()

		if strings.ToLower(text) == "выход" {
			break
		}

		list.AddTask(text)
		fmt.Println("Текущий список задач:")
		list.PrintTasks()
	}

	fmt.Println("\nВаш финальный список задач:")
	list.PrintTasks()
}
 
