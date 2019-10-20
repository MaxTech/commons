package commons

import (
    "bytes"
    "github.com/google/uuid"
    "strings"
)

func GenUUIDString(_splitChar string) string {
    newUuid := uuid.New().String()
    if _splitChar == "-" {
        return newUuid
    }
    return strings.Replace(newUuid, "-", _splitChar, -1)
}

func GenUUIDBytes() *bytes.Buffer {
    newUuid := uuid.New()
    return bytes.NewBuffer(newUuid[:])
}
