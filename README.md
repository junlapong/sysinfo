# sysinfo

Prints some system information (CPU usage, temperature, ..) and updates it in real time.

> Only on Linux-based OS since it uses files in `/proc/` and `/sys/`.

To run it directly on your machine:

```sh
go run sysinfo.go
```

```
Nbr of CPU(s): 4
CPU 1: 4.9%   
CPU 2: 2.0%  
CPU 3: 14.6%  
CPU 4: 3.9%   
Memory Usage: 30.6%  
CPU Temperature: 44℃
```

> Ctrl+C to exit the program.

### Docker

If you don't have Go on your machine, you can also use it with Docker, by using the image from my Docker Hub:

```sh
docker run -it jbdrvl/sysinfo:latest
```

Or you can build the image yourself with the Dockerfile:

```sh
git pull https://github.com/jbdrvl/sysinfo.git sysinfo
cd ./sysinfo
docker build -t sysinfo . && docker run -it sysinfo
```

### Requirement: `goterm`

`sysinfo` uses [goterm](https://github.com/buger/goterm) to display the information. To get it:

```sh
go get github.com/buger/goterm
```
