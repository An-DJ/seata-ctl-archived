# Introduction

This is a CLI tool for [Seata](https://github.com/seata/seata) named `seata-ctl`.

```shell
$ seata-ctl -h
seata-ctl is a CLI tool for Seata

Usage:
  seata-ctl [flags]
  seata-ctl [command]

Available Commands:
  version     Print the version number of seata-ctl

Flags:
  -h, --help              help for seata-ctl
      --ip string         Seata Server IP (default "127.0.0.1")
      --password string   Password (default "seata")
      --port int          Seata Server Admin Port (default 7091)
      --username string   Username (default "seata")

Use "seata-ctl [command] --help" for more information about a command.
```
# How to use

## Login

```shell
$ seata-ctl --ip 127.0.0.1 --port 7091 --username seata --password seata
127.0.0.1:7091 > # input command here
```

## Help

```shell
127.0.0.1:7091 > help
Usage:
  [command] [flag] 

Available Commands:
  get         Get the resource
  help        Help about any command
  quit        Quit the session
  reload      Reload the configuration
  set         Set the resource
  try         Try if this node is ready
```

## Get

```shell
127.0.0.1:7091 > get -h    
Get the resource

Usage:
   get [flags]
   get [command]

Available Commands:
  config        Get the configuration
  config-center Get the config-center configuration
  registry      Get the registry configuration
  status        Get the status

Flags:
  -h, --help   help for get

Use "get [command] --help" for more information about a command.
```

e.g. Get the status of the Seata server cluster:

```shell
127.0.0.1:7091 > get status
+-------+--------------------+--------+
| TYPE  | ADDRESS            | STATUS |
+-------+--------------------+--------+
| nacos | 192.168.163.1:7091 | ok     |
+-------+--------------------+--------+
| nacos | 192.168.163.2:7091 | ok     |
+-------+--------------------+--------+
```

e.g. Get the configuration `server.servicePort`:

```shell
127.0.0.1:7091 > get config --key server.servicePort   
+--------------------+-------+
| KEY                | VALUE |
+--------------------+-------+
| server.servicePort | 8091  |
+--------------------+-------+
```

## Set

```shell
127.0.0.1:7091 > set -h                    
Set the resource

Usage:
   set [flags]
   set [command]

Available Commands:
  config        Set the configuration
  config-center Set the config-center configuration
  registry      Set the registry configuration

Flags:
  -h, --help   help for set

Use "set [command] --help" for more information about a command.
```

e.g. Set the registry type to `eureka`:

```shell
127.0.0.1:7091 > set registry --key registry.type --value eureka
+---------------+--------+
| KEY           | VALUE  |
+---------------+--------+
| registry.type | eureka |
+---------------+--------+
```

You can found that the Seata server is registered at `eureka` registry.

e.g. Set the configuration center to `nacos`

```shell
127.0.0.1:7091 > set config-center --key config.type --value nacos 
+-------------+-------+
| KEY         | VALUE |
+-------------+-------+
| config.type | nacos |
+-------------+-------+
```

You can found that the configuration in `nacos` is loaded.

e.g. Set a configuration item which can be dynamically configured (such as `server.undo.logSaveDays`):

```shell
set config --key server.undo.logSaveDays --value 5
+-------------------------+-------+
| KEY                     | VALUE |
+-------------------------+-------+
| server.undo.logSaveDays | 5     |
+-------------------------+-------+
```

## Try

Try to submit an example transaction to check if the server is ok:

```shell
127.0.0.1:7091 > try
Try an example txn successfully, xid=192.168.163.1:8091:522856277732237313
```

## Reload

TBD

## Quit

Quit the session:

```shell
127.0.0.1:7091 > quit
Quit the session
```