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
	"github.com/xiaoma20082008/httl"
	"io"
	"os"
)

type FileResource struct {
	BaseResource
	path string
}

func (fr *FileResource) Open() (io.ReadCloser, error) {
	if _, e := os.Stat(fr.path); e != nil {
		return nil, e
	}
	file, e := os.OpenFile(fr.path, os.O_RDWR, 0777)
	if e != nil {
		return nil, e
	}
	return file, nil
}

func NewFileResource(e httl.Engine, name string, locale string, encoding string, path string) *FileResource {
	fr := &FileResource{
		path: path,
	}
	fr.engine = e
	fr.name = name
	fr.locale = locale
	fr.encoding = encoding
	fr.lastModified = -1
	return fr
}
