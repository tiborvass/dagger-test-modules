package main

import (
	"context"
	"errors"
	"fmt"
)

// +check
func (*RootMod) CheckPass(ctx context.Context) error {
	fmt.Println("passing")
	return nil
}

// +check
func (*RootMod) CheckFail(ctx context.Context) error {
	return errors.New("expected failure")
}
