package types

const (
	// ModuleName defines the module name
	ModuleName = "estore"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_estore"
)

var (
	ParamsKey = []byte("p_estore")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
