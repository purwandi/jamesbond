builder: 
	rm -rf build wrapper app
	garble -tiny -literals -seed=random build -o wrapper cmder/*
	go build -o app rest/*

run:
	./wrapper ./app