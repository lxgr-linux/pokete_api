#!/usr/bin/sh

echo ":: Downloading..."
for file in poketes attacks
do
	echo " -> $file.py"
	curl https://raw.githubusercontent.com/lxgr-linux/pokete/master/pokete_data/${file}.py > ./${file}.py
done
