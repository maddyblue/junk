set -ex

docker run --rm -it --privileged \
	-v /dev:/dev \
	-v $(pwd)/mjibson:/qmk_firmware/keyboards/ferris/keymaps/mjibson \
	qmkfm/qmk_firmware \
	make ferris/sweep:mjibson$1

#docker run --rm -it -v $(pwd)/mjibson:/qmk_firmware/keyboards/ferris/keymaps/mjibson qmkfm/qmk_firmware make ferris/sweep:mjibson:dfu-split-left
