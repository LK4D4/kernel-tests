## What is it

* kernel config for building static (without modules) kernel, which will be able to do
basic container manipulations.
* tool for generating libvirt config for starting ubuntu vm with direct-boot kernel

## Example commands:

Build kernel with `linux-next.config`:
```
cp linux-next.config ~/project/linux-next/.config
make CC="ccache gcc" -j4"
```
Create image:
```
sudo virt-builder -m 2048 --size 20G --format qcow2 -o ubuntu.qcow2 ubuntu-14.04
```
Generate config:
```
go run genconfig.go -kernel ~/project/linux-next/arch/x86/boot/bzImage -image ./ubuntu.qcow2 -mount kernel:~/project/linux-next > ubuntu.xml

```
Run vm:
```
sudo virsh create --console --autodestroy ubuntu.xml
```

In vm you can run `setup_trace.sh` script for setting up kernel tracing for current
process.
