package commons

import "encoding/hex"

func HexEncodeBytesToBytes(src []byte) []byte {
    dst := make([]byte, hex.EncodedLen(len(src)))
    _ = hex.Encode(dst, src)
    return dst
}

func HexDecodeBytesToBytes(src []byte) []byte {
    dst := make([]byte, hex.EncodedLen(len(src)))
    _, _ = hex.Decode(dst, src)
    return dst
}
