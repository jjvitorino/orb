/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package policies

const CurrentRPCSchemaVersion = "1.0"

type RPC struct {
	SchemaVersion string      `json:"schema_version"`
	Func          string      `json:"func"`
	Payload       interface{} `json:"payload"`
}