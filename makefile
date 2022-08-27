run_api:
	go run cmd/apiserver/main.go cmd/apiserver/submit_job.go
build_api:
	go cmd/apiserver/main.go cmd/apiserver/submit_job.go

run_processor:
	go run cmd/processor/main.go