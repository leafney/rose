package rose

import "encoding/json"

func JsonMarshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func JsonUnMarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func JsonMarshalStr(v interface{}) (string, error) {
	bt, err := json.Marshal(v)
	return string(bt), err
}

func JsonUnMarshalStr(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}
