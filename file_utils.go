package commons

import (
    "bytes"
    "log"
    "os"
    "strings"
)

func CheckAndMakeDirs(dirPath string) {
    checkPath := strings.ToLower(dirPath)
    _, err := os.Stat(checkPath)
    if err != nil && os.IsNotExist(err) {
        err = os.MkdirAll(checkPath, os.ModePerm)
        if err != nil {
            log.Println("create dir error:", err.Error())
        }
    }
}

func FillByte(fixByteCount int, fileBuffer *bytes.Buffer) {
    for i := 0; i < fixByteCount; i++ {
        fileBuffer.Write([]byte{0x0})
    }
}

func CheckExists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }
    if os.IsNotExist(err) {
        return false
    }
    return false
}

func GetFileExtName(fileName string) string {
    fileNameSplit := strings.Split(fileName, ".")
    extName := fileNameSplit[len(fileNameSplit)-1]
    return extName
}
