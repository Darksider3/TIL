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
printf "\n\n### Files: ${countFile}, Directorys: ${countDirs}, both: ${countAll}\n" >> README.md
git add .
```

Full List with correct indentation for directorys/subdirs:
You will need here mddir from the NPM-Manager. It inputs to the beginning and ending of the file three backticks.
```sh
# md-tree!
sed -i '/\# Project tree/Q' README.md
tree=$(tree -tf --noreport -I '*~' --charset ascii $1 |
       sed -e 's/| \+/  /g' -e 's/[|`]-\+/ */g' -e 's:\(* \)\(\(.*/\)\([^/]\+\)\):\1[\4](\2):g')
countDirs=$(find . -type d -not -path "*/.git/*" | wc -l)
countDirs=$(($countDirs-1))
countFile=$(find . -type f -not -path "*/.git/*" | wc -l)
countFile=$(($countFile-1))
countAll=$(($countDirs+$countFile))
mddir > /dev/null
sed -i '1s/^/\`\`\`\n/' directoryList.md
printf "\`\`\`" >> directoryList.md
printf "# Project tree\n\n${tree}\n" >> README.md
printf "\n\n### Files: ${countFile}, Directorys: ${countDirs}, both: ${countAll}\n" >> README.md

printf "\n\n Full list at [directoryList.md](./directoryList.md)" >> README.md
git add .
```

## My actual pre-commit file
```
#!/bin/sh
#
# An example hook script to verify what is about to be committed.
# Called by "git commit" with no arguments.  The hook should
# exit with non-zero status after issuing an appropriate message if
# it wants to stop the commit.
#
# To enable this hook, rename this file to "pre-commit".

if git rev-parse --verify HEAD >/dev/null 2>&1
then
	against=HEAD
else
	# Initial commit: diff against an empty tree object
	against=4b825dc642cb6eb9a060e54bf8d69288fbee4904
fi

# If you want to allow non-ASCII filenames set this variable to true.
allownonascii=$(git config --bool hooks.allownonascii)

# Redirect output to stderr.
exec 1>&2

# Cross platform projects tend to avoid non-ASCII filenames; prevent
# them from being added to the repository. We exploit the fact that the
# printable range starts at the space character and ends with tilde.
if [ "$allownonascii" != "true" ] &&
	# Note that the use of brackets around a tr range is ok here, (it's
	# even required, for portability to Solaris 10's /usr/bin/tr), since
	# the square bracket bytes happen to fall in the designated range.
	test $(git diff --cached --name-only --diff-filter=A -z $against |
	  LC_ALL=C tr -d '[ -~]\0' | wc -c) != 0
then
	cat <<\EOF
Error: Attempt to add a non-ASCII file name.

This can cause problems if you want to work with people on other platforms.

To be portable it is advisable to rename the file.

If you know what you are doing you can disable this check using:

  git config hooks.allownonascii true
EOF
	exit 1
fi

# If there are whitespace errors, print the offending file names and fail.
# exec git diff-index --check --cached $against --



# md-tree!
sed -i '/\# Project tree/Q' README.md
tree=$(tree -tf --noreport -I '*~' --charset ascii $1 |
       sed -e 's/| \+/  /g' -e 's/[|`]-\+/ */g' -e 's:\(* \)\(\(.*/\)\([^/]\+\)\):\1[\4](\2):g')
countDirs=$(find . -type d -not -path "*/.git/*" | wc -l)
countDirs=$(($countDirs-1))
countFile=$(find . -type f -not -path "*/.git/*" | wc -l)
countFile=$(($countFile-1))
countAll=$(($countDirs+$countFile))
mddir > /dev/null
sed -i '1s/^/\`\`\`\n/' directoryList.md
printf "\`\`\`" >> directoryList.md
printf "# Project tree\n\n${tree}\n" >> README.md
printf "\n\n### Files: ${countFile}, Directorys: ${countDirs}, both: ${countAll}\n" >> README.md

printf "\n\n Full list at [directoryList.md](./directoryList.md)" >> README.md
git add .
```