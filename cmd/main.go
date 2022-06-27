package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"tutero_assignment/pkg/src/graph"
	"tutero_assignment/pkg/step"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "find_node",
		Action: func(*cli.Context) error {
			if _, err := drive(); err != nil {
				return err
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func drive() (int, error) {
	g, err := graph.Random(func(opts *graph.RandomOptions) {
		opts.MinPerRank = 15
		opts.MaxPerRank = 20
		opts.MinRanks = 8
		opts.MaxRanks = 10
		opts.Percent = rand.Float32()
	})
	g.AddNode("A") // Ensure at-least one node
	if err != nil {
		return 0, err
	}

	stepper := step.New()
	nodes := g.Nodes()
	sentinel := nodes[rand.Intn(len(nodes))]

	steps := 0
	for {
		submitted, err := stepper.Step(*g)
		if err != nil {
			return 0, err
		}
		if submitted == sentinel {
			fmt.Printf("Found target node in %d steps!\n", steps)
			return steps, nil
		}
		truncated := false
		for _, child := range g.Children(submitted) {
			// found in children
			if child == sentinel {
				truncated = true
				parents := g.Parents(submitted)
				for _, parent := range parents {
					g.RemoveNode(parent)
				}
				break
			}
		}
		if !truncated {
			for _, parent := range g.Parents(submitted) {
				// found in parents
				if parent == sentinel {
					truncated = true
					children := g.Children(submitted)
					for _, child := range children {
						g.RemoveNode(child)
					}
					break
				}
			}
		}
		g.RemoveNode(submitted)
		steps++
	}
}
