# sysinfo

Prints some system information (CPU usage, temperature, ..) and updates it in real time.

To run it:

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

> Ctrl+C to exit it.

### Requirement: `goterm`

`sysinfo` uses [goterm](https://github.com/buger/goterm) to display the information. To get it:

```sh
go get github.com/buger/goterm
```
