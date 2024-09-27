// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"
	"fmt"

	"github.com/Alarak-Alex/unioffice/schema/soo/dml"
)

type CT_TLIterateIntervalPercentage struct {
	// Value
	ValAttr dml.ST_PositivePercentage
}

func NewCT_TLIterateIntervalPercentage() *CT_TLIterateIntervalPercentage {
	ret := &CT_TLIterateIntervalPercentage{}
	return ret
}

func (m *CT_TLIterateIntervalPercentage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLIterateIntervalPercentage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_PositivePercentage(attr.Value)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TLIterateIntervalPercentage: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TLIterateIntervalPercentage and its children
func (m *CT_TLIterateIntervalPercentage) Validate() error {
	return m.ValidateWithPath("CT_TLIterateIntervalPercentage")
}

// ValidateWithPath validates the CT_TLIterateIntervalPercentage and its children, prefixing error messages with path
func (m *CT_TLIterateIntervalPercentage) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
