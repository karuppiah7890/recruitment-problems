# Linter

This problem has been inspired by real life tools and software.

We want to design a linter. In programming, linter is a tool that lints source
code, that is checks source code. It checks source code for syntax error,
styling issues based on the idiomatic ways of the programming language and
some style rules.

You are going to design a linter for a simple English sentences. Given an
English sentence / paragraph as input to the linter, it has to look for issues/
errors based on the rules. Rules are defined below

1. Whenever a sentence starts with the word "because", warn the user saying it's
usually not a good thing to start sentences with "because".
2. Whenever there's a full stop ('.') or a comma (',') , it should be followed
by a space (' '). Throw an error if that's not the case
3. There must always be a full stop at the end of the sentence / paragraph.
Throw an error if that's not the case
4. The first word of every sentence must have it's first letter capitalised
(upper case). Throw an error if that's not the case.

Things to note - A full stop ('.') denotes the end of a sentence and the start
of the next sentence, if there's one.

The linter is a command line tool. It takes input from file, the command line
tool is give the file path as input. 

```bash
# example
$ linter input.text
```

Once the linter finds the warnings and errors, it has to list them along with
the count of errors and warnings.

How to show the errors or warnings - 
1. Show it as list
2. For each item in the list, tell if it's an error or a warning
3. Mention what rule was not followed

When there is even one error, it has to exit the program with a non-zero exit
code to show that the linting failed or the lint process was a failure. When
there no errors, irrespective of the count of warnings, exit the program with
exit code zero to show the linting was successful. 

Example:

```bash
$ cat input.txt
This is the first sentence. it is different from a line
as lines are separated by new lines and sentences are
separated by dots or full stops. Because of this they
are different

$ linter input.text
* [ERR] The first word of every sentence must have first letter capitalised
* [WARN] It is usually not a good thing to start sentences with "because"

1 Error
1 Warning

$ echo $?
1
```

Tough Extension: When showing errors or warnings, tell the position of the character in the paragraph about where the error or warning occurred. Position has to be of the form "1st line 50th character", that was an example. This will be helpful for the user to correct the issue to fix the warnings or errors. And line is defined based on new line character ('\n') in the paragraphs. 

Example: 

```bash
$ cat input.txt
This is the first sentence. it is different from a line
as lines are separated by new lines and sentences are
separated by dots or full stops. Because of this they
are different

$ linter input.text
* [ERR] The first word of every sentence must have first letter capitalised.
Issue is at 1st line 29th character

* [WARN] It is usually not a good thing to start sentences with "because".
Issue is at 3rd line 34th character

1 Error
1 Warning

$ echo $?
1
```

Another extension (easy): When showing error or warning, apart from saying position, also show a maximum of 5 characters before and a maximum of 5 characters after that position, based on the availability of characters. And also highlight the place where the error has occurred using '[ ]' symbols. 

Example:

```bash
$ cat input.txt
This is the first sentence. it is different from a line
as lines are separated by new lines and sentences are
separated by dots or full stops. Because of this they
are different

$ linter input.text
* [ERR] The first word of every sentence must have first letter capitalised.
Issue is at 1st line 29th character

"nce. [i]t is "

* [WARN] It is usually not a good thing to start sentences with "because".
Issue is at 3rd line 34th character

"ops. [Because] of t"

1 Error
1 Warning

$ echo $?
1
```

Looks like in the above case, it makes sense to highlight the whole word
`Because` instead of just `B`. May be the question could be changed to allow
different kind of rules - some based on characters, some based on words. For
characters, it's a single position to point out and show error. For words, it's
a set of positions to point out the whole word - in fact it's a more general
version and a word with one character is a specific thing ;)

And a few characters before and after does not look that great. May be we can
convert this question as - show a maximum of 2 words before and 2 words after
or something like that. 

Example:

```bash
$ cat input.txt
This is the first sentence. it is different from a line
as lines are separated by new lines and sentences are
separated by dots or full stops. Because of this they
are different

$ linter input.text
* [ERR] The first word of every sentence must have first letter capitalised.
Issue is at 1st line 29th character

"first sentence. [i]t is different"

* [WARN] It is usually not a good thing to start sentences with "because".
Issue is at 3rd line 34th character

"full stops. [Because] of this"

1 Error
1 Warning

$ echo $?
1
```

That looks neater ;) :)

Easy extension - provide options for the tool through flags (options).
Define a toggle flag called strict. A toggle flag denotes a feature that
can be toggled on and off by using the flag and it takes no arguments.

And when strict is enabled, it means that the linting process is strict and
does not allow even warnings, let alone errors. So, if there's one or more
warnings, or even one or more errors, it makes the program exit with non-zero
exit code. Only when there are zero errors and zero warnings, linter exits with
exit code zero when strict is enabled.

Example

```bash
# without strict flag
$ cat input.txt
This is the first sentence. It is different from a line
as lines are separated by new lines and sentences are
separated by dots or full stops. Because of this they
are different

$ linter input.text
* [WARN] It is usually not a good thing to start sentences with "because".
Issue is at 3rd line 34th character

"full stops. [Because] of this"

0 Error
1 Warning

# notice the exit code here
$ echo $?
0

# with strict flag
$ linter --strict input.txt
* [WARN] It is usually not a good thing to start sentences with "because".
Issue is at 3rd line 34th character

"full stops. [Because] of this"

0 Error
1 Warning

# now, notice the exit code here!
$ echo $?
1
```
