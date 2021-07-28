package parser

import (
	"fmt"
	"io/ioutil"

	"github.com/fikriauliya/yamlexplorer/entity"
	"gopkg.in/yaml.v2"
)

func extractHeader(mapSlice yaml.MapSlice) ([]string, error) {
	var dataSlice []interface{}
	for _, v := range mapSlice {
		if v.Key == "data" {
			dataSlice = v.Value.([]interface{})
			break
		}
	}
	if len(dataSlice) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	headerSlice := dataSlice[0].(yaml.MapSlice)
	res := make([]string, len(headerSlice))
	for i, val := range headerSlice {
		res[i] = val.Key.(string)
	}
	return res, nil
}

func extractBody(header []string, mapSlice yaml.MapSlice) ([]([]string), error) {
	var dataSlice []interface{}
	for _, v := range mapSlice {
		if v.Key == "data" {
			dataSlice = v.Value.([]interface{})
			break
		}
	}
	if len(dataSlice) == 0 {
		return nil, fmt.Errorf("no data found")
	}
	body := make([]([]string), len(dataSlice))
	headerOrder := make(map[string]int)
	for i, h := range header {
		headerOrder[h] = i
	}
	for i, row := range dataSlice {
		row := row.(yaml.MapSlice)
		body[i] = make([]string, len(row))
		for _, v := range row {
			var newValue string
			switch v.Value.(type) {
			case string:
				newValue = v.Value.(string)
			case int:
				newValue = fmt.Sprint(v.Value.(int))
			}
			j := headerOrder[v.Key.(string)]
			body[i][j] = newValue
		}
	}
	return body, nil
}

func readAndUnmarshallYAML(path string) (yaml.MapSlice, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read the file: %s", err)
	}
	var m yaml.MapSlice
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall YAML: %s", err)
	}
	return m, err
}

func ParseYAML(path string) (*entity.Table, error) {
	m, err := readAndUnmarshallYAML(path)
	if err != nil {
		return nil, err
	}

	header, err := extractHeader(m)
	if err != nil {
		return nil, err
	}
	body, err := extractBody(header, m)
	if err != nil {
		return nil, err
	}
	return &entity.Table{
		Header: header,
		Body:   body}, nil
}
