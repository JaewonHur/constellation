/*
Copyright (c) Edgeless Systems GmbH

SPDX-License-Identifier: AGPL-3.0-only
*/

package cmd

import (
	"context"
	"errors"
	"testing"

	"github.com/edgelesssys/constellation/v2/internal/attestation/measurements"
	"github.com/edgelesssys/constellation/v2/internal/cloud/cloudprovider"
	"github.com/edgelesssys/constellation/v2/internal/config"
	"github.com/edgelesssys/constellation/v2/internal/constants"
	"github.com/edgelesssys/constellation/v2/internal/file"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpgradeExecute(t *testing.T) {
	testCases := map[string]struct {
		upgrader stubUpgrader
		wantErr  bool
	}{
		"success": {
			upgrader: stubUpgrader{},
		},
		"upgrade error": {
			upgrader: stubUpgrader{err: errors.New("error")},
			wantErr:  true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)
			cmd := newUpgradeExecuteCmd()
			cmd.Flags().String("config", constants.ConfigFilename, "") // register persistent flag manually

			handler := file.NewHandler(afero.NewMemMapFs())
			cfg := defaultConfigWithExpectedMeasurements(t, config.Default(), cloudprovider.Azure)
			require.NoError(handler.WriteYAML(constants.ConfigFilename, cfg))

			err := upgradeExecute(cmd, tc.upgrader, handler)
			if tc.wantErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
		})
	}
}

type stubUpgrader struct {
	err error
}

func (u stubUpgrader) Upgrade(context.Context, string, measurements.M) error {
	return u.err
}
