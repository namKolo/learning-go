package messenger

import (
	"encoding/json"
	"encoding/xml"
)

type Message struct {
	Body   []byte
	Format string
}

type MessageBuilder interface {
	SetRecipient(recipient string)
	SetText(text string)
	Message() (*Message, error)
}

type JSONMessageBuilder struct {
	recipient string
	text      string
}

func (b *JSONMessageBuilder) SetRecipient(recipient string) {
	b.recipient = recipient
}
func (b *JSONMessageBuilder) SetText(text string) {
	b.text = text
}

func (b *JSONMessageBuilder) Message() (*Message, error) {
	m := make(map[string]string)
	m["recipient"] = b.recipient
	m["message"] = b.text

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return &Message{Body: data, Format: "JSON"}, nil
}

type XMLMessageBuilder struct {
	recipient string
	text      string
}

func (b *XMLMessageBuilder) SetRecipient(recipient string) {
	b.recipient = recipient
}
func (b *XMLMessageBuilder) SetText(text string) {
	b.text = text
}

func (b *XMLMessageBuilder) Message() (*Message, error) {
	type XMLMessage struct {
		Recipient string `xml:"recipient"`
		Text      string `xml:"body"`
	}

	m := XMLMessage{
		Recipient: b.recipient,
		Text:      b.text,
	}

	data, err := xml.Marshal(m)
	if err != nil {
		return nil, err
	}

	return &Message{Body: data, Format: "XML"}, nil
}
