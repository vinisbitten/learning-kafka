package mail

type Mail struct {
	body []byte
	sender Sender
	receiver Receiver
}

type Sender struct {
	Id []byte
	Password []byte
}

type Receiver struct {
	host []byte
	port []byte
}