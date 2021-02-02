/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package framework

import (
	l2SvcAttPars "github.com/Nordix/eno/pkg/l2serviceattachmentparser"
	"github.com/Nordix/eno/pkg/render"
)

//type IpamCni interface {
//
//}

type Cnier interface {
	HandleCni(sattp *l2SvcAttPars.L2SrvAttParser, d *render.RenderData) error
}

type createCnierFunc func() Cnier