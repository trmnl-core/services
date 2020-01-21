FROM micro/go-micro
ADD . /services
COPY entrypoint.sh /
WORKDIR /
RUN chmod 755 entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]
