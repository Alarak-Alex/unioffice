// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/Alarak-Alex/unioffice/schema/soo/dml"
)

func TestCT_TextParagraphPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_TextParagraphProperties()
	if v == nil {
		t.Errorf("dml.NewCT_TextParagraphProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextParagraphProperties should validate: %s", err)
	}
}

func TestCT_TextParagraphPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextParagraphProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextParagraphProperties()
	xml.Unmarshal(buf, v2)
}
