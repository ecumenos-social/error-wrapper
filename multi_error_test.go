package errorwrapper_test

import (
	"reflect"
	"testing"

	errorwrapper "github.com/ecumenos-social/error-wrapper"
)

var e4 error

func TestDefaultMultiError(t *testing.T) {
	tests := []struct {
		name string
		want errorwrapper.MultiError
	}{
		{
			name: "test-1",
			want: errorwrapper.DefaultMultiError(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errorwrapper.DefaultMultiError(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultMultiError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiError_IsNil(t *testing.T) {
	type fields struct {
		errs []error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "test-1",
			fields: fields{[]error{e4, e4, e4}},
			want:   true,
		},
		{
			name:   "test-2",
			fields: fields{[]error{e4, e1, e4}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := errorwrapper.DefaultMultiError()
			e.AddErrors(tt.fields.errs...)
			if got := e.IsNil(); got != tt.want {
				t.Errorf("MultiError.IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMultiError(t *testing.T) {
	type args struct {
		errs []error
	}
	tests := []struct {
		name string
		args args
		want errorwrapper.MultiError
	}{
		{
			name: "test-1",
			args: args{[]error{e1, e2, e3}},
			want: errorwrapper.NewMultiError(e1, e2, e3).(errorwrapper.MultiError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errorwrapper.NewMultiError(tt.args.errs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMultiError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multierror_AddErrors(t *testing.T) {
	type args struct {
		errs []error
	}
	tests := []struct {
		name string
		e    errorwrapper.MultiError
		args args
		want errorwrapper.MultiError
	}{
		{
			name: "test-1",
			args: args{[]error{e1, e2, e3, e4, e0, e4}},
			want: errorwrapper.NewMultiError(e1, e2, e3, e0).(errorwrapper.MultiError),
		},
		{
			name: "test-2",
			args: args{[]error{e1, e2, e3, e4, e0, e4}},
			want: errorwrapper.NewMultiError(e1, e2, e3, e0).(errorwrapper.MultiError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e = errorwrapper.DefaultMultiError()
			if got := tt.e.AddErrors(tt.args.errs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("multierror.AddErrors() = %v, want %v", got, tt.want)
			}
		})
	}
}
