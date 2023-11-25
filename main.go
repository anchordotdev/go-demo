package main

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"html/template"

	"io"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"

	_ "github.com/anchordotdev/anchor-go"
)

type IndexData struct {
	Time       string
	Configured bool
	URL        string
	Message    string
}

func main() {
	// setup the API server
	srv := &http.Server{

		Handler: http.HandlerFunc(api),
		// Handler: http.HandlerFunc(root),
	}

	// load the secret portion of the ACME EAB token
	acmeKey, err := base64.RawURLEncoding.DecodeString(os.Getenv("ACME_HMAC_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	// configure TLS via ACME provisioned certificates
	cfg := &tls.Config{
		GetCertificate: (&autocert.Manager{
			Prompt:      autocert.AcceptTOS,
			HostPolicy:  autocert.HostWhitelist(os.Getenv("HOST")),
			RenewBefore: 336 * time.Hour, // 14 days

			Client: &acme.Client{
				DirectoryURL: os.Getenv("ACME_DIRECTORY_URL"),
			},

			ExternalAccountBinding: &acme.ExternalAccountBinding{
				KID: os.Getenv("ACME_KID"),
				Key: acmeKey,
			},
		}).GetCertificate,
	}

	// provision a certificate and create the TLS listener
	ln, _ := tls.Listen("tcp", os.Getenv("ADDR"), cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Start the https server
	log.Fatal(srv.Serve(ln))
}

func api(w http.ResponseWriter, r *http.Request) {
	switch r.RequestURI {
	case "/":
		fmt.Println("request to / recieved")
		tmpl, _ := template.ParseFiles("index.html")
		configured, url, message := backend()
		data := IndexData{
			Time:       time.Now().Format("Mon Jan 2 15:04:05 MST 2006"),
			Configured: configured,
			URL:        url,
			Message:    message,
		}
		//render and serve index.html on "/"
		tmpl.Execute(w, data)
	case "/api":
		fmt.Println("request to /api recieve")
		response := `{"name":"go-demo", "message":"Hello from go-demo backend API"}`
		//render and serve json response for mock API endpoint
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(response))
	default:
		fmt.Println("unknown request: " + r.RequestURI)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found\n"))
	}
}

func backend() (bool, string, string) {
	url, ok := os.LookupEnv("BACKEND_URL")
	if ok {
		//ping backend
		message := ping_backend(url)
		return ok, url, message
	} else {
		return ok, "", ""
	}
}

func ping_backend(url string) string {
	// load the Localhost CA certificates.
	//pki.Init()

	// configure http client to use the anchor CA certificates.
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				//RootCAs: anchor.Certs.CertPool(),
			},
		},
	}

	res, _ := client.Get(url)
	body, _ := io.ReadAll(res.Body)
	return string(body)
}
