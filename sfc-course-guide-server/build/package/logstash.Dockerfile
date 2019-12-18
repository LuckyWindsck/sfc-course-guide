ARG ELK_VERSION

# https://github.com/elastic/logstash-docker
FROM docker.elastic.co/logstash/logstash:${ELK_VERSION}

# Add your logstash plugins setup here
# Example: RUN logstash-plugin install logstash-filter-json
RUN logstash-plugin install logstash-input-jdbc
COPY ./configs/logstash/driver/${LOGSTASH_JDBC_DRIVER} /usr/share/logstash/logstash-core/lib/jars/${LOGSTASH_JDBC_DRIVER}
