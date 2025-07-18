package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"todo-cli/pkgs/actor"
	"todo-cli/pkgs/storage"
	"todo-cli/pkgs/task"
)

func main() {
	t := flag.String("task", "", "The task description")
	ac := flag.String("action", "add", "The action to perform")
	flag.Parse()

	storager := storage.NewStorage()
	if err := actor.NewActor().Validate(actor.Action(*ac)); err != nil {
		panic(err)
	}

	switch actor.Action(*ac) {
	case actor.Add:
		task, err := task.NewTask(t)
		if err != nil {
			fmt.Println(err)
			return
		}

		storager.Add(*task)
	case actor.Remove:
		storager.Remove(*t)
	case actor.MarkCompleted:
		storager.Remove(*t)
		task, err := task.NewTask(t)
		if err != nil {
			fmt.Println(err)
			return
		}
		task.MarkCompleted()

		storager.Add(*task)
	case actor.MarkIncomplete:
		storager.Remove(*t)
		task, err := task.NewTask(t)
		if err != nil {
			fmt.Println(err)
			return
		}
		task.MarkIncomplete()

		storager.Add(*task)
	default:
		fmt.Println("Invalid action")
	}

	b, _ := json.Marshal(storager.GetData())
	os.WriteFile("result.json", b, 0644)
}
