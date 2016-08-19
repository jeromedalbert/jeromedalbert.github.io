---
layout: post
title: "JSONB + Rails to store glorified strings in your DB"
date: 2016-06-27 08:48
comments: true
categories: 
---

<img class="header" src="https://c2.staticflickr.com/8/7364/27257296733_bc2c0d0de2_m.jpg" title="MacGyver would be proud of this." />

Everyone likes a good old quick and dirty hack every once in a while.
For example, who hasn't stored — or at least hasn't been tempted by storing — a list of values as a simple string in the database?

Let's see when this could be a good choice, and how to use Postgres JSONB to make our lazy decision nicer<!--more--> to implement!

## The case for storing strings in DB

Say we are working on a simplified app where products have predefined available sizes.
When buying a product, users can choose between these available sizes.
For example a T-shirt could come in S, M, L and XL sizes.
A pack of cookies could have "tiny", "reasonable" and "super-size-me" editions. Etc.
When trying to implement this, we are faced with a choice.

1. We could "just wing it™" and store a comma-separated string.
   Our T-shirt would have the `"S,M,L,XL"` string in an `available_sizes` field.
   When editing this product, the front-end would send this string and the back-end would store it as is.
   When displaying these sizes to our users in a select box, our code would `split(',')` to be able to iterate on them.

2. Another possibly cleaner solution would be to have a separate `sizes` table.
   A product has many sizes and a size can belong to many products.
   This is what they usually teach you in traditional Computer Science classes.

Which solution is best?
As usual it's all about trade-off and it depends on your context.
Solution 2 has no data duplication and is possibly easier to query for stats and whatnot,
but it involves two additional tables:
`sizes` + a table to store our many-to-many relationship.
For a simple feature like this, especially if the future is uncertain, I would argue that solution 1 is often good enough.
It is the quickest to implement with minimal effort, and easily modifiable in the future if needed.
If one day people need to do aggregate queries for stats purposes (inefficient/cumbersome and sometimes impossible with strings),
one can always improve later with a migration.
Either by implementing the second solution with additional tables, or by using some form of noSQL storage, depending on the needs.

## A nicer way with JSONB

So you've decided to store good ol' comma-separated strings in your app.
You probably didn't need to read a blog post for that!
But some of you may not know that the Postgres JSONB format could fit the bill quite nicely as a glorified string.

JSONB allows you to store JSON as is, and Rails has very good support for it.
How good? Well, it means that your back-end can literally store the `["S", "M", "L", "XL"]` array, which is valid JSON.
When retrieving the value from the database, Rails presents it to you as a Ruby array, so you can directly iterate on it without doing any `split`!
It's arrays all the way down, no strings attached (bad pun intended).

All you need is Postgres 9.4+ and a migration like this:

```ruby
class AddSizesToProducts < ActiveRecord::Migration
  def change
    add_column :products, :sizes, :jsonb, default: []
  end
end
```

And you're good to go.

*EDIT: As Hauleth has pointed out in the comments, Postgres has an <a
href="https://www.postgresql.org/docs/current/static/arrays.html"
target="_blank">array</a> data type exactly for this use case. The
corresponding migration would then be:*

```ruby
class AddSizesToProducts < ActiveRecord::Migration
  def change
    add_column :products, :sizes, :text, array: true, default: []
  end
end
```

## Conclusion

With this technique you get the advantages of a quick and easy implementation without the uneasy feeling of handling dirty strings.
If you want to drink more of this JSONB kool-aid, I recommend reading
<a href="https://blog.codeship.com/unleash-the-power-of-storing-json-in-postgres" target="_blank">this post from the Codeship blog</a>,
especially if you are planning to store javascript objects (pro-tip: you can use `default: {}` in your migrations if needed).
In the meantime, happy hacking!
