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
	"github.com/Alarak-Alex/unioffice/schema/soo/ofc/math"
)

type CT_Tc struct {
	// Table Cell Identifier
	IdAttr *string
	// Table Cell Properties
	TcPr              *CT_TcPr
	EG_BlockLevelElts []*EG_BlockLevelElts
}

func NewCT_Tc() *CT_Tc {
	ret := &CT_Tc{}
	return ret
}

func (m *CT_Tc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	e.EncodeToken(start)
	if m.TcPr != nil {
		setcPr := xml.StartElement{Name: xml.Name{Local: "w:tcPr"}}
		e.EncodeElement(m.TcPr, setcPr)
	}
	for _, c := range m.EG_BlockLevelElts {
		c.MarshalXML(e, xml.StartElement{})
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Tc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
	}
lCT_Tc:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tcPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "tcPr"}:
				m.TcPr = NewCT_TcPr()
				if err := d.DecodeElement(m.TcPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "altChunk"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "altChunk"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmp := NewCT_AltChunk()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				tmpblocklevelelts.AltChunk = append(tmpblocklevelelts.AltChunk, tmp)
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXml"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXml"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmpcontentblockcontent.CustomXml = NewCT_CustomXmlBlock()
				if err := d.DecodeElement(tmpcontentblockcontent.CustomXml, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sdt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "sdt"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmpcontentblockcontent.Sdt = NewCT_SdtBlock()
				if err := d.DecodeElement(tmpcontentblockcontent.Sdt, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "p"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "p"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmp := NewCT_P()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				tmpcontentblockcontent.P = append(tmpcontentblockcontent.P, tmp)
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tbl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "tbl"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmp := NewCT_Tbl()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				tmpcontentblockcontent.Tbl = append(tmpcontentblockcontent.Tbl, tmp)
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "proofErr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "proofErr"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.ProofErr = NewCT_ProofErr()
				if err := d.DecodeElement(tmprunlevelelts.ProofErr, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "permStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "permStart"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermStart = NewCT_PermStart()
				if err := d.DecodeElement(tmprunlevelelts.PermStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "permEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "permEnd"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermEnd = NewCT_Perm()
				if err := d.DecodeElement(tmprunlevelelts.PermEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ins"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "ins"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Ins = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Ins, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "del"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "del"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Del = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Del, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFrom"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveFrom"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveFrom = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveFrom, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveTo"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveTo"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveTo = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveTo, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookmarkStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bookmarkStart"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkStart = NewCT_Bookmark()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookmarkEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bookmarkEnd"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFromRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveFromRangeStart"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFromRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveFromRangeEnd"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveToRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveToRangeStart"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveToRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveToRangeEnd"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "commentRangeStart"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeStart = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "commentRangeEnd"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlInsRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlInsRangeStart"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlInsRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlInsRangeEnd"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlDelRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlDelRangeStart"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlDelRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlDelRangeEnd"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveFromRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlMoveFromRangeStart"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveFromRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlMoveFromRangeEnd"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveToRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlMoveToRangeStart"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveToRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlMoveToRangeEnd"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "oMathPara"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "oMathPara"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMathPara = math.NewOMathPara()
				if err := d.DecodeElement(tmpmathcontent.OMathPara, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "oMath"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "oMath"}:
				tmpblocklevelelts := NewEG_BlockLevelElts()
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMath = math.NewOMath()
				if err := d.DecodeElement(tmpmathcontent.OMath, &el); err != nil {
					return err
				}
				m.EG_BlockLevelElts = append(m.EG_BlockLevelElts, tmpblocklevelelts)
				tmpblocklevelelts.EG_ContentBlockContent = append(tmpblocklevelelts.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			default:
				unioffice.Log("skipping unsupported element on CT_Tc %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Tc
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Tc and its children
func (m *CT_Tc) Validate() error {
	return m.ValidateWithPath("CT_Tc")
}

// ValidateWithPath validates the CT_Tc and its children, prefixing error messages with path
func (m *CT_Tc) ValidateWithPath(path string) error {
	if m.TcPr != nil {
		if err := m.TcPr.ValidateWithPath(path + "/TcPr"); err != nil {
			return err
		}
	}
	for i, v := range m.EG_BlockLevelElts {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_BlockLevelElts[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
