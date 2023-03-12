/**
 * @Author:      leafney
 * @Date:        2023-03-12 20:08
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import "testing"

func TestSortedMap(t *testing.T) {
	//m := map[string]string{
	//	"name": "liming",
	//	"size": "29",
	//	"page": "2",
	//	"tag":  "hello",
	//}
	//a := SortedMap(m)

	a := SortedMapNew()
	a.Push("web", "true")
	a.Push("apple", "yes")
	t.Log(a)
	t.Log(a.Len())

}
