package commons

import "encoding/hex"

type hexUtils struct {
}

var HexUtils *hexUtils

func (hu *hexUtils) HexEncodeBytesToBytes(src []byte) []byte {
    dst := make([]byte, hex.EncodedLen(len(src)))
    hex.Encode(dst, src)
    return dst
}

func (hu *hexUtils) HexDecodeBytesToBytes(src []byte) []byte {
    dst := make([]byte, hex.EncodedLen(len(src)))
    hex.Decode(dst, src)
    return dst
}
