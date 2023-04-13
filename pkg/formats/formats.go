package formats

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type typeMatch struct {
	pattern string
	method  string
	key     string
	params  []string
	value   string
}

func FindAllReplaces(sampleJson string, attributes []string, values []string) string {
	patterns := []string{
		`\$(splitOrNull)\((\w+),(\d+),(\d+)\)`,
		`\$(int)\((\w+)\)`,
		`\$(stringOrEmpty)\((\w+)\)`,
		`\$(date)\((\w+),"(.*)","(.*)"\)`,
	}

	for _, pattern := range patterns {
		r, _ := regexp.Compile(pattern)
		matches := r.FindAllStringSubmatch(sampleJson, 100)

		for _, v := range matches {
			tm := &typeMatch{
				pattern: v[0],
				method:  v[1],
				key:     v[2],
				params:  v[2:],
			}
			// Get the true value by attribute key
			tm.updateValue(attributes, values)

			result := tm.doIt()
			sampleJson = strings.Replace(sampleJson, tm.pattern, result, 1)
		}
	}

	// Find and replace all params
	for _, v := range attributes {

		tm := &typeMatch{
			pattern: "$" + v,
			method:  "stringOrEmpty",
			key:     v,
			params:  []string{},
		}
		// Get the true value by attribute key
		tm.updateValue(attributes, values)

		result := tm.doIt()
		sampleJson = strings.Replace(sampleJson, tm.pattern, result, 1)
	}

	return sampleJson
}

func (t *typeMatch) doIt() string {
	switch t.method {
	case "splitOrNull":
		if t.value != "" {
			startIndex, _ := strconv.Atoi(t.params[1])
			endIndex, _ := strconv.Atoi(t.params[2])
			return `"` + substr(t.value, startIndex, endIndex) + `"`
		}
		return "null"
	case "int":
		if t.value != "" {
			floatValue, _ := strconv.ParseFloat(t.value, 64)
			intValue := math.Ceil(floatValue)
			return fmt.Sprintf("%.f", intValue)
		}
		return "null"
	case "stringOrEmpty":
		if t.value != "" {
			return t.value
		}
		break
	case "date":
		if t.value != "" {
			format := t.params[1]
			letters := map[string]string{
				"d": "02",
				"m": "01",
				"y": "2006",
				"h": "15",
				"i": "04",
				"s": "05",
			}

			for l, d := range letters {
				format = strings.Replace(format, l, d, 10)
			}

			newDate, _ := time.Parse(time.DateOnly, t.value)
			if t.params[2] != "0" {
				days, _ := strconv.Atoi(t.params[2])
				newDate = newDate.AddDate(0, 0, days)
			}
			return `"` + newDate.Format(format) + `"`
		}
		return "null"
	}
	return ""
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	a := string(asRunes[start : start+length])

	return a
}

func (t *typeMatch) updateValue(attributes []string, values []string) {
	for index, v := range attributes {
		if v == t.key {
			t.value = values[index]
			return
		}
	}
	t.value = ""
}
