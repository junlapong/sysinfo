package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
	"strings"
	"runtime"
	"time"
	"regexp"
	tm "github.com/buger/goterm"
)

type cpu_val struct {
	idle int
	total int
}

func main() {
	tm.Clear()
	num_CPU := runtime.NumCPU()
	tm.MoveCursor(1,1)
	tm.Println("Nbr of CPU(s):", num_CPU)
	file, err := os.Open("/proc/stat")
	check(err)
	reader := bufio.NewReader(file)
	reader.ReadString('\n')
	var cpu_infos_prev = make([]cpu_val, num_CPU)
	for i:=0; i<num_CPU; i++ {
		cpu_line, err := reader.ReadString('\n')
		check(err)
		idle, total := getValues(cpu_line)
		var new_cpu_val = cpu_val {idle, total}
		cpu_infos_prev = append(cpu_infos_prev, new_cpu_val)
	}
	var mem_percentage float64
	var temp int
	for {
		time.Sleep(1 * time.Second)
		file, err = os.Open("/proc/stat")
		check(err)
		reader = bufio.NewReader(file)
		reader.ReadString('\n')
		for i:=0; i<num_CPU; i++ {
			if i==0 {
				mem_percentage = getMemInfo()
				tm.MoveCursor(1, num_CPU+2)
				tm.Printf("Memory Usage: %.1f%%  ", mem_percentage)
				temp = getTemp()
				tm.MoveCursor(1, num_CPU+3)
				tm.Printf("CPU Temperature: %d\u2103 ", temp)
			}
			cpu_line, err := reader.ReadString('\n')
			check(err)
			idle, total := getValues(cpu_line)
			diff_total := total - cpu_infos_prev[i].total + 1 // to avoid div by 0
			diff_idle := idle - cpu_infos_prev[i].idle
			percentage := 100*(float64(diff_total - diff_idle)/float64(diff_total))
			cpu_infos_prev[i].total = total
			cpu_infos_prev[i].idle = idle
			tm.MoveCursor(1, i+2)
			tm.Printf("CPU %d: %.1f%%  ", i+1, percentage)
		}
		tm.Printf("\n\n\n")
		tm.Flush()
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/* CPU USAGE */
func getValues(cpu string) (int, int) {
	values := strings.Split(cpu, " ")
	total := 0
	idle := 0
	for i:=1; i<len(values); i++ {
		new_val, err := strconv.Atoi(strings.Replace(values[i], "\n", "", 1))
		check(err)
		total += new_val
		if i==4 {
			idle = new_val
		}
	}
	return idle, total
}

/* MEMORY USAGE */
func getMemInfo() float64 {
	file, err := os.Open("/proc/meminfo")
	check(err)
	rd := bufio.NewReader(file)
	total, err := rd.ReadString('\n')
	check(err)
	rgx_num, err := regexp.Compile("[0-9]+")
	check(err)
	rgx_unit, err := regexp.Compile("[GMk ]B")
	check(err)
	total_num := rgx_num.FindString(total)
	total_unit := rgx_unit.FindString(total)
	total_float, err := strconv.ParseFloat(total_num, 64)
	check(err)
	total_float, err = toBytes(total_float, total_unit)
	check(err)
	rd.ReadString('\n')
	check(err)
	available, err := rd.ReadString('\n')
	check(err)
	available_float, err := strconv.ParseFloat(rgx_num.FindString(available), 64)
	check(err)
	available_unit := rgx_unit.FindString(available)
	available_float, err = toBytes(available_float, available_unit)
	check(err)
	return (1 - (available_float/total_float))*100.0
}

func toBytes(num float64, unit string) (float64, error) {
	if unit == " B" {
		return num, nil
	}
	if unit == "kB" {
		return num*1000.0, nil
	}
	if unit == "MB" {
		return num*1000000.0, nil
	}
	if unit == "GB" {
		return num*1000000000.0, nil
	}
	return 0, fmt.Errorf("UNKNOWN %s UNIT - Add it to the toBytes function?\n", unit)
}

/* TEMPERATURE */
func getTemp() int {
	file, err := os.Open("/sys/class/thermal/thermal_zone0/temp")
	check(err)
	reader := bufio.NewReader(file)
	temp, err := reader.ReadString('\n')
	check(err)
	temp_val, err := strconv.Atoi(strings.Replace(temp, "\n", "", 1))
	check(err)
	return temp_val/1000
}
