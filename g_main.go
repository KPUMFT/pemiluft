package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var dryRun bool
	dryRun = true

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "dry-run",
				Usage:       "run target without actually execute it",
				Value:       false,
				Destination: &dryRun,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "allUp",
				Usage: "create all target",
				Action: func(c *cli.Context) error {
					return AllUp(dryRun)
				},
			},
			{
				Name:  "allDown",
				Usage: "drop all target",
				Action: func(c *cli.Context) error {
					return AllDown(dryRun)
				},
			},
			{
				Name:  "daftar_pemilih",
				Usage: "manage table daftar_pemilih",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "create table daftar_pemilih",
						Action: func(c *cli.Context) error {
							return Target001Up(dryRun)
						},
					},
					{
						Name:  "down",
						Usage: "drop table daftar_pemilih",
						Action: func(c *cli.Context) error {
							return Target001Down(dryRun)
						},
					},
				},
			},
			{
				Name:  "daftar_pilihan",
				Usage: "manage table daftar_pilihan",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "create table daftar_pilihan",
						Action: func(c *cli.Context) error {
							return Target002Up(dryRun)
						},
					},
					{
						Name:  "down",
						Usage: "drop table daftar_pilihan",
						Action: func(c *cli.Context) error {
							return Target002Down(dryRun)
						},
					},
				},
			},
			{
				Name:  "ref_calon_dpm",
				Usage: "manage table ref_calon_dpm",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "create table ref_calon_dpm",
						Action: func(c *cli.Context) error {
							return Target003Up(dryRun)
						},
					},
					{
						Name:  "down",
						Usage: "drop table ref_calon_dpm",
						Action: func(c *cli.Context) error {
							return Target003Down(dryRun)
						},
					},
				},
			},
			{
				Name:  "ref_calon_gubernur",
				Usage: "manage table ref_calon_gubernur",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "create table ref_calon_gubernur",
						Action: func(c *cli.Context) error {
							return Target004Up(dryRun)
						},
					},
					{
						Name:  "down",
						Usage: "drop table ref_calon_gubernur",
						Action: func(c *cli.Context) error {
							return Target004Down(dryRun)
						},
					},
				},
			},
			{
				Name:  "ref_calon_himatro",
				Usage: "manage table ref_calon_himatro",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "create table ref_calon_himatro",
						Action: func(c *cli.Context) error {
							return Target005Up(dryRun)
						},
					},
					{
						Name:  "down",
						Usage: "drop table ref_calon_himatro",
						Action: func(c *cli.Context) error {
							return Target005Down(dryRun)
						},
					},
				},
			},
			{
				Name:  "ref_calon_hmti",
				Usage: "manage table ref_calon_hmti",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "create table ref_calon_hmti",
						Action: func(c *cli.Context) error {
							return Target006Up(dryRun)
						},
					},
					{
						Name:  "down",
						Usage: "drop table ref_calon_hmti",
						Action: func(c *cli.Context) error {
							return Target006Down(dryRun)
						},
					},
				},
			},
			{
				Name:  "ref_dapil",
				Usage: "manage table ref_dapil",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "create table ref_dapil",
						Action: func(c *cli.Context) error {
							return Target007Up(dryRun)
						},
					},
					{
						Name:  "down",
						Usage: "drop table ref_dapil",
						Action: func(c *cli.Context) error {
							return Target007Down(dryRun)
						},
					},
				},
			},
			{
				Name:  "vw_hasil_dpm",
				Usage: "manage view vw_hasil_dpm",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "create view vw_hasil_dpm also export data if specified",
						Action: func(c *cli.Context) error {
							return Target008Up(dryRun)
						},
					},
					{
						Name:  "down",
						Usage: "drop view vw_hasil_dpm",
						Action: func(c *cli.Context) error {
							return Target008Down(dryRun)
						},
					},
				},
			},
			{
				Name:  "vw_hasil_gubernur",
				Usage: "manage view vw_hasil_gubernur",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "create view vw_hasil_gubernur also export data if specified",
						Action: func(c *cli.Context) error {
							return Target009Up(dryRun)
						},
					},
					{
						Name:  "down",
						Usage: "drop view vw_hasil_gubernur",
						Action: func(c *cli.Context) error {
							return Target009Down(dryRun)
						},
					},
				},
			},
			{
				Name:  "vw_hasil_hmp",
				Usage: "manage view vw_hasil_hmp",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "create view vw_hasil_hmp also export data if specified",
						Action: func(c *cli.Context) error {
							return Target010Up(dryRun)
						},
					},
					{
						Name:  "down",
						Usage: "drop view vw_hasil_hmp",
						Action: func(c *cli.Context) error {
							return Target010Down(dryRun)
						},
					},
				},
			},
			{
				Name:  "vw_pilihan",
				Usage: "manage view vw_pilihan",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "create view vw_pilihan also export data if specified",
						Action: func(c *cli.Context) error {
							return Target011Up(dryRun)
						},
					},
					{
						Name:  "down",
						Usage: "drop view vw_pilihan",
						Action: func(c *cli.Context) error {
							return Target011Down(dryRun)
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		exitWithError(err)
	}
}
