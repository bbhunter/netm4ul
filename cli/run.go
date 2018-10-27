package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/netm4ul/netm4ul/core/api"
	"github.com/netm4ul/netm4ul/core/communication"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// var CLIprojectName string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run scan on the defined target",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		createSessionBase()
		cliSession.Config.Algorithm.Mode = cliMode
		if cliProjectName != "" {
			createProject(cliProjectName)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("Too few arguments ! Expecting target.")
		}

		targets, err := parseTargets(args)
		if err != nil {
			log.Errorf("Error while parsing targets : %s\n", err.Error())
		}

		log.Debugf("targets : %+v", targets)
		log.Debugf("CLIModules : %+v", cliModules)
		log.Debugf("Modules : %+v", cliSession.Config.Modules)
		log.Debugf("CLIMode : %+v", cliMode)
		log.Debugf("Mode : %+v", cliSession.Config.Algorithm.Mode)

		// if len(CLImodules) > 0 {
		// 	mods, err := parseModules(CLImodules, CLISession)
		// 	if err != nil {
		// 		log.Errorf(err.Error())
		// 	}
		// 	addModules(mods, CLISession)
		// }

		if len(cliModules) > 0 {
			fmt.Println("Running only specified module(s) :", cliModules)
			runSpecifiedModules(targets, cliModules)
			return
		}

		runModules(targets)
	},
}

func createProject(project string) {
	url := "http://" + cliSession.Config.Server.IP + ":" + strconv.FormatUint(uint64(cliSession.Config.API.Port), 10) +
		"/api/v1/projects"
	jsonInput, err := json.Marshal(project)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonInput))
	if err != nil {
		log.Fatal(err)
	}

	res, err := json.Marshal(resp.Body)
	if err != nil {
		fmt.Println("Received invalid json !", err)
	}
	fmt.Println(res)
}

func runModules(targets []communication.Input) {
	url := "http://" + cliSession.Config.Server.IP + ":" + strconv.FormatUint(uint64(cliSession.Config.API.Port), 10) +
		"/api/v1/projects/" +
		cliSession.Config.Project.Name +
		"/run"

	jsonInput, err := json.Marshal(targets)
	if err != nil {
		fmt.Printf("Error : %s\n", err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonInput))
	if err != nil {
		fmt.Printf("Error : %s\n", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Session-Token", cliSession.Config.API.Token)

	var res api.Result
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error : %s\n", err.Error())
		return
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&res)
	if err != nil {
		fmt.Println("Received invalid json !", err)
		return
	}
	defer req.Body.Close()

	if res.Status == "error" {
		fmt.Printf("Could not execute command : %s\n", res.Message)
		return
	}

	fmt.Printf("Command sent ! (%+v)\n", res)
}

func runSpecifiedModules(targets []communication.Input, modules []string) {

	url := "http://" + cliSession.Config.Server.IP + ":" + strconv.FormatUint(uint64(cliSession.Config.API.Port), 10) +
		"/api/v1/projects/" +
		cliSession.Config.Project.Name +
		"/run"

	jsonInput, err := json.Marshal(targets)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range modules {
		r, err := http.Post(url+"/"+m, "application/json", bytes.NewBuffer(jsonInput))
		if err != nil {
			log.Fatal(err)

		}
		var res api.Result
		err = json.NewDecoder(r.Body).Decode(&res)
		if err != nil {
			fmt.Println("Received invalid json !", err)
		}
		if res.Code == api.CodeOK {
			fmt.Println(res.Message, res.Data)
		} else {
			fmt.Printf("An error occurred : [%d], %s\n", res.Code, res.Message)
		}
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().StringArrayVar(&cliModules, "modules", []string{}, "Set custom module(s)")
	runCmd.PersistentFlags().StringVarP(&cliMode, "mode", "m", DefaultMode, "Use predefined mode")
	runCmd.PersistentFlags().StringVarP(&cliProjectName, "project", "p", "", "Set project name")
}
