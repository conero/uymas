package str

import "strings"

// JsonTagName get json name omit option like `json:"name,omitempty"` and so on.
func JsonTagName(name string) string {
	if name == "" {
		return ""
	}

	// trim `,` like option name `omitempty`
	idx := strings.Index(name, ",")
	if idx > -1 {
		name = name[:idx]
	}

	return name
}
