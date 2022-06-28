package xsnowflake

import (
	"testing"
)

func TestSnowflake(t *testing.T) {

	// init
	err := InitWorker(1)
	if err != nil {
		t.Error(err)
		return
	}

	// getNextId
	for i := 0; i < 100; i++ {
		if id, err := GetNextId(); err != nil {
			t.Error(err)
		} else {
			t.Log(id)
		}
	}

}
