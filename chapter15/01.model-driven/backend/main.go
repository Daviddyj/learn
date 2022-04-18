package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/client"
	clientv3 "go.etcd.io/etcd/client/v3"
	"learn/chapter15/01.model-driven/models"
	"learn/pkg/dockertool"
	"log"
)

func main() {
	backendName := "daiyijie"

	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"http://localhost:2379"}})
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	dockercli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		watcher := cli.Watch(ctx, backendName)
		for respData := range watcher {
			evs := respData.Events
			whetherBreak := false
			for _, ev := range evs {
				rawBackend := ev.Kv.Value

				backend := &models.RankServiceBackend{}
				json.Unmarshal(rawBackend, backend) //todo handle

				ids, _ := dockertool.List(ctx, dockercli, map[string]string{"nginx": "latest"})
				backend.Status.RunningCount = len(ids)

				if backend.Expected.Count != backend.Status.RunningCount {
					//todo调用docke创建新的容器实例 或删除一些，并更新状态
					ip, err := dockertool.Run(ctx, dockercli, map[string]string{"nginx": "latest"}, "nginx:latest", nil)
					if err != nil {
						//todo handle err
					} else {
						backend.Status.RunningCount++
						backend.Status.InstanceIPs = append(backend.Status.InstanceIPs, ip)
						rawData, _ := json.Marshal(backend)
						cli.Put(ctx, backendName, string(rawData))
					}
				} else {
					fmt.Println("已经满足预期:", backend.Expected.Count)
				}
			}
			if whetherBreak {
				break
			}
		}
	}()
	<-ctx.Done()
}
