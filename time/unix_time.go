/**
 * @Author:      leafney
 * @Date:        2021-09-07 11:12
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package time

import "time"

func AddDateUnix(year, month, day int) int64 {
	return time.Now().AddDate(year, month, day).Unix()
}
