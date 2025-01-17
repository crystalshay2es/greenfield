package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/bnb-chain/greenfield/testutil/keeper"
	"github.com/bnb-chain/greenfield/x/challenge/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestGetChallengeId(t *testing.T) {
	keeper, ctx := keepertest.ChallengeKeeper(t)
	keeper.SaveChallenge(ctx, types.Challenge{
		Id:            100,
		ExpiredHeight: 1000,
	})
	require.True(t, keeper.GetChallengeId(ctx) == 100)
}

func TestAttestedChallengeIds(t *testing.T) {
	keeper, ctx := keepertest.ChallengeKeeper(t)
	params := types.DefaultParams()
	params.AttestationKeptCount = 5
	keeper.SetParams(ctx, params)

	keeper.AppendAttestChallengeId(ctx, 1)
	keeper.AppendAttestChallengeId(ctx, 2)
	keeper.AppendAttestChallengeId(ctx, 3)
	require.Equal(t, []uint64{1, 2, 3}, keeper.GetAttestChallengeIds(ctx))

	keeper.AppendAttestChallengeId(ctx, 4)
	keeper.AppendAttestChallengeId(ctx, 5)
	keeper.AppendAttestChallengeId(ctx, 6)
	require.Equal(t, []uint64{2, 3, 4, 5, 6}, keeper.GetAttestChallengeIds(ctx))

	params.AttestationKeptCount = 8
	keeper.SetParams(ctx, params)
	keeper.AppendAttestChallengeId(ctx, 7)
	keeper.AppendAttestChallengeId(ctx, 8)
	require.Equal(t, []uint64{2, 3, 4, 5, 6, 7, 8}, keeper.GetAttestChallengeIds(ctx))

	params.AttestationKeptCount = 3
	keeper.SetParams(ctx, params)
	keeper.AppendAttestChallengeId(ctx, 9)
	require.Equal(t, []uint64{7, 8, 9}, keeper.GetAttestChallengeIds(ctx))

	params.AttestationKeptCount = 5
	keeper.SetParams(ctx, params)
	keeper.AppendAttestChallengeId(ctx, 10)
	require.Equal(t, []uint64{7, 8, 9, 10}, keeper.GetAttestChallengeIds(ctx))
}
