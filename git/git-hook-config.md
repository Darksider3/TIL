# Configure git's hook directory!

```
ln -s .git/hooks/ Hooks/ # link hook directory into add-able directory
git config core.hooksPath Hooks/ #use from now on $gitDir/Hook as your hook dir!
```

