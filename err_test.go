/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2023-06-17 08:57
 * @Description:
 */

package rose

import (
	"fmt"
	"testing"
)

func TestErrTry(t *testing.T) {

	err := ErrTry(func() {

		fmt.Println("hello")
		panic("not a world")
	})

	t.Log(err)

	fmt.Println("ok")
}
