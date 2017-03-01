# *.dat fail, APT

APT has an ability to lock all the files it actually uses.

However, this is in some cases pretty f\*cked up.
(Or is it debconf? Both are quiet shitty here)
If we discontinue an apt-process with an SIGKILL, it just leaves the locks around.
That way it could following error happen:
``` shell
root@user-desktop:/home/user# apt-get install flashybrid
Reading package lists... Done
Building dependency tree
Reading state information... Done
flashybrid is already the newest version.
0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
1 not fully installed or removed.
After this operation, 0B of additional disk space will be used.
debconf: DbDriver "config": /var/cache/debconf/config.dat is locked by another process: Resource temporarily unavailable
Setting up flashybrid (0.15+nmu2) ...
debconf: DbDriver "config": /var/cache/debconf/config.dat is locked by another process: Resource temporarily unavailable
dpkg: error processing flashybrid (--configure):
subprocess post-installation script returned error exit status 1
Errors were encountered while processing:
flashybrid
E: Sub-process /usr/bin/dpkg returned an error code (1)
```
(Source: http://askubuntu.com/questions/136881/debconf-dbdriver-config-config-dat-is-locked-by-another-process-resource-t)

To resolve that, we need at best `fuser` and `kill`. And root.

First, get the process holding the lock: 
``` shell
sudo fuser -v /var/cache/debconf/config.dat
USER        PID ACCESS COMMAND
/var/cache/debconf/config.dat:
root      18210 F.... dpkg-preconfigu
```
Note the PID, and kill it properly: 
``` shell
sudo kill PID
sudo kill -9 PID  # if the first doesn't work
```



In my situation it was a pending libc6. That's even worse.

So: Hitting CTRL+C 30 times isn't a good option, obviously.