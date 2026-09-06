package twofer

import "fmt"

// ShareWith returns a sentence sharing a cookie with the named person,
// or with "you" if no name is given.
func ShareWith(name string) string {
	if name != "" {
		return fmt.Sprintf("One for %s, one for me.", name)

	}
	return "One for you, one for me."

}
