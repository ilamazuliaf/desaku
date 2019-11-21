run:
	go build
	rm nohup.out
	nohup ./desaku &

pull:
	git pull origin master