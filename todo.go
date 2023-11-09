package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
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

func (todos *Todos) Print() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *todos {
		idx++
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprint(idx)},
			{Text: item.Name},
			{Text: fmt.Sprint(item.Done)},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CreatedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: "Your todos are here"},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}
