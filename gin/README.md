- POST /
  - 実行例
```
curl --location 'http://localhost:8080/?id=0001' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "name": "hoge",
    "mail": "aaa@test.com"
}'
```
- Get /:name/:id
  - 実行例
```
http://localhost:8080/hoge/0123456789
```
