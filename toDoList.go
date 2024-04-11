package main

import (
	"fmt"
	"os"
	"strings"
)

type activity struct {
	name    string
	checked bool
}

var activitiesList []activity

func creatTask(taskName string) {
	activitiesList = append(activitiesList, activity{name: taskName, checked: false})
	choice()
}

func checkTask(taskIndex int) {
	if activitiesList[taskIndex].checked == false {
		activitiesList[taskIndex].checked = true
		activitiesList[taskIndex].name += " -> X"
	} else {
		activitiesList[taskIndex].checked = false
		newName := activitiesList[taskIndex].name[:(len(activitiesList[taskIndex].name) - 5)]
		activitiesList[taskIndex].name = newName
	}
	choice()
}

func deleteTask(taskIndex int) {
	activitiesList = append(activitiesList[:taskIndex], activitiesList[taskIndex+1:]...)
	choice()
}

func showList() {
	for i := 0; i < len(activitiesList); i++ {
		fmt.Printf("%d. %s\n", i+1, activitiesList[i].name)
	}
}

func restart() {
	choice()
}

func choice() {
	choice := 0
	fmt.Printf("\n==============================\n")
	fmt.Println("To-do List!")
	fmt.Println("Your current To-Do List:")
	showList()
	fmt.Printf("==============================\n\n")
	fmt.Println("Instructions:")
	fmt.Println("1. Create a task")
	fmt.Println("2. Check a task")
	fmt.Println("3. Delete a task")
	fmt.Println("4. Exit application")
	fmt.Printf("\nYour Choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var taskName = ""
		fmt.Printf("All right! Your new task's name: ")
		fmt.Scan(&taskName)
		taskName = strings.TrimSpace(taskName)
		creatTask(taskName)
	case 2:
		fmt.Println("Ok, let's check it!")
		fmt.Print("Give me your task's number: ")
		var taskIndex = 0
		fmt.Scan(&taskIndex)
		taskIndex--
		if taskIndex+1 > 0 && taskIndex+1 <= len(activitiesList) {
			checkTask(taskIndex)
		} else {
			fmt.Println("Sorry, wrong number! Let's try again!")
			restart()
		}
	case 3:
		fmt.Println("Ok, let's delete it!")
		fmt.Print("Give me your task's number: ")
		var taskIndex = 0
		fmt.Scan(&taskIndex)
		taskIndex--
		if taskIndex+1 > 0 && taskIndex+1 <= len(activitiesList) {
			deleteTask(taskIndex)
		} else {
			fmt.Println("Sorry, wrong number! Let's try again!")
			restart()
		}
	case 4:
		fmt.Println("All right, see you later!")
		os.Exit(0)
	default:
		fmt.Println("Option doesn't exist! Try again!")
		restart()
	}

}

func main() {
	choice()
}
