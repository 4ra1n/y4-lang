package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/google/go-github/v57/github"
	"golang.org/x/net/proxy"
	"golang.org/x/oauth2"
)

const socksProxy = "127.0.0.1:10808"

func main() {
	data, err := os.ReadFile("token.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	token := string(data)

	owner := "4ra1n"
	repo := "y4-lang"
	workflowName := "y4-lang"
	workflowFileName := "y4-lang.yml"
	artifactName := "build-artifact"

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	dialer, err := proxy.SOCKS5("tcp",
		socksProxy, nil, proxy.Direct)
	if err != nil {
		fmt.Println("Error creating dialer:", err)
		return
	}
	t := tc.Transport.(*oauth2.Transport)
	t.Base = &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
	}
	tc.Transport = t

	client := github.NewClient(tc)

	workflows, _, err := client.Actions.ListWorkflows(ctx, owner, repo, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	var workflowID int64
	for _, workflow := range workflows.Workflows {
		if *workflow.Name == workflowName {
			workflowID = *workflow.ID
			break
		}
	}
	if workflowID == 0 {
		fmt.Println("workflow not found")
		return
	}
	_, err = client.Actions.CreateWorkflowDispatchEventByID(
		ctx, owner, repo, workflowID, github.CreateWorkflowDispatchEventRequest{
			Ref: "master",
		})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Println("workflow dispatched successfully")

	var runID int64
	for {
		runs, _, err := client.Actions.ListWorkflowRunsByFileName(
			ctx, owner, repo, workflowFileName, &github.ListWorkflowRunsOptions{
				Status: "completed",
			})
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(runs.WorkflowRuns) > 0 {
			runID = *runs.WorkflowRuns[0].ID
			break
		}
		fmt.Println("workflow not complete")
		time.Sleep(5 * time.Second)
	}

	artifacts, _, err := client.Actions.ListWorkflowRunArtifacts(
		ctx, owner, repo, runID, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	var artifactID int64
	for _, artifact := range artifacts.Artifacts {
		if *artifact.Name == artifactName {
			artifactID = *artifact.ID
			break
		}
	}

	if artifactID == 0 {
		fmt.Println("Artifact not found")
		return
	}

	artifactUrl, _, err := client.Actions.DownloadArtifact(
		ctx, owner, repo, artifactID, 10)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Println("get build file success")
	fileName := "build.zip"
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := os.Stat(fileName); err == nil {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		fmt.Printf("file %s successfully deleted.\n", fileName)
	} else if os.IsNotExist(err) {
		fmt.Printf("file %s does not exist\n", fileName)
	} else {
		fmt.Printf("error: %v\n", err)
	}
	err = downloadFile(fileName, artifactUrl.String())
	fmt.Println("download build file success")
}

func downloadFile(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
