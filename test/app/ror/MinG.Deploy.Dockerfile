FROM ruby:2.5.3-alpine

VOLUME /tmp

ARG PROJECT_NAME=app

COPY . /app

ENV RAILS_ENV=production
ENV RAILS_PORT=3000

RUN rake db:migrate
RUN rake assets:precompile

ENTRYPOINT ["rails" server" "-e" "$RAILS_ENV" "-p" "$RAILS_PORT"]