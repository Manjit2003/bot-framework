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

type TextMessageParams struct {
	To         string
	Text       string
	PreviewUrl string
	Header     string
	Footer     string
}

type ListRow struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"` // optional
}

type ListSection struct {
	Title string    `json:"title"`
	Rows  []ListRow `json:"rows"` // max 10 rows per section
}

type ListData struct {
	ButtonText string        `json:"button"`
	Sections   []ListSection `json:"sections"`
}

type ListMessageParams struct {
	*TextMessageParams
	ListData *ListData
}

type DocumentData struct {
	DocumentName string `json:"documentName"`
	DocumentUrl  string `json:"documentUrl"`
}

type DocumentMessageParams struct {
	*TextMessageParams
	DocumentData *DocumentData
}

/*{
  "buttons": [
    {
      "type": "reply",
      "reply": {
        "id": "id1",
        "title": "Button1"
      }
    },
    {
      "type": "reply",
      "reply": {
        "id": "id2",
        "title": "Button2"
      }
    },
    {
      "type": "reply",
      "reply": {
        "id": "id3",
        "title": "Button3"
      }
    }
  ]
}*/

type ButtonContent struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type Button struct {
	Type  string        `json:"type"`
	Reply ButtonContent `json:"reply"`
}

type ButtonData struct {
	*TextMessageParams
	ButtonTitle string   `json:"-"`
	Buttons     []Button `json:"buttons"`
}
