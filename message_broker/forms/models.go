package forms

type Message struct {
	From        string   `json:"adress"`
	To          string   `json:"destination"`
	Message     string   `json:"message"`
}

type User struct {
	Id          string   `json:"_id"`
	Nickname    string   `json:"nickname"`
	Hashed_pwd  string   `json:"hashed_pwd"`
	MessageBox  string   `json:"messageBox"`
}

type Request map[string]interface{}