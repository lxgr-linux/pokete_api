for file in poketes attacks
do
	curl https://raw.githubusercontent.com/lxgr-linux/pokete/master/pokete_data/${file}.py > ./${file}.py
done
go run poke_api.go
