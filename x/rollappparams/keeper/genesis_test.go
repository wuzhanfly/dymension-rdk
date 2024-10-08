package keeper_test

import (
	"testing"

	testkeepers "github.com/dymensionxyz/dymension-rdk/testutil/keepers"
	utils "github.com/dymensionxyz/dymension-rdk/testutil/utils"
	"github.com/dymensionxyz/dymension-rdk/x/rollappparams/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	app := utils.Setup(t, false)
	k, ctx := testkeepers.NewTestRollappParamsKeeperFromApp(app)

	state := types.GenesisState{}

	state.Params = types.NewParams("mock", "5f8393904fb1e9c616fe89f013cafe7501a63f86")

	k.InitGenesis(ctx, &state)
	got := k.ExportGenesis(ctx)
	require.NotNil(t, got)

	require.Equal(t, state.Params, got.Params)
}
