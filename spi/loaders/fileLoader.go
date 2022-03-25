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

package loaders

import (
	"github.com/xiaoma20082008/httl-go"
	"github.com/xiaoma20082008/httl-go/spi/loaders/resources"
	"os"
)

type FileLoader struct {
	BaseLoader
}

func (l *FileLoader) Exists(name string, locale string, encoding string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func (l *FileLoader) Load(name string, locale string, encoding string) (httl.Resource, error) {
	if encoding == "" {
		encoding = l.encoding
	}
	if l.Exists(name, locale, encoding) {
		return resources.NewFileResource(l.engine, name, locale, encoding, name), nil
	}
	return nil, nil
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}
