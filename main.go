package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	/*
		filename := "Bagara - Live @ Radio Intense 23.12.2022 Melodic Techno & Progressive House DJ Mix 4K.mp3"
		encodedFilename := url.PathEscape(filename)
		log.Println(encodedFilename)
	*/

	// TODO: perform network discovery to find IP address of speaker from a provided speaker name

	// TODO: find a free port to use and spawn a http server
	// TODO: create a server that serves files from the filename directory

	/*
		url, err := url.Parse("http://192.168.1.134:1400")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(url)
	*/
	speakerIp := "192.168.1.134"
	sendRequest(speakerIp)
}

func sendRequest(speakerIp string) {
	client := http.Client{}

	service := "DeviceProperties"

	speakerUrl := fmt.Sprintf("http//:%s:1400", speakerIp)
	requestUrl := fmt.Sprintf("%s/%s/Control", speakerUrl, service)

	body := []byte(`<?xml version="1.0" ?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
  <s:Body>
    <u:GetHouseholdID xmlns:u="urn:schemas-upnp-org:service:DeviceProperties:1"/>
  </s:Body>
</s:Envelope>`)
	bodyBuf := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", requestUrl, bodyBuf)
	if err != nil {
		log.Fatalf("Failed to make request: %s, %v", requestUrl, err)
	}

	req.Header = http.Header{
		"Charset":      {"utf-8"},
		"Content-Type": {"text/xml"},
		"SOAPACTION": {
			"urn:schemas-upnp-org:service:DeviceProperties:1#GetHouseholdID",
		},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to handle response of %v, %v", req, err)
	}

	var resB []byte
	n, err := res.Body.Read(resB)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	log.Printf("Read %d bytes from response", n)
	log.Println(string(resB))
}
