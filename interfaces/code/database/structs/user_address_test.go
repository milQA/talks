package structs_test

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"code/database/structs"
)

func TestUserAddress_Scan(t *testing.T) {
	tests := []struct {
		name string
		arg  struct {
			value interface{}
		}
		want struct {
			ua  *structs.UserAddress
			err error
		}
	}{
		{
			name: "nil value",
			arg: struct {
				value interface{}
			}{value: nil},
			want: struct {
				ua  *structs.UserAddress
				err error
			}{ua: new(structs.UserAddress), err: nil},
		},
		{
			name: "string null value",
			arg: struct {
				value interface{}
			}{value: "null"},
			want: struct {
				ua  *structs.UserAddress
				err error
			}{ua: new(structs.UserAddress), err: nil},
		},
		{
			name: "bytes null value",
			arg: struct {
				value interface{}
			}{value: []byte("null")},
			want: struct {
				ua  *structs.UserAddress
				err error
			}{ua: new(structs.UserAddress), err: nil},
		},
		{
			name: "bad string value",
			arg: struct {
				value interface{}
			}{value: []byte("boo")},
			want: struct {
				ua  *structs.UserAddress
				err error
			}{
				ua:  new(structs.UserAddress),
				err: fmt.Errorf("cannot unmarshal UserAddress: invalid character 'b' looking for beginning of value"),
			},
		},
		{
			name: "bad int value",
			arg: struct {
				value interface{}
			}{value: 42},
			want: struct {
				ua  *structs.UserAddress
				err error
			}{
				ua:  new(structs.UserAddress),
				err: structs.ErrCannotCastUserAddress,
			},
		},
		{
			name: "bytes empty value",
			arg: struct {
				value interface{}
			}{value: []byte("{}")},
			want: struct {
				ua  *structs.UserAddress
				err error
			}{ua: new(structs.UserAddress), err: nil},
		},
		{
			name: "string empty value",
			arg: struct {
				value interface{}
			}{value: "{}"},
			want: struct {
				ua  *structs.UserAddress
				err error
			}{ua: new(structs.UserAddress), err: nil},
		},
		{
			name: "string empty value",
			arg: struct {
				value interface{}
			}{value: `{"city":"","street":"","home":0,"flat":0}`},
			want: struct {
				ua  *structs.UserAddress
				err error
			}{ua: &structs.UserAddress{Valid: true}, err: nil},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ua := new(structs.UserAddress)
			err := ua.Scan(tt.arg.value)

			if !equalError(err, tt.want.err) {
				t.Errorf("check error: %+v", err)
			}

			if !reflect.DeepEqual(ua, tt.want.ua) {
				t.Errorf("want: %+v, got: %+v", tt.want.ua, ua)
			}
		})
	}
}

func TestUserAddress_Value(t *testing.T) {
	tests := []struct {
		name string
		arg  struct {
			ua *structs.UserAddress
		}
		want struct {
			value driver.Value
			err   error
		}
	}{
		{
			name: "invalid UserAddress",
			arg: struct {
				ua *structs.UserAddress
			}{ua: new(structs.UserAddress)},
			want: struct {
				value driver.Value
				err   error
			}{value: nil, err: nil},
		},
		{
			name: "valid UserAddress",
			arg: struct {
				ua *structs.UserAddress
			}{ua: &structs.UserAddress{Valid: true}},
			want: struct {
				value driver.Value
				err   error
			}{value: []byte(`{"city":"","street":"","home":0,"flat":0}`), err: nil},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			value, err := tt.arg.ua.Value()

			if !equalError(err, tt.want.err) {
				t.Errorf("check error: %+v", err)
			}

			if !reflect.DeepEqual(value, tt.want.value) {
				t.Errorf("want: %d, got: %d", tt.want.value, value)
			}
		})
	}
}

func equalError(a, b error) bool {
	if a == nil {
		return b == nil
	}
	if b == nil {
		return a == nil
	}
	return a.Error() == b.Error() || errors.Is(a, b)
}
