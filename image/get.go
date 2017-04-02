package image

import(
    "os"
    "io"
    "net/http"
    "log"
)

func Get(url *string, path string) {
    img, err := os.Create(path)
    if err != nil {
        log.Fatal(err)
    }
    defer img.Close()

    resp, err := http.Get(*url)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    if _, err := io.Copy(img, resp.Body); err != nil {
        log.Fatal(err)
    }

}