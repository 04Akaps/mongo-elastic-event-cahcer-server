```
docker run -d -p 9200:9200 -p 9300:9300 \
-e "discovery.type=single-node" \
-e "ELASTIC_USERNAME=<사용할 이름>" \
-e "ELASTIC_PASSWORD=<사용할 패스워드>" \
--name elasticsearch-docker \
docker.elastic.co/elasticsearch/elasticsearch:7.14.0
```

```
docker run -d -p 5601:5601 --link elasticsearch-docker:elasticsearch --name kibana-docker docker.elastic.co/kibana/kibana:7.14.0
```

```
curl -u <사용자이름>:<비밀번호> -X GET "http://localhost:9200/<인덱스이름>/_search"
```