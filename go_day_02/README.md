## Command line utilities

### Finding Things

Run "godoc" to access the following links.
- [Cmd Package](http://localhost:6060/pkg/main/myFind/cmd/app)
- [Internal Package Run](http://localhost:6060/pkg/main/myFind/internal/app/run)
- [Internal Package Calculations](http://localhost:6060/pkg/main/myFind/internal/app/walker)

This utility replicates the functionality of the `find` command, 
enabling users to locate entries of various types, 
including directories, regular files, and symbolic links,
by specifying a path and a set of command-line options.

To use the program, execute it as follows:

```bash
# Finding all files/directories/symlinks recursively in directory /foo
./myFind /foo
```

Use -sl, -d, or -f to selectively print only symlinks, directories, or files. 
You can specify one, two, or all three of them. For instance:

Specify `-sl`, `-d`, or `-f` to print only symlinks, directories, or files.
You can specify one, two, or all three of them. For instance:

```bash
./myFind -f -sl /path/to/directory
./myFind -d /path/to/other/directory
```

Additionally, an extra option, -ext, has been introduced. 
This option is ONLY active when -f is specified,
allowing you to filter and print only files with a specific extension.

The program excludes files and directories that the current user lacks access permissions for, 
avoiding runtime errors by omitting these cases from the output.


### Counting Things

Run "godoc" to access the following links.
- [Cmd Package](http://localhost:6060/pkg/main/myWc/cmd/app)
- [Internal Package Run](http://localhost:6060/pkg/main/myWc/internal/app/run)
- [Internal Package Calculations](http://localhost:6060/pkg/main/myWc/internal/app/counter)

The utility presented here mirrors the functionality of the `wc` command.

Let's assume the files are utf-8 encoded text files, 
signifying that the program can handle texts in both English and Russian. 
It disregards punctuation and considers spaces as the sole word delimiters.

To enhance usability, three mutually exclusive flags have been implemented:
- `-l` for counting lines;
- `-m` for counting characters;
- `-w` for counting words.

The program is executed using the following format:
```bash
# Counting words
~$ ./myWc -w input.txt
# Counting lines
~$ ./myWc -l input2.txt input3.txt
# Counting characters
~$ ./myWc -m input4.txt input5.txt input6.txt
```


### Running Things

Run "godoc" to access the following links.
- [Cmd Package](http://localhost:6060/pkg/main/myXargs/cmd/app)
- [Internal Package Run](http://localhost:6060/pkg/main/myXargs/internal/app/run)

There is a tool similar to `xargs`. 
It treats all parameters as a command, such as 'wc -l' or 'ls -la', 
and constructs a command by appending all lines received from the program's stdin as arguments, 
then executes it. For example, running:

```bash
~$ echo -e "/a\n/b\n/c" | ./myXargs ls -la
```

It is equivalent to executing:

```bash
~$ ls -la /a /b /c
```

You can test this tool in conjunction with previous ones, so running:

```bash
~$ ./myFind -f -ext 'log' /path/to/some/logs | ./myXargs ./myWc -l
```

will calculate line counts for all ".log" files in the /path/to/some/logs directory recursively.


### Archiving Things

Run "godoc" to access the following links.
- [Cmd Package](http://localhost:6060/pkg/main/myRotate/cmd/app)
- [Internal Package Run](http://localhost:6060/pkg/main/myRotate/internal/app/run)
- [Internal Package Calculations](http://localhost:6060/pkg/main/myRotate/internal/app/archiver)
- 
The most recent addition to the tools is the log rotation utility.

"Log rotation" is a process in which the old log file is archived 
and stored separately to prevent logs from accumulating indefinitely in a single file. 
Its functionality is demonstrated as follows:

```bash
# Creates a file /path/to/logs/some_application_1600785299.tag.gz
# where 1600785299 is a UNIX timestamp obtained from `some_application.log`'s
~$ ./myRotate /path/to/logs/some_application.log
```

```bash
# Creates two tar.gz files, each with a timestamp (one for each log),
# and places them in the /data/archive directory
~$ ./myRotate -a /data/archive /path/to/logs/some_application.log /path/to/logs/other_application.log
```
