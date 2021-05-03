package main
// import
import ("net/http"
		"html/template")
// pages
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/proj", projPage)
	http.ListenAndServe(":2048", nil)
}
// home page
func home(page http.ResponseWriter, r *http.Request) {
	tahoma := "Eto sdelano dlia raboti saita"
	templ, _ := template.ParseFiles("pgs/home.html")
	templ.Execute(page, tahoma)
}
// projects list
func projPage(page http.ResponseWriter, r *http.Request) {
	
}
