package commons

import (
    "bytes"
    "fmt"
    "os"
    "strings"
    "time"
)

func CheckAndMakeDirs(_dirPath string) {
    checkPath := strings.ToLower(_dirPath)
    _, err := os.Stat(checkPath)
    if err != nil && os.IsNotExist(err) {
        err = os.MkdirAll(checkPath, os.ModePerm)
        if err != nil {
            _, _ = fmt.Fprintln(os.Stderr, time.Now().Format(time.RFC3339Nano), "[ERROR]", "create dir error:", err.Error())
        }
    }
}

func FillByte(_fixByteCount int, _fileBuffer *bytes.Buffer) {
    for i := 0; i < _fixByteCount; i++ {
        _fileBuffer.Write([]byte{0x0})
    }
}

func CheckExists(_path string) bool {
    _, err := os.Stat(_path)
    if err == nil {
        return true
    }
    if os.IsNotExist(err) {
        return false
    }
    return false
}

func GetFileExtName(_fileName string) string {
    fileNameSplit := strings.Split(_fileName, ".")
    extName := fileNameSplit[len(fileNameSplit)-1]
    return extName
}
