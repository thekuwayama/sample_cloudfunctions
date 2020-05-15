package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "os"
)

type response struct {
	Method string              `json:"method"`
	Body   string              `json:"body"`
	Header map[string][]string `json:"header"`
}

func ParseHttpRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Error: read HTTP Request Body")
		return
	}

	res := response{
		Method: r.Method,
		Body:   string(body),
		Header: r.Header,
	}

	b, err := json.MarshalIndent(res, "", "\t")
	if err != nil {
		fmt.Fprint(w, "Error: json marshal")
		return
	}

	err = PostToSlack("#sakaitest", "cloudfunction", "```\n"+string(b)+"\n```", ":ghost:")
	if err != nil {
		fmt.Fprint(w, "Error: HTTP Post")
		return
	}

	fmt.Fprint(w, "```\n"+string(b)+"\n```")
}

type payload struct {
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Text     string `json:"text"`
	Icon     string `json:"icon_emoji"`
}

const webhookUrl = "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"

func PostToSlack(channel, username, text, icon string) error {
	p := payload{
		Channel:  channel,
		Username: username,
		Text:     text,
		Icon:     icon,
	}
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

	_, err = http.Post(webhookUrl, "application/json", bytes.NewReader(b))
	return err
}

// func main() {
//    	http.HandleFunc("/", ParseHttpRequest)
//    	err := http.ListenAndServe(":8080", nil)
//    	if err != nil {
//    		fmt.Fprintf(os.Stderr, "Error: starting http server %v", err)
//    		return
//  	}
// }
