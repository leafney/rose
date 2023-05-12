package rose

import (
	"bytes"
	"encoding/gob"
)

func GobSerializeErr(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GobSerialize Gob 序列化
func GobSerialize(data interface{}) []byte {
	bt, _ := GobSerializeErr(data)
	return bt
}

func GobDeserializeErr(data []byte, obj interface{}) error {
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(obj)
	if err != nil {
		return err
	}
	return nil
}

// GobDeserialize 反序列化
func GobDeserialize(data []byte, obj interface{}) {
	_ = GobDeserializeErr(data, obj)
}
