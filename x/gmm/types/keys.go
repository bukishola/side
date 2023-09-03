package types

import fmt "fmt"

const (
	// ModuleName defines the module name
	ModuleName = "gmm"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_gmm"

	Version = "side-1"

	GAMMTokenPrefix = "side/gamm/"
)

var (
	// KeyPoolsPrefix defines prefix to store pools.
	KeyPoolsPrefix            = []byte{0x02}
	KeyCurrentPoolCountPrefix = []byte{0x03}
	KeyPoolIdToCountPrefix    = []byte{0x04}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetPoolShareDenom(poolId string) string {
	return fmt.Sprintf("%s%d", GAMMTokenPrefix, poolId)
}
