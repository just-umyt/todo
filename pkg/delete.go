package pkg

import (
	"errors"
	"fmt"
	"os"
)

func (t *Todos) Delete(index int) {
	ls := *t
	if index <= 0 || index > len(ls) {

		fmt.Fprintln(os.Stderr, errors.New("invalid index"))
		os.Exit(1)
	}

	*t = append(ls[:index-1], ls[index:]...)

}
