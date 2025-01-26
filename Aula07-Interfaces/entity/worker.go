package entity

import "time"

type Worker struct {
	Base
	Company string
	Role    string
	Salary  float64
}

func (worker *Worker) Deactivate() {
	worker.Ativo = false
	worker.UpdatedAt = time.Now()
	worker.DeletedAt = time.Now()

}

func (worker *Worker) Activate() {
	worker.Ativo = true
	worker.UpdatedAt = time.Now()
	worker.DeletedAt = time.Time{}
}
