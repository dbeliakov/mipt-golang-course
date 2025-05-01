package formvalues

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Person struct {
	Name    string   `form:"name"`
	Age     int      `form:"age"`
	Active  bool     `form:"active"`
	Hobbies []string `form:"hobbies"`
}

type Simple struct {
	Title string
	Count int
}

func TestUnpack(t *testing.T) {
	tests := []struct {
		name    string
		form    url.Values
		target  interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name: "all basic types",
			form: url.Values{
				"name":   {"Alice"},
				"age":    {"30"},
				"active": {"true"},
			},
			target: &Person{},
			want: &Person{
				Name:   "Alice",
				Age:    30,
				Active: true,
			},
		},
		{
			name: "string slice",
			form: url.Values{
				"hobbies": {"reading", "coding"},
			},
			target: &Person{},
			want: &Person{
				Hobbies: []string{"reading", "coding"},
			},
		},
		{
			name: "struct tags override",
			form: url.Values{
				"name": {"Bob"},
				"age":  {"25"},
			},
			target: &Person{},
			want: &Person{
				Name: "Bob",
				Age:  25,
			},
		},
		{
			name: "default field names",
			form: url.Values{
				"title": {"Test"},
				"count": {"42"},
			},
			target: &Simple{},
			want: &Simple{
				Title: "Test",
				Count: 42,
			},
		},
		{
			name: "invalid int",
			form: url.Values{
				"age": {"thirty"},
			},
			target:  &Person{},
			wantErr: true,
		},
		{
			name: "invalid bool",
			form: url.Values{
				"active": {"yes"},
			},
			target:  &Person{},
			wantErr: true,
		},
		{
			name: "unsupported type",
			form: url.Values{
				"title": {"Test"},
			},
			target:  &struct{ Title float64 }{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &http.Request{Form: tt.form}
			err := Unpack(req, tt.target)

			if tt.wantErr {
				require.Error(t, err, "Unpack() should fail")
				return
			}

			require.NoError(t, err, "Unpack() should succeed")
			assert.Equal(t, tt.want, tt.target, "Unpack() result mismatch")
		})
	}
}

func TestUnpack_EdgeCases(t *testing.T) {
	t.Run("empty form", func(t *testing.T) {
		var p Person
		req := &http.Request{Form: url.Values{}}
		require.NoError(t, Unpack(req, &p))
		assert.Equal(t, Person{}, p)
	})

	t.Run("nil target", func(t *testing.T) {
		req := &http.Request{Form: url.Values{"name": {"Alice"}}}
		err := Unpack(req, nil)
		require.Error(t, err)
	})

	t.Run("non-pointer target", func(t *testing.T) {
		req := &http.Request{Form: url.Values{"name": {"Alice"}}}
		var p Person
		err := Unpack(req, p)
		require.Error(t, err)
	})
}
