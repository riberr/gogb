package gameboy

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

/*
 uses blargg's cpu_instrs test roms and gameboy-logs
 https://gbdev.gg8.se/files/roms/blargg-gb-tests/
 https://github.com/wheremyfoodat/Gameboy-logs/tree/master
*/

func TestCpuOutputBlargg01(t *testing.T) {
	testRom(
		"../third_party/gb-test-roms/cpu_instrs/individual/",
		"01-special.gb",
		"../third_party/gameboy-doctor/truth/zipped/cpu_instrs/1.log",
		false,
		false,
		t,
	)
}

/*
Fails at "timer doesn't work". Probably need to have proper timing

// VALUE AT FF05 SHOULD BE 04, not 00. 04 is interrupt bit 2 TIMER
*/
func TestCpuOutputBlargg02(t *testing.T) {
	testRom(
		"../third_party/gb-test-roms/cpu_instrs/individual/",
		"02-interrupts.gb",
		// https://github.com/robert/gameboy-doctor/pull/11
		"../third_party/Blargg2LYStubbed/EpicLogReformat.txt",
		false,
		false,
		t,
	)
}

func TestCpuOutputBlargg03(t *testing.T) {
	testRom(
		"../third_party/gb-test-roms/cpu_instrs/individual/",
		"03-op sp,hl.gb",
		"../third_party/gameboy-doctor/truth/zipped/cpu_instrs/3.log",
		true,
		false,
		t,
	)
}

func TestCpuOutputBlargg04(t *testing.T) {
	testRom(
		"../third_party/gb-test-roms/cpu_instrs/individual/",
		"04-op r,imm.gb",
		"../third_party/gameboy-doctor/truth/zipped/cpu_instrs/4.log",
		false,
		false,
		t,
	)
}

func TestCpuOutputBlargg05(t *testing.T) {
	testRom(
		"../third_party/gb-test-roms/cpu_instrs/individual/",
		"05-op rp.gb",
		"../third_party/gameboy-doctor/truth/zipped/cpu_instrs/5.log",
		false,
		false,
		t,
	)
}

func TestCpuOutputBlargg06(t *testing.T) {
	testRom(
		"../third_party/gb-test-roms/cpu_instrs/individual/",
		"06-ld r,r.gb",
		"../third_party/gameboy-doctor/truth/zipped/cpu_instrs/6.log",
		false,
		false,
		t,
	)
}

func TestCpuOutputBlargg07(t *testing.T) {
	testRom(
		"../third_party/gb-test-roms/cpu_instrs/individual/",
		"07-jr,jp,call,ret,rst.gb",
		"../third_party/gameboy-doctor/truth/zipped/cpu_instrs/7.log",
		false,
		true,
		t,
	)
}

func TestCpuOutputBlargg08(t *testing.T) {
	testRom(
		"../third_party/gb-test-roms/cpu_instrs/individual/",
		"08-misc instrs.gb",
		"../third_party/gameboy-doctor/truth/zipped/cpu_instrs/8.log",
		false,
		false,
		t,
	)
}

func TestCpuOutputBlargg09(t *testing.T) {
	testRom(
		"../third_party/gb-test-roms/cpu_instrs/individual/",
		"09-op r,r.gb",
		"../third_party/gameboy-doctor/truth/zipped/cpu_instrs/9.log",
		false,
		false,
		t,
	)
}

func TestCpuOutputBlargg10(t *testing.T) {
	testRom(
		"../third_party/gb-test-roms/cpu_instrs/individual/",
		"10-bit ops.gb",
		"../third_party/gameboy-doctor/truth/zipped/cpu_instrs/10.log",
		false,
		false,
		t,
	)
}

func TestCpuOutputBlargg11(t *testing.T) {
	testRom(
		"../third_party/gb-test-roms/cpu_instrs/individual/",
		"11-op a,(hl).gb",
		"../third_party/gameboy-doctor/truth/zipped/cpu_instrs/11.log",
		false,
		false,
		t,
	)
}

func testRom(
	romPath string,
	romName string,
	logPath string,
	debug bool,
	ignoreLogResult bool,
	t *testing.T,
) {
	// SETUP
	logFile, err := os.Open(logPath)
	if err != nil {
		t.Fatalf("Error opening file: %v", err)
	}
	defer logFile.Close()

	log := bufio.NewReader(logFile)

	gb := New(false)

	if !gb.Bus.LoadCart(romPath, romName) {
		t.Fatalf("error loading rom")
	}

	nrOfLines, err := lineCounter(logPath)
	nrOfLines++ // first line is line 1
	if err != nil {
		t.Fatal(err)
	}

	// RUN TEST
	i := 1
	// loop until we reach end of log file

	//var output string

	// step cpu until we reach the next fetch
	for {
		gb.Step()
		if gb.Cpu.NewLog {
			gb.Cpu.NewLog = false

			/*
				if gb.Cpu.GetState() == cpu.FetchOpCode && gb.Cpu.Cycle == 0 {
					//output = gb.Cpu.GetInternalString()
					break
				}
			*/

			logLine, _, err := log.ReadLine()
			if err != nil {
				if err.Error() == "EOF" {
					println(i)
					fmt.Printf("\n")
					break
				}

				fmt.Println("Error reading line:", err)
				return
			}

			if debug {
				println(strings.Trim(gb.Cpu.Log, "\n"))
			}
			if !ignoreLogResult {
				if strings.Trim(string(logLine), "\n") != strings.Trim(gb.Cpu.Log, "\n") {
					t.Fatalf("%v/%v: not equal!\ngot: \n%v\nwant: \n%v", i, nrOfLines, gb.Cpu.Log, string(logLine))
				}
			}
			i++
		}

		/*
			// print serial output
			res := gb.SerialLink.GetLog()
			if res != "" {
				println(strings.Trim(res, "\n"))
			}
		*/

	}

	// ASSERT
	res := gb.SerialLink.GetLog()
	if strings.Trim(res[len(res)-7:], "\n") != "Passe" && !strings.Contains(res, "Passed") && !strings.Contains(res, "Passe") {
		t.Fatalf("%v did not return 'Passed'\n", romName)
	}
}

func lineCounter(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

// https://github.com/robert/gameboy-doctor/pull/11
func GenerateReformatedLogFor02Interrupt(t *testing.T) {
	// SETUP
	logFile, err := os.Open("../../third_party/Blargg2LYStubbed/EpicLog.txt")
	if err != nil {
		t.Fatalf("Error opening file: %v", err)
	}
	defer logFile.Close()
	log := bufio.NewReader(logFile)

	output, err := os.Create("../../third_party/Blargg2LYStubbed/EpicLogReformat.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer output.Close()

	for {
		logLine, _, err := log.ReadLine()
		if err != nil {

			if err.Error() == "EOF" {
				fmt.Printf("\n")
				break
			}

			fmt.Println("Error reading line:", err)
			return
		}

		//println(string(logLine))
		out := fmt.Sprintf("A:%s F:%s B:%s C:%s D:%s E:%s H:%s L:%s SP:%s PC:%s PCMEM:%s,%s,%s,%s\n",
			logLine[3:5], logLine[9:11], logLine[15:17], logLine[21:23], logLine[27:29], logLine[33:35], logLine[39:41], logLine[45:47], logLine[52:56], logLine[64:68],
			logLine[70:72], logLine[73:75], logLine[76:78], logLine[79:81],
		)
		output.WriteString(out)
	}
}
