package users

type User struct {
	Id         int
	Name       string
	UserConfig UserConfig
	PasswordId int
}

type UserConfig struct {
	Id        int
	UserTheme UserTheme
}

type UserTheme struct {
	Id           int
	Name         string
	ColorPalette ColorPalette
}

type ColorPalette struct {
	Id              int
	PrimaryColor    string
	SecondaryColor  string
	TertiaryColor   string
	QuaternaryColor string
}
