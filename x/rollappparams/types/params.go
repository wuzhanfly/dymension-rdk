package types

import (
	"fmt"
	"regexp"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/dymensionxyz/dymint/da/registry"
	"github.com/dymensionxyz/gerr-cosmos/gerrc"
)

const (
	// length of the version commit string.
	VersionLength = 40
	// Data availability used by the RollApp. Default value used is mock da.
	DefaultDA = "mock"
)

// Parameter store keys.
var (
	KeyDa      = []byte("da")
	KeyVersion = []byte("version")

	// git commit for the version used for the rollapp binary. it must be overwritten in the build process
	Version = "<version>"
	// default max block size accepted (equivalent to block max size it can fit into a celestia blob).
	// regexp used to validate version commit
	VersionRegExp = regexp.MustCompile(`^[a-z0-9]*$`)
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params object
func NewParams(
	da string,
	version string,
) Params {
	return Params{
		Da:      da,
		Version: version,
	}
}

// DefaultParams returns default x/rollappparams module parameters.
func DefaultParams() Params {
	return Params{
		Da:      DefaultDA,
		Version: Version,
	}
}

func (p Params) Validate() error {
	err := ValidateDa(p.Da)
	if err != nil {
		return err
	}
	err = ValidateVersion(p.Version)
	if err != nil {
		return err
	}
	return nil
}

func ValidateDa(i any) error {
	if registry.GetClient(i.(string)) == nil {
		return fmt.Errorf("invalid DA type: DA %s: %w", i, gerrc.ErrInvalidArgument)
	}

	return nil

}

func ValidateVersion(i any) error {

	version, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid version type param type: %w", gerrc.ErrInvalidArgument)
	}
	if len(version) != VersionLength {
		return fmt.Errorf("invalid version length: param length: %d accepted: %d: %w", len(version), VersionLength, gerrc.ErrInvalidArgument)
	}
	if !VersionRegExp.MatchString(version) {
		return fmt.Errorf("invalid version: it must be alphanumeric %w", gerrc.ErrInvalidArgument)
	}

	return nil
}

// Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDa, &p.Da, ValidateDa),
		paramtypes.NewParamSetPair(KeyVersion, &p.Version, ValidateVersion),
	}
}
