package types

const (
	// ModuleName defines the module name.
	ModuleName = "rollappparams"
	// StoreKey defines the primary store key.
	StoreKey = ModuleName
	// QuerierRoute is the querier route for the params store.
	QuerierRoute = StoreKey
)

var StateKey = []byte{0x01}
