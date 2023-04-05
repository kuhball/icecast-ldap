package main

import (
	"crypto/tls"
	"github.com/go-ldap/ldap/v3"
	"log"
	"net/http"
	"os"
)

type user struct {
	name     string
	password string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// uses env variables for binding to ldap server
func ldapCheck(user user) bool {
	var l *ldap.Conn

	if os.Getenv("ICECAST_AUTH_LDAP_SECURE") != "" {
		tlsConfig := &tls.Config{InsecureSkipVerify: true}
		var err error
		l, err = ldap.DialTLS("tcp", os.Getenv("ICECAST_AUTH_LDAP_SRV")+":636", tlsConfig)
		check(err)
	} else {
		var err error
		l, err = ldap.Dial("tcp", os.Getenv("ICECAST_AUTH_LDAP_SRV")+":389")
		check(err)
	}

	err := l.Bind("uid="+user.name+","+os.Getenv("ICECAST_AUTH_LDAP_DN"), user.password)
	if err != nil {
		// error in ldap bind
		log.Println(err)
		return false
	}
	// successful bind

	return true
}

//parses request and handles response for icecast
func handler(w http.ResponseWriter, r *http.Request) {
	var user user

	user.name = r.FormValue("user")
	user.password = r.FormValue("pass")

	if ldapCheck(user) {
		w.Header().Set("icecast-auth-user", "1")
		return
	} else {
		w.Header().Add("icecast-auth-user", "0")
		w.Header().Add("Icecast-Auth-Message", "error")
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":1337", nil)
	check(err)
}
