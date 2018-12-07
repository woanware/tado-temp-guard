package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	flags "github.com/jessevdk/go-flags"
	cron "github.com/robfig/cron"
	viper "github.com/spf13/viper"
)

// Background: http://blog.scphillips.com/posts/2017/01/the-tado-api-v2/

// ##### Structs ##############################################################

type Options struct {
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`
}

// ##### Constants ############################################################

const APP_NAME string = "tado-temp-guard (tts)"
const APP_VERSION string = "0.0.1"

const URL_HOME_ID string = "https://my.tado.com/api/v2/me?username=%s&password=%s"
const URL_ZONES string = "https://my.tado.com/api/v2/homes/%d/zones?username=%s&password=%s"
const URL_ZONE string = "https://my.tado.com/api/v2/homes/%d/zones/%d/state?username=%s&password=%s"
const URL_OVERLAY string = "https://my.tado.com/api/v2/homes/%d/zones/%d/overlay?username=%s&password=%s"

// ##### Variables ############################################################

var (
	options      Options
	configReader *viper.Viper
	config       *Config
	cronner      *cron.Cron
	home         Home
	zones        Zones
)

// ##### Methods ##############################################################

func main() {

	fmt.Println(fmt.Sprintf("\n%s v%s - woanware\n", APP_NAME, APP_VERSION))

	parseCommandLine()
	initialiseConfiguration()
	config = parseConfiguration()
	if config == nil {
		return
	}

	err := getHome()
	if err != nil {
		fmt.Printf("Error retrieving home: %v", err)
		return
	}

	err = getZones()
	if err != nil {
		fmt.Printf("Error retrieving zones: %v", err)
		return
	}

	cronner = cron.New()
	cronner.AddFunc(fmt.Sprintf("@every %dm", config.IntervalMinutes), checkTemperature)
	cronner.Start()

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

//
func parseCommandLine() {

	var parser = flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}

//
func getHome() error {

	resp, err := http.Get(fmt.Sprintf(URL_HOME_ID, config.UserName, config.Password))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &home)
	if err != nil {
		return err
	}

	if options.Verbose == true {
		prettyPrint, err := json.MarshalIndent(home, "", "  ")
		if err != nil {
			fmt.Printf("Error displaying JSON (Home): %v", err)
		}
		fmt.Printf("Home:\n%s\n\n", (string(prettyPrint)))
	}

	return nil
}

//
func getZones() error {

	if len(home.Homes) > 1 {
		return errors.New("Too many homes!")
	}

	resp, err := http.Get(fmt.Sprintf(URL_ZONES, home.Homes[0].ID, config.UserName, config.Password))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &zones)
	if err != nil {
		return err
	}

	if options.Verbose == true {
		prettyPrint, err := json.MarshalIndent(zones, "", "  ")
		if err != nil {
			fmt.Printf("Error displaying JSON (Zones): %v", err)
		}
		fmt.Printf("Zones:\n%s\n\n", (string(prettyPrint)))
	}

	return nil
}

//
func checkTemperature() {

	for _, z := range zones {
		resp, err := http.Get(fmt.Sprintf(URL_ZONE, home.Homes[0].ID, z.ID, config.UserName, config.Password))
		if err != nil {
			fmt.Printf("Error retrieving zone data: %v", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		var zone Zone
		err = json.Unmarshal(body, &zone)
		if err != nil {
			fmt.Printf("Error reading zone data: %v", err)
			continue
		}

		if options.Verbose == true {
			prettyPrint, err := json.MarshalIndent(zone, "", "  ")
			if err != nil {
				fmt.Printf("Error displaying JSON (Zone): %v", err)
			}
			fmt.Printf("Zone:\n%s\n\n", (string(prettyPrint)))
		}

		if strings.ToUpper(zone.Setting.Type) != "HEATING" {
			continue
		}

		if strings.ToUpper(zone.Setting.Power) != "ON" {
			continue
		}

		if zone.Setting.Temperature.Celsius <= config.MaxTemperatureCelsius {
			continue
		}

		fmt.Printf("Temperature has been set TOOOOOO HIGH: %2.1f", zone.Setting.Temperature.Celsius)

		setTemperature(z.ID)
	}

}

//
func setTemperature(zoneID int) {

	temp := fmt.Sprintf(JSON_OVERLAY, config.MaxTemperatureCelsius)

	client := &http.Client{}
	request, err := http.NewRequest("PUT", fmt.Sprintf(URL_OVERLAY, home.Homes[0].ID, zoneID, config.UserName, config.Password), strings.NewReader(temp))
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error retrieving overlay data: %v", err)
		return
	}

	if response.StatusCode != 200 {
		fmt.Printf("Error retrieving overlay data: %v", err)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	var overlay Overlay
	err = json.Unmarshal(body, &overlay)
	if err != nil {
		fmt.Printf("Error reading overlay data: %v", err)
		return
	}

	if options.Verbose == true {
		prettyPrint, err := json.MarshalIndent(overlay, "", "  ")
		if err != nil {
			fmt.Printf("Error displaying JSON (OVerlay): %v", err)
		}
		fmt.Printf("Overlay:\n%s\n\n", (string(prettyPrint)))
	}

	if overlay.Setting.Temperature.Celsius != config.MaxTemperatureCelsius {
		fmt.Println("Unable to set temperature")
	} else {
		fmt.Printf("Temperature set to: %2.1f", overlay.Setting.Temperature.Celsius)
	}
}

//
func reloadConfig() {

	config = parseConfiguration()
}
