package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Name string
	Done bool
}

type ToDoList struct {
	Tasks []Task
}

// Добавление задачи
func (l *ToDoList) AddTask(taskName string) {
	l.Tasks = append(l.Tasks, Task{Name: taskName})
}

// Печать задач
func (l *ToDoList) PrintTasks() {
	if len(l.Tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
	}
	for i, task := range l.Tasks {
		doneSymbol := " "
		if task.Done {
			doneSymbol = "x"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, doneSymbol, task.Name)
	}
}

// Отметка задачи как выполненной
func (l *ToDoList) MarkTaskDone(index int) {
	if index >= 0 && index < len(l.Tasks) {
		l.Tasks[index].Done = true
	} else {
		fmt.Println("Неверный индекс задачи")
	}
}

// Удаление задачи
func (l *ToDoList) RemoveTask(index int) {
	if index >= 0 && index < len(l.Tasks) {
		l.Tasks = append(l.Tasks[:index], l.Tasks[index+1:]...)
	} else {
		fmt.Println("Неверный индекс задачи")
	}
}

// Редактирование задачи
func (l *ToDoList) EditTask(index int, newName string) {
	if index >= 0 && index < len(l.Tasks) {
		l.Tasks[index].Name = newName
	} else {
		fmt.Println("Неверный индекс задачи")
	}
}

// Сохранение задач в файл
func (l *ToDoList) SaveToFile(filename string) error {
	data, err := json.Marshal(l.Tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// Загрузка задач из файла
func (l *ToDoList) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &l.Tasks)
}

func main() {
	list := ToDoList{}
	scanner := bufio.NewScanner(os.Stdin)

	// Загрузка задач из файла при запуске
	if err := list.LoadFromFile("tasks.json"); err != nil {
		fmt.Println("Не удалось загрузить задачи:", err)
	}

	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать задачи")
		fmt.Println("3. Отметить задачу выполненной")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Редактировать задачу")
		fmt.Println("6. Сохранить задачи")
		fmt.Println("7. Выход")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("Введите задачу:")
			scanner.Scan()
			taskName := scanner.Text()
			list.AddTask(taskName)
		case "2":
			fmt.Println("Текущий список задач:")
			list.PrintTasks()
		case "3":
			fmt.Println("Введите номер задачи для отметки выполненной:")
			scanner.Scan()
			var index int
			fmt.Sscanf(scanner.Text(), "%d", &index)
			list.MarkTaskDone(index - 1)
		case "4":
			fmt.Println("Введите номер задачи для удаления:")
			scanner.Scan()
			var index int
			fmt.Sscanf(scanner.Text(), "%d", &index)
			fmt.Printf("Вы уверены, что хотите удалить задачу '%s'? (y/n)\n", list.Tasks[index-1].Name)
			scanner.Scan()
			if strings.ToLower(scanner.Text()) == "y" {
				list.RemoveTask(index - 1)
			}
		case "5":
			fmt.Println("Введите номер задачи для редактирования:")
			scanner.Scan()
			var index int
			fmt.Sscanf(scanner.Text(), "%d", &index)
			fmt.Println("Введите новое название задачи:")
			scanner.Scan()
			newName := scanner.Text()
			list.EditTask(index - 1, newName)
		case "6":
			if err := list.SaveToFile("tasks.json"); err != nil {
				fmt.Println("Ошибка при сохранении задач:", err)
			} else {
				fmt.Println("Задачи успешно сохранены.")
			}
		case "7":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверный выбор. Пожалуйста, выберите действие от 1 до 7.")
		}
	}
}
