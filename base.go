package barobill

import (
	"encoding/xml"
)

type Envelope[T any] struct {
	XMLName xml.Name `xml:"soap12:Envelope"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Soap12  string   `xml:"xmlns:soap12,attr"`
	Body    Body[T]  `xml:"soap12:Body"`
}

type ResEnvelope[T any] struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body[T]  `xml:"Body"`
}

type Body[T any] struct {
	Namespace string `xml:"-"`
	Body      T
}

func NewEnvelope[T any](t T) *Envelope[T] {
	return &Envelope[T]{
		Soap12: "http://www.w3.org/2003/05/soap-envelope",
		Xsi:    "http://www.w3.org/2001/XMLSchema-instance",
		Xsd:    "http://www.w3.org/2001/XMLSchema",
		Body: Body[T]{
			Body: t,
		},
	}
}
