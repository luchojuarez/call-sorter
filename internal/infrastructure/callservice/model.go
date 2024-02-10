package callservice

import "time"

const (
	CallTypeNational      = "INTERNATIONAL"
	CallTypeInternational = "NATIONAL"
	CallTypeFriends       = "FRIENDS"
)

type Model struct {
	OriginNumber    string
	RecipientNumber string
	Duration        int
	Date            time.Time
	CallType        string
}
