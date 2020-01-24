package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/devfacet/gocmd"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type RecentSnapshot struct {
	Hash   string `json:"hash"`
	Height int    `json:"height"`
}

type RecentSnapshotAlignment struct {
	Height       int
	Hashes       int
	UniqueHashes int
}

func getRecentSnapshots(ip string, snapshots chan []RecentSnapshot) {
	response, err := http.Get("http://" + ip + ":9000/snapshot/recent")
	check(err)

	responseData, err := ioutil.ReadAll(response.Body)
	check(err)
	recentSnapshots := make([]RecentSnapshot, 0)
	json.Unmarshal(responseData, &recentSnapshots)

	snapshots <- recentSnapshots
}

func getHosts() []string {
	hosts := os.Getenv("HOSTS_FILE")
	if hosts == "" {
		log.Fatal("Error loading $HOSTS_FILE")
	}

	file, err := os.Open(hosts)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ips := make([]string, 0)
	for scanner.Scan() {
		ips = append(ips, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ips
}

func uniqueHashes(snapshots []RecentSnapshot) []RecentSnapshot {
	var unique []RecentSnapshot

	for _, v := range snapshots {
		skip := false
		for _, u := range unique {
			if v.Hash == u.Hash {
				skip = true
				break
			}
		}
		if !skip {
			unique = append(unique, v)
		}
	}

	return unique
}

func main() {
	flags := struct {
		Help           bool     `short:"h" long:"help" description:"Display usage" global:"true"`
		Version        bool     `short:"v" long:"version" description:"Display version"`
		VersionEx      bool     `long:"vv" description:"Display version (extended)"`
		CheckAlignment struct{} `command:"check-alignment" description:"Checks alignment on the cluster"`
		CheckBalance   struct{} `command:"check-balance" description:"Checks address balance on the cluster"`
	}{}

	gocmd.HandleFlag("CheckAlignment", func(cmd *gocmd.Cmd, args []string) error {
		ips := getHosts()
		fmt.Println("Checking alignment on nodes:", len(ips))

		recentSnapshotsChan := make(chan []RecentSnapshot, len(ips))
		defer close(recentSnapshotsChan)

		for _, ip := range ips {
			go getRecentSnapshots(ip, recentSnapshotsChan)
		}

		recentSnapshots := make([]RecentSnapshot, 0)
		for range ips {
			recentSnapshots = append(recentSnapshots, <-recentSnapshotsChan...)
		}

		allRecentSnapshots := make(map[int][]RecentSnapshot)
		uniqueRecentSnapshots := make([]RecentSnapshotAlignment, 0)

		for _, snapshot := range recentSnapshots {
			allRecentSnapshots[snapshot.Height] = append(allRecentSnapshots[snapshot.Height], snapshot)
		}

		for height, snapshots := range allRecentSnapshots {
			uniqueRecentSnapshots = append(uniqueRecentSnapshots, RecentSnapshotAlignment{ height, len(snapshots), len(uniqueHashes(snapshots)) })
		}

		sort.Slice(uniqueRecentSnapshots, func(i, j int) bool {
			return uniqueRecentSnapshots[i].Height < uniqueRecentSnapshots[j].Height
		})

		aligned := true

		for _, alignment := range uniqueRecentSnapshots {
			if alignment.UniqueHashes > 1 {
				aligned = false
			}
			fmt.Println("Height:", alignment.Height, "\t Hashes:", alignment.Hashes, "\t Unique hashes:", alignment.UniqueHashes)
		}

		if !aligned {
			log.Fatal("Cluster is misaligned!")
		} else {
			fmt.Println("Cluster is aligned!")
		}

		return nil
	})

	gocmd.HandleFlag("CheckBalance", func(cmd *gocmd.Cmd, args []string) error {
		return nil
	})

	gocmd.New(gocmd.Options{
		Name:        "cl-tools",
		Version:     "0.0.1",
		Description: "Constellation command line tools",
		Flags:       &flags,
		ConfigType:  gocmd.ConfigTypeAuto,
	})
}
