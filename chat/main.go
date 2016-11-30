package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		w.Write(
			[]byte(`
				<html>
					<head>
						<title>Chat application</title>
					</head>
					<body>
						This is a chat application.
					</body>
				</html>
			`),
		)
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
