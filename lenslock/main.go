package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site </h1> ")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>Contact Page</h1><p>To get in touch email me at 
	<a href="mailto:mathematics06physics@gmail.com">mathematics06physics@gmail.com</a></p>`)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `<div>
	<h1>FAQ Page</h1>
	<ul>
		<li>
			<p>Q: Is there a free version?</p>
			<p>A: Yes! we offer a free trial for 30days on any paid plans</p>
		</li>
		<li>
			<p>Q: What are your support hours?</p>
			<p>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends</p>
		</li>
		<li>
			<p>Q: How do I contact support?</p>
			<p>A: Email us - <a href="mailto:support@lenslock.com">support@lenslock.com</a></p>
		</li>
		</ul>
	</div>`
	fmt.Fprintf(w, html)
}

func notFounfHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprint(w, r.URL.Path, "\n")
// 	// fmt.Fprint(w, r.URL.RawPath, "\n")

// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		//TODO: handler page not found error
// 		notFounfHandler(w, r)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		//TODO: handler page not found error
		notFounfHandler(w, r)
	}
}

func main() {
	var router Router

	fmt.Println("starting the server on :3000")

	http.ListenAndServe(":3000", router)
}
