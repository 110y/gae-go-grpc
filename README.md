## Usage

```sh
cp .env.sample .env
docker-compose up -d
jo -p message=foo | curl -s -X POST localhost:20001/echo -d @- | jq .
{
  "message": "HELLO: foo"
}
```
