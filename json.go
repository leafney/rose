package rose

import "encoding/json"

// JsonMarshal
func JsonMarshal(v interface{}) []byte {
	bt, _ := json.Marshal(v)
	return bt
}

// JsonMarshalWithErr
func JsonMarshalWithErr(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// JsonMarshalStr
func JsonMarshalStr(v interface{}) string {
	bt, _ := json.Marshal(v)
	return string(bt)
}

// JsonMarshalStrWithErr
func JsonMarshalStrWithErr(v interface{}) (string, error) {
	bt, err := json.Marshal(v)
	return string(bt), err
}

// JsonUnMarshal
func JsonUnMarshal(data []byte, v interface{}) {
	json.Unmarshal(data, v)
}

// JsonUnMarshalWithErr
func JsonUnMarshalWithErr(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// JsonUnMarshalStr
func JsonUnMarshalStr(s string, v interface{}) {
	json.Unmarshal([]byte(s), v)
}

// JsonUnMarshalStrWithErr
func JsonUnMarshalStrWithErr(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}
