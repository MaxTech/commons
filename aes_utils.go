package commons

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "errors"
    "io"
)

//使用PKCS7进行填充，IOS也是7
func PKCS7Padding(_cipherText []byte, _blockSize int) []byte {
    padding := _blockSize - len(_cipherText)%_blockSize
    padText := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(_cipherText, padText...)
}

func PKCS7UnPadding(_originData []byte) []byte {
    length := len(_originData)
    unPadding := int(_originData[length-1])
    return _originData[:(length - unPadding)]
}

//aes加密，填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
func AesCBCEncrypt(_rawData, _key []byte) ([]byte, error) {
    block, err := aes.NewCipher(_key)
    if err != nil {
        return nil, err
    }

    //填充原文
    blockSize := block.BlockSize()
    _rawData = PKCS7Padding(_rawData, blockSize)
    //初始向量IV必须是唯一，但不需要保密
    cipherText := make([]byte, blockSize+len(_rawData))
    //block大小 16
    iv := cipherText[:blockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }

    //block大小和初始向量大小一定要一致
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(cipherText[blockSize:], _rawData)

    return cipherText, nil
}

func AesCBCDecrypt(_encryptData, _key []byte) ([]byte, error) {
    block, err := aes.NewCipher(_key)
    if err != nil {
        return nil, err
    }

    blockSize := block.BlockSize()

    if len(_encryptData) < blockSize {
        return nil, errors.New("ciphertext too short")
    }
    iv := _encryptData[:blockSize]
    _encryptData = _encryptData[blockSize:]

    // CBC mode always works in whole blocks.
    if len(_encryptData)%blockSize != 0 {
        return nil, errors.New("ciphertext is not a multiple of the block size")
    }

    mode := cipher.NewCBCDecrypter(block, iv)

    // CryptBlocks can work in-place if the two arguments are the same.
    mode.CryptBlocks(_encryptData, _encryptData)
    //解填充
    _encryptData = PKCS7UnPadding(_encryptData)
    return _encryptData, nil
}

func ZeroPadding(_cipherText []byte, _blockSize int) []byte {
    padding := _blockSize - len(_cipherText)%_blockSize
    padtext := bytes.Repeat([]byte{0}, padding)
    return append(_cipherText, padtext...)
}

func ZeroUnPadding(_originData []byte) []byte {
    length := len(_originData)
    unPadding := int(_originData[length-1])
    return _originData[:(length - unPadding)]
}
