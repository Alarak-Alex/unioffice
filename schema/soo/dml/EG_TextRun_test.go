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

func TestEG_TextRunConstructor(t *testing.T) {
	v := dml.NewEG_TextRun()
	if v == nil {
		t.Errorf("dml.NewEG_TextRun must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_TextRun should validate: %s", err)
	}
}

func TestEG_TextRunMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_TextRun()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_TextRun()
	xml.Unmarshal(buf, v2)
}
