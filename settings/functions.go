package settings

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// read file
func (s *Struct) readFileContent() []byte {
	content, _ := ioutil.ReadFile(s.ConfigFilePath)
	return content
}

// Write value to yaml file
func (s *Struct) Write(key string, value interface{}) error {
	content := s.readFileContent()
	var data map[string]interface{}
	err := yaml.Unmarshal(content, &data)
	if err != nil {
		return err
	}
	if data == nil {
		data = make(map[string]interface{})
	}
	data[key] = value
	content, err = yaml.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(s.ConfigFilePath, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Delete  value from yaml file
func (s *Struct) Delete(key string) error {
	content := s.readFileContent()
	var data map[string]interface{}
	err := yaml.Unmarshal(content, &data)
	if err != nil {
		return err
	}
	delete(data, key)
	content, err = yaml.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(s.ConfigFilePath, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Update value in yaml file
func (s *Struct) Update(key string, value string) error {
	return s.Write(key, value)
}
