api:
	go run cmd/apiserver/main.go cmd/apiserver/submit_job.go
build_api:
	go cmd/apiserver/main.go cmd/apiserver/submit_job.go

processor:
	go run cmd/processor/main.go
processor_reset:
	mv 907292c5-b00e-4133-ae81-492743177605.json 907292c5-b00e-4133-ae81-492743177605.json.pending