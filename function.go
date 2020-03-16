package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type request struct {
	Name string `json:"name"`
}

type response struct {
	Age uint8  `json:"age"`
	Sex string `json:"sex"`
}

func ListOfName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	in := []request{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "[{\"age\":0,\"sex\":\"unknown\"}]")
		return
	}

	err = json.Unmarshal([]byte(body), &in)
	if err != nil {
		fmt.Fprint(w, "[{\"age\":0,\"sex\":\"unknown\"}]")
		return
	}

	out := []response{}
	for _, req := range in {
		if req.Name == "Alice" {
			out = append(out, response{10, "female"})
		} else if req.Name == "Bob" {
			out = append(out, response{20, "male"})
		} else if req.Name == "Charlie" {
			out = append(out, response{30, "male"})
		} else {
			out = append(out, response{0, "unknown"})
		}
	}

	fmt.Fprint(w, "[")
	for i, res := range out {
		bytes, _ := json.Marshal(res)
		fmt.Fprintf(w, "%s", string(bytes))
		if i != len(out)-1 {
			fmt.Fprintf(w, ",")
		}
	}
	fmt.Fprint(w, "]")
}

// func main() {
// 	http.HandleFunc("/", ListOfName)
// 	http.ListenAndServe(":8080", nil)
// }
