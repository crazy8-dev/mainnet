// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/mainnet/blob/master/LICENSE.md.

// +build functest

package functest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetInfo(t *testing.T) {
	info := getInfo(t)
	require.NotNil(t, info)
	require.NotEqual(t, "", info.RootDomain)
	require.NotEqual(t, "", info.RootMember)
	require.NotEqual(t, "", info.NodeDomain)
}
