FROM debian:buster-slim

VOLUME ["/code"]

WORKDIR /code/
#COPY 目录
COPY cloudopscmd /code/cloudopscmd

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置编码
ENV LANG="en_US.UTF-8" \
  LANGUAGE="en_US:en" \
  LC_ALL="en_US.UTF-8"
ENTRYPOINT [ "/code/cloudopscmd httpServer" ]