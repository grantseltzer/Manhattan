package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/docker/engine-api/types"
)

const default_seccomp_profile = "/etc/sysconfig/docker-seccomp-profile.json"

func main() {
	configFile, configError := os.Open(default_seccomp_profile)
	if configError != nil {
		fmt.Println("Error opening default configuration: ", default_seccomp_profile)
		fmt.Println("You can specify a configuration file with -location")
		os.Exit(-1)
	}

	var SeccompProfile types.Seccomp

	jsonParser := json.NewDecoder(configFile)
	parseError := jsonParser.Decode(&SeccompProfile)
	if parseError != nil {
		fmt.Println("Error parsing configuration file")
		os.Exit(-2)
	}

	kill := flag.String("kill", "", "Respond to system call with KILL")
	trap := flag.String("trap", "", "Respond to system call with TRAP")
	errno := flag.String("errno", "", "Respond to system call with ERRNO")
	trace := flag.String("trace", "", "Respond to system call with TRACE")
	allow := flag.String("allow", "", "Respond to system call with ALLOW")

	location := flag.String("location", default_seccomp_profile,
		"Set the location for the exported seccomp profile.")
	name := flag.String("name", defaultTime(), "Set name of output file")

	flag.Parse()

	parseSysCallFlag("kill", *kill, &SeccompProfile)
	parseSysCallFlag("trap", *trap, &SeccompProfile)
	parseSysCallFlag("errno", *errno, &SeccompProfile)
	parseSysCallFlag("trace", *trace, &SeccompProfile)
	parseSysCallFlag("allow", *allow, &SeccompProfile)

	fullPath := parseLocation(*location, *name)

	fmt.Println(fullPath)
	// TODO:
	// - Write SeccompProfile back to a file
	// - Add feature to set default action
	// - Allow user to enter in a custom location for the exported profile
	//   otherwise place it in the same directory as default_seccomp_profile
	// - Go over code to make sure all errors are properly checked
	// - Write unit tests

}
