package oauth2_logger_proxy

import (
	"github.com/officeadminsorted/oauth2/v5/server"
	"log"
	"net/http"
)

func ProxyClientFormHandler(handler func(r *http.Request) (string, string, error)) server.ClientInfoHandler {
	return func(r *http.Request) (clientID, clientSecret string, err error) {
		log.Printf("ClientFormHandler(%vs)", r.Form)
		clientID, clientSecret, err = handler(r)
		log.Printf("ClientFormHandler() = %#v, %#v, %#v", clientID, clientSecret, err)
		return
	}
}
