package commons

import "encoding/json"

func ParseJsonString(_data interface{}) string {
    jsonB, _ := json.Marshal(_data)
    if jsonB == nil {
        return ""
    }
    return string(jsonB)
}
