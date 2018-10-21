package structs

type Questions struct {
	Title     string `json:"title"`
	Questions []struct {
		Text    string   `json:"text"`
		Type    string   `json:"type"`
		Answers []string `json:"answers"`
		Answer  string   `json:"answer"`
	} `json:"questions"`
}

type Results []struct {
	User   string `json:"user"`
	Result string `json:"result"`
}

type Result struct {
	User   string `json:"user"`
	Result string `json:"result"`
}
