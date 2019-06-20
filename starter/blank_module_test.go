package starter

import (
	//"github.com/cosmos/cosmos-sdk/types/module"
	"encoding/json"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

type mockKeeper struct {
	mock.Mock
}

func TestNewBLankModule(t *testing.T) {

	mod := NewBlankModule(modName, mockKeeper{})

	require.Equal(t, mod.Name(), modName, "blank module should store its name")

}

func TestUnimplementedFuncs(t *testing.T) {
	mod := NewBlankModule(modName, mockKeeper{})

	assert.Panics(t, func() { mod.NewQuerierHandler() })

	assert.Panics(t, func() { mod.NewHandler() })

	require.Equal(t, mod.DefaultGenesis(), json.RawMessage(nil), "blank module returns nil for DefaultGenesis")

}
