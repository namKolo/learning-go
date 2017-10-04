package messenger

type Sender struct{}

func (s *Sender) BuildMessage(builder MessageBuilder) (*Message, error) {
	builder.SetRecipient("MR KOlo")
	builder.SetText("Hihihi")
	return builder.Message()
}
