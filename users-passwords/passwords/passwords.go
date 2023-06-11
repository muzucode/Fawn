package passwords

type Password struct {
	Id               int
	Salt             string
	HashingAlgorithm string
	Hash             string
}
