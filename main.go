package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Lutfen bir komut girin..")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(9999)
		load := loadTask()
		proccess := os.Args[2]
		load = append(load, Task{ID: n, Description: proccess, Status: "in-progress"})
		saveTask(load)
		fmt.Printf("%v Eklendi", proccess)
	case "update":
		load := loadTask()
		proccess := os.Args[2]
		proccessInt, err := strconv.Atoi(proccess)
		if err != nil {
			fmt.Println("Gecerli id girin")
		}

		updateContent := os.Args[3]

		for i, v := range load {
			if v.ID == proccessInt {
				load[i].Description = updateContent
				saveTask(load)
			}
		}
	case "list":
		laod := loadTask()
		for _, v := range laod {
			fmt.Printf("ID: %v \n Descrition: %v \n Status: %v \n CreatedAt: %v \n UpdatedAt: %v \n --- \n", v.ID, v.Description, v.Status, v.CreatedAt, v.UpdatedAt)
		}
	case "delete":
		load := loadTask()
		proccess := os.Args[2]
		proccessInt, err := strconv.Atoi(proccess)
		if err != nil {
			fmt.Println("Gecerli id girin")
		}

		target := proccessInt
		index := -1

		for i, v := range load {
			if v.ID == target {
				index = i
				break
			}
		}

		load = append(load[:index], load[index+1:]...)
		saveTask(load)

	case "mark-in-progress":
		load := loadTask()
		proccess := os.Args[2]
		proccessInt, err := strconv.Atoi(proccess)
		if err != nil {
			fmt.Println("Gecerli id girin")
		}
		for i, v := range load {
			if proccessInt == v.ID {
				load[i].Status = "in-progress"
				saveTask(load)
			}
		}
	case "mark-done":
		load := loadTask()
		proccess := os.Args[2]
		proccessInt, err := strconv.Atoi(proccess)
		if err != nil {
			fmt.Println("Gecerli id girin")
		}
		for i, v := range load {
			if proccessInt == v.ID {
				load[i].Status = "Done"
				saveTask(load)
			}
		}
	default:
		fmt.Println("Bilinmeyen komut!")

	}

}

func loadTask() []Task {
	if _, err := os.Stat("task.json"); os.IsNotExist(err) {
		f, err := os.Create("task.json")
		if err != nil {
			fmt.Println("dosya olusturmada hata", err)
		}
		f.WriteString("[]")
	}

	data, err := os.ReadFile("task.json")
	if err != nil {
		fmt.Println("dosya okumada hata", err)
	}
	var tasks []Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("unmarshall da sikinti", err)
	}

	return tasks
}

func saveTask(task []Task) {
	data, err := json.MarshalIndent(task, "", " ")
	if err != nil {
		fmt.Println("marshal da hata", err)
	}

	os.WriteFile("task.json", data, 0644)
}
