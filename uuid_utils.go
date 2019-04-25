package commons

import (
    "bytes"
    "github.com/google/uuid"
    "strings"
)

func GenUUIDString(spliteChar string) string {
    uuid := uuid.New().String()
    if spliteChar == "-" {
        return uuid
    }
    return strings.Replace(uuid, "-", spliteChar, -1)
}

func GenUUIDBytes() *bytes.Buffer {
    return bytes.NewBuffer(uuid.New())
}
