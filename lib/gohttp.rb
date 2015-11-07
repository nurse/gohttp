require "gohttp/version"
begin
  require "gohttp/#{RUBY_VERSION[/\d+\.\d+/]}/gohttp"
rescue LoadError
  require "gohttp/gohttp"
end

class Gohttp
  # Your code goes here...
end
