// Copyright 2016 Palantir Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package distgo

import (
	"github.com/nmiyake/pkg/errorstringer"
	"github.com/palantir/pkg/cli"
	"github.com/palantir/pkg/cli/cfgcli"

	"github.com/palantir/godel/apps/distgo/cmd/artifacts"
	"github.com/palantir/godel/apps/distgo/cmd/build"
	"github.com/palantir/godel/apps/distgo/cmd/dist"
	"github.com/palantir/godel/apps/distgo/cmd/products"
	"github.com/palantir/godel/apps/distgo/cmd/publish"
	"github.com/palantir/godel/apps/distgo/cmd/run"
)

func App() *cli.App {
	app := cli.NewApp(cfgcli.Handler(), cli.DebugHandler(errorstringer.StackWithInterleavedMessages))
	app.Name = "distgo"
	app.Usage = "Build, run, test and publish products in a Go project"
	app.Subcommands = []cli.Command{
		products.Command(),
		artifacts.Command(),
		build.Command(),
		run.Command(),
		dist.Command(),
		publish.Command(),
	}
	return app
}
