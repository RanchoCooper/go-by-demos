package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/spf13/cast"
    "go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
    "go.etcd.io/etcd/clientv3"
)

/**
 * @author Rancho
 * @date 2021/12/12
 */

func main() {
    cli, err := clientv3.New(clientv3.Config{
        Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
        DialTimeout: 5 * time.Second,
    })
    if err != nil {
        panic(err)
    }
    defer cli.Close()

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    resp, err := cli.Put(ctx, "rancho", "cooper")
    if err != nil {
        switch err {
        case context.Canceled:
            log.Fatalf("ctx is canceled by another routine: %v", err)
        case context.DeadlineExceeded:
            log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
        case rpctypes.ErrEmptyKey:
            log.Fatalf("client-side error: %v", err)
        default:
            log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
        }
    }
    fmt.Println(cast.ToString(resp))
}
