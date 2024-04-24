/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2024-04-24 08:33
 * @Description:
 */

package rose

import (
	"os"
	"path/filepath"
)

// DIsExist 判断目录路径是否存在
func DIsExist(path string) bool {
	exist, err := DIsExistE(path)
	if err != nil {
		return false
	}
	return exist
}

// DIsExistE 判断目录路径是否存在，抛出异常
func DIsExistE(path string) (bool, error) {
	dirPath := filepath.Dir(path)
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

// DEnsurePathExist 确保目录路径存在 (判断路径是否存在，如果不存在则自动创建所需目录路径
func DEnsurePathExist(path string) error {
	dirPath := filepath.Dir(path)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
