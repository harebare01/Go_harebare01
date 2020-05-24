package main
import (
  "net/http"
  "log"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(200)
  w.Header().Set("(Content-Type", "text/html; charset=utf8")

  str := `
<!DOCTYPE html>
<html lang="ja">
<head><meta charset="utf-8">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.8.0/css/bulma.min.css">
  <title>Dockerデモ</title>
</head>
<body>
  <div class="has-text-centered">
  <dev class="is-size-1 has-text-danger">二階堂春雄with CI/CD</dev>
  </dev>
</body>
</html>`
  w.Write([]byte(str))
}

func main() {
  http.HandleFunc("/", rootHandler)
  log.Fatal(http.ListenAndServe(":80", nil))
}
