// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"github.com/Alarak-Alex/unioffice/schema/soo/ofc/docPropsVTypes"
)

func TestArrayConstructor(t *testing.T) {
	v := docPropsVTypes.NewArray()
	if v == nil {
		t.Errorf("docPropsVTypes.NewArray must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.Array should validate: %s", err)
	}
}

func TestArrayMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewArray()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewArray()
	xml.Unmarshal(buf, v2)
}
