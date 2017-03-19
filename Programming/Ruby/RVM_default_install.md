# Installing Ruby Version manager(RVM)

I had recently the much pleasent experience to contribute to an repository here on github, again. 
This time, i wanted to give an introduction in the installation of [tomato.es](https://tomato.es).
First off, i started to install Ruby trough my Packet Manager(APT).
Did work. But was too old. What the heck should we do?
In one issue they mentioned `RVM`, which could handle the Ruby Versions without even installing it globally. It get installed trough a ~/.rvm directory and adding it to path with an line in the .bashrc/.bash\_profile.

It gets easy now:

```sh 
gpg --keyserver hkp://keys.gnupg.net --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3
\curl -sSL https://get.rvm.io | bash -s stable
```
(Source: [RVM.io](https://rvm.io/))

Want a all-in-one Version? Append `--ruby`:
```sh
\curl -sSL https://get.rvm.io | bash -s stable --ruby
```

I wish sometimes, python has something like this. `virtualenv` is nice, no doubt about it. But the possibility to fastly install a different version into a user-home directory, without the need to populate it across all ENVs, is nice. And getting it maintained. :smile: