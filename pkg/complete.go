package pkg

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func (t *Todos) Complete(index int) {
	ls := *t
	if index <= 0 || index > len(ls) {
		fmt.Fprintln(os.Stderr, errors.New("invalid index"))
		os.Exit(1)
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

}
