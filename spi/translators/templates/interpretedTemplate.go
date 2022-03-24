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
	"io"
)

type InterpretedTemplate struct {
	BaseTemplate
}

func (t *InterpretedTemplate) Evaluate(o map[string]any) (string, error) {
	sw := StringWriter{}
	defer sw.Close()
	if t.Render(o, &sw) != nil {
		return "", nil
	}
	return sw.ToString(), nil
}

func (t *InterpretedTemplate) Render(o map[string]any, w io.Writer) error {
	v := InterpretedVisitor{}
	return t.Accept(&v)
}

func NewInterpretedTemplate() *InterpretedTemplate {
	return &InterpretedTemplate{}
}
