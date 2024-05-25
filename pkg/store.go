package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func (t *Todos) Store(filename string) {
	data, err := json.Marshal(t)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)

	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.New("fayly ady yalnys"))
		os.Exit(1)
	}
}
