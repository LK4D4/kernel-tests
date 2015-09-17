package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	kernel  = flag.String("kernel", "", "path to kernel binary")
	img     = flag.String("image", "", "path to vm image")
	imgType = flag.String("image-type", "qcow2", "type of vm image")
	mount   = flag.String("mount", "", "directory which can be mounted inside vm over 9p in for name:hostpath")
)

var tmpl = `<domain type='kvm' xmlns:qemu='http://libvirt.org/schemas/domain/qemu/1.0'>
  <name>Ubuntu</name>
  <memory unit='KiB'>1048576</memory>
  <currentMemory unit='KiB'>1048576</currentMemory>
  <os>
    <type arch='x86_64' machine='pc'>hvm</type>
	<kernel>%s</kernel>
	<cmdline>root=/dev/vda1 ro console=tty0 console=ttyS0,115200 LANG=en_US.UTF-8</cmdline>
  </os>
  <features>
    <acpi/>
    <pae/>
  </features>
  <clock offset='utc'/>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>destroy</on_crash>
  <devices>
    <emulator>/usr/bin/qemu-system-x86_64</emulator>
    <disk type='file' device='disk'>
      <driver name='qemu' type='%s'/>
      <source file='%s'/>
      <target dev='vda' bus='virtio'/>
    </disk>
    <controller type='usb' index='0'/>
    <controller type='pci' index='0' model='pci-root'/>
    <input type='mouse' bus='ps2'/>
    <input type='keyboard' bus='ps2'/>
    <video>
      <model type='cirrus' vram='16384' heads='1'/>
    </video>
	<serial type='pty'>
	 <source path='/dev/pts/2'/>
	 <target port='0'/>
	</serial>
	<console type='file' tty='/dev/pts/2'>
	 <source path='/dev/pts/2'/>
	 <target port='0'/>
	</console>
    <memballoon model='none'/>
    <interface type='network'>
     <source network='default'/>
    </interface>
	%s
    <graphics type='vnc' port='-1'/>
  </devices>
</domain>`

func main() {
	flag.Parse()
	var mountCfg string
	if *mount != "" {
		parts := strings.Split(*mount, ":")
		if len(parts) != 2 {
			fmt.Printf("wrong 'mount' format: %s, expected name:hostpath form", *mount)
		}
		mountCfg = fmt.Sprintf(`<filesystem type='mount' accessmode='passthrough'>
      <source dir='%s'/>
      <target dir='%s'/>
    </filesystem>`, parts[1], parts[0])
	}
	fmt.Printf(tmpl, *kernel, *imgType, *img, mountCfg)
}
