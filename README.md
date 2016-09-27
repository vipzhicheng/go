Go
==

What is Go?
-----------

This Go is not the programming language from Google. This is a shell command which can help manage multiple remote servers SSH access easily, and do not need to remember so many accounts for those servers.

OS requirements
---------------

The command tested on MacOSX, but should be working on most linux distributions. It depends on expect on your local system, and the remote servers need to enable SSH.

If you feels something wrong, please [file an issue](https://github.com/vipzhicheng/go/issues/new) on this project, because I do not test the scripts on every case, but I am happy to help you out.

Usage
-----

I am just showing the usage for MacOSX users, other OS usage maybe different, but the idea is same.

### Check out code from Github

```
$ mkdir ~/bin
$ cd ~/bin
$ git clone https://github.com/vipzhicheng/go.git
$ cd go
$ cp .go.conf.example ~/.go.conf
$ chmod a+x go
$ chmod a+x ssh-expect
```

### Set PATH in .bash_profile

```
export PATH=~/bin/go:~/bin:$PATH
```

```
$ source ~/.bash_profile
```

### Set ~/.go.conf, you can see demo settings as follows.

```
# IP USER:PASS LABEL

192.168.1.7:22000 user1:pass1 label:7
192.168.1.8 user2:pass2 label:8
192.168.1.9 user3::absolute_private_file_path label:9
```

You can ignore port setting if you are using default port(22) in remote server.

### How to use this command

```
$ go label

Found follow servers: (Which one do you want to connect?)
[1] user1@192.168.1.7 label:7
[2] user2@192.168.1.8 label:8
Please choose by ID:
1

Logging into user1@192.168.1.7 ...
spawn ssh user1@192.168.1.7 -p 22000
user1@192.168.1.7's password:
Last login: Mon Mar 10 18:35:02 2014 from 192.168.1.6
$
```

### Options

There is only one option, -g, with this option, you can add -D7070 to the connection.

### Change log

2015-06-01
1. Ignore comments lines
2. Add support for secret file

Inspiration & Thanks
--------------------

I know it must be somewhere about the situation of manage multiple SSH accesses via expect. Then I found [this](http://imbugs.com/blog/articles/99.html), which is written in Chinese. Thanks for the code, most of this project is from that code, but some features I need are missing. so I added them into this project.


