# Shell

Shell refers to a terminal shell like `bash`, `zsh`, `fish` etc. This problem
has been inspired by the existing shells.

Implement a shell. The following are the features that need to be
implemented, in an incremental manner!

It can do things like mkdir, touch, ls.

It does not work with the real file system. It's an in memory file system.
Once the program exits, it loses the file system we created.

mkdir makes exactly one directory. And gives error if directory already exists. 

touch works at one level, that is, in the current directory

touch works at multiple levels, that is nested directories

ls never gives error. ls shows list of files and directories. 

ls -l shows the differentiation between files and directories. 

ls -a to show hidden files. Hidden files are files starting with `.` and
this also includes the current working director `.` and the parent directory `..`

Implement mkdir -p which makes directories at different levels if they are not
present. Does not give error when directory already exists.

Implement echo to print strings

Implement echo to print shell environment variable `$?` to show status code / 
exit code of previous command. For errors it's 1 and for success it's 0.
No string interpolation needed, it would look like

```
$ echo $?
1
```

Implement cd. cd does not work for files but only for directories. 

cd .. and cd . should work

Implement cp and mv command for files. It will give error for directories. 
It can take exactly one source and one destination. 

cp and mv can take multiple sources and one destination. But only for files.

Implement cp and mv command for directories when -r flag is given, support
multiple sources and directories

Implement comments using `#`. Anything that starts with `#` should be ignored.
