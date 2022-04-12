FROM golang:1.18 as build

WORKDIR bin

# cache dependencies
ADD go.mod go.sum ./
RUN go mod download

# build
ADD . .
RUN go build -o /main ./cmd/collection_agent_github

# copy artifacts to a clean image
FROM public.ecr.aws/lambda/go:1
COPY --from=build /main ${LAMBDA_TASK_ROOT}
CMD [ "/main" ]