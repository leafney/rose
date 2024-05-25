/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2024-05-25 01:37
 * @Description:
 */

package rose

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	FlowBase = 1024
	FlowMB   = FlowBase
	FlowGB   = FlowBase * FlowBase
)

func formatFloat(value float64) string {
	if math.Mod(value, 1.0) == 0 {
		return fmt.Sprintf("%.0f", value)
	}
	return fmt.Sprintf("%.2f", value)
}

// FlowKBtoMB 流量转换 KB -> MB 示例：1500KB -> 1.46MB
func FlowKBtoMB(kb float64) string {
	return fmt.Sprintf("%.2fMB", kb/FlowBase)
}

// FlowKBtoGB 流量转换 KB -> GB 示例：2500000KB -> 2.38GB
func FlowKBtoGB(kb float64) string {
	return fmt.Sprintf("%.2fGB", kb/FlowBase/FlowBase)
}

// FlowKBtoGMB 流量转换 返回: 总 GB，总 MB,整数 GB+剩余 MB 示例：2500000KB -> 2.38GB 2441.41MB 2GB 393.41MB
func FlowKBtoGMB(kb float64) (gb, mb, gmb string) {

	theGb := math.Round(kb / FlowGB)
	theMb := math.Mod(kb, FlowGB) / FlowMB

	if theGb >= 1.0 {
		if theMb >= 1.0 {
			gmb = fmt.Sprintf("%vGB %vMB", formatFloat(theGb), formatFloat(theMb))
		} else {
			gmb = fmt.Sprintf("%vGB", formatFloat(theGb))
		}
	} else {
		gmb = fmt.Sprintf("%vMB", formatFloat(theMb))
	}

	// 总 GB
	gb = fmt.Sprintf("%.2fGB", kb/FlowGB)
	// 总 MB
	mb = fmt.Sprintf("%.2fMB", kb/FlowMB)

	return
}

func FlowMBtoKB(mb float64) string {
	return fmt.Sprintf("%.2fKB", mb*FlowBase)
}

func FlowMBtoMB(mb float64) string {
	return fmt.Sprintf("%.2fMB", mb)
}

func FlowMBtoGB(mb float64) string {
	return fmt.Sprintf("%.2fGB", mb/FlowBase)
}

// FlowMBtoGMB 流量转换 返回: 总 GB，总 MB,整数 GB+剩余 MB 示例：
func FlowMBtoGMB(mb float64) (gb, mmb, gmb string) {
	theGb := mb / FlowBase

	//if mb >= FlowBase {
	//	theMb := math.Mod(mb, FlowBase)
	//	if theMb == 0 {
	//		gmb = fmt.Sprintf("%.2fGB", theGb)
	//	} else {
	//		gmb = fmt.Sprintf("%.2fGB %.2fMB", theGb, theMb)
	//	}
	//} else {
	//	gmb = fmt.Sprintf("%.2fMB", mb)
	//}

	if theGb >= 1 {
		wholeGb := int(theGb)
		remainderMb := int((theGb - float64(wholeGb)) * FlowBase)
		if remainderMb == 0 {
			gmb = fmt.Sprintf("%dGB", wholeGb)
		} else {
			gmb = fmt.Sprintf("%dGB %dMB", wholeGb, remainderMb)
		}
	} else {
		gmb = fmt.Sprintf("%.0fMB", mb)
	}

	gb = fmt.Sprintf("%.2fGB", theGb)
	mmb = fmt.Sprintf("%.2fMB", mb)

	return
}

func FlowGBtoKB(gb float64) string {
	return fmt.Sprintf("%.2fKB", gb*FlowBase*FlowBase)
}

func FlowGBtoMB(gb float64) string {
	return fmt.Sprintf("%.2fMB", gb*FlowBase)
}

// FlowGBtoGB GB -> GB 示例：2.5 -> 2.50GB
func FlowGBtoGB(gb float64) string {
	return fmt.Sprintf("%.2fGB", gb)
}

// FlowGBtoGMB GB -> 流量转换 返回: 总 GB，总 MB,整数 GB+剩余 MB 示例：
func FlowGBtoGMB(gb float64) (ggb, mb, gmb string) {
	wholeGb := int(gb)
	theMb := (gb - float64(wholeGb)) * FlowBase

	if wholeGb >= 1 {
		if theMb == 0 {
			gmb = fmt.Sprintf("%dGB", wholeGb)
		} else {
			gmb = fmt.Sprintf("%.dGB %.0fMB", wholeGb, theMb)
		}
	} else {
		gmb = fmt.Sprintf("%.0fMB", gb*FlowBase)
	}

	ggb = fmt.Sprintf("%.2fGB", gb)
	mb = fmt.Sprintf("%.2fMB", gb*FlowBase)

	return
}

// FlowParseToKB 解析流量并转换为 KB 示例：1.1MB1.5kb -> 1127.9 -> 1127.90KB
// 支持流量单位 GB、MB、KB，不区分大小写；支持带有空格如：2GB 10MB
func FlowParseToKB(flow string) (float64, error) {
	//parts := strings.Fields(flow)
	//totalKB := 0.0
	//
	//for _, part := range parts {
	//	value, err := strconv.ParseFloat(part[:len(part)-2], 64)
	//	if err != nil {
	//		return 0, err
	//	}
	//	unit := strings.ToUpper(part[len(part)-2:])
	//	switch unit {
	//	case "GB":
	//		totalKB += value * FlowGB
	//	case "MB":
	//		totalKB += value * FlowMB
	//	case "KB":
	//		totalKB += value
	//	default:
	//		return 0, fmt.Errorf("unsupported unit: %s", unit)
	//	}
	//}

	re := regexp.MustCompile(`(\d+(\.\d+)?)([gmk]b)`)
	matches := re.FindAllStringSubmatch(strings.ToLower(flow), -1)
	totalKB := 0.0

	for _, match := range matches {
		value, err := strconv.ParseFloat(match[1], 64)
		if err != nil {
			return 0, err
		}

		unit := match[3]
		switch unit {
		case "gb":
			totalKB += value * FlowGB
		case "mb":
			totalKB += value * FlowMB
		case "kb":
			totalKB += value
		default:
			return 0, fmt.Errorf("unsupported unit: %s", unit)
		}
	}

	return totalKB, nil
}
