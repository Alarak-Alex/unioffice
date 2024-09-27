// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/Alarak-Alex/unioffice/schema/soo/wml"
)

func TestCT_Base64BinaryConstructor(t *testing.T) {
	v := wml.NewCT_Base64Binary()
	if v == nil {
		t.Errorf("wml.NewCT_Base64Binary must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Base64Binary should validate: %s", err)
	}
}

func TestCT_Base64BinaryMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Base64Binary()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Base64Binary()
	xml.Unmarshal(buf, v2)
}
