NAME:
  minio - Cloud Storage Server.

DESCRIPTION:
  Minio is an Amazon S3 compatible object storage server. Use it to store photos, videos, VMs, containers, log files, or any blob of data as objects.

USAGE:
  minio [FLAGS] COMMAND [ARGS...]

COMMANDS:
  server   Start object storage server.
  gateway  Start object storage gateway.
  update   Check for a new software update.
  version  Print version.
  
FLAGS:
  --config-dir value, -C value  Path to configuration directory. (default: "/root/.minio")
  --quiet                       Disable startup information.
  --json                        Output server logs and startup information in json format.
  --help, -h                    Show help.
  
VERSION:
  2018-03-16T22:52:12Z