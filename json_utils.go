package commons

import "encoding/json"

func ParseJsonString(data interface{}) string {
    jsonB, _ := json.Marshal(data)
    if jsonB == nil {
        return ""
    }
    return string(jsonB)
}
