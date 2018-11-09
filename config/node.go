package config

type Node struct {
	Host string

	Type             string
	StartBlockHeight int64
	Confirmations    int64
}
