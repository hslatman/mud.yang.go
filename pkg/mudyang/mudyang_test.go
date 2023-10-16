package mudyang_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hslatman/mud.yang.go/pkg/mudyang"
)

func TestMudfile_ΛValidate(t *testing.T) {
	t.Parallel()
	leafRefOptions := &ytypes.LeafrefOptions{
		IgnoreMissingData: false,
		Log:               true,
	}
	type args struct {
		opts []ygot.ValidationOption
	}
	tests := []struct {
		name              string
		filepath          string
		args              args
		wantUnmarshalErr  bool
		wantValidationErr bool
	}{
		{
			name:     "ok/amazon-echo",
			filepath: "./../../examples/amazonEchoMud.json",
			args: args{
				opts: []ygot.ValidationOption{leafRefOptions},
			},
		},
		{
			name:     "ok/lightbulb-2000",
			filepath: "./../../examples/lightbulb2000.json",
			args: args{
				opts: []ygot.ValidationOption{leafRefOptions},
			},
		},
		{
			name:     "ok/wemo-switch",
			filepath: "./../../examples/wemoswitchMud.json",
			args: args{
				opts: []ygot.ValidationOption{leafRefOptions},
			},
		},
		{
			name:     "ok/sbom/mixed",
			filepath: "./../../examples/sbom/mixedExample.json",
			args: args{
				opts: []ygot.ValidationOption{leafRefOptions},
			},
		},
		{
			name:     "ok/sbom/complete",
			filepath: "./../../examples/sbom/completeExample.json",
			args: args{
				opts: []ygot.ValidationOption{leafRefOptions},
			},
		},
		{
			name:     "ok/tls/example",
			filepath: "./../../examples/tls/example.json",
			args: args{
				opts: []ygot.ValidationOption{leafRefOptions},
			},
		},
		{
			name:     "fail/amazon-echo",
			filepath: "./../../examples/invalidAmazonEchoMud.json",
			args: args{
				opts: []ygot.ValidationOption{leafRefOptions},
			},
			wantUnmarshalErr: true,
		},
	}
	for _, tt := range tests {
		tc := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data, err := os.ReadFile(tc.filepath)
			require.NoError(t, err)

			var v map[string]interface{}
			err = json.Unmarshal(data, &v)
			require.NoError(t, err)

			mud := &mudyang.Mudfile{}
			err = mudyang.Unmarshal(data, mud)
			if tc.wantUnmarshalErr {
				assert.Error(t, err)
				return
			}

			err = mud.ΛValidate(tc.args.opts...)
			if tc.wantValidationErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
