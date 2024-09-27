// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
)

type CT_DynamicFilter struct {
	// Dynamic filter type
	TypeAttr ST_DynamicFilterType
	// Value
	ValAttr *float64
	// ISO Value
	ValIsoAttr *time.Time
	// Max Value
	MaxValAttr *float64
	// Max ISO Value
	MaxValIsoAttr *time.Time
}

func NewCT_DynamicFilter() *CT_DynamicFilter {
	ret := &CT_DynamicFilter{}
	ret.TypeAttr = ST_DynamicFilterType(1)
	return ret
}

func (m *CT_DynamicFilter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	if m.ValIsoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "valIso"},
			Value: fmt.Sprintf("%v", *m.ValIsoAttr)})
	}
	if m.MaxValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxVal"},
			Value: fmt.Sprintf("%v", *m.MaxValAttr)})
	}
	if m.MaxValIsoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxValIso"},
			Value: fmt.Sprintf("%v", *m.MaxValIsoAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DynamicFilter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TypeAttr = ST_DynamicFilterType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
			continue
		}
		if attr.Name.Local == "valIso" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.ValIsoAttr = &parsed
			continue
		}
		if attr.Name.Local == "maxVal" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.MaxValAttr = &parsed
			continue
		}
		if attr.Name.Local == "maxValIso" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.MaxValIsoAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_DynamicFilter: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DynamicFilter and its children
func (m *CT_DynamicFilter) Validate() error {
	return m.ValidateWithPath("CT_DynamicFilter")
}

// ValidateWithPath validates the CT_DynamicFilter and its children, prefixing error messages with path
func (m *CT_DynamicFilter) ValidateWithPath(path string) error {
	if m.TypeAttr == ST_DynamicFilterTypeUnset {
		return fmt.Errorf("%s/TypeAttr is a mandatory field", path)
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	return nil
}
