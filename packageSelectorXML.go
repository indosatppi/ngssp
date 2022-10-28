package ngssp

import "encoding/xml"

type NgsspSelectorEnvelope struct {
	XMLName                     xml.Name                    `xml:"Envelope"`
	NgsspVariantPackageSelector NgsspVariantPackageSelector `xml:"Body>NgsspVariantPackageSelector"`
}

type NgsspVariantPackageSelector struct {
	CoherenceKey            string
	NgsspPackageKey         string
	NgsspVariantCommand     NgsspVariantCommand
	ListOfNgsspVariantRules []NgsspVariantRule `xml:"ListOfNgsspVariantRules>NgsspVariantRule"`
}

type NgsspVariantCommand struct {
	Command  string
	Channels []Channel `xml:"Channels>Channel"`
}

type Channel struct {
	Name       string
	Shortcodes []string `xml:"Shortcodes>Shortcode"`
}

type NgsspVariantRule struct {
	RuleId      string
	RuleDesc    string
	RuleVersion int
	NgsspRule   []NgsspRule
}

type NgsspRule struct {
	Type       string
	Operation  string
	Values     []string `xml:"Values>Value"`
	JoinMethod string
}
