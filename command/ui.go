package command

import (
	"encoding/json"
	"fmt"
)

type ui struct {
	output string
}

func (u *ui) PrettyPrint(output interface{}) {
	switch u.output {
	case "json":
		prettyJSON, _ := json.MarshalIndent(output, "", "    ")
		fmt.Printf("%s\n", prettyJSON)
		break
	case "yaml":
	case "table":
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
