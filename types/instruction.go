package types

import "github.com/j0nnyboi/safecoin-go-sdk/common"

type CompiledInstruction struct {
	ProgramIDIndex int
	Accounts       []int
	Data           []byte
}

type Instruction struct {
	ProgramID common.PublicKey
	Accounts  []AccountMeta
	Data      []byte
}

type AccountMeta struct {
	PubKey     common.PublicKey
	IsSigner   bool
	IsWritable bool
}
