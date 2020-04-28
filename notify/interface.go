package notify

type Notifyer interface {
	Endpoint() string
	MaxChars() string
	Send() error
}