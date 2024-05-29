package formatting

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TitleCase(str string) string {
	titleCase := cases.Title(language.English)
	return titleCase.String(str)
}

// will turn into ex: hello_sekai
func ToFieldFormat(str string) (resp string) {
	// Replace spaces with underscores
	resp = strings.ReplaceAll(str, " ", "_")

	// Convert string to lowercase
	resp = strings.ToLower(resp)
	return
}

// will turn into ex: Hello world
func FromFieldFormat(str string) (resp string) {
	// Replace spaces with underscores
	resp = strings.ReplaceAll(str, "_", " ")

	// Convert string to Title case
	resp = TitleCase(resp)
	return
}

// will turn into ex: hello-sekai
func ToParamFormat(str string) (resp string) {
	// Replace spaces with underscores
	resp = strings.ReplaceAll(str, " ", "-")

	// Convert string to lowercase
	resp = strings.ToLower(resp)
	return
}

// From hello-world will turn into ex: Hello world
func FromParamFormat(str string) (resp string) {
	// Replace spaces with underscores
	resp = strings.ReplaceAll(str, "-", " ")
	// Convert string to Title case
	resp = TitleCase(resp)
	return
}
