require 'bundler'
Bundler::GemHelper.install_tasks
require 'rspec/core/rake_task'
require 'rake/extensiontask'

RSpec::Core::RakeTask.new(:spec)
task :default => [:compile, :spec]

spec = eval File.read('gohttp.gemspec')
Rake::ExtensionTask.new('gohttp', spec) do |ext|
  ext.ext_dir = 'ext/gohttp'
  ext.lib_dir = File.join(*['lib', 'gohttp', ENV['FAT_DIR']].compact)
  ext.source_pattern = "*.{c,cpp,go}"
end

RSpec::Core::RakeTask.new :spec  do |spec|
  spec.pattern = 'spec/**/*_spec.rb'
end
