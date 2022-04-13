FROM golang:1.18 as build

WORKDIR bin

# cache dependencies
ADD go.mod go.sum ./
RUN go mod download

# build
ADD . .
RUN go build -o /main ./cmd/pipeline/signal_orchestrator

# copy artifacts to a clean image
FROM public.ecr.aws/lambda/go:1
ARG aws_access_key
ARG aws_secret_key
ENV AWS_ACCESS_KEY_ID=$aws_access_key
ENV AWS_SECRET_ACCESS_KEY=$aws_secret_key

COPY --from=build /main ${LAMBDA_TASK_ROOT}
CMD [ "/main" ]