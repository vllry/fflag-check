About
===
fflag-check is the exposed service for reading fflag feature flag values.

API
---
See fflag-check-api for GRPC protobuffers.

    func query(account string, flag string) bool {
     
        // Set up a connection to the server.
        conn, err := grpc.Dial("HOST:PORT", grpc.WithInsecure())
        if err != nil {
            log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()
        c := fflagcheckapi.NewFeatureFlagClient(conn)
     
        // Contact the server and print out its response.
        r, err := c.GetFlag(context.Background(), &fflagcheckapi.FlagQuery{AccountId: account, FlagName: flag})
        if err != nil {
            log.Fatalf("could not get: %v", err)
        }
     
        return r.Value
    }


Running
===

