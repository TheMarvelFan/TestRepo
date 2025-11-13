// twofer defines who the two cookies are for.
package twofer

import "fmt"

// ShareWith returns who the two cookies are for, while the second one will always be for me.
func ShareWith(name string) string {
    if name == "" {
		return "One for you, one for me."
    }

    return fmt.Sprintf("One for %s, one for me.", name)
}
