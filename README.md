Checking http requests.

### Request
```bash
curl -X POST \
-H "Content-Type: application/json" \
-d '{
  "id":"ID",
  "password":"PASSWORD"
}' \
"http://localhost:8000/test?aaa=bbb"
```

### Console
```bash
===== Remote Address =====
[::1]:54808
===== Method =====
POST
===== Host =====

===== URI =====
/test?aaa=bbb
===== Path =====
/test
===== RawQuery =====
aaa=bbb
===== Header =====
[Accept]:[*/*]
[Content-Type]:[application/json]
[Content-Length]:[40]
[User-Agent]:[curl/7.64.1]
===== Form =====
[aaa]:[bbb]
===== Body =====
{
  "id":"ID",
  "password":"PASSWORD"
}
```