// Click here and
package main

import (
	"fmt"
	"time"
)

type task struct {
	description string
	status      string
	createdDate time.Time
}

var taskList []task

func addTask() {
	var newTask task
	fmt.Println("Введите описание задачи:")
	fmt.Scanln(&newTask.description)
	newTask.status = "не выполнено"
	newTask.createdDate = time.Now()
	taskList = append(taskList, newTask)
}

func viewTasks() {
	if len(taskList) == 0 {
		fmt.Println("Список задач пуст")
		return
	}
	for i, t := range taskList {
		fmt.Printf("Задача %d: %s\nСтатус: %s\nДата создания: %s\n", i+1, t.description, t.status, t.createdDate.Format("02.01.2006 15:04:05"))
	}
}

func deleteTask() {
	var taskIndex int
	fmt.Println("Введите номер задачи, которую нужно удалить:")
	fmt.Scanln(&taskIndex)
	if taskIndex <= 0 || taskIndex > len(taskList) {
		fmt.Println("Задача не найдена")
		return
	}
	taskList = append(taskList[:taskIndex-1], taskList[taskIndex:]...)
	fmt.Println("Задача успешно удалена")
}

func changeTaskStatus() {
	var taskIndex int
	fmt.Println("Введите номер задачи, статус которой нужно изменить:")
	fmt.Scanln(&taskIndex)
	if taskIndex <= 0 || taskIndex > len(taskList) {
		fmt.Println("Задача не найдена")
		return
	}
	taskList[taskIndex-1].status = "выполнено"
	fmt.Println("Статус задачи успешно изменен")
}

func main() {
	fmt.Println("Программа управления списком задач")
	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Просмотреть список задач")
		fmt.Println("3. Удалить задачу")
		fmt.Println("4. Изменить статус задачи")
		fmt.Println("5. Выйти из программы")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask()
		case 2:
			viewTasks()
		case 3:
			deleteTask()
		case 4:
			changeTaskStatus()
		case 5:
			fmt.Println("Программа завершена")
			return
		default:
			fmt.Println("Неверный выбор")
		}
	}
}
