package templatestests

// DO NOT EDIT!
// This code is generated with http://github.com/gabstv/gowrap tool
// using ../templates/grpc_validate template

//go:generate gowrap gen -p github.com/gabstv/gowrap/templates_tests -i InterfaceWithValidtableArg -t ../templates/grpc_validate -o interface_with_grpc_validation.go

import (
	"context"

	grpc_codes "google.golang.org/grpc/codes"
	grpc_status "google.golang.org/grpc/status"
)

// InterfaceWithValidtableArgWithGRPCValidation implements InterfaceWithValidtableArg interface instrumented with GRPC request validation
type InterfaceWithValidtableArgWithGRPCValidation struct {
	InterfaceWithValidtableArg
}

// NewInterfaceWithValidtableArgWithGRPCValidation returns InterfaceWithValidtableArgWithGRPCValidation
func NewInterfaceWithValidtableArgWithGRPCValidation(base InterfaceWithValidtableArg) InterfaceWithValidtableArgWithGRPCValidation {
	return InterfaceWithValidtableArgWithGRPCValidation{
		InterfaceWithValidtableArg: base,
	}
}

// Method implements InterfaceWithValidtableArg
func (_d InterfaceWithValidtableArgWithGRPCValidation) Method(ctx context.Context, r *ValidatableRequest) (err error) {
	if _v, _ok := interface{}(r).(interface{ Validate() error }); _ok {
		if err = _v.Validate(); err != nil {
			err = grpc_status.Error(grpc_codes.InvalidArgument, err.Error())
			return
		}
	}

	return _d.InterfaceWithValidtableArg.Method(ctx, r)
}
