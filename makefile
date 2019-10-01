export MSG=YES

timing_service: 

	cd ./timingService && go build && ./timingService

circuit_service:
	cd ./circuitService && go build && ./circuitService