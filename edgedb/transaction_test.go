// This source file is part of the EdgeDB open source project.
//
// Copyright 2020-present EdgeDB Inc. and the EdgeDB authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package edgedb

import (
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactionSaves(t *testing.T) {
	ctx := context.Background()
	tx, err := client.Transaction(ctx)
	require.Nil(t, err)

	err = tx.Start(ctx)
	require.Nil(t, err)

	name := "test" + strconv.Itoa(rand.Int())
	// todo maybe clean up the random entry :thinking:
	err = tx.Query(
		ctx,
		"INSERT User{ name := <str>$0 }",
		(*interface{})(nil),
		name,
	)
	assert.Nil(t, err)

	err = tx.Commit(ctx)
	require.Nil(t, err)

	var result string
	err = client.QueryOne(
		ctx,
		`
		SELECT User.name
		FILTER User.name = <str>$0;
		`,
		&result,
		name,
	)

	assert.Nil(t, err)
	assert.Equal(t, name, result)
}

func TestTransactionRollsBack(t *testing.T) {
	ctx := context.Background()
	tx, err := client.Transaction(ctx)
	assert.Nil(t, err)

	err = tx.Start(ctx)
	require.Nil(t, err)

	name := "test" + strconv.Itoa(rand.Int())
	// todo maybe clean up the random entry :thinking:
	err = tx.Query(
		ctx,
		"INSERT User{ name := <str>$0 }",
		(*interface{})(nil),
		name,
	)
	assert.Nil(t, err)

	err = tx.RollBack(ctx)
	require.Nil(t, err)

	var result string
	err = client.QueryOne(
		ctx,
		`
		SELECT User.name
		FILTER User.name = <str>$0;
		`,
		&result,
		name,
	)

	assert.Equal(t, ErrorZeroResults, err)
}
