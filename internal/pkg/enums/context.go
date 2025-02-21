package enums

type ContextKey string

const (
	ContextKeySession ContextKey = "session"
	ContextKeyTraceID ContextKey = "event_id_key"
)

func (key ContextKey) ToString() string {
	return string(key)
}
