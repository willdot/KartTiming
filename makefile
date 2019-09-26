export MSG=YES

timing_service: 

	cd ./timingService && go build && ./timingService

circuit_serice:
	cd ./circuitService && go build && ./circuitService