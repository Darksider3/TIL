# adding a pre-commit-hook to insert an file-structure tree into the README.md or refreshing it


Just put the following into the `.git/hooks/pre-commit file:`

``` sh
# md-tree!
sed -i '/\# Project tree/Q' README.md
tree=$(tree -tf --noreport -I '*~' --charset ascii $1 |
       sed -e 's/| \+/  /g' -e 's/[|`]-\+/ */g' -e 's:\(* \)\(\(.*/\)\([^/]\+\)\):\1[\4](\2):g')
countDirs=$(find . -type d -not -path "*/.git/*" | wc -l)
countDirs=$(($countDirs-1))
countFile=$(find . -type f -not -path "*/.git/*" | wc -l)
countFile=$(($countFile-1))
countAll=$(($countDirs+$countFile))
printf "# Project tree\n\n${tree}\n" >> README.md
git add .
```

The counter-variables are here for a specific use: If you wan't to list them too!

Full with counter: 
``` sh
# md-tree!
sed -i '/\# Project tree/Q' README.md
tree=$(tree -tf --noreport -I '*~' --charset ascii $1 |
       sed -e 's/| \+/  /g' -e 's/[|`]-\+/ */g' -e 's:\(* \)\(\(.*/\)\([^/]\+\)\):\1[\4](\2):g')
countDirs=$(find . -type d -not -path "*/.git/*" | wc -l)
countDirs=$(($countDirs-1))
countFile=$(find . -type f -not -path "*/.git/*" | wc -l)
countFile=$(($countFile-1))
countAll=$(($countDirs+$countFile))
printf "# Project tree\n\n${tree}\n" >> README.md
printf " ### Files: ${countFile}, Directorys: ${countDirs}, both: ${countAll}\n" >> README.md
git add .
```