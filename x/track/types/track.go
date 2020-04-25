package types

import (
	"fmt"
	"github.com/bitsongofficial/go-bitsong/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
	"strings"
	"time"
)

/************************************
 * Track
 ************************************/

type Track struct {
	Address       crypto.Address `json:"address" yaml:"address"`
	Path          string         `json:"path" yaml:"path"`
	Rewards       TrackRewards   `json:"rewards" yaml:"rewards"`
	RightsHolders RightsHolders  `json:"rights_holders" yaml:"rights_holders"`
	Owner         sdk.AccAddress `json:"owner" yaml:"owner"`
	Totals        TrackTotals    `json:"totals" yaml:"totals"`
	CreatedAt     time.Time      `json:"created_at" yaml:"created_at"`
}

func NewTrack(path string, rewards TrackRewards, rightsHolders RightsHolders, owner sdk.AccAddress) Track {
	return Track{
		Path:          path,
		Rewards:       rewards,
		RightsHolders: rightsHolders,
		Owner:         owner,
		Totals: TrackTotals{
			Streams:  0,
			Rewards:  sdk.NewCoin(types.BondDenom, sdk.ZeroInt()),
			Accounts: 0,
		},
	}
}

func (t Track) Validate() error {
	if len(strings.TrimSpace(t.Path)) == 0 {
		return fmt.Errorf("track path cannot be empty")
	}

	if len(t.Path) > MaxPathLength {
		return fmt.Errorf("track path cannot be longer than %d characters", MaxPathLength)
	}

	if err := t.Rewards.Validate(); err != nil {
		return err
	}

	if err := t.RightsHolders.Validate(); err != nil {
		return err
	}

	if t.Owner == nil {
		return fmt.Errorf("invalid track owner: %s", t.Owner)
	}

	return nil
}

// nolint
func (t Track) String() string {
	return fmt.Sprintf(`Address: %s
Path: %s
Rewards - %s
Rights Holders
%s
Created At: %s
Owner: %s
Totals
%s`,
		t.Address.String(), t.Path, t.Rewards.String(), t.RightsHolders,
		t.CreatedAt, t.Owner.String(), t.Totals.String(),
	)
}

func (t Track) Equals(track Track) bool {
	return t.Address.String() == track.Address.String() &&
		t.Path == track.Path &&
		t.Rewards.Equals(track.Rewards) &&
		t.RightsHolders.Equals(track.RightsHolders) &&
		t.Owner.Equals(track.Owner)
}

/************************************
 * Tracks
 ************************************/

// Tracks is an array of track
type Tracks []Track

// nolint
func (t Tracks) String() string {
	out := "Address - Owner\n"
	for _, track := range t {
		out += fmt.Sprintf("%s - %s\n",
			track.Address, track.Owner.String())
	}
	return strings.TrimSpace(out)
}
