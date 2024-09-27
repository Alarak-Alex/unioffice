// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/Alarak-Alex/unioffice"
)

type CT_Shape3D struct {
	ZAttr            *ST_Coordinate
	ExtrusionHAttr   *int64
	ContourWAttr     *int64
	PrstMaterialAttr ST_PresetMaterialType
	BevelT           *CT_Bevel
	BevelB           *CT_Bevel
	ExtrusionClr     *CT_Color
	ContourClr       *CT_Color
	ExtLst           *CT_OfficeArtExtensionList
}

func NewCT_Shape3D() *CT_Shape3D {
	ret := &CT_Shape3D{}
	return ret
}

func (m *CT_Shape3D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ZAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "z"},
			Value: fmt.Sprintf("%v", *m.ZAttr)})
	}
	if m.ExtrusionHAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "extrusionH"},
			Value: fmt.Sprintf("%v", *m.ExtrusionHAttr)})
	}
	if m.ContourWAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "contourW"},
			Value: fmt.Sprintf("%v", *m.ContourWAttr)})
	}
	if m.PrstMaterialAttr != ST_PresetMaterialTypeUnset {
		attr, err := m.PrstMaterialAttr.MarshalXMLAttr(xml.Name{Local: "prstMaterial"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.BevelT != nil {
		sebevelT := xml.StartElement{Name: xml.Name{Local: "a:bevelT"}}
		e.EncodeElement(m.BevelT, sebevelT)
	}
	if m.BevelB != nil {
		sebevelB := xml.StartElement{Name: xml.Name{Local: "a:bevelB"}}
		e.EncodeElement(m.BevelB, sebevelB)
	}
	if m.ExtrusionClr != nil {
		seextrusionClr := xml.StartElement{Name: xml.Name{Local: "a:extrusionClr"}}
		e.EncodeElement(m.ExtrusionClr, seextrusionClr)
	}
	if m.ContourClr != nil {
		secontourClr := xml.StartElement{Name: xml.Name{Local: "a:contourClr"}}
		e.EncodeElement(m.ContourClr, secontourClr)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Shape3D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "z" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.ZAttr = &parsed
			continue
		}
		if attr.Name.Local == "extrusionH" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.ExtrusionHAttr = &parsed
			continue
		}
		if attr.Name.Local == "contourW" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.ContourWAttr = &parsed
			continue
		}
		if attr.Name.Local == "prstMaterial" {
			m.PrstMaterialAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_Shape3D:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "bevelT"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "bevelT"}:
				m.BevelT = NewCT_Bevel()
				if err := d.DecodeElement(m.BevelT, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "bevelB"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "bevelB"}:
				m.BevelB = NewCT_Bevel()
				if err := d.DecodeElement(m.BevelB, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extrusionClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extrusionClr"}:
				m.ExtrusionClr = NewCT_Color()
				if err := d.DecodeElement(m.ExtrusionClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "contourClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "contourClr"}:
				m.ContourClr = NewCT_Color()
				if err := d.DecodeElement(m.ContourClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Shape3D %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Shape3D
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Shape3D and its children
func (m *CT_Shape3D) Validate() error {
	return m.ValidateWithPath("CT_Shape3D")
}

// ValidateWithPath validates the CT_Shape3D and its children, prefixing error messages with path
func (m *CT_Shape3D) ValidateWithPath(path string) error {
	if m.ZAttr != nil {
		if err := m.ZAttr.ValidateWithPath(path + "/ZAttr"); err != nil {
			return err
		}
	}
	if m.ExtrusionHAttr != nil {
		if *m.ExtrusionHAttr < 0 {
			return fmt.Errorf("%s/m.ExtrusionHAttr must be >= 0 (have %v)", path, *m.ExtrusionHAttr)
		}
		if *m.ExtrusionHAttr > 27273042316900 {
			return fmt.Errorf("%s/m.ExtrusionHAttr must be <= 27273042316900 (have %v)", path, *m.ExtrusionHAttr)
		}
	}
	if m.ContourWAttr != nil {
		if *m.ContourWAttr < 0 {
			return fmt.Errorf("%s/m.ContourWAttr must be >= 0 (have %v)", path, *m.ContourWAttr)
		}
		if *m.ContourWAttr > 27273042316900 {
			return fmt.Errorf("%s/m.ContourWAttr must be <= 27273042316900 (have %v)", path, *m.ContourWAttr)
		}
	}
	if err := m.PrstMaterialAttr.ValidateWithPath(path + "/PrstMaterialAttr"); err != nil {
		return err
	}
	if m.BevelT != nil {
		if err := m.BevelT.ValidateWithPath(path + "/BevelT"); err != nil {
			return err
		}
	}
	if m.BevelB != nil {
		if err := m.BevelB.ValidateWithPath(path + "/BevelB"); err != nil {
			return err
		}
	}
	if m.ExtrusionClr != nil {
		if err := m.ExtrusionClr.ValidateWithPath(path + "/ExtrusionClr"); err != nil {
			return err
		}
	}
	if m.ContourClr != nil {
		if err := m.ContourClr.ValidateWithPath(path + "/ContourClr"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
