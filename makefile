api:
	go run cmd/apiserver/main.go cmd/apiserver/submit_job.go cmd/apiserver/get_job.go
build_api:
	go cmd/apiserver/main.go cmd/apiserver/submit_job.go cmd/apiserver/get_job.go

processor:
	go run cmd/processor/main.go

cleaner:
	go run cmd/cleaner/main.go