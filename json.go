package rose

import (
	"encoding/json"
)

// JsonMarshal
func JsonMarshal(v interface{}) []byte {
	bt, _ := json.Marshal(v)
	return bt
}

// JsonMarshalErr
func JsonMarshalErr(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// JsonMarshalStr
func JsonMarshalStr(v interface{}) string {
	bt, _ := json.Marshal(v)
	return string(bt)
}

// JsonMarshalStrErr
func JsonMarshalStrErr(v interface{}) (string, error) {
	bt, err := json.Marshal(v)
	return string(bt), err
}

// JsonUnMarshal
func JsonUnMarshal(data []byte, v interface{}) {
	json.Unmarshal(data, v)
}

// JsonUnMarshalErr
func JsonUnMarshalErr(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// JsonUnMarshalStr
func JsonUnMarshalStr(s string, v interface{}) {
	json.Unmarshal([]byte(s), v)
}

// JsonUnMarshalStrErr
func JsonUnMarshalStrErr(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}
