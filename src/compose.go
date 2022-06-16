package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"
)

// CONFIGURATION -----
// Constants
const CONFIG_FILE string = "config.yaml"

// GENERATION -----

// Templates types
// type Inventory struct {
// 	Material string
// 	Count    uint
// }

// @todo crézer un  packahe qui genere vmt des moduels et compagnie ?
// puis l'utiliser pour copmposer l'application, moins poirri que là

//go:embed templates/*.tmpl
var mainTemplate embed.FS

func main() {

	c := NewConfiguration(CONFIG_FILE)
	c.Dump()

	// Crawling workspaces
	for workspace, configuration := range c.Workspaces {
		// configuration.Domain
		// configuration.Provider
		// configuration.Modules

		// @todo inclure les infos, génrer un fichier par environnement ? faut que ça s'inclue bien
		// Hmmm

		tmpl := template.Must(template.ParseFS(mainTemplate, "templates/*.tmpl"))

		err := tmpl.ExecuteTemplate(os.Stdout, "main.tf.tmpl", configuration)
		if err != nil {
			panic(err)
		}

// 		{{ template "modules.tf.tmpl" . }}

// {{ template "output.tf.tmpl" . }}



		// Crawling modules
		for moduleName, variables := range configuration.Modules {

			// - generate modules list wiuth variables
			fmt.Printf("workspace:%s\n moduleName:%s\n variables:%s\n\n", workspace, moduleName, variables)

		}
	}

	// fmt.Printf("Hello world!")

	// // Read configuration file
	// configFile, readError := ioutil.ReadFile(CONFIG_FILE)
	// if readError != nil {
	// 	log.Fatal(readError)
	// }

	// // Parse YAML
	// configuration := make(map[string]User)
	// unmarshalError := yaml.Unmarshal(configFile, &configuration)
	// if unmarshalError != nil {
	// 	log.Fatal(unmarshalError)
	// }

	// // Parse configuration (DUMP)
	// for key, value := range configuration {
	// 	fmt.Printf("%s: %s\n", key, value)
	// }

	// // ---

	// // Template generation
	// // for each module, include module by name and generate each configfuration from configuration array, without rearranging file
	// sweaters := Inventory{"wool2", 10}
	// tmpl, err := template.ParseFiles("./templates/main.tf.tmpl")
	// if err != nil {
	// 	panic(err)
	// }
	// err = tmpl.Execute(os.Stdout, sweaters)
	// if err != nil {
	// 	panic(err)
	// }
}

// -----------

// Yaml types
// https://devopssec.fr/article/interfaces-golang
// https://docs.microsoft.com/fr-fr/learn/modules/go-methods-interfaces/1-methods
type ConfigurationInterface interface { // création de L'interface Forme
	Air() float64       // signature de la méthode Air()
	Perimetre() float64 // signature de la méthode Perimetre()
}
