package servers

type Server struct {
	Id                  int
	Name                string
	AddressIPv4         string
	AddressIPv6         string
	Description         string
	DistributionName    string
	DistributionVersion string
	PrivateKeyPath      string
	GroupId             int
}
