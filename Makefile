all:
	cd /root/linux-stable && make -j8
	virsh create --console --autodestroy fedora.xml
