package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/getlantern/systray"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

var mCPU *systray.MenuItem
var mRAM *systray.MenuItem
var mDisk *systray.MenuItem
var mNetwork *systray.MenuItem
var mUptime *systray.MenuItem

var prevNetIO net.IOCountersStat

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	iconData, err := ioutil.ReadFile("pumpkin.ico")
	if err != nil {
		log.Fatalf("Icon could not be loaded: %v", err)
	}
	systray.SetIcon(iconData)

	lang := getSystemLanguage()
	if strings.HasPrefix(lang, "tr") {
		systray.SetTitle("Sistem Kullanım Bilgileri")
		systray.SetTooltip("Sistem bilgilerini gösterir")
	} else {
		systray.SetTitle("System Usage Info")
		systray.SetTooltip("Displays system information")
	}

	mCPU = systray.AddMenuItem("", "")
	mRAM = systray.AddMenuItem("", "")
	mDisk = systray.AddMenuItem("", "")
	mNetwork = systray.AddMenuItem("", "")
	mUptime = systray.AddMenuItem("", "")
	mQuit := systray.AddMenuItem("Exit", "Exit the application")

	go func() {
		netIO, err := net.IOCounters(false)
		if err != nil || len(netIO) == 0 {
			log.Fatalf("Failed to get network I/O counters: %v", err)
		}
		prevNetIO = netIO[0]

		for {
			updateUsage(lang)
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func updateUsage(lang string) {
	cpuPercent, _ := cpu.Percent(0, false)
	v, _ := mem.VirtualMemory()
	diskUsage, _ := disk.Usage("/")
	netIO, _ := net.IOCounters(false)
	uptime := getUptime()

	var netSent float64
	var netRecv float64
	if len(netIO) > 0 {
		netSent = float64(netIO[0].BytesSent - prevNetIO.BytesSent)
		netRecv = float64(netIO[0].BytesRecv - prevNetIO.BytesRecv)
		prevNetIO = netIO[0]
	}

	cpuText := ""
	ramText := ""
	diskText := ""
	networkText := ""
	uptimeText := ""

	if strings.HasPrefix(lang, "tr") {
		cpuText = fmt.Sprintf("CPU: %.2f%%", cpuPercent[0])
		ramText = fmt.Sprintf("RAM: %.2f%%", v.UsedPercent)
		diskText = fmt.Sprintf("Disk: %.2f GB (%.2f%%)", float64(diskUsage.Used)/1024/1024/1024, diskUsage.UsedPercent)
		networkText = fmt.Sprintf("Network: %.2f MB/s (%.2f GB)", (netSent+netRecv)/1024/1024, (netSent+netRecv)/1024/1024/1024)
		uptimeText = fmt.Sprintf("Çalışma Süresi: %s", uptime)
	} else {
		cpuText = fmt.Sprintf("CPU: %.2f%%", cpuPercent[0])
		ramText = fmt.Sprintf("RAM: %.2f%%", v.UsedPercent)
		diskText = fmt.Sprintf("Disk: %.2f GB (%.2f%%)", float64(diskUsage.Used)/1024/1024/1024, diskUsage.UsedPercent)
		networkText = fmt.Sprintf("Network: %.2f MB/s (%.2f GB)", (netSent+netRecv)/1024/1024, (netSent+netRecv)/1024/1024/1024)
		uptimeText = fmt.Sprintf("Uptime: %s", uptime)
	}

	mCPU.SetTitle(cpuText)
	mRAM.SetTitle(ramText)
	mDisk.SetTitle(diskText)
	mNetwork.SetTitle(networkText)
	mUptime.SetTitle(uptimeText)
}

func onExit() {
	playExitSound("song.wav")
}

func playExitSound(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open sound file: %v", err)
	}
	defer f.Close()

	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatalf("Failed to decode sound file: %v", err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)

	duration := float64(streamer.Len()) / float64(format.SampleRate)
	time.Sleep(time.Duration(duration * float64(time.Second)))
}

func getSystemLanguage() string {
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = os.Getenv("LC_ALL")
	}
	if lang == "" {
		lang = "en"
	}
	return lang
}

func getUptime() string {
	if isWindows() {
		cmd := exec.Command("powershell", "-Command", "(Get-CimInstance -ClassName Win32_OperatingSystem).LastBootUpTime")
		output, err := cmd.Output()
		if err != nil {
			return "N/A"
		}
		return strings.TrimSpace(string(output))
	}

	data, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return "N/A"
	}
	return strings.TrimSpace(string(data))
}

func isWindows() bool {
	return strings.Contains(strings.ToLower(os.Getenv("OS")), "windows")
}
