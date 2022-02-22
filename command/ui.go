package command

import (
	"encoding/json"
	"fmt"
	"os"
)

type ui struct {
	output string
}

func (u *ui) PrettyPrint(output interface{}) {
	switch u.output {
	case "json":
		u.printJson(output)
		break
	case "yaml":
		u.printYaml(output)
		break
	case "table":
		u.printTable(output)
		break
	default:
		fmt.Printf("Unsupported output format: %s\n", u.output)
	}

}

func NewUI() *ui {
	return &ui{output: "json"}
}

func (u *ui) Raw(out string) {
	fmt.Println(out)
}

func (u *ui) printJson(output interface{}) {
	prettyJSON, _ := json.MarshalIndent(output, "", "    ")
	fmt.Printf("%s\n", prettyJSON)
}

func (u *ui) printYaml(output interface{}) {
	fmt.Println("Yaml printing not yet supported")
	os.Exit(1)
}

func (u *ui) printTable(output interface{}) {
	fmt.Println("Table printing not yet supported")
	os.Exit(1)
}
