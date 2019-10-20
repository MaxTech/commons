package commons

import (
    "crypto/sha1"
    "crypto/sha256"
    "encoding/hex"
)

func BytesSha256ToBytes(_srcBytes []byte) ([]byte, error) {
    shaBytes := make([]byte, 0)
    newHash := sha256.New()
    _, err := newHash.Write(_srcBytes)
    if err == nil {
        shaBytes = newHash.Sum(nil)
    }
    return shaBytes, err
}

func BytesSha256ToHexString(_srcBytes []byte) (string, error) {
    hexString := ""
    shaBytes, err := BytesSha256ToBytes(_srcBytes)
    if err == nil {
        hexString = hex.EncodeToString(shaBytes)
        return "", err
    }
    return hexString, err
}

func StringSha256ToHexString(_srcString string) (string, error) {
    hexString := ""
    shaBytes, err := BytesSha256ToBytes([]byte(_srcString))
    if err == nil {
        hexString = hex.EncodeToString(shaBytes)
    }
    return hexString, err
}

func BytesSha1ToBytes(_srcBytes []byte) ([]byte, error) {
    shaBytes := make([]byte, 0)
    newHash := sha1.New()
    _, err := newHash.Write(_srcBytes)
    if err == nil {
        shaBytes = newHash.Sum(nil)
    }
    return shaBytes, err
}

func BytesSha1ToHexString(_srcBytes []byte) (string, error) {
    hexString := ""
    shaBytes, err := BytesSha1ToBytes(_srcBytes)
    if err == nil {
        hexString = hex.EncodeToString(shaBytes)
        return "", err
    }
    return hexString, err
}

func StringSha1ToHexString(_srcString string) (string, error) {
    hexString := ""
    shaBytes, err := BytesSha1ToBytes([]byte(_srcString))
    if err == nil {
        hexString = hex.EncodeToString(shaBytes)
    }
    return hexString, err
}
