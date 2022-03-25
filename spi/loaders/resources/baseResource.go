/*
 *  Copyright 2022 HTTL-Go Team
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package resources

import (
	"github.com/xiaoma20082008/httl-go"
	"github.com/xiaoma20082008/httl-go/utils"
	"io"
)

type BaseResource struct {
	engine       httl.Engine
	name         string
	locale       string
	encoding     string
	lastModified uint64
	httl.Resource
}

func (r *BaseResource) Name() string         { return r.name }
func (r *BaseResource) Encoding() string     { return r.encoding }
func (r *BaseResource) Locale() string       { return r.locale }
func (r *BaseResource) LastModified() uint64 { return r.lastModified }
func (r *BaseResource) Length() uint64       { return -1 }
func (r *BaseResource) Source() (string, error) {
	reader, err := r.Open()
	defer (*reader).Close()
	if err != nil {
		return "", err
	}
	return utils.ReadFully(reader)
}
func (r *BaseResource) Open() (*io.ReadCloser, error) { panic("") }
func (r *BaseResource) Engine() httl.Engine           { return r.engine }
