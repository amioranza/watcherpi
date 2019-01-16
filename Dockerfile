FROM armhf/debian

LABEL maintainer="amioranza@mdcnet.ninja"
LABEL description="watcherpi"

WORKDIR /

COPY watcherpi /watcherpi

ENTRYPOINT [ "/watcherpi" ]