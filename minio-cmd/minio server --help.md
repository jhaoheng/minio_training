NAME:
  minio server - Start object storage server.

USAGE:
  minio server [FLAGS] DIR1 [DIR2..]
  minio server [FLAGS] DIR{1...64}

DIR:
  DIR points to a directory on a filesystem. When you want to combine
  multiple drives into a single large system, pass one directory per
  filesystem separated by space. You may also use a '...' convention
  to abbreviate the directory arguments. Remote directories in a
  distributed setup are encoded as HTTP(s) URIs.


FLAGS:
  --address value               Bind to a specific ADDRESS:PORT, ADDRESS can be an IP or hostname. (default: ":9000")
  --config-dir value, -C value  Path to configuration directory. (default: "/root/.minio")
  --quiet                       Disable startup information.
  --json                        Output server logs and startup information in json format.
  --help, -h                    Show help.
  
ENVIRONMENT VARIABLES:
  ACCESS:
     MINIO_ACCESS_KEY: Custom username or access key of minimum 3 characters in length.
     MINIO_SECRET_KEY: Custom password or secret key of minimum 8 characters in length.

  BROWSER:
     MINIO_BROWSER: To disable web browser access, set this value to "off".

  REGION:
     MINIO_REGION: To set custom region. By default it is "us-east-1".

  UPDATE:
     MINIO_UPDATE: To turn off in-place upgrades, set this value to "off".

EXAMPLES:
  1. Start minio server on "/home/shared" directory.
      $ minio server /home/shared

  2. Start minio server bound to a specific ADDRESS:PORT.
      $ minio server --address 192.168.1.101:9000 /home/shared

  3. Start minio server on a 12 disks server.
      $ minio server /mnt/export1/ /mnt/export2/ /mnt/export3/ /mnt/export4/ \
          /mnt/export5/ /mnt/export6/ /mnt/export7/ /mnt/export8/ /mnt/export9/ \
          /mnt/export10/ /mnt/export11/ /mnt/export12/

  4. Start distributed minio server on a 4 node setup with 1 drive each. Run following commands on all the 4 nodes.
      $ export MINIO_ACCESS_KEY=minio
      $ export MINIO_SECRET_KEY=miniostorage
      $ minio server http://192.168.1.11/mnt/export/ http://192.168.1.12/mnt/export/ \
          http://192.168.1.13/mnt/export/ http://192.168.1.14/mnt/export/

  5. Start minio server on 64 disks server.
      $ minio server /mnt/export{1...64}

  6. Start distributed minio server on an 8 node setup with 8 drives each. Run following command on all the 8 nodes.
      $ export MINIO_ACCESS_KEY=minio
      $ export MINIO_SECRET_KEY=miniostorage
      $ minio server http://node{1...8}.example.com/mnt/export/{1...8}