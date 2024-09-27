// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/Alarak-Alex/unioffice/schema/soo/dml/chart"
)

func TestCT_PieSerConstructor(t *testing.T) {
	v := chart.NewCT_PieSer()
	if v == nil {
		t.Errorf("chart.NewCT_PieSer must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PieSer should validate: %s", err)
	}
}

func TestCT_PieSerMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PieSer()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PieSer()
	xml.Unmarshal(buf, v2)
}
