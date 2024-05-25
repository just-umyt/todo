package pkg

import (
	"fmt"
	"time"

	"github.com/alexeyco/simpletable"
)

func (t *Todos) Print() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "Created At"},
			{Align: simpletable.AlignCenter, Text: "Completed At"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		task := item.Task
		if item.Done {
			task = "✅" + item.Task
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: fmt.Sprintf("%t", item.Done)},
			{Text: item.CreatedAt.Format(time.ANSIC)},
			{Text: item.CompletedAt.Format(time.ANSIC)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: fmt.Sprintf("You have  %d pending todos", t.CountPending())},
	}}

	table.SetStyle(simpletable.StyleUnicode)
	table.Print()
}

func (t *Todos) CountPending() (total int) {
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}
	return
}
