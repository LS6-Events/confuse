package confuse

import "strings"

func extractValuesFromTag(tag string) (name string, options map[string]string) {
	split := strings.Split(tag, ",")
	name = split[0]
	options = make(map[string]string)

	for _, option := range split[1:] {
		optionSplit := strings.Split(option, "=")
		options[optionSplit[0]] = optionSplit[1]
	}

	return name, options
}
