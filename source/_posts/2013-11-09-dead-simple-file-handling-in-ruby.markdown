---
layout: post
title: "Dead simple file handling in Ruby"
date: 2013-11-10 18:52
comments: true
categories: 
---

{% img right margintop http://upload.wikimedia.org/wikipedia/commons/3/36/Ken_n_dennis.jpg 200 The first hipsters. %}

**Common scenario**  
You are writing a basic Ruby script that needs to read or write in some text files.  
So you search "ruby write file" and "ruby read file" on Google.  
You end up on the usual tutorials and Stack Overflow answers that tell you to open files with `r` or `w` mode, "1970 C style"<!--more-->.

It's OK, it works.
But you may think:  
"I'm in the 21<sup>st</sup> century goddammit. I just want to read/write in a file, spare me the details!

Isn't Ruby supposed to be high level?  
Isn't Ruby supposed to be extremely terse?  
Isn't Ruby supposed to be super expressive as well?

File writing and reading should be simple as hell, right?"

Worry not, fellow citizen. Have a look at the code below.  
In all of these examples, Ruby automatically opens and closes the file. No need to think about it.

## Read a whole file

<div class="nolinenos">
``` ruby
content = File.read('input.txt')
```
</div>

Done.

## Write to a file
<div class="nolinenos">
``` ruby
File.write('output.txt', 'Hello world')
```
</div>

'Nuff said.

## What if I want to ...?

The above code covers 80%\* of the needs for quick Ruby scripts.
Feel free to skip the rest of this post.  
Just for the sake of being comprehensive, here are the other 20%\*, slightly more complex because used less often.

#### Read a file line by line (useful if you don't want to load a big file entirely in memory):

<div class="nolinenos">
``` ruby
File.open('input.txt', 'r').each_line { |line| puts line }
```
</div>
Another way to do it is `File.foreach('input.txt') { |line| puts line }`, though I don't find `foreach` very Ruby idiomatic.

#### Put the lines of a file in an array of lines (I've never needed to do it, but this is for reference):

<div class="nolinenos">
``` ruby
lines = File.readlines('input.txt')
```
</div>

#### Append some text at the end of a file:

<div class="nolinenos">
``` ruby
File.open('file.txt', 'a') { |file| file.write('some text') }
```
</div>
 
You can also do `file << 'some text'`. Or `file.puts('some text')` if you want to insert the text with a `\n` in the end.

<div class="references">
*These numbers are completely made up. Do not trust the Internet.
</div>
