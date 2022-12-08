package mail

type Mail struct {
	Body     string
	Receiver Receiver
	Sender   Sender
	Smtp     SmtpServerConf
	Subject  string
}

type Sender struct {
	Id       string
	Password string
}

type Receiver struct {
	Id string
}

type SmtpServerConf struct {
	Host string
	Port string
}
