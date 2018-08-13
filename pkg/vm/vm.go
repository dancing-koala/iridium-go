package vm

type VM struct {
	registers [32]int32
}

func New() *VM {
	return &VM{}
}
