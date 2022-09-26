// The AGPLv3 License (AGPLv3)

// Copyright (c) 2022 ZHAO Zhenhua <zhao.zhenhua@gmail.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package did

import (
	"encoding/json"
	"reflect"
	"testing"
)

// func TestString(t *testing.T) {
// 	type fields struct {
// 		Scheme           string
// 		Method           string
// 		MethodSpecificID string
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   string
// 	}{
// 		// TODO: Add test cases.
// 		{"case1", fields{Scheme: "did", Method: "example", MethodSpecificID: "1234567890"}, "did:example:1234567890"},
// 		{"case2", fields{Scheme: "did", Method: "example", MethodSpecificID: "1234567890abcdef"}, "did:example:1234567890abcdef"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := &DID{
// 				Scheme:           tt.fields.Scheme,
// 				Method:           tt.fields.Method,
// 				MethodSpecificID: tt.fields.MethodSpecificID,
// 			}
// 			if got := d.String(); got != tt.want {
// 				t.Errorf("DID.String() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestParse(t *testing.T) {
	type args struct {
		did string
	}
	tests := []struct {
		name    string
		args    args
		want    *DID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.did)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshalDidURL(t *testing.T) {
	did := DID{Scheme: "did", Method: "example", MethodSpecificID: "1234567890"}
	data, _ := json.Marshal(did)
	t.Logf("%s", data)
}
