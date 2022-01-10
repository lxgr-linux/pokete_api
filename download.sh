#!/usr/bin/sh

echo ":: Downloading..."
for file in poketes attacks types
do
	echo " -> pokete_$file.py"
	curl https://raw.githubusercontent.com/lxgr-linux/pokete/master/pokete_data/${file}.py > ./pokete_${file}.py
done
