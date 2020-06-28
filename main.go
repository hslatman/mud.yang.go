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
	mud := &mudyang.Mudfile{}
	if err := mudyang.Unmarshal([]byte(json), mud); err != nil {
		panic(fmt.Sprintf("Can't unmarshal JSON: %v", err))
	}

	println(*mud.Mud.MudUrl)
	println(*mud.Mud.MudVersion)
	println(mud.Mud.MudSignature)

	for k, v := range mud.Acls.Acl {
		println(k, v)
	}

	// jsonString, err := ygot.EmitJSON(mud, &ygot.EmitJSONConfig{
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
