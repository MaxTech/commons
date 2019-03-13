package commons

import (
    "strconv"
    "strings"
    "time"
)

func TrimAllEmpty(s string) string {
    return strings.Replace(s, " ", "", -1)
}

func Str2uint64(s string) uint64 {
    i, _ := strconv.ParseInt(s, 10, 64)
    return uint64(i)
}

func Str2time(s string) time.Time {
    t, _ := time.Parse("2006-01-02", s)
    return t
}

func Str2int64(s string) int64 {
    i, _ := strconv.ParseInt(s, 10, 64)
    return i
}

func Uint64ToStr(i uint64) string {
    return strconv.FormatUint(i,10)
}

func Int2str(i int) string {
    return strconv.Itoa(i)
}

func Float64Tostr(v float64) string {
    return strconv.FormatFloat(v, 'f', -1, 64)
}

func TrimCommaSplitStr(s string) []string {
    return strings.Split(strings.Trim(s,","),",")
}

func TimeStamp2StrData(stamp string) string {
    timeLayout := "2006-01-02"
    return timeStamp2Str(Str2int64(stamp), timeLayout)
}
func TimeStamp2StrDataTime(stamp string) string {
    timeLayout := "2006-01-02 15:04:05"
    return timeStamp2Str(Str2int64(stamp), timeLayout)
}

func timeStamp2Str(stamp int64, timeLayout string) string {
    return time.Unix(stamp, 0).Format(timeLayout)
}

func str2timeStamp(str string, timeLayout string) int64 {
    //timeLayout := "2006-01-02 15:04:05"
    loc, _ := time.LoadLocation("Local")
    theTime, _ := time.ParseInLocation(timeLayout, str, loc)
    return theTime.Unix()
}
