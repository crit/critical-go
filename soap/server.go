package soap

type Server interface {
	Ping() bool
	WSDL() string
	SetWSDL() string
}
