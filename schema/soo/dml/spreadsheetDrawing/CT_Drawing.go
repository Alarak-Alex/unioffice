// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package spreadsheetDrawing

import (
	"encoding/xml"
	"fmt"

	"github.com/Alarak-Alex/unioffice"
)

type CT_Drawing struct {
	EG_Anchor []*EG_Anchor
}

func NewCT_Drawing() *CT_Drawing {
	ret := &CT_Drawing{}
	return ret
}

func (m *CT_Drawing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.EG_Anchor != nil {
		for _, c := range m.EG_Anchor {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Drawing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Drawing:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "twoCellAnchor"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "twoCellAnchor"}:
				tmpanchor := NewEG_Anchor()
				tmpanchor.TwoCellAnchor = NewCT_TwoCellAnchor()
				if err := d.DecodeElement(tmpanchor.TwoCellAnchor, &el); err != nil {
					return err
				}
				m.EG_Anchor = append(m.EG_Anchor, tmpanchor)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "oneCellAnchor"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "oneCellAnchor"}:
				tmpanchor := NewEG_Anchor()
				tmpanchor.OneCellAnchor = NewCT_OneCellAnchor()
				if err := d.DecodeElement(tmpanchor.OneCellAnchor, &el); err != nil {
					return err
				}
				m.EG_Anchor = append(m.EG_Anchor, tmpanchor)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "absoluteAnchor"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "absoluteAnchor"}:
				tmpanchor := NewEG_Anchor()
				tmpanchor.AbsoluteAnchor = NewCT_AbsoluteAnchor()
				if err := d.DecodeElement(tmpanchor.AbsoluteAnchor, &el); err != nil {
					return err
				}
				m.EG_Anchor = append(m.EG_Anchor, tmpanchor)
			default:
				unioffice.Log("skipping unsupported element on CT_Drawing %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Drawing
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Drawing and its children
func (m *CT_Drawing) Validate() error {
	return m.ValidateWithPath("CT_Drawing")
}

// ValidateWithPath validates the CT_Drawing and its children, prefixing error messages with path
func (m *CT_Drawing) ValidateWithPath(path string) error {
	for i, v := range m.EG_Anchor {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_Anchor[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
