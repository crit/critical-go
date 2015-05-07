package auth

type Guard struct {
	user                    Authenticatable
	userSet                 bool
	lastAttempted           Authenticatable
	viaRemember             bool
	provider                UserProvider
	session                 Session
	cookie                  Cookie
	request                 Request
	events                  EventDispatcher
	loggedOut               bool
	tokenRetrievalAttempted bool
}

func NewGuard(provider UserProvider, session Session, request Request) Guard {
	g := Guard{}
	g.SetProvider(provider)
	g.SetSession(session)
	g.SetRequest(request)

	return g
}

func (g *Guard) SetProvider(provider UserProvider) {
	g.provider = provider
}

func (g *Guard) SetSession(session Session) {
	g.session = session
}

func (g *Guard) SetRequest(request Request) {
	g.request = request
}

func (g *Guard) Check() bool {
	_, err := g.User()

	return err == nil
}

func (g *Guard) User() (user Authenticatable, err error) {
	if g.loggedOut {
		return user, ErrorLoggedOut
	}

	if g.userSet {
		return g.user, nil
	}

	id := g.session.Get(g.getName())
}

func (g *Guard) getName() string {

}
