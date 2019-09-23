# The git credential store - cache version, not stored on disk

Git is able to store your credentials, either on your disk - with `store` - or your memory, as a cache
To use the credential cache globally, activate it in your `~/.gitconfig` with:
```
git config --global credential.helper cache
git config --global credential.helper 'cache --timeout=3600'
```
For 3600seconds save in your memory :)
