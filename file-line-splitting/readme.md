# File line splitting

Given a file with many lines of text, the user wants to have a maximum length of
80 characters per line. But the file may contain lines that are longer than 80
characters. So, write a program to help the user cut the lines in the file and
put them in the next line and maintain the maximum length in all lines.The user
does not want to cut words in the line while cutting lines, so make sure the
program does not cut words. Words are separated by a space.

Extension: if the whole file has been read to implement the program, optimize
for it to work ith larger files. In case of larger files, it's not good to read
the whole file into the program memory (RAM), so optimize the solution.
