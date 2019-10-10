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

package gpio

import (
	"github.com/spf13/cobra"
	"go.eqrx.net/mauzr/pkg/program"
)

// SubCommand creates a cobra command for this driver.
func SubCommand(p *program.Program) *cobra.Command {
	command := cobra.Command{
		Use:   "gpio",
		Short: "Expose a GPIO driver",
		Long:  "Expose a GPIO driver via REST.",
		Run: func(cmd *cobra.Command, args []string) {
			p.Mux.Handle("/input", InputHandler())
			p.Mux.Handle("/output", OutputHandler())
		},
	}
	return &command
}