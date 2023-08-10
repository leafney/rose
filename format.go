/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2023-08-10 15:06
 * @Description:
 */

package rose

import (
	"fmt"
	"strings"
)

type FmtItems map[string]interface{}

// FmtString Similar to formatting strings with named parameters in Python: `replace {msg} with {data}` {"msg":"hello","data":123}
func FmtString(template string, items FmtItems) string {
	for key, value := range items {
		template = strings.ReplaceAll(template, "{"+key+"}", fmt.Sprintf("%v", value))
	}
	return template
}

// FmtStringWith Similar to formatting strings with named parameters in Python;
// prefix default value `{`, suffix default value `}`
func FmtStringWith(prefix, suffix, template string, items FmtItems) string {
	if len(prefix) == 0 {
		prefix = "{"
	}
	if len(suffix) == 0 {
		suffix = "}"
	}

	for key, value := range items {
		template = strings.ReplaceAll(template, fmt.Sprintf("%s%s%s", prefix, key, suffix), fmt.Sprintf("%v", value))
	}
	return template
}
