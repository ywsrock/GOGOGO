package data

import "encoding/xml"

type Ca struct {
	XMLName    xml.Name `xml:"a,omitempty" json:"a,omitempty"`
	Attrhref   string   `xml:"href,attr"  json:",omitempty"`
	Attrtarget string   `xml:"target,attr"  json:",omitempty"`
	string     string   `xml:",chardata" json:",omitempty"`
}

type Cchannel struct {
	XMLName   xml.Name   `xml:"channel,omitempty" json:"channel,omitempty"`
	Citem     []*Citem   `xml:"item,omitempty" json:"item,omitempty"`
	Clanguage *Clanguage `xml:"language,omitempty" json:"language,omitempty"`
	Ctitle    *Ctitle    `xml:"title,omitempty" json:"title,omitempty"`
}

type Cdescription struct {
	XMLName xml.Name `xml:"description,omitempty" json:"description,omitempty"`
	Ca      *Ca      `xml:"a,omitempty" json:"a,omitempty"`
	Cfont   *Cfont   `xml:"font,omitempty" json:"font,omitempty"`
}

type Cfont struct {
	XMLName   xml.Name `xml:"font,omitempty" json:"font,omitempty"`
	Attrcolor string   `xml:"color,attr"  json:",omitempty"`
	string    string   `xml:",chardata" json:",omitempty"`
}

type Citem struct {
	XMLName      xml.Name      `xml:"item,omitempty" json:"item,omitempty"`
	Cdescription *Cdescription `xml:"description,omitempty" json:"description,omitempty"`
	Clink        *Clink        `xml:"link,omitempty" json:"link,omitempty"`
	CpubDate     *CpubDate     `xml:"pubDate,omitempty" json:"pubDate,omitempty"`
	Csource      *Csource      `xml:"source,omitempty" json:"source,omitempty"`
	Ctitle       *Ctitle       `xml:"title,omitempty" json:"title,omitempty"`
}

type Clanguage struct {
	XMLName xml.Name `xml:"language,omitempty" json:"language,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type Clink struct {
	XMLName xml.Name `xml:"link,omitempty" json:"link,omitempty"`
}

type CpubDate struct {
	XMLName xml.Name `xml:"pubDate,omitempty" json:"pubDate,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type Csource struct {
	XMLName xml.Name `xml:"source,omitempty" json:"source,omitempty"`
	Attrurl string   `xml:"url,attr"  json:",omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type Ctitle struct {
	XMLName xml.Name `xml:"title,omitempty" json:"title,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}
