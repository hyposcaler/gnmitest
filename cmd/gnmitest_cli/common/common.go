/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package common contains functions to run suite text proto and to write
// report proto into file.
package common

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"

	gtpb "github.com/openconfig/gnmitest/proto/gnmitest"
	rpb "github.com/openconfig/gnmitest/proto/report"
	spb "github.com/openconfig/gnmitest/proto/suite"
)

var (
	// Dialer to use while calling into gnmitest service. It can be overridden
	// by an external package to use a custom dialer.
	Dialer = grpc.Dial
	// Opts is a list of grpc.DialOption passed the the Dialer.  It can be
	// overridden by an external package to use custom options.
	Opts = []grpc.DialOption{grpc.WithInsecure()}
)

// runSuite dials into given address(a) and calls Run RPC of gnmitest service
// with the given suite proto(s). It returns a test report if running suite is
// successful. Otherwise an error is returned.
func runSuite(ctx context.Context, a string, s *spb.Suite) (*rpb.Report, error) {
	conn, err := Dialer(a, Opts...)
	if err != nil {
		return nil, fmt.Errorf("dial func failed for %q; %v", a, err)
	}
	defer conn.Close()

	cl := gtpb.NewGNMITestClient(conn)

	rep, err := cl.Run(ctx, s)
	if err != nil {
		return nil, fmt.Errorf("running Suite failed; %v", err)
	}
	return rep, nil
}

// suite function unmarshals the suite proto text file from the local testdata
// directory in current working directory. It returns error if it fails to
// unmarshal text proto file.
func suite(f string) (*spb.Suite, error) {
	absPath, err := filepath.Abs(f)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for %q; %v", f, err)
	}

	b, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read %q; %v", absPath, err)
	}

	s := &spb.Suite{}
	if err = proto.UnmarshalText(string(b), s); err != nil {
		return nil, fmt.Errorf("failed to unmarshal %v; %v", string(b), err)
	}

	return s, nil
}

// writeReport writes given report proto into a file with the given filepath.
func writeReport(filePath string, r *rpb.Report) error {
	p, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(p, []byte(proto.MarshalTextString(r)), 0644); err != nil {
		return fmt.Errorf("failed to write %q file; %v", p, err)
	}

	return nil
}

// Run function runs the Suite proto given as textproto and writes Report proto
// into a file as textproto.
func Run(address, suiteFile, reportFile string) (*rpb.Report, error) {
	switch {
	case address == "":
		return nil, errors.New("gnmitest endpoint address must be set")
	case suiteFile == "":
		return nil, errors.New("suite text proto file must be set")
	case reportFile == "":
		return nil, errors.New("report text proto file must be set")
	}

	// Read suite proto file.
	s, err := suite(suiteFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read %q suite text proto file; %v", suiteFile, err)
	}

	// Run suite proto file.
	r, err := runSuite(context.Background(), address, s)
	if err != nil {
		return nil, fmt.Errorf("running suite proto file failed; %v", err)
	}

	if err := writeReport(reportFile, r); err != nil {
		return nil, fmt.Errorf("failed to write %q report text proto file; %v", reportFile, r)
	}

	return r, nil
}
