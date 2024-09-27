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
	"strconv"

	"github.com/Alarak-Alex/unioffice"
)

type CT_NumPicBullet struct {
	// Picture Numbering Symbol ID
	NumPicBulletIdAttr int64
	// Picture Numbering Symbol Properties
	Pict    *CT_Picture
	Drawing *CT_Drawing
}

func NewCT_NumPicBullet() *CT_NumPicBullet {
	ret := &CT_NumPicBullet{}
	return ret
}

func (m *CT_NumPicBullet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:numPicBulletId"},
		Value: fmt.Sprintf("%v", m.NumPicBulletIdAttr)})
	e.EncodeToken(start)
	if m.Pict != nil {
		sepict := xml.StartElement{Name: xml.Name{Local: "w:pict"}}
		e.EncodeElement(m.Pict, sepict)
	}
	if m.Drawing != nil {
		sedrawing := xml.StartElement{Name: xml.Name{Local: "w:drawing"}}
		e.EncodeElement(m.Drawing, sedrawing)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NumPicBullet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "numPicBulletId" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.NumPicBulletIdAttr = parsed
			continue
		}
	}
lCT_NumPicBullet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "pict"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "pict"}:
				m.Pict = NewCT_Picture()
				if err := d.DecodeElement(m.Pict, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "drawing"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "drawing"}:
				m.Drawing = NewCT_Drawing()
				if err := d.DecodeElement(m.Drawing, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_NumPicBullet %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NumPicBullet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NumPicBullet and its children
func (m *CT_NumPicBullet) Validate() error {
	return m.ValidateWithPath("CT_NumPicBullet")
}

// ValidateWithPath validates the CT_NumPicBullet and its children, prefixing error messages with path
func (m *CT_NumPicBullet) ValidateWithPath(path string) error {
	if m.Pict != nil {
		if err := m.Pict.ValidateWithPath(path + "/Pict"); err != nil {
			return err
		}
	}
	if m.Drawing != nil {
		if err := m.Drawing.ValidateWithPath(path + "/Drawing"); err != nil {
			return err
		}
	}
	return nil
}
