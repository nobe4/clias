package generators

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/nobe4/clias/internal/aliases"
)

var (
	ErrUnknownGenerator  = errors.New("unknown generator")
	ErrInvalidTemplate   = errors.New("invalid template")
	ErrExecutingTemplate = errors.New("error executing template")
)

func Generate(which string, a aliases.Aliases, w io.Writer) error {
	for g := range strings.SplitSeq(which, ",") {
		t, ok := tmpls[g]
		if !ok {
			return fmt.Errorf("%w: %s", ErrUnknownGenerator, g)
		}

		if err := gen(t, a, w); err != nil {
			return err
		}
	}

	return nil
}

func gen(c string, a aliases.Aliases, w io.Writer) error {
	t, err := template.New("").Parse(c)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidTemplate, err)
	}

	if err := t.Execute(w, a); err != nil {
		return fmt.Errorf("%w: %w", ErrExecutingTemplate, err)
	}

	return nil
}
