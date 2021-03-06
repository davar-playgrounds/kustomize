// Copyright 2021 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package build

import (
	"github.com/spf13/pflag"
)

func AddFunctionFlags(set *pflag.FlagSet) {
	set.BoolVar(
		&theFlags.fnOptions.EnableExec, "enable-exec", false, /*do not change!*/
		"enable support for exec functions -- note: exec functions run arbitrary code -- do not use for untrusted configs!!! (Alpha)")
	set.BoolVar(
		&theFlags.fnOptions.EnableStar, "enable-star", false,
		"enable support for starlark functions. (Alpha)")
	set.BoolVar(
		&theFlags.fnOptions.Network, "network", false,
		"enable network access for functions that declare it")
	set.StringVar(
		&theFlags.fnOptions.NetworkName, "network-name", "bridge",
		"the docker network to run the container in")
	set.StringArrayVar(
		&theFlags.fnOptions.Mounts, "mount", []string{},
		"a list of storage options read from the filesystem")
	set.StringArrayVarP(
		&theFlags.fnOptions.Env, "env", "e", []string{},
		"a list of environment variables to be used by functions")
}
