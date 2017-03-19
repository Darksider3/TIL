# Intro

Today i wanted to try something new on my Windows machine, and thought of a "workflow" starting-icon for automated starting of a set of specific executables.
This where in this test-case the Bash on Windows and Tagspaces. And both of them should contain the workflow-name "TIL" here.

This wasn't that easy, honestly. Without much explanation about what-the-heck google has no answer, i decided to use a batch script here. Not good with it. Never used it for more then some 'lil work. Nothing special things.

It pretty fast popped up the `start` command, which accept the following input strings:
```bat
start "TITLE" LOCATION(in double-qoutes if it got some special chars) ARGUMENTS
```

Easy.
## But... Why is `start` confusing then?

Did u notice that i only used the double-qoutes on the title? Here's why:

`start` avoids it here to start any program that is qouted, execpt these ones that are in directorys with special characters, and won't accept the arguments in qoutes as well from me.
Tried everything i imagine:
1
```bat
start "foo" C:\bin\helloworld.exe "Hi"
```
2
```bat
start "foo" C:\bin\helloworld.exe "H i"
```
3
```bat
start "foo" C:\bin\helloworld.exe "H" "I"
```
4
```bat
start "foo" C:\bin\helloworld.exe Hi
```
5
```bat
start "foo" C:\bin\helloworld.exe H i
```

The funny: It accepted the second one, and the last two(4&5). But not 1 and 3.
Why the hell...?

But i don't want to dig further there... it seems to be pretty messy.

## START has some magic path-arguments!

Just discovered /D`Path`, where `Path` means a directory inside the actual ENV, what on Windows includes the CWD.

You can just do, of course, a cd to come up a directory.
Or you can use a clean option! I choose the clean one.

```bat
start "foo" /DC:\bin helloworld.exe
```
And it even solves my qoute-problem. :smile: