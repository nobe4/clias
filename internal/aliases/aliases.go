package aliases

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type (
	Alias   map[string][]string
	Aliases map[string]Alias
)

func Parse(in io.Reader) (Aliases, error) {
	a := Aliases{}

	if err := json.NewDecoder(in).Decode(&a); err != nil {
		return nil, fmt.Errorf("failed to decode the aliases: %w", err)
	}

	return a, nil
}

func (a Aliases) Find(binary string, args []string) []string {
	joinedArgs := strings.Join(args, " ")

	if alias, ok := a[binary]; ok {
		if args, ok := alias[joinedArgs]; ok {
			return args
		}
	}

	return nil
}

func (a Aliases) List(w io.Writer) {
	x, y := []string{}, []string{}

	for binary, alias := range a {
		for name, args := range alias {
			x = append(x, fmt.Sprintf("%s %s", binary, name))
			y = append(y, fmt.Sprintf("%s %s", binary, strings.Join(args, " ")))
		}
	}

	maxX := 0
	for _, x := range x {
		if l := len(x); l > maxX {
			maxX = l
		}
	}

	for i, x := range x {
		fmt.Fprintf(w, "%-*s => %s\n", maxX, x, y[i])
	}
}
