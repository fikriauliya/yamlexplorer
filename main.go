package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type table struct {
	headers []string
	values  []([]string)
}

func parseFile(path string) (*table, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read the file: %s", err)
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall YAML: %s", err)
	}

	var headers []string
	var values []([]string)

	for _, v := range m {
		rows := v.([]interface{})
		values = make([]([]string), len(rows))

		firstRow := true
		for i, row := range rows {
			row := row.(map[interface{}]interface{})
			if firstRow {
				headers = make([]string, len(row))
				for header := range row {
					headers = append(headers, header.(string))
				}
				firstRow = false
			}
			values[i] = make([]string, len(headers))
			for _, value := range row {
				values[i] = append(values[i], value.(string))
			}
		}
		break
	}
	return &table{
		headers: headers,
		values:  values}, nil
}

func main() {
	path := "sample/index.yaml"
	table, err := parseFile(path)
	if err != nil {
		fmt.Println("Error parsing file: %s: %s", path, err)
		return
	}
	fmt.Println(table)
}
