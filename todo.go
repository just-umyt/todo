package todo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Name        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
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
	ls := *todos
	
	if in <= 0 || in > len(ls) {
		return errors.New("invalid index")
	}

	*todos = append(ls[:in-1], ls[in:]...)

	return nil
}

func (todos *Todos) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
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
