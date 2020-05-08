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
  <html>
    <meta charset='UTF-8'>
    <body>
       <h1>Hello Docker <font color=‘green’>二階堂</font>   on CentOS7 .</h1>
    </body>
  </html>`

  w.Write([]byte(str))
}

func main() {
  http.HandleFunc("/", rootHandler)
  log.Fatal(http.ListenAndServe(":80", nil))
}

