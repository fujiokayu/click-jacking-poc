# click-jacking-poc

## run web server

### go

```shell
go run server.go
```

### python

```shell
python -m http.server
```

## how to use

1. `git clone https://github.com/fujiokayu/click-jacking-poc.git && cd click-jacking-poc`
2. run a local web server at server.go or any way you like.
3. open your local web server in your browser  
  http://localhost:3000/
4. set the URL you wish to verify in the textarea.
5. click the 'load url' button.
6. if you have not done anything to prevent Click Jacking, your website will be displayed with opacity set to 0.1.

## references

- [X-Frame-Options - MDN Web Docs](https://developer.mozilla.org/ja/docs/Web/HTTP/Headers/X-Frame-Options)
- [CSP: frame-ancestors - MDN Web Docs](https://developer.mozilla.org/ja/docs/Web/HTTP/Headers/Content-Security-Policy/frame-ancestors)
