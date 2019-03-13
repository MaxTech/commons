package commons

import (
    "crypto/sha1"
    "crypto/sha256"
    "encoding/hex"
)

type hashUtils struct {
}

var HashUtils *hashUtils

func (hu *hashUtils)BytesSha256ToBytes(srcBytes []byte) (shaBytes []byte, err error) {
    hasher := sha256.New()
    _, err = hasher.Write(srcBytes)
    if err == nil {
        shaBytes = hasher.Sum(nil)
    }
    return shaBytes, err
}

func (hu *hashUtils)BytesSha256ToHexString(srcBytes []byte) (hexString string, err error) {
    shaBytes, err := hu.BytesSha256ToBytes(srcBytes)
    if err == nil {
        hexString = hex.EncodeToString(shaBytes)
        return "", err
    }
    return hexString, err
}

func (hu *hashUtils)StringSha256ToHexString(srcString string) (hexString string, err error) {
    shaBytes, err := hu.BytesSha256ToBytes([]byte(srcString))
    if err == nil {
        hexString = hex.EncodeToString(shaBytes)
    }
    return hexString, err
}

func (hu *hashUtils)BytesSha1ToBytes(srcBytes []byte) (shaBytes []byte, err error) {
    hasher := sha1.New()
    _, err = hasher.Write(srcBytes)
    if err == nil {
        shaBytes = hasher.Sum(nil)
    }
    return shaBytes, err
}

func (hu *hashUtils)BytesSha1ToHexString(srcBytes []byte) (hexString string, err error) {
    shaBytes, err := hu.BytesSha1ToBytes(srcBytes)
    if err == nil {
        hexString = hex.EncodeToString(shaBytes)
        return "", err
    }
    return hexString, err
}

func (hu *hashUtils)StringSha1ToHexString(srcString string) (hexString string, err error) {
    shaBytes, err := hu.BytesSha1ToBytes([]byte(srcString))
    if err == nil {
        hexString = hex.EncodeToString(shaBytes)
    }
    return hexString, err
}
