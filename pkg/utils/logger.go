package utils

import "log"

func LogInfo(host string, msg string) {

	prefix := "host " + host + ": "

	log.Println(prefix + msg)

}

func LogFatal(host string, msg string, err error) {
	prefix := "host " + host + ": "
	if err != nil {
		log.Fatalln(prefix+msg, err)
	} else {
		log.Fatalln(prefix + msg)
	}


}
