package copperclient

import "encoding/json"

func marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func unmarshal(body []byte, v interface{}) error {
	return json.Unmarshal(body, v)
}
