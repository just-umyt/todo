package pkg

import (
	"time"

	"github.com/just-umyt/cli/models"
)

type Todos []models.Item

func (t *Todos) Add(task string) {
	todo := models.Item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}
