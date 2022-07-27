// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"reflect"

	"github.com/banzaicloud/logging-operator/pkg/sdk/logging/model/syslogng/config/render"
)

func destinationDefStmt(name string, body render.Renderer) render.Renderer {
	return braceDefStmt("destination", name, body)
}

func isActiveDestinationDriver(f Field) bool {
	return hasDestDriverTag(f) && f.Meta.Type.Kind() == reflect.Pointer && !f.Value.IsNil()
}

func hasDestDriverTag(f Field) bool {
	return structFieldSettings(f.Meta).Has("dest-drv")
}
