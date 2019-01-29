# sysinfo

Prints some system information (CPU usage, temperature, ..) and updates it in real time.

> Only on Linux-based OS since it uses files in `/proc/` and `/sys/`.

## Running the script

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

(Or `go build sysinfo.go && ./sysinfo`)

__Requirement:__ `goterm`

`sysinfo` uses [goterm](https://github.com/buger/goterm) to display the information. To get it:

```sh
go get github.com/buger/goterm
```

### Using Docker

If you don't have Go on your machine, you can also use this script through Docker, by pulling the image from my Docker Hub:

```sh
docker run -it jbdrvl/sysinfo:latest
```

Or you can build the image yourself with the Dockerfile:

```sh
git pull https://github.com/jbdrvl/sysinfo.git
cd sysinfo
docker build -t sysinfo . && docker run -it sysinfo
```
