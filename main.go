package main
import (
	"net/http"
	"strings"
	"fmt"
	"time"
)
func sayHello(w http.ResponseWriter, r *http.Request) {
	// http://127.0.0.1:8080/?key=good
	r.ParseForm()  // parse arguments, you have to call this by yourself
	parsed := ""
	parsed += "r.URL.Path: " + r.URL.Path + "<br/>"
	parsed += "r.URL.Scheme: " + r.URL.Scheme + "<br/>"
	for k, v := range r.Form {
		parsed += k + ": " + strings.Join(v, "") + "<br/>"
	}

	serverTime := "<p>" + time.Now().String() + "</p><br/>"

	_, err := r.Cookie("username")
	if err!=nil {
		fmt.Println(err)
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "username", Value: "golang1234", Expires: expiration}
		http.SetCookie(w, &cookie)
	}

	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<body>
<h2>Cookie Content</h2>
<script>document.write(document.cookie);</script><h2>HTTP GET Content</h2><p>` + parsed + `</p>`+serverTime+`</body>
</html> `)
}
func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
