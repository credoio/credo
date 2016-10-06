jsonpathos
==========

command line tool to grok JSON: not quite the jsonpath or jq , but easier to use

Example:

```
$ docker ps

CONTAINER ID        IMAGE                 COMMAND                CREATED             STATUS              PORTS                                            NAMES
a6a2ef96b816        zettio/weave:latest   /home/weave/weaver -   2 days ago          Up 2 days           0.0.0.0:6783->6783/tcp, 0.0.0.0:6783->6783/udp   weave               

$ docker inspect a6a2 | head -60

docker inspect a6a2 | head  -60
[{
    "Args": [
        "-iface",
        "ethwe",
        "-wait",
        "5",
        "-name",
        "7a:7e:d7:43:3c:da"
    ],
    "Config": {
        "AttachStderr": false,
        "AttachStdin": false,
        "AttachStdout": false,
        "Cmd": [
            "-name",
            "7a:7e:d7:43:3c:da"
        ],
        "CpuShares": 0,
        "Cpuset": "",
        "Domainname": "",
        "Entrypoint": [
            "/home/weave/weaver",
            "-iface",
            "ethwe",
            "-wait",
            "5"
        ],
        "Env": [
            "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
        ],
        "ExposedPorts": {
            "6783/tcp": {},
            "6783/udp": {}
        },
        "Hostname": "a6a2ef96b816",
        "Image": "zettio/weave",
        "Memory": 0,
        "MemorySwap": 0,
        "NetworkDisabled": false,
        "OnBuild": null,
        "OpenStdin": false,
        "PortSpecs": null,
        "StdinOnce": false,
        "Tty": false,
        "User": "",
        "Volumes": null,
        "WorkingDir": "/home/weave"
    },
    "Created": "2014-11-08T08:09:56.020127035Z",
    "Driver": "aufs",
    "ExecDriver": "native-0.2",
    "HostConfig": {
        "Binds": null,
        "ContainerIDFile": "",
        "Dns": null,
        "DnsSearch": null,
        "Links": null,
        "LxcConf": [],
        "NetworkMode": "bridge",
        "PortBindings": {

$ docker inspect a6a2 | go run main.go - print /[0]/Config/Env

["PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"]
```
