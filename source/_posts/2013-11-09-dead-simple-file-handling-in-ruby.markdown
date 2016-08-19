---
layout: post
title: "Dead simple file handling in Ruby"
date: 2013-11-10 18:52
comments: true
categories: 
---

<img class="header margintop" src="http://farm4.staticflickr.com/3753/11699963024_bda6a4c9b8_m.jpg" title="The first hipsters. These guys developed the C programming language." /> 

**Common scenario**  
You are writing a basic Ruby script that needs to read or write in some text files.  
So you search "ruby write file" and "ruby read file" on Google.  
You end up on the usual tutorials and Stack Overflow answers that tell you to open files with `r` or `w` mode, "1970 C style"<!--more-->.

It's OK, it works. But you may think:  
"I'm in the 21<sup>st</sup> century goddammit. I just want to read/write in a file, spare me the details!"

Worry not, fellow citizen. Have a look at the code below.  
In all of these examples, Ruby automatically opens and closes the file. No need to think about it.

## Read a whole file

``` ruby
content = File.read('input.txt')
```

Done.

## Write to a file

``` ruby
File.write('output.txt', 'Hello world')
```

'Nuff said.

## What if I want to ...?

The above code covers 80%\* of the needs for quick Ruby scripts. Feel free to skip the rest of this post.  
Just for the sake of being comprehensive, here are the other 20%\*, slightly more complex because used less often.

#### Read a file line by line (useful if you don't want to load a big file entirely in memory):

``` ruby
File.open('input.txt', 'r').each_line { |line| puts line }
```
Another way to do it is `File.foreach('input.txt') { |line| puts line }`, though
I don't find `foreach` very Ruby idiomatic.

#### Put the lines of a file in an array of lines (I've never needed to do it, but this is for reference):

``` ruby
lines = File.readlines('input.txt')
```

#### Append some text at the end of a file:

``` ruby
File.open('file.txt', 'a') { |file| file.write('some text') }
```
 
You can also do `file << 'some text'`.
Or `file.puts('some text')` if you want to insert the text with a `\n` in the end.

<div class="references">
*These numbers are completely made up. Do not trust the Internet.
</div>
