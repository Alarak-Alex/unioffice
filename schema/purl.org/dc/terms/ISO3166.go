// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package terms

import (
	"encoding/xml"
	"fmt"
)

type ISO3166 struct {
}

func NewISO3166() *ISO3166 {
	ret := &ISO3166{}
	return ret
}

func (m *ISO3166) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "ISO3166"
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *ISO3166) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing ISO3166: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the ISO3166 and its children
func (m *ISO3166) Validate() error {
	return m.ValidateWithPath("ISO3166")
}

// ValidateWithPath validates the ISO3166 and its children, prefixing error messages with path
func (m *ISO3166) ValidateWithPath(path string) error {
	return nil
}
