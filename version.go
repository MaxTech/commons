package commons

import "fmt"

var Version string

func init() {
    Version = fmt.Sprintf(
        "|- %s module:\t\t\t%s",
        "commons",
        "0.0.1")
}
