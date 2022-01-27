package soaphandler

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"text/template"
)

type Request struct {
	//Values are set in below fields as per the request
	Blz string
}

func populateRequest() *Request {
	req := Request{}
	req.Blz = "61480001"

	return &req
}

func generateSOAPRequest(req *Request) (*http.Request, error) {
	// Using the var getTemplate to construct request
	template, err := template.New("InputRequest").Parse(getTemplate)
	if err != nil {
		fmt.Println("Error while marshling object. %s ", err.Error())
		return nil, err
	}

	doc := &bytes.Buffer{}
	// Replacing the doc from template with actual req values
	err = template.Execute(doc, req)
	if err != nil {
		fmt.Println("template.Execute error. %s ", err.Error())
		return nil, err
	}

	buffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(buffer)
	err = encoder.Encode(doc.String())
	if err != nil {
		fmt.Println("encoder.Encode error. %s ", err.Error())
		return nil, err
	}

	r, err := http.NewRequest(http.MethodPost, "http://www.thomas-bayer.com/axis2/services/BLZService", bytes.NewBuffer([]byte(doc.String())))
	if err != nil {
		fmt.Println("Error making a request. %s ", err.Error())
		return nil, err
	}

	return r, nil
}
