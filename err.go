/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2023-06-17 08:53
 * @Description:
 */

package rose

import (
	"fmt"
)

// ErrTry capture possible panic exceptions through recover
func ErrTry(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error [%v]", r)
		}
	}()

	fn()

	return err
}
