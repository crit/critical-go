package soap

type Client interface {
	Server() Server
}

func NewClient(wsdl string) {}
