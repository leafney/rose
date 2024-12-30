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

// DirExists 判断目录路径是否存在
func DirExists(path string) bool {
	// 获取路径的状态信息
	info, err := os.Stat(path)
	if err != nil {
		// 如果路径不存在或发生其他错误
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	// 判断是否为目录
	return info.IsDir()
}

// Deprecated: use DirExists replaced
// DIsExist 判断目录路径是否存在
func DIsExist(path string) bool {
	exist, err := DIsExistE(path)
	if err != nil {
		return false
	}
	return exist
}

// Deprecated: use DirExists replaced
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

// DirExistsEnsure 确保目录路径存在 (判断路径是否存在，如果不存在则自动创建所需目录路径
func DirExistsEnsure(path string) error {
	dirPath := filepath.Dir(path)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
