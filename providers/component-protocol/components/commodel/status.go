// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commodel

import (
	"fmt"
)

// UnifiedStatus .
type UnifiedStatus int

// IUnifiedStatus .
type IUnifiedStatus interface {
	fmt.Stringer
}

// UnifiedStatus .
const (
	ErrorStatus UnifiedStatus = iota
	WarningStatus
	SuccessStatus
	ProcessingStatus
	DefaultStatus
)

// String .
func (g UnifiedStatus) String() string {
	switch g {
	case 0:
		return "error"
	case 1:
		return "warning"
	case 2:
		return "success"
	case 3:
		return "processing"
	case 4:
		return "default"
	default:
		return "default"
	}
}
