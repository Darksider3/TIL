With `tar czf` you can compress any directory, you have access to, to a .tar.gz file.

Not that `tar` is some command with cryptic parameters...
```sh
tar czf /backup/backup.tar.gz /
```

## Verbosity first
If you want verbose output, i.e. the status of the compression/decompression, simply add `--verbose`, `-v` or `v` to the parameters.

## The actual thing!

`tar` is obvious. `c` stands for `--create`, `z` for `--gzip` and `f` for `--file=ARCHIVE`, and the `/` is obv. the directory you want to compress.
Expanded command is, therefore:
```
tar --create --gzip --file=/backup/backup.tar.gz /
```
Ofcourse, this won't work if you didn't exclude `/backup`, because the file changes every time tar finds a new file and compressed it.

We need to exclude. `--exclude`!
Some directorys, that i tend to exclude everytime:
```sh
tar czf /backup/backup.tar.gz --exclude=/backup --exclude=/sys --exclude=/dev --exclude=/proc --exclude=/mnt --exclude=/media /
```
Backup's just the entire partition, but explicity won't touch any devices and mounted(in `/media` or `/mnt`) partitions.

## Play the backup back!

To do so, we need to boot into another distro, use here an livecd or some other installed linux.

Now we mount the "broken" partition, and decompress the backup into `/` of it.

```sh
mount /dev/sda1 /mnt/
tar xf /mnt/backup.tar.gz -C /mnt
```
Expanded `tar` here: `x`=`--extract` or `--get`, `f` is here again the .tar.gz file and `-C` means that we define a specific directory where we want to extract the file to(`--directory=DIR`):
```sh
tar --extract --file /mnt/backup.tar.gz --directory=/mnt
```
The `gzip` here is detected by `tar` automatically. But, if you want it to: Add `--gzip` or `-z`/`z`:
```sh
tar xfz $tar.gz.File -C $OutputDir
```
And we restored all directorys. Except, surely, the excluded ones.
