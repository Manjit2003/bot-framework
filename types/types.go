package types

type Host string

type ReqParams struct {
	Params map[string]string
}

type ServiceConfig struct {
	Host       Host
	Token      string
	UserId     string
	Password   string
	WaBaNumber string
}

type SendTextMessageParams struct {
	To         string
	Text       string
	PreviewUrl string
	Header     string
	Footer     string
}
