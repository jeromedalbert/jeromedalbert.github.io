---
layout: post
title: "A diagram of the Ruby Core object model"
date: 2013-10-05 20:26
comments: true
categories: 
---

I am a visual learner. A picture is worth a thousand words.

When I started learning Ruby, I could not find a decent diagram that would just sum up the Ruby classes, modules, and their hierarchy.

So I made mine.<!--more-->  
By hand.  
On an A4 sheet.  

Here is the result (click for a full size view):

<a href="http://farm6.staticflickr.com/5443/10075536704_84aa13676a_o.jpg" target="_blank">{% img http://farm6.staticflickr.com/5443/10075536704_b78362422d_b.jpg %}</img></a>

Most of Ruby's guts at one glance: pretty neat, huh ?

 - Modules are on the top left.
 - The usual types are here: strings, arrays, hashes, numbers, "booleans" (actually `TrueClass` and `FalseClass`).
 - The error hierarchy is on the right.
 - Other classes are in the middle.

Note: To have a decent layout, I only drew 70% of the <a href="http://ruby-doc.org/core-2.0" target="_blank">Ruby 2 Core classes</a>. All of the important ones are there, except `Random`, `Time`, `GC` and `ObjectSpace`: when I noticed, it was too late. :-)

## What can we "draw" from this ?

Well, this model is pretty much what we would have expected:

 - Arrays are enumerable, numbers are comparable.
 - "Everything is an object", or more precisely, every box is this diagram is a BasicObject.
 - Object includes the Kernel module, so Object has all the instance methods of Kernel, like `puts`, `gets`, `exit`.  
 When you write Ruby code, you are inside a `main` object, that's why you can just write `puts 'Hello world'` instead of `Kernel.puts 'Hello world'`.

So, everything makes sense and is quite simple actually !  
Except maybe the fact that a module is a class, and class inherits from module. Don't worry, you will eventually wrap your head around it.

## Going further

In addition to Ruby Core, the <a href="http://www.ruby-doc.org/stdlib-2.0.0/" target="_blank">Ruby Standard Library</a> contains many other very useful classes and modules, such as `Date`, `Forwardable`, `SimpleDelegator`, `Net::HTTP`, `OpenStruct`, etc.

This diagram does not explain the Ruby method lookup path, which takes eigenclasses, included modules and superclasses into account. I recommend reading this <a href="https://github.com/elm-city-craftworks/practicing-ruby-manuscripts/blob/master/articles/v1/001-method-lookup.md" target="_blank">enlightening article</a> from Practicing Ruby.
