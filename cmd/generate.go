package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/genutil"
	"github.com/openconfig/ygot/ygen"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates new mudyang model(s)",
	Long:  `Generates a new mudyang implementation using the MUD and dependent YANG specifications`,
	Run: func(cmd *cobra.Command, args []string) {

		// SOURCE: https://github.com/openconfig/ygot/blob/master/generator/generator.go#L298
		// No changes to the implementation has been made, apart from how the arguments
		// are provided to the actual code generation
		// See license in vendor/github.com/openconfig/ygot/LICENSE

		modsExcluded := []string{}
		skipEnumDedup := false
		ignoreCircDeps := false
		generateFakeRoot := true
		fakeRootName := "mudfile"
		shortenEnumLeafNames := false
		packageName := "mudyang"
		generateSchema := true

		ygotImportPath := genutil.GoDefaultYgotImportPath
		ytypesImportPath := genutil.GoDefaultYtypesImportPath
		goyangImportPath := genutil.GoDefaultGoyangImportPath

		generateRename := false
		addAnnotations := false
		annotationPrefix := ygen.DefaultAnnotationPrefix
		generateGetters := false
		generateDelete := false
		generateAppend := false
		generateLeafGetters := false
		includeModelData := false
		generateSimpleUnions := true

		compressPaths := false
		excludeState := false
		preferOperationalState := false

		compressBehaviour, _ := genutil.TranslateToCompressBehaviour(compressPaths, excludeState, preferOperationalState)

		cg := ygen.NewYANGCodeGenerator(&ygen.GeneratorConfig{
			Caller: "github.com/hslatman/mud.yang.go",
			ParseOptions: ygen.ParseOpts{
				ExcludeModules:        modsExcluded,
				SkipEnumDeduplication: skipEnumDedup,
				YANGParseOptions: yang.Options{
					IgnoreSubmoduleCircularDependencies: ignoreCircDeps,
				},
			},
			TransformationOptions: ygen.TransformationOpts{
				CompressBehaviour:    compressBehaviour,
				GenerateFakeRoot:     generateFakeRoot,
				FakeRootName:         fakeRootName,
				ShortenEnumLeafNames: shortenEnumLeafNames,
			},
			PackageName:        packageName,
			GenerateJSONSchema: generateSchema,
			GoOptions: ygen.GoOpts{
				YgotImportPath:       ygotImportPath,
				YtypesImportPath:     ytypesImportPath,
				GoyangImportPath:     goyangImportPath,
				GenerateRenameMethod: generateRename,
				AddAnnotationFields:  addAnnotations,
				AnnotationPrefix:     annotationPrefix,
				GenerateGetters:      generateGetters,
				GenerateDeleteMethod: generateDelete,
				GenerateAppendMethod: generateAppend,
				GenerateLeafGetters:  generateLeafGetters,
				IncludeModelData:     includeModelData,
				GenerateSimpleUnions: generateSimpleUnions,
			},
		})

		includePaths := []string{
			"yang/", // TODO: this on its own does not seem to work
		}

		generateModules := []string{
			"yang/ietf-packet-fields@2019-03-04.yang", // TODO: we currently provide these as modules, but I think importing should be enough?
			"yang/ietf-ethertypes@2019-03-04.yang",
			"yang/ietf-acldns.yang",
			"yang/ietf-inet-types.yang",
			"yang/ietf-access-control-list.yang",
			"yang/iana-tls-profile@2020-11-02.yang",            // NOTE: sourced from https://raw.githubusercontent.com/YangModels/yang/3af23949e11a2acd2f36df1dc0afca73ffe118ac/experimental/ietf-extracted-YANG-modules/iana-tls-profile@2020-11-02.yang
			"yang/ietf-acl-tls@2020-11-02.yang",                // NOTE: sourced from https://raw.githubusercontent.com/YangModels/yang/3af23949e11a2acd2f36df1dc0afca73ffe118ac/experimental/ietf-extracted-YANG-modules/ietf-acl-tls@2020-11-02.yang
			"yang/iana-hash-algs.yang",                         // NOTE: sourced from https://raw.githubusercontent.com/YangModels/yang/3af23949e11a2acd2f36df1dc0afca73ffe118ac/experimental/ietf-extracted-YANG-modules/iana-hash-algs@2020-03-08.yang
			"yang/ietf-netconf-acm.yang",                       // NOTE: sourced from https://raw.githubusercontent.com/huawei/yang/855d2d384d49fea03872e75fcea4d40619cf3528/network-router/8.20.0/atn980b/ietf-netconf-acm.yang
			"yang/ietf-crypto-types@2021-09-14.yang",           // NOTE: sourced from https://yangcatalog.org/YANG-modules/
			"yang/ietf-mud-transparency@2021-10-xx.yang",       // NOTE: currently sourced from https://github.com/elear/mud-sbom/commit/e8a1280a15f742c333f6222068df69c99f328de2
			"yang/ietf-ol@2021-05-21.yang",                     // NOTE: sourced from https://yangcatalog.org/YANG-modules/
			"yang/iana-opsawg-mud-tls-profile@2019-06-12.yang", // NOTE: sourced from https://raw.githubusercontent.com/YangModels/yang/3af23949e11a2acd2f36df1dc0afca73ffe118ac/experimental/ietf-extracted-YANG-modules/iana-opsawg-mud-tls-profile@2019-06-12.yang
			"yang/ietf-mud-tls@2020-10-19.yang",                // NOTE: sourced from https://raw.githubusercontent.com/YangModels/yang/3af23949e11a2acd2f36df1dc0afca73ffe118ac/experimental/ietf-extracted-YANG-modules/ietf-mud-tls@2020-10-19.yang
			"yang/ietf-mud@2019-01-28.yang",
		}

		generatedGoCode, errs := cg.GenerateGoCode(generateModules, includePaths)
		if errs != nil {
			fmt.Println(errs)
			return
		}

		var outfh *os.File
		ocStructsOutputFile := "pkg/mudyang/mudyang.go"

		// Assign the newly created filehandle to the outfh, and ensure
		// that it is synced and closed before exit of main.
		outfh = genutil.OpenFile(ocStructsOutputFile)
		defer genutil.SyncFile(outfh)

		writeGoCodeSingleFile(outfh, generatedGoCode)
	},
}

