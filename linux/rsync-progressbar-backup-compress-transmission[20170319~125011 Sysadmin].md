Needed it to recursivly copy an Windows User directory to new HDD.

```sh
rsync --progress -z -P $SOURCE $DEST
```

Will backup $SOURCE recursivly to $DEST and compress(`-z`) it during the transmission. Of course, it get decompressed on the destination.
This is especially usefull if you can only use some slow connection, like USB2 or a 1 Mbit-DSL.

`--progress` will activate output the actual file it is compressing and transfering. 
Beware: This will spam your Commandline-Window pretty fast with very much output :wink: You are "warned".
