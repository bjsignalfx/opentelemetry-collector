// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ballastextension

import (
	"go.opentelemetry.io/collector/config"
)

// Config has the configuration for the fluentbit extension.
type Config struct {
	*config.ExtensionSettings `mapstructure:"-"`

	// SizeMiB is the size, in MiB, of the memory ballast
	// to be created for this process.
	SizeMiB uint32 `mapstructure:"size_mib"`
}
