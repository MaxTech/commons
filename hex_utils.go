package commons

import "encoding/hex"

func HexEncodeBytesToBytes(_src []byte) []byte {
    dst := make([]byte, hex.EncodedLen(len(_src)))
    _ = hex.Encode(dst, _src)
    return dst
}

func HexDecodeBytesToBytes(_src []byte) []byte {
    dst := make([]byte, hex.EncodedLen(len(_src)))
    _, _ = hex.Decode(dst, _src)
    return dst
}
