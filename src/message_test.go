package gomq

import (
	"reflect"
	"testing"
)

func Test_NewMessage(t *testing.T) {
	type args struct {
		payload []byte
	}

	tests := []struct {
		name    string
		args    args
		expect  *Message
		wantErr bool
	}{
		{
			name:    "creates valid message",
			args:    args{[]byte("hello, world")},
			wantErr: false,
			expect:  &Message{make(map[string]string, 0), []byte("hello, world")},
		},
		{
			name:    "returns err if no payload data",
			args:    args{[]byte("")},
			wantErr: true,
			expect:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			msg, err := NewMessage(test.args.payload)

			if test.wantErr != (err != nil) {
				if err != nil {
					t.Error(err)
				} else {
					t.Error("expected an error")
				}

				return
			}

			if !reflect.DeepEqual(msg, test.expect) {
				t.Errorf("expected %v to be %v", msg, test.expect)
			}
		})
	}
}

func Test_NewMessageWithHeaders(t *testing.T) {
	type args struct {
		payload []byte
		headers map[string]string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		expect  *Message
	}{
		{
			name:    "creates valid message",
			args:    args{[]byte("hello, world"), map[string]string{"abc": "def", "ghi": "jkl"}},
			wantErr: false,
			expect:  &Message{map[string]string{"abc": "def", "ghi": "jkl"}, []byte("hello, world")},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			msg, err := NewMessageWithHeaders(test.args.payload, test.args.headers)

			if test.wantErr != (err != nil) {
				if err != nil {
					t.Error(err)
				} else {
					t.Error("expected an error")
				}

				return
			}

			if !reflect.DeepEqual(msg, test.expect) {
				t.Errorf("expected %v to be %v", msg, test.expect)
			}
		})
	}
}
