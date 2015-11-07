require 'mkmf'
find_executable('go')
$objs = []
def $objs.empty?; false ;end
create_makefile("gohttp/gohttp")
case `#{CONFIG['CC']} --version`
when /Free Software Foundation/
  ldflags = '-Wl,--unresolved-symbols=ignore-all'
when /clang/
  ldflags = '-undefined dynamic_lookup'
end
File.open('Makefile', 'a') do |f|
  f.write <<eom.gsub(/^ {8}/, "\t")
$(DLLIB): Makefile $(srcdir)/gohttp.go $(srcdir)/wrapper.go
        CGO_CFLAGS='$(INCFLAGS)' CGO_LDFLAGS='#{ldflags}' \
          go build -p 4 -buildmode=c-shared -o $(DLLIB) $(srcdir)/gohttp.go $(srcdir)/wrapper.go
eom
end