// SOURCE: https://github.com/openconfig/ygot/blob/master/generator/generator.go#L98
// No changes to the implementation have been made.
// See license in vendor/github.com/openconfig/ygot/LICENSE
//
// writeGoCodeSingleFile takes a ygen.GeneratedGoCode struct and writes the Go code
// snippets contained within it to the io.Writer, w, provided as an argument.
// The output includes a package header which is generated.
func writeGoCodeSingleFile(w io.Writer, goCode *ygen.GeneratedGoCode) error {
	// Write the package header to the supplier writer.
	fmt.Fprint(w, goCode.CommonHeader)
	fmt.Fprint(w, goCode.OneOffHeader)

	// Write the returned Go code out. First the Structs - which is the struct
	// definitions for the generated YANG entity, followed by the enumerations.
	for _, snippet := range goCode.Structs {
		fmt.Fprintln(w, snippet.String())
	}

	for _, snippet := range goCode.Enums {
		fmt.Fprintln(w, snippet)
	}

	// Write the generated enumeration map out.
	fmt.Fprintln(w, goCode.EnumMap)

	// Write the schema out if it was received.
	if len(goCode.JSONSchemaCode) > 0 {
		fmt.Fprintln(w, goCode.JSONSchemaCode)
	}

	if len(goCode.EnumTypeMap) > 0 {
		fmt.Fprintln(w, goCode.EnumTypeMap)
	}

	return nil
}
