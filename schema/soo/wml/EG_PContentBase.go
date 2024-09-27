// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/Alarak-Alex/unioffice"
)

type EG_PContentBase struct {
	CustomXml *CT_CustomXmlRun
	FldSimple []*CT_SimpleField
	Hyperlink *CT_Hyperlink
}

func NewEG_PContentBase() *EG_PContentBase {
	ret := &EG_PContentBase{}
	return ret
}

func (m *EG_PContentBase) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CustomXml != nil {
		secustomXml := xml.StartElement{Name: xml.Name{Local: "w:customXml"}}
		e.EncodeElement(m.CustomXml, secustomXml)
	}
	if m.FldSimple != nil {
		sefldSimple := xml.StartElement{Name: xml.Name{Local: "w:fldSimple"}}
		for _, c := range m.FldSimple {
			e.EncodeElement(c, sefldSimple)
		}
	}
	if m.Hyperlink != nil {
		sehyperlink := xml.StartElement{Name: xml.Name{Local: "w:hyperlink"}}
		e.EncodeElement(m.Hyperlink, sehyperlink)
	}
	return nil
}

func (m *EG_PContentBase) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_PContentBase:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXml"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXml"}:
				m.CustomXml = NewCT_CustomXmlRun()
				if err := d.DecodeElement(m.CustomXml, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "fldSimple"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "fldSimple"}:
				tmp := NewCT_SimpleField()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.FldSimple = append(m.FldSimple, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hyperlink"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "hyperlink"}:
				m.Hyperlink = NewCT_Hyperlink()
				if err := d.DecodeElement(m.Hyperlink, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on EG_PContentBase %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_PContentBase
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_PContentBase and its children
func (m *EG_PContentBase) Validate() error {
	return m.ValidateWithPath("EG_PContentBase")
}

// ValidateWithPath validates the EG_PContentBase and its children, prefixing error messages with path
func (m *EG_PContentBase) ValidateWithPath(path string) error {
	if m.CustomXml != nil {
		if err := m.CustomXml.ValidateWithPath(path + "/CustomXml"); err != nil {
			return err
		}
	}
	for i, v := range m.FldSimple {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/FldSimple[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.Hyperlink != nil {
		if err := m.Hyperlink.ValidateWithPath(path + "/Hyperlink"); err != nil {
			return err
		}
	}
	return nil
}
