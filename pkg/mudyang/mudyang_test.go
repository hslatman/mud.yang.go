package mudyang_test

import (
	"os"
	"testing"

	"github.com/hslatman/mud.yang.go/pkg/mudyang"
	"github.com/openconfig/ygot/ytypes"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	json, err := os.ReadFile("./../../examples/amazonEchoMud.json")
	require.NoError(t, err)

	mud := &mudyang.Mudfile{}
	err = mudyang.Unmarshal([]byte(json), mud)
	require.NoError(t, err)

	options := &ytypes.LeafrefOptions{
		IgnoreMissingData: false,
		Log:               true,
	}
	err = mud.ΛValidate(options)
	assert.NoError(t, err)
}
