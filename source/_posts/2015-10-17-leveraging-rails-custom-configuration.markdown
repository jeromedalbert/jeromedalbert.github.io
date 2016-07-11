---
layout: post
title: "Leveraging Rails' custom configuration"
date: 2015-10-17 16:52
comments: true
categories: 
---

<img class="header" src="https://c2.staticflickr.com/6/5641/22236919056_b4cd72cd09_m.jpg" title="Pimp my Rails" />

These days, I use environment variables for credentials and other sensitive data
(`secrets.yml` could do the job but is too boilerplate-y for my taste). Dummy
environment values are stored in my project
<a href="https://github.com/bkeepers/dotenv" target="_blank">dotenv</a>
file and production values are deployed to the PaaS server of my choice with a
simple `heroku config:set` / `eb setenv` / etc.

But what about non-sensitive configuration values such as emails or your company
address?<!--more-->

Storing them in environment variables both in your project and in your server is
overkill and not really compliant with
<a href="http://12factor.net/config" target="_blank">12 factor apps</a>.
Those values are best stored in your project source code, in some sort of
configuration file.

Rails 4.2 introduced a
<a href="http://guides.rubyonrails.org/configuring.html#custom-configuration"
target="_blank">custom configuration object</a>
that can be accessed with `Rails.configuration.x`, and
<a href="http://edgeguides.rubyonrails.org/4_2_release_notes.html#railties-notable-changes"
target="_blank">Rails.application.config_for</a>
loads YAML file values depending on the environment. By combining those two
features, you can have all your custom configuration in a handy all-in-one YAML
file!

Say your project is called MyApp:

- Create a `my_app.yml` file in the `config` directory. It would look something
  like this:

```yaml
# MyApp custom configuration.
# Do not put sensitive data in this file (use environment variables instead).

default: &default
  emails:
    support: support@my-app.com
    marketing: marketing@my-app.com
  company_address: 221B Baker Street

production:
  <<: *default

development:
  <<: *default
  emails:
    support: support+dev@my-app.com

test:
  <<: *default
```

- Add this line in your `config/application.rb` file:

```ruby
# ...

module MyApp
  class Application < Rails::Application
    # ...

    # Custom configuration
    config.x = Hashie::Mash.new(Rails.application.config_for(:my_app))
  end
end

```

You can see that I wrap the loaded configuration hash in a
<a href="https://github.com/intridea/hashie#mash" target="_blank">hashie Mash</a>
object. I just use it for convenience so that I can access configuration
variables with dots rather than brackets (it's like `OpenStruct` with nested
hash support). To use it, add `gem 'hashie'` to your `Gemfile`.

- You can now access your configuration variables anywhere in your app like so:
```ruby
Rails.configuration.x.emails.support
```  

I hope that in future releases, Rails will improve its secret and configuration
handling so that they can be even easier to use out of the box. But for now,
this will do!
