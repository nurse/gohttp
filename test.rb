#!/usr/bin/env ruby

require 'gohttp'
require 'minitest/autorun'

class TestGohttp < Minitest::Test
  def test_hello
    assert_output("hello, world\n") do
      Gohttp.hello
    end
  end

  def test_hello
    o = Object.new
    def o.to_str; "https://例え.テスト/ファイル.html"; end
    g = Gohttp.new
    assert_equal "https://%E4%BE%8B%E3%81%88.%E3%83%86%E3%82%B9%E3%83%88/%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB.html", g.get_url(o)
  end

  def test_get
    r = Gohttp.get('https://example.com/index.html')
    assert_kind_of Response, r
    assert_equal '200 OK', r.status
    assert_equal 200, r.status_code
    assert_equal 'HTTP/1.1', r.proto
    assert_equal 1270, r.size
    assert_equal [], r.transfer_encoding
    hdr = r.header
    assert_kind_of Hash, hdr
    assert_equal %w[max-age=604800], hdr["Cache-Control"]
    assert_equal %w[text/html], hdr["Content-Type"]
    b = r.body
    assert_match /\A#<GoIO:\*http.bodyEOFSignal:0x\h+>/, b.to_s
    assert_match /Example Domain/, b.read
    assert_equal true, b.close
    req = r.request
    assert_equal 'GET', req.method
    assert_equal 'HTTP/1.1', req.proto
    assert_equal Hash[], req.header
    assert_nil req.body
    assert_equal 0, req.content_length
    assert_equal [], req.transfer_encoding
    assert_equal 'example.com', req.host
    assert_equal '', req.remote_addr
    assert_equal '', req.request_uri
  end

  def test_head
    r = Gohttp.head('https://example.com/index.html')
    assert_kind_of Response, r
    assert_equal '200 OK', r.status
    assert_equal 200, r.status_code
    assert_equal 'HTTP/1.1', r.proto
    assert_equal 1270, r.size
    assert_equal [], r.transfer_encoding
    hdr = r.header
    assert_kind_of Hash, hdr
    assert_equal %w[max-age=604800], hdr["Cache-Control"]
    assert_equal %w[text/html], hdr["Content-Type"]
    b = r.body
    assert_match /\A#<GoIO:\*http.bodyEOFSignal>/, b.to_s
    assert_match '', b.read
    assert_equal true, b.close
    req = r.request
    assert_equal 'HEAD', req.method
    assert_equal 'HTTP/1.1', req.proto
    assert_equal Hash[], req.header
    assert_nil req.body
    assert_equal 0, req.content_length
    assert_equal [], req.transfer_encoding
    assert_equal 'example.com', req.host
    assert_equal '', req.remote_addr
    assert_equal '', req.request_uri
  end

  def test_head
    r = Gohttp.post_form('http://nalsh.jp/env.cgi', {'foo' => 'bar'})
    assert_kind_of Response, r
    assert_equal '200 OK', r.status
    assert_equal 200, r.status_code
    assert_equal 'HTTP/1.1', r.proto
    assert_equal -1, r.size
    assert_equal %w[chunked], r.transfer_encoding
    hdr = r.header
    assert_kind_of Hash, hdr
    assert_equal nil, hdr["Cache-Control"]
    assert_equal %w[text/html], hdr["Content-Type"]
    b = r.body
    assert_match /\A#<GoIO:\*http.bodyEOFSignal:0x\h+>/, b.to_s
    assert_match '', b.read
    assert_equal true, b.close
    req = r.request
    assert_equal 'POST', req.method
    assert_equal 'HTTP/1.1', req.proto
    assert_equal Hash["Content-Type"=>["application/x-www-form-urlencoded"]], req.header
    assert_match /\A#<GoIO:ioutil.nopCloser>\z/, req.body.to_s
    assert_equal 7, req.content_length
    assert_equal [], req.transfer_encoding
    assert_equal 'nalsh.jp', req.host
    assert_equal '', req.remote_addr
    assert_equal '', req.request_uri
  end
end

class TestGourl < Minitest::Test
  def test_escape
    assert_equal '%E3%81%82+%2B',  Gourl.query_escape('あ +')
  end

  def test_escape
    assert_equal 'あ  ', Gourl.query_unescape('%E3%81%82%20+')
    assert_raises(ArgumentError){ Gourl.query_unescape('%%81%82%20+') }
  end
end
