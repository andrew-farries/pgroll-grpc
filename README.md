# pgroll-grpc

A simple pgroll gRPC server that can start, complete and roll back migrations.

## Instructions

* Have a `postgres` instance running on `:5432`
* `make run`
* Invoke the `Start`, `Complete` and `Rollback` gRPC methods with a tool like `evans` or `grpcurl`.

Example message payloads are in the `/examples` dir.

eg:

```bash
evans --proto proto/pgroll.proto cli call pgroll.PGRoll.Start --file examples/05_sql.json
echo '{}' | evans --proto proto/pgroll.proto cli call pgroll.PGRoll.Complete
```
