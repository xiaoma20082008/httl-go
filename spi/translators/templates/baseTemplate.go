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

package templates

import (
	"github.com/xiaoma20082008/httl"
	"io"
	"strings"
)

type BaseTemplate struct {
	name         string
	encoding     string
	locale       string
	lastModified uint64
	length       uint64
	resource     httl.Resource
	node         httl.Node
	parent       httl.Template

	httl.Template
}

func (t *BaseTemplate) IsMacro() bool { return false }
func (t *BaseTemplate) Render(o map[string]any, w io.Writer) error {
	panic("TODO IMPLEMENTATION")
}
func (t *BaseTemplate) Evaluate(o map[string]any) (string, error) {
	panic("TODO IMPLEMENTATION")
}

func (t *BaseTemplate) Name() string                 { return t.name }
func (t *BaseTemplate) Encoding() string             { return t.encoding }
func (t *BaseTemplate) Locale() string               { return t.locale }
func (t *BaseTemplate) LastModified() uint64         { return t.lastModified }
func (t *BaseTemplate) Length() uint64               { return t.length }
func (t *BaseTemplate) Source() (string, error)      { return t.resource.Source() }
func (t *BaseTemplate) Open() (io.ReadCloser, error) { return t.resource.Open() }
func (t *BaseTemplate) Engine() httl.Engine          { return t.resource.Engine() }

func (t *BaseTemplate) Accept(v httl.Visitor) error {
	if v.Visit(t) != nil {
		for _, node := range t.Children() {
			node.Accept(v)
		}
	}
	return nil
}
func (t *BaseTemplate) Parent() httl.Node     { return t.parent }
func (t *BaseTemplate) Children() []httl.Node { return nil }

type StringWriter struct {
	buf strings.Builder
	io.WriteCloser
}

func (sw *StringWriter) Close() error {
	sw.buf.Reset()
	return nil
}
func (sw *StringWriter) Write(p []byte) (n int, err error) {
	return sw.buf.Write(p)

}
func (sw *StringWriter) ToString() string {
	return sw.buf.String()
}
