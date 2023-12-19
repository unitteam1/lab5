FROM scratch

ARG misha_remeslo=mymainappforlab

WORKDIR /app

COPY ${misha_remeslo} /app/main

EXPOSE 8080

CMD ["./main"]