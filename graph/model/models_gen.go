// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type NewPageInput struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewUserInput struct {
	Name string `json:"name"`
}

type MarkupLanguage string

const (
	MarkupLanguageMarkdown MarkupLanguage = "MARKDOWN"
	MarkupLanguageASCIIDoc MarkupLanguage = "ASCIIDOC"
)

var AllMarkupLanguage = []MarkupLanguage{
	MarkupLanguageMarkdown,
	MarkupLanguageASCIIDoc,
}

func (e MarkupLanguage) IsValid() bool {
	switch e {
	case MarkupLanguageMarkdown, MarkupLanguageASCIIDoc:
		return true
	}
	return false
}

func (e MarkupLanguage) String() string {
	return string(e)
}

func (e *MarkupLanguage) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MarkupLanguage(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MarkupLanguage", str)
	}
	return nil
}

func (e MarkupLanguage) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
