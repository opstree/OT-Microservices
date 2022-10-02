FROM docker.elastic.co/elasticsearch/elasticsearch:7.17.6

ENV ES_JAVA_OPTS "-Xms512m -Xmx512m"

ENV ELASTIC_PASSWORD "elastic"

COPY ./elasticsearch.yml /usr/share/elasticsearch/config/elasticsearch.yml
