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
	"archive/zip"
	"github.com/xiaoma20082008/httl-go"
	"io"
)

type ZipResource struct {
	path string
	BaseResource
}

func (r *ZipResource) LastModified() uint64 { return r.lastModified }

func (r *ZipResource) Open() (*io.ReadCloser, error) {
	zip.OpenReader(r.path)
	return nil, nil
}

func NewZipResource(e httl.Engine, name string, locale string, encoding string, path string) *ZipResource {
	r := ZipResource{
		path: path,
	}
	r.engine = e
	r.name = name
	r.locale = locale
	r.encoding = encoding
	r.lastModified = -1
	return &r
}
