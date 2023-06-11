package environments

type Environment struct {
	Id        int
	Name      string
	ApiKey    string
	DebugMode bool
	Position  int
	CreatedAt []uint8
	UpdatedAt []uint8
}
