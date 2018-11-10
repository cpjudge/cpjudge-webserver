FROM gobuffalo/buffalo:v0.12.7

RUN mkdir -p $GOPATH/src/github.com/cpjudge/cpjudge_webserver
WORKDIR $GOPATH/src/github.com/cpjudge/cpjudge_webserver

# Uncomment to run the binary in "production" mode:
# ENV GO_ENV=production

# Bind the app to 0.0.0.0 so it can be seen from outside the container
ENV ADDR=0.0.0.0

EXPOSE 3000

CMD buffalo dev run
