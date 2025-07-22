package cmd

import (
	"flag"
	"fmt"
	"todo-cli/configs"
	"todo-cli/pkgs/actor"
	"todo-cli/pkgs/storage"
	"todo-cli/pkgs/task"
)

var (
	errConfigMissing = fmt.Errorf("Missing config")
)

type App struct {
	config *configs.Config
}

func NewApp(config *configs.Config) *App {
	if config == nil {
		panic(errConfigMissing)
	}

	return &App{
		config: config,
	}
}

func (a *App) Start() error {
	t := flag.String("task", "", "The task description")
	ac := flag.String("action", "add", "The action to perform")
	flag.Parse()

	if err := actor.NewActor().Validate(actor.Action(*ac)); err != nil {
		return err
	}

	storager := storage.NewStorage(a.config)

	switch actor.Action(*ac) {
	case actor.Add:
		task, err := task.NewTask(t)
		if err != nil {
			return err
		}

		storager.Add(*task)
	case actor.Remove:
		storager.Remove(*t)
	case actor.MarkCompleted:
		storager.Remove(*t)
		task, err := task.NewTask(t)
		if err != nil {
			return err
		}
		task.MarkCompleted()

		storager.Add(*task)
	case actor.MarkIncomplete:
		storager.Remove(*t)
		task, err := task.NewTask(t)
		if err != nil {
			return err
		}
		task.MarkIncomplete()

		storager.Add(*task)
	default:
		fmt.Println("Invalid action")
	}

	if err := storager.Save(); err != nil {
		return err
	}

	return nil
}
