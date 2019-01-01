package slack

// ResponseType is a type of slack message type
type ResponseType string

const (
	// InChannel = everyone in the channel can view the message
	InChannel ResponseType = "in_channel"
	// Ephemeral = only the person who triggered the command can view the message
	Ephemeral ResponseType = "ephemeral"
)

// Response is a Slack Response struct
type Response struct {
	Text         string       `json:"text"`
	ResponseType ResponseType `json:"response_type,omitempty"`
}
