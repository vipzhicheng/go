Go
==

What is Go?
-----------

This Go is not the programming language which is created by Google. This is a shell command which can help managing multiple remote servers SSH accounts easily, so you do not need to remember so many accounts for those servers.

There are so many solutions which can help do similar things. I just hope you like this simple way.

OS requirements
---------------

The project is just tested on MacOSX, but should work on most linux distributions. It depends on expect command and bash environment, and the remote servers need to enable sshd.

If you feels something wrong, please [file an issue](https://github.com/vipzhicheng/go/issues/new) on this project, because I do not test the scripts on every case, but I am happy to help you out.

Features
--------

* Multi accounts configuration in one file, easy format
* Support comments to group accounts
* Support multi keywords filter
* Support select by dynamic number
* Support different port for each ssh account
* Support private key file
* Support -g option to attach -D7070

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
# IP USER:PASS:PRIVATE_KEY_FILE LABEL

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

* -g, with this option, you can add -D7070 to the connection.
* -h, help info, command format.

### support multi filter words

This is a very neat feature, if you need to manage tens of servers.

```
$ go foo bar
```

Inspiration & Thanks
--------------------

This is a feature that I need, I know it must be somewhere which I can find out the situation of managing multiple SSH accesses via expect. Then I found [this](http://imbugs.com/blog/articles/99.html), which is written in Chinese. Thanks for the code, most of idea is from that code, then I rewrite most of the code, add some new features.


