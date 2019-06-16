package logger

import "log"

func Log(host string, msg string) {

	prefix := "host " + host + ": "

	log.Println(prefix + msg)

}

func Fatal(host string, msg string, err error) {
	prefix := "host " + host + ": "
	if err != nil {
		log.Fatalln(prefix+msg, err)
	} else {
		log.Fatalln(prefix + msg)
	}

}
