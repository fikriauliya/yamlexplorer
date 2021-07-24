package parser

import (
	"fmt"
	"io/ioutil"

	"github.com/fikriauliya/yamlexplorer/entity"
	"gopkg.in/yaml.v2"
)

func extractHeader(rows []interface{}) []string {
	var headers []string
	for _, row := range rows {
		row := row.(map[interface{}]interface{})
		headers = make([]string, len(row))
		for header := range row {
			headers = append(headers, header.(string))
		}
	}
	return headers
}

func extractBody(rows []interface{}) []([]string) {
	body := make([]([]string), len(rows))
	for i, row := range rows {
		row := row.(map[interface{}]interface{})
		body[i] = make([]string, len(row))
		for _, value := range row {
			body[i] = append(body[i], value.(string))
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

	for _, v := range *m {
		rows := v.([]interface{})

		header := extractHeader(rows)
		body := extractBody(rows)
		return &entity.Table{
			Header: header,
			Body:   body}, nil

	}
	return nil, fmt.Errorf("invalid YAML")
}
