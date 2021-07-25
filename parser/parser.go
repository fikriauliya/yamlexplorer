package parser

import (
	"fmt"
	"io/ioutil"

	"github.com/fikriauliya/yamlexplorer/entity"
	"gopkg.in/yaml.v2"
)

func extractHeader(rows []interface{}, order []string) []string {
	headers := make(map[string]bool)
	headerRow := rows[0].(map[interface{}]interface{})

	for header := range headerRow {
		headers[header.(string)] = true
	}
	result := make([]string, 0)
	if order != nil {
		for _, o := range order {
			if _, ok := headers[o]; ok {
				result = append(result, o)
			}
		}
	} else {
		for k := range headers {
			result = append(result, k)
		}
	}
	return result
}

func extractBody(header []string, rows []interface{}) []([]string) {
	body := make([]([]string), len(rows))
	headerOrder := make(map[string]int)
	for i, h := range header {
		headerOrder[h] = i
	}
	for i, row := range rows {
		row := row.(map[interface{}]interface{})
		body[i] = make([]string, len(row))
		for k, value := range row {
			var newValue string
			switch value.(type) {
			case string:
				newValue = value.(string)
			case int:
				newValue = fmt.Sprint(value.(int))
			}
			j := headerOrder[k.(string)]
			body[i][j] = newValue
		}
	}
	return body
}

func readAndUnmarshallYAML(path string) (*map[interface{}]interface{}, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read the file: %s", err)
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall YAML: %s", err)
	}
	return &m, err
}

func ParseYAML(path string) (*entity.Table, error) {
	m, err := readAndUnmarshallYAML(path)
	if err != nil {
		return nil, err
	}

	var order []string
	meta, ok := (*m)[interface{}("_meta")]
	if ok {
		meta := meta.(map[interface{}]interface{})
		o, ok := meta[interface{}("order")]
		if ok {
			o := o.([]interface{})
			order = make([]string, len(o))
			for i, item := range o {
				order[i] = item.(string)
			}
		}
	}

	for k, v := range *m {
		if k != "_meta" {
			rows := v.([]interface{})

			header := extractHeader(rows, order)
			body := extractBody(header, rows)
			return &entity.Table{
				Header: header,
				Body:   body}, nil
		}
	}
	return nil, fmt.Errorf("invalid YAML")
}
