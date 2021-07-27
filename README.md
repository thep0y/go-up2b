## go-up2b

Support the following image beds:
- sm.ms
- imgtu.com
- gitee.com
- github.com
- ipfs

### Commands

```
up2b can upload images to the specified image bed

Usage:
  up2b [command]

Available Commands:
  choose      Switch the image bed to be used
  current     Show image bed in use
  help        Help about any command
  list        List all available image beds
  save        Save the configuration of one or more image beds,
	      and select the first one as the default image bed
  upload      Upload multiple images, the maximum is 10
  version     Display the current version number

Flags:
  -h, --help   help for up2b

Use "up2b [command] --help" for more information about a command.
```

### How to use

#### Use go runtime environment

If you have set up the go runtime environment on your computer, and `$GOPATH/bin` has been added to the environment variable, you can install and use it directly.

```shell
go get github.com/thep0y/go-up2b
```

Configure in `typora` as shown in the figure below:

![截屏2021-06-24 13.13.44](https://cdn.jsdelivr.net/gh/thep0y/image-bed/md/1624511830524171.png)

#### Do not use go operating environment

If you have not set up the go runtime environment, you only need to download the binary file corresponding to the system.

You need to find the binary file corresponding to the system in the latest release and download it.

Add absolute path of binary file in `typora`:

![截屏2021-06-24 13.25.05](https://cdn.jsdelivr.net/gh/thep0y/image-bed/md/1624512338974944.png)

#### Detailed introduction of each command

You can use the `-h` or `--help` flags after each command to view detailed introductions and examples.

Such as:

```shell
$ go-up2b save -h
Save the configuration of one or more image beds, and select the first one as the default image bed.

Usage:
  up2b save [flags]

Examples:
  One configuration or multiple configurations can be saved, and the image bed corresponding to the first configuration passed in is used as the used image bed by default.
  The format of the configuration is %d %s, such as:
  
          up2b save 1 "username password" 0 "token" ...

  This command will use [ imgtu.com ] as the default image bed.
  The configuration information must be enclosed in double quotation marks, and each field is separated by a space.
  The configuration information format of each image bed is as follows:
	- 0:sm.ms      => "token"
	- 1:imgtu.com  => "username password"
	- 2:gitee.com  => "token username repo folder"
	- 3:github.com => "token username repo folder"
  - 4:ipfs       => ""


Flags:
  -h, --help   help for save
```

