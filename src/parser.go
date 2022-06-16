package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	ConfigFilePath string
	ConfigFile     []byte

	Workspaces map[string]Workspace
}

// YAML type mapping
type Workspace struct {
	Provider string                       `yaml:"provider"`
	Domain   string                       `yaml:"domain"`
	Modules  map[string]map[string]string `yaml:"modules"`
}

// PUBLIC - Create a configuration object
func NewConfiguration(file string) *Configuration {
	c := new(Configuration)
	c.ConfigFilePath = file
	c.ConfigFile = c.readConfigurationFile(file)
	c.Workspaces = c.importWorkspaces(c.ConfigFile)

	return c
}

// PUBLIC - Dump the content of the current configuration
func (c *Configuration) Dump() {
	s, _ := json.MarshalIndent(c, "", "\t")
	fmt.Printf("%s", s)
}

// Parse file content and map the YAML with the Workspace structure
func (c *Configuration) importWorkspaces(file []byte) map[string]Workspace {
	w := make(map[string]Workspace)
	unmarshalError := yaml.Unmarshal(file, &w)
	if unmarshalError != nil {
		log.Fatal(unmarshalError)
	}

	return w
}

// Read configuration file and return content as []byte
func (c *Configuration) readConfigurationFile(filePath string) []byte {
	fileContent, readError := ioutil.ReadFile(filePath)
	if readError != nil {
		log.Fatal(readError)
	}

	return fileContent
}
