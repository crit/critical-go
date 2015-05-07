package auth

type Authenticatable interface{}
type UserProvider interface{}
type Cookie interface{}
type Request interface{}
type EventDispatcher interface{}

type Session interface {
	Get(name string) (id int)
}
