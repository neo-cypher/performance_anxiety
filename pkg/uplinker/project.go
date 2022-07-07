// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package uplinker

import (
	"context"
	"github.com/zeebo/errs"
	"log"
	"storj.io/uplink"

	"storj.io/common/telemetry"
)

var Error = errs.Class("Project")

func telemetryServer(ctx context.Context, address string) (*telemetry.Server, error) {
	server, err := telemetry.Listen(address)
	if err != nil {
		return nil, Error.New("failed to listen on %s: %v", address, err)
	}
	defer ctx.Check(server.Close)
	return server, nil
}
func Project(accessGrant string) {
	ctx := context.Background()
	server, err := telemetryServer(ctx, "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	access, err := uplink.ParseAccess(accessGrant)
	if err != nil {
		log.Fatalf("could not parse access grant: %v", err)
	}

	project, err := uplink.OpenProject(ctx, access)
	if err != nil {
		log.Fatalf("could not open project: %v", err)
	}
	defer project.Close()

}
