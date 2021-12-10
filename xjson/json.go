package xjson

import "encoding/json"

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func UnMarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func MarshalStr(v interface{}) string {
	bt, _ := json.Marshal(v)
	return string(bt)
}

func MarshalStrErr(v interface{}) (string, error) {
	bt, err := json.Marshal(v)
	return string(bt), err
}

func UnMarshalStr(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}
