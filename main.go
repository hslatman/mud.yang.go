package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/hslatman/mud.yang.go/pkg/mudyang"
)

var fname = flag.String("f", "", "File(path) to MUD file to read")

func main() {

	flag.Parse()

	json, _ := ioutil.ReadFile(*fname)
	loadd := &mudyang.Mudfile{}
	if err := mudyang.Unmarshal([]byte(json), loadd); err != nil {
		panic(fmt.Sprintf("Can't unmarshal JSON: %v", err))
	}

	fmt.Println(*loadd.Mud.MudUrl)
	fmt.Println(*loadd.Mud.MudVersion)
	fmt.Println(loadd.Mud.MudSignature)

	for k, v := range loadd.Acls.Acl {
		fmt.Println(k, v)
	}

	// jsonString, err := ygot.EmitJSON(loadd, &ygot.EmitJSONConfig{
	// 	Format: ygot.RFC7951,
	// 	Indent: "  ",
	// 	RFC7951Config: &ygot.RFC7951JSONConfig{
	// 		AppendModuleName: true,
	// 	},
	// })

	// if err != nil {
	// 	panic(fmt.Sprintf("JSON error: %v", err))
	// }

	// fmt.Println(jsonString)

}
