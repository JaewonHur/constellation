package core

import (
	"errors"
	"io"
	"testing"

	"github.com/edgelesssys/constellation/cli/file"
	"github.com/edgelesssys/constellation/coordinator/attestation/vtpm"
	"github.com/edgelesssys/constellation/coordinator/state"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestAdvanceState(t *testing.T) {
	someErr := errors.New("failed")

	testCases := map[string]struct {
		initialState        state.State
		newState            state.State
		openTPMErr          error
		expectErr           bool
		expectOpenTPMCalled bool
	}{
		"init -> coordinator": {
			initialState:        state.AcceptingInit,
			newState:            state.ActivatingNodes,
			expectOpenTPMCalled: true,
		},
		"init -> node": {
			initialState:        state.AcceptingInit,
			newState:            state.IsNode,
			expectOpenTPMCalled: true,
		},
		"init -> failed": {
			initialState: state.AcceptingInit,
			newState:     state.Failed,
		},
		"uninit -> init": {
			initialState: state.Uninitialized,
			newState:     state.AcceptingInit,
		},
		"openTPM error": {
			initialState:        state.AcceptingInit,
			newState:            state.ActivatingNodes,
			openTPMErr:          someErr,
			expectErr:           true,
			expectOpenTPMCalled: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			openTPMCalled := false
			openTPM := func() (io.ReadWriteCloser, error) {
				openTPMCalled = true
				if tc.openTPMErr != nil {
					return nil, tc.openTPMErr
				}
				return vtpm.OpenSimulatedTPM()
			}

			core, err := NewCore(&stubVPN{}, nil, nil, nil, nil, nil, zaptest.NewLogger(t), openTPM, nil, file.NewHandler(afero.NewMemMapFs()))
			require.NoError(err)
			assert.Equal(state.Uninitialized, core.GetState())
			core.state = tc.initialState

			err = core.AdvanceState(tc.newState, []byte("secret"), []byte("cluster"))
			assert.Equal(tc.expectOpenTPMCalled, openTPMCalled)

			if tc.expectErr {
				assert.Error(err)
				assert.Equal(tc.initialState, core.GetState())
				return
			}
			require.NoError(err)

			assert.Equal(tc.newState, core.GetState())
		})
	}
}
