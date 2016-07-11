---
layout: post
title: 'Ruby: How to iterate "the right way"'
date: 2014-01-05 14:12
comments: true
categories: 
---

<img class="header"
src="http://farm8.staticflickr.com/7377/11753946913_0bbf38c9bf_m.jpg" title=
"Froot loops. Loop, iteration... Got it? Hahaha. So hilarious. We're having a ball here."
/>

<div class="left-header">
{% blockquote Professor David West, <a href="http://vimeo.com/77415896#t=48m04s" target="_blank">OOP is Dead! Long Live OODD!</a> %}
If any of you have written code in the last year that had an explicit loop [...], you don't understand collections.
{% endblockquote %}
</div>

You may be baffled by this quote if you come from a C-flavored language such as
C++ and Java<!--more-->, where explicit loops like `for` and `foreach` are part
of your everyday life.

Luckily, collection methods come built-in with Ruby. Once you realize how
powerful they are compared to explicit loops, there is no going back!

## each

Let's begin with the collection method that has the least added value.  
`each` is the equivalent of a `for` loop. Use it when you need to iterate on a
collection with side effects.

``` ruby
['a', 'b', 'c'].each do |e|
  # do side effects, such as printing to the console, writing to a file, persisting in database, etc.
end
```

For your information, Ruby also has a `for` statement (nobody uses it though).

## map

*(alias to avoid: collect)*  
Whenever you need to transform some values into some other values, `map` is your
friend.

``` ruby
[1, 2, 3].map { |e| e * 2 } # returns [2, 4, 6]

['a', 'b', 'c'].map { |e| e.upcase } # returns ['A', 'B', 'C']

['a', 'b', 'c'].map(&:upcase) # same result as before, for experienced lazy Ruby programmers
```

`map` is your bread and butter. I probably use it more than `each`.

## select

*(alias to avoid: find_all)*  
Very useful when you need to filter (i.e. "select") multiple values.

``` ruby
[1, 2, 3, 4].select { |e| e % 2 == 0 } # returns [2, 4]
```

## reject

The contrary of `select`

``` ruby
[1, 2, 3, 4].reject { |e| e % 2 == 0 } # returns [1, 3]
```

## partition

`select` + `reject`

``` ruby
[2, 3, 4, 5].partition { |e| e.even? } # returns [[2, 4], [3, 5]]
```

## find

*(alias to avoid: detect)*  
Very useful when you need to find a single value.

``` ruby
[4, 6, 8, 13].find { |e| e > 7 } # Returns 8 (the first found element)
```

## reduce

*(alias to avoid: inject)*  
`reduce` is very important in functionnal programming, where it is also known as
`fold`. Indeed, it can be used to build everything else: `map`, `select`,
`find`, `min`, `max`, sums, etc.  
The idea is to use an accumulator that will contain the final result. This
accumulator can be anything: a number, an array, a hash, etc.

``` ruby
[1, 2, 3].reduce(0) { |acc, e| acc + e } # returns 6 (0 + 1 + 2 + 3).
                                         # 0 is the initial value of the accumulator.

[1, 2, 3].reduce { |acc, e| acc + e } # same result as before. When we omit the accumulator initial
                                      # value, the first element of the array is chosen.
                                      # So in our case, 1 is the initial value of the accumulator.
                                      # What is computed is 1 + 2 + 3.

[1, 2, 3].reduce(:+) # same result as before, for experienced lazy Ruby programmers.
                     # Reduce accepts :+, but you can also use &:+

[2, 3, 4].reduce { |acc, e| acc * e } # returns 24 (2 * 3 * 4)

[0, 2, 3, 4].reduce { |acc, e| acc * e } # returns 0 (0 * 2 * 3 * 4)
```

I rarely use `reduce` in practice, but it is fun and good to know for your
computer science culture.

## all?

``` ruby
[2, 4, 6].all? { |e| e.even? } # returns true
```

Self-explanatory.

## any?

``` ruby
[3, 8, 42].any? { |e| e > 10 } # returns true
```

Self-explanatory.

By the way, I like using the "no block" form of `any?` to ask if something "has
any" significant element:
``` ruby
[3, 4].any? # returns true
[].any? # returns false
[nil].any? # returns false
[false].any? # returns false
```

I find it more expressive than `!some_array.empty?`. This is just a matter of
personal taste: feel free to use whichever you want.

##  times

This is a funny one!

``` ruby
3.times { puts 'Hello world!' } # prints 'Hello world!' 3 times

3.times { |i| puts i } # prints '0', '1' and '2'
```

In practice I don't use it very often though.

## Sorting methods

``` ruby
[7, 2, 5].sort # returns [2, 5, 7]
['c', 'b', 'a'].sort # returns ['a', 'b', 'c']

employees.sort_by {|e| e.last_name} # sort your employees by last name
```

## But... I want indexes!

On top of the current element, you also need the current index? Worry not,
fellow citizen. Ruby has it all.

``` ruby
['a', 'b', 'c'].each_with_index do |e, i|
  # do stuff
end

['a', 'b', 'c'].map.with_index do |e, i|
  # do stuff
end
```

In practice I don't use `each_with_index` very often. I never had to use
`map.with_index`, but I put it for the sake of being comprehensive.

## Explicit loops

Collection methods will probably cover 90% of your needs. What is certain is
that you won't need to use a `for` loop ever again. 

However, there are cases when the number of iterations in not known in advance:
if you work on an algorithm or low-level code, for example. This is a typical
job for `while` and `until`. In this case, these explicit loops are ok to use.

``` ruby
finished = false
until finished
  # do stuff
end

x = 100
while x > 0
  # do stuff
end
```

You may want to use an infinite loop: for example the main infinite loop of a
video game. Here is the syntax:

``` ruby
loop do
 # do stuff.
 # You can still exit the loop with break.
end
```

Same as `while true`, but terser.

## Going further

Ruby collection methods become even more powerful if you combine them:

``` ruby
['coconut', 'lemon', 'banana', 'apple'].select { |e| e.size > 5 }
                                       .map    { |e| e.upcase }
                                       .sort

# returns ['BANANA', 'COCONUT']
```

Don't do this to excess though!

This article covers the main collection methods: feel free to dive into the
<a href="http://ruby-doc.org/core" target="_blank">Ruby Core documentation</a>
for more.  
Finally, if you are wondering why I avoid aliases such as *collect* and
*inject*, you can read this community-driven
<a href="https://github.com/bbatsov/ruby-style-guide#naming" target="_blank">Ruby style guide</a>
and submit a bug if you don't agree. :-)
