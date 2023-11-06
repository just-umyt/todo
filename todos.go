package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type item struct {
	Name        string    `json : "Name"`
	Done        bool      `json : "Done"`
	CreatedAt   time.Time `json : "CreatedAt"`
	CompletedAt time.Time `json : "CompletedAt"`
}

type Todos []item

func (todos *Todos) Add(s string) {
	newItem := item{
		Name:        s,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*todos = append(*todos, newItem)
}

func (todos *Todos) Complete(in int) error {
	if in <= 0 || in > len(*todos) {
		return errors.New("invalid index")
	}
	(*todos)[in-1].Done = true
	(*todos)[in-1].CompletedAt = time.Now()
	return nil
}

func (todos *Todos) Delete(in int) error {
	if in <= 0 || in > len(*todos) {
		return errors.New("invalid index")
	}

	*todos = append((*todos)[:in-1], (*todos)[in:]...)

	return nil
}

func (todos *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return errors.New("cant open the file")
	}

	if len(file) == 0 {
		return errors.New("pustoy json")
	}

	if er := json.Unmarshal(file, todos); er != nil {
		return er
	}

	return nil
}

func (todos *Todos) Store(filename string) error {
	data, err := json.Marshal(todos)
	if err != nil {
		return nil
	}

	return os.WriteFile(filename, data, 0644)

}
