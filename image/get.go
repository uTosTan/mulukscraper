package image

import(
    "os"
    "io"
    "net/http"
    "log"
)

func Get(url string, path string) {
    img, _ := os.Create(path)
    defer img.Close()

    resp, _ := http.Get(url)
    defer resp.Body.Close()

    if _, err := io.Copy(img, resp.Body); err != nil {
        log.Fatal(err)
    }

}