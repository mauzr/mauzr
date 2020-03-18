/*
Copyright 2019 Alexander Sowitzki.

GNU Affero General Public License version 3 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://opensource.org/licenses/AGPL-3.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package play

import (
	"context"
	"time"

	"go.eqrx.net/mauzr/pkg/io/rest"
)

const (
	PartChangeForm = `
<!DOCTYPE html>
<html>
<head>
	<meta name="viewport" content="width=device-width, initial-scale=1" />
</head>
<body>
	<form method="get">
    <input type="radio" name="stance" value="off" checked> Off<br>
    <input type="radio" name="stance" value="default"> Default<br>
    <input type="radio" name="stance" value="bright"> Bright<br>
		<input type="radio" name="stance" value="alert"> Alert<br>
		<input type="radio" name="stance" value="rainbow"> Rainbow<br>
		<input type="radio" name="stance" value="theme"> Theme<br>
		<br>
		<input type="submit" value="Submit">
	</form>
</body>
</html>
`
)

// AddPartChangerEndpoint that will listen for part change requests.
func AddPartChangerEndpoint(c rest.REST, path string, changers ...Changer) {
	c.Endpoint(path, PartChangeForm, func(query *rest.Request) {
		args := struct {
			Stance string `json:"stance"`
		}{}
		if err := query.Args(&args); err != nil {
			return
		}
		ctx, cancel := context.WithTimeout(query.Ctx, 3*time.Second)
		defer cancel()
		for _, changer := range changers {
			query.RequestError = changer.ChangePart(ctx, args.Stance)
		}
	})
}