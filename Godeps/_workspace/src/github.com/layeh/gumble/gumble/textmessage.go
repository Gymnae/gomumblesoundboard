package gumble

import (
	"io"

	"github.com/layeh/gumble/gumble/MumbleProto"
)

// TextMessage is a chat message that can be received from and sent to the
// server.
type TextMessage struct {
	// User who sent the message (can be nil).
	Sender *User

	// Users that receive the message.
	Users []*User

	// Channels that receive the message.
	Channels []*Channel

	// Channels that receive the message and send it recursively to sub-channels.
	Trees []*Channel

	// Chat message.
	Message string
}

func (pm *TextMessage) writeTo(client *Client, w io.Writer) (int64, error) {
	packet := MumbleProto.TextMessage{
		Message: &pm.Message,
	}
	if pm.Users != nil {
		packet.Session = make([]uint32, len(pm.Users))
		for i, user := range pm.Users {
			packet.Session[i] = user.session
		}
	}
	if pm.Channels != nil {
		packet.ChannelId = make([]uint32, len(pm.Channels))
		for i, channel := range pm.Channels {
			packet.ChannelId[i] = channel.id
		}
	}
	if pm.Trees != nil {
		packet.TreeId = make([]uint32, len(pm.Trees))
		for i, channel := range pm.Trees {
			packet.TreeId[i] = channel.id
		}
	}
	proto := protoMessage{&packet}
	return proto.writeTo(client, w)
}
