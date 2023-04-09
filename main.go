package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "", "CSV file path")
	jsonSample := flag.String("sample", "", "Json sample path")
	isCompact := flag.Bool("compact", false, "Compact view")
	outputFile := flag.String("output", "output.json", "Output file")
	showResult := flag.Bool("show", false, "Show result in stdout")
	flag.Parse()

	data, err := parse(*path, *jsonSample, *isCompact)
	if err != nil {
		fmt.Printf("Some error: \n %s", err)
		return
	}

	if *showResult {
		fmt.Println(data)
		return
	}

	f, _ := os.Create(*outputFile)
	_, err = f.WriteString(data)
	fmt.Printf("Result in %s \n", *outputFile)
	return

}

func parse(path string, jsonSample string, isCompact bool) (string, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("error opening %s\n", path)
	}

	reader := csv.NewReader(csvFile)
	template := readTemplateJson(jsonSample)

	entries, err := proceed(reader, template)

	result, err := json.MarshalIndent(entries, "", " ")
	if err != nil {
		return "", fmt.Errorf("Marshal error %s\n", err)
	}

	if isCompact {
		dst := &bytes.Buffer{}
		err = json.Compact(dst, result)
		return dst.String(), nil
	}

	return string(result), nil
}

func readTemplateJson(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

func proceed(reader *csv.Reader, sample string) ([]map[string]interface{}, error) {
	var (
		entries    []map[string]interface{}
		attributes []string
		firstRow   = false
	)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("failed to parse csv: %s", err)
		}

		if !firstRow {
			attributes = row
			firstRow = true
			continue
		}

		newTemplate := sample
		for i, value := range row {
			attribute := attributes[i]
			newTemplate = strings.Replace(newTemplate, "$"+attribute, ""+value+"", 1)
		}

		myStoredVariable := map[string]interface{}{}
		err = json.Unmarshal([]byte(newTemplate), &myStoredVariable)

		if err != nil {
			return nil, fmt.Errorf("failed to parse csv: %s", err)
		}

		entries = append(entries, myStoredVariable)
	}

	return entries, nil
}
