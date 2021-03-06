// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the URIs of this project regarding your
// rights to use or distribute this software.

package assemblers_test

import (
	"os"
	"testing"

	"github.com/singularityware/singularity/src/pkg/build/assemblers"
	"github.com/singularityware/singularity/src/pkg/build/sources"
	"github.com/singularityware/singularity/src/pkg/build/types"
	"github.com/singularityware/singularity/src/pkg/test"
)

const (
	assemblerDockerDestDir = "/tmp/docker_alpine_assemble_test"
	assemblerShubDestDir   = "/tmp/shub_alpine_assemble_test"
)

// TestAssembler sees if we can build a SIF image from a docke based kitchen to /tmp
func TestSandboxAssemblerDocker(t *testing.T) {
	test.DropPrivilege(t)
	defer test.ResetPrivilege(t)

	b, err := types.NewBundle("sbuild-sandboxAssembler")
	if err != nil {
		return
	}

	b.Recipe, err = types.NewDefinitionFromURI(assemblerDockerURI)
	if err != nil {
		t.Fatalf("unable to parse URI %s: %v\n", assemblerDockerURI, err)
	}

	ocp := &sources.OCIConveyorPacker{}

	if err := ocp.Get(b); err != nil {
		t.Fatalf("failed to Get from %s: %v\n", assemblerDockerURI, err)
	}

	_, err = ocp.Pack()
	if err != nil {
		t.Fatalf("failed to Pack from %s: %v\n", assemblerDockerURI, err)
	}

	a := &assemblers.SandboxAssembler{}

	err = a.Assemble(b, assemblerDockerDestDir)
	if err != nil {
		t.Fatalf("failed to assemble from %s: %v\n", assemblerDockerURI, err)
	}

	defer os.RemoveAll(assemblerDockerDestDir)
}
func TestSandboxAssemblerShub(t *testing.T) {
	test.DropPrivilege(t)
	defer test.ResetPrivilege(t)

	b, err := types.NewBundle("sbuild-sandboxAssembler")
	if err != nil {
		return
	}

	b.Recipe, err = types.NewDefinitionFromURI(assemblerShubURI)
	if err != nil {
		t.Fatalf("unable to parse URI %s: %v\n", assemblerShubURI, err)
	}

	scp := &sources.ShubConveyorPacker{}

	if err := scp.Get(b); err != nil {
		t.Fatalf("failed to Get from %s: %v\n", assemblerShubURI, err)
	}

	_, err = scp.Pack()
	if err != nil {
		t.Fatalf("failed to Pack from %s: %v\n", assemblerShubURI, err)
	}

	a := &assemblers.SIFAssembler{}

	err = a.Assemble(b, assemblerShubDestDir)
	if err != nil {
		t.Fatalf("failed to assemble from %s: %v\n", assemblerShubURI, err)
	}

	defer os.RemoveAll(assemblerShubDestDir)
}
