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

	"github.com/Alarak-Alex/unioffice/schema/soo/dml"
)

type WdCT_EffectExtent struct {
	LAttr dml.ST_Coordinate
	TAttr dml.ST_Coordinate
	RAttr dml.ST_Coordinate
	BAttr dml.ST_Coordinate
}

func NewWdCT_EffectExtent() *WdCT_EffectExtent {
	ret := &WdCT_EffectExtent{}
	return ret
}

func (m *WdCT_EffectExtent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "l"},
		Value: fmt.Sprintf("%v", m.LAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "t"},
		Value: fmt.Sprintf("%v", m.TAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
		Value: fmt.Sprintf("%v", m.RAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "b"},
		Value: fmt.Sprintf("%v", m.BAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_EffectExtent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "l" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.LAttr = parsed
			continue
		}
		if attr.Name.Local == "t" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.TAttr = parsed
			continue
		}
		if attr.Name.Local == "r" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.RAttr = parsed
			continue
		}
		if attr.Name.Local == "b" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.BAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing WdCT_EffectExtent: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the WdCT_EffectExtent and its children
func (m *WdCT_EffectExtent) Validate() error {
	return m.ValidateWithPath("WdCT_EffectExtent")
}

// ValidateWithPath validates the WdCT_EffectExtent and its children, prefixing error messages with path
func (m *WdCT_EffectExtent) ValidateWithPath(path string) error {
	if err := m.LAttr.ValidateWithPath(path + "/LAttr"); err != nil {
		return err
	}
	if err := m.TAttr.ValidateWithPath(path + "/TAttr"); err != nil {
		return err
	}
	if err := m.RAttr.ValidateWithPath(path + "/RAttr"); err != nil {
		return err
	}
	if err := m.BAttr.ValidateWithPath(path + "/BAttr"); err != nil {
		return err
	}
	return nil
}
