# Gohttp

A demo to implement a extension library with Go.

## Installation

Add this line to your application's Gemfile:

```ruby
gem 'gohttp'
```

And then execute:

    $ bundle

Or install it yourself as:

    $ gem install gohttp

## Usage

```ruby
require 'gohttp'
r = Gohttp.get('https://example.com/index.html')
puts r.body.read
```

## Contributing

1. Fork it ( https://github.com/nurse/gohttp/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
