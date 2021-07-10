/*
 * Copyright (c) 2019 VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package component

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/vmware-tanzu/octant/pkg/action"

	"github.com/stretchr/testify/require"

	"github.com/vmware-tanzu/octant/internal/testutil"
)

func TestMetadata_UnmarshalJSON(t *testing.T) {
	data, err := ioutil.ReadFile(filepath.Join("testdata", "metadata.json"))
	require.NoError(t, err)

	got := Metadata{}
	require.NoError(t, got.UnmarshalJSON(data))

	expected := Metadata{
		Type: "type",
		Title: []TitleComponent{
			NewText("title"),
		},
		Accessor: "accessor",
	}
	require.Equal(t, expected, got)
}

func TestContentResponse_Add(t *testing.T) {
	tests := []struct {
		name       string
		components []Component
		wanted     []Component
	}{
		{
			name:       "in general",
			components: []Component{NewText("test")},
			wanted:     []Component{NewText("test")},
		},
		{
			name:       "with nil components",
			components: []Component{nil, NewText("test"), nil},
			wanted:     []Component{NewText("test")},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cr := NewContentResponse(TitleFromString("cr"))
			cr.Add(test.components...)
			testutil.AssertJSONEqual(t, test.wanted, cr.Components)
		})
	}
}

func TestContentResponse_AddTitleComponents(t *testing.T) {
	tests := []struct {
		name       string
		components []Component
		wanted     []Component
	}{
		{
			name:       "with no components",
			components: []Component{},
			wanted:     nil,
		},
		{
			name:       "with nil components",
			components: []Component{nil, NewButton("btn", nil), nil},
			wanted:     []Component{NewButton("btn", nil)},
		},
		{
			name:       "in general",
			components: []Component{NewIcon("car")},
			wanted:     []Component{NewIcon("car")},
		},
		{
			name: "with many components",
			components: []Component{
				NewIcon("car"),
				NewButton("btn", action.Payload{}),
				NewButton("btn", action.Payload{}),
				NewButton("btn", action.Payload{}),
				NewButton("btn", action.Payload{}),
				NewLink("link", "", ""),
			},
			wanted: []Component{
				NewIcon("car"),
				NewButton("btn", action.Payload{}),
				NewButton("btn", action.Payload{}),
				NewButton("btn", action.Payload{}),
				NewButton("btn", action.Payload{}),
				NewLink("link", "", ""),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cr := NewContentResponse(TitleFromString("cr"))
			cr.AddTitleComponents(test.components...)
			testutil.AssertJSONEqual(t, test.wanted, cr.TitleComponents)
		})
	}
}
