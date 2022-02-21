package command

import (
	"encoding/json"
	"fmt"
)

type ui struct {
}

func (u *ui) PrettyPrint(output interface{}) {
	prettyJSON, _ := json.MarshalIndent(output, "", "    ")
	fmt.Printf("%s\n", prettyJSON)
}

func NewUI() *ui {
	return &ui{}
}

func (u *ui) Raw(out string) {
	fmt.Println(out)
}
