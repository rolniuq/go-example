package saga

import "fmt"

type Worker struct {
	Action       func() error
	Compensation func() error
}

type Saga struct {
	workers []Worker
}

func NewSaga() *Saga {
	return &Saga{
		workers: make([]Worker, 0),
	}
}

func (s *Saga) AddWorker(worker Worker) *Saga {
	s.workers = append(s.workers, worker)

	return s
}

func (s *Saga) Run() error {
	var executed []Worker

	for _, worker := range s.workers {
		if err := worker.Action(); err != nil {
			for i := len(executed) - 1; i >= 0; i-- {
				if err := executed[i].Compensation(); err != nil {
					return err
				}
			}

			return fmt.Errorf("Saga rolling back %v", err)
		}
		executed = append(executed, worker)
	}

	return nil
}
