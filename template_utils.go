package commons

import (
    "fmt"
    "github.com/gin-contrib/multitemplate"
    "path/filepath"
    "runtime"
    "strings"
)

func CreateCommonRender(templatesDir string) multitemplate.Renderer {
    renderer := multitemplate.NewRenderer()

    layouts, err := filepath.Glob(fmt.Sprintf("%s/%s", templatesDir, "layouts/*.html"))
    if err != nil {
        panic(err.Error())
    }

    contents, err := filepath.Glob(fmt.Sprintf("%s/%s", templatesDir, "contents/**/*.html"))
    if err != nil {
        panic(err.Error())
    }

    for _, content := range contents {
        dirList := make([]string, 0)
        switch runtime.GOOS {
        case "windows":
            dirList = strings.Split(content, `\`)
        default:
            dirList = strings.Split(content, `/`)
        }
        dirName := dirList[len(dirList)-2]

        includes, err := filepath.Glob(fmt.Sprintf("%s/%s/%s/%s", templatesDir, "includes", dirName, "*.html"))
        if err != nil {
            panic(err.Error())
        }

        layoutCopy := make([]string, len(layouts))
        copy(layoutCopy, layouts)

        includeCopy := make([]string, len(includes))
        copy(includeCopy, includes)

        files := append(layoutCopy, includeCopy...)
        files = append(files, content)

        renderer.AddFromFiles(dirName, files...)
    }

    defaults, err := filepath.Glob(fmt.Sprintf("%s/%s", templatesDir, "default/*.html"))
    if err != nil {
        panic(err.Error())
    }

    for _, page := range defaults {
        renderer.AddFromFiles(filepath.Base(page), page)
    }

    return renderer
}

