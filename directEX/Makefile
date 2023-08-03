all:
	go run produsers/main.go &
	go run consumers/consumerAll/main.go &
	go run consumers/consumerError/main.go &
	go run consumers/consumerInfo/main.go &
	go run consumers/consumerWarning/main.go &

kill:
	killall main