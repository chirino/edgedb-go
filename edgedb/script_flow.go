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
	"fmt"
	"net"

	"github.com/edgedb/edgedb-go/edgedb/protocol"
	"github.com/edgedb/edgedb-go/edgedb/protocol/message"
)

func scriptFlow(ctx context.Context, conn net.Conn, query string) (err error) {
	msg := []byte{message.ExecuteScript, 0, 0, 0, 0}
	protocol.PushUint16(&msg, 0) // no headers
	protocol.PushString(&msg, query)
	protocol.PutMsgLength(msg)

	rcv, err := writeAndRead(ctx, conn, msg)
	if err != nil {
		return err
	}

	for len(rcv) > 0 {
		bts := protocol.PopMessage(&rcv)
		mType := protocol.PopUint8(&bts)

		switch mType {
		case message.CommandComplete:
		case message.ReadyForCommand:
		case message.ErrorResponse:
			return decodeError(&bts)
		default:
			panic(fmt.Sprintf("unexpected message type: 0x%x", mType))
		}
	}

	return nil
}