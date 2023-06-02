package rose

import (
	"archive/zip"
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// FIsExist whether the file spcified by the given path is exists
func FIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)

}

// FIsDir whether the specified path is a directory
func FIsDir(path string) bool {
	fio, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}

	if err != nil {
		return false
	}

	return fio.IsDir()
}

// FRemove 移除指定路径文件
func FRemove(path string) error {
	return os.Remove(path)
}

// FZip zip the file and save it to destPath
func FZip(filePath string, destPath string) error {
	zipFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(filePath)+"/")

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return nil
}

// FUnZip unzip the file and save it to destPath
func FUnZip(zipFile string, destPath string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		path := filepath.Join(destPath, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// FReadFile 直接读取文件（适用于小文件读取）
func FReadFile(filePath string) (string, error) {
	res, err := os.ReadFile(filePath)
	return string(res), err
}

// FWriteFile 直接写入文件（适用于小文件写入）
func FWriteFile(filePath string, data string, append bool) error {
	flag := os.O_WRONLY | os.O_CREATE
	if append {
		flag |= os.O_APPEND
	} else {
		flag |= os.O_TRUNC
	}
	f, err := os.OpenFile(filePath, flag, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

func FReadBigFile(filePath string, writer io.Writer) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, 4096)
	for {
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := writer.Write(buffer[:n]); err != nil {
			return err
		}
	}
	return nil
}

// FWriteBigFile 写入大文件
func FWriteBigFile(filePath string, data string, append bool) error {
	//file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	return err
	//}
	//defer file.Close()
	//
	//buf := bufio.NewWriter(file)
	//_, err = buf.WriteString(data)
	//
	//io.WriteString()
	//
	//err = buf.Flush()
	//return err

	flag := os.O_WRONLY | os.O_CREATE
	if append {
		flag |= os.O_APPEND
	} else {
		flag |= os.O_TRUNC
	}
	file, err := os.OpenFile(filePath, flag, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return nil
}
