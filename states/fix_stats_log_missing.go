package states

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	datapbv2 "github.com/milvus-io/birdwatcher/proto/v2.2/datapb"
	"github.com/milvus-io/birdwatcher/states/etcd/common"
	"github.com/spf13/cobra"
	clientv3 "go.etcd.io/etcd/client/v3"
	"path"
)

func FixStatsLogMissing(cli clientv3.KV, basePath string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fix_stats_log",
		Short: "fix stats log missing",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			get, err := cmd.Flags().GetBool("get")
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			getAndUpdate, err := cmd.Flags().GetBool("getAndUpdate")
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			cnt := 0
			err = common.ForeachSegmentsV2(ctx, cli, basePath, func(segment datapbv2.SegmentInfo) {
				for fidx, fBinlogs := range segment.GetStatslogs() {
					key := path.Join(basePath, "datacoord-meta/statslog",
						fmt.Sprintf("%d/%d/%d/%d", segment.CollectionID, segment.PartitionID, segment.ID, fBinlogs.FieldID))

					var newBinlogs []*datapbv2.Binlog
					var binlog *datapbv2.Binlog
					cntLogID0 := 0
					for _, binlog = range fBinlogs.Binlogs {
						if binlog.LogID == 0 {
							if get || getAndUpdate {
								resp, err := cli.Get(ctx, key)
								if err != nil {
									fmt.Printf("failed to delete statslogs from etcd for segment %d, err: %s\n", segment.ID, err.Error())
									return
								}
								fmt.Println("key:", key, " has already exists: ", resp.Count == 1)
							}
							cntLogID0++
						} else {
							newBinlogs = append(newBinlogs, &datapbv2.Binlog{
								EntriesNum:    binlog.EntriesNum,
								TimestampFrom: binlog.TimestampFrom,
								TimestampTo:   binlog.TimestampTo,
								LogSize:       binlog.LogSize,
								LogID:         binlog.LogID,
								LogPath:       binlog.LogPath,
							})
						}
					}

					segment.GetStatslogs()[fidx].Binlogs = newBinlogs
					if len(newBinlogs) > 0 && cntLogID0 > 0 && getAndUpdate {
						bs, err := proto.Marshal(&segment)
						if err != nil {
							fmt.Printf("failed to marshal segment %d, err: %s\n", segment.ID, err.Error())
							return
						}

						_, err = cli.Put(ctx, key, string(bs))
						if err != nil {
							fmt.Printf("failed to put stats log of segment: %s, err: %s\n", key, err.Error())
							return
						}
						fmt.Println("key:", key, " put new segment done")
					}
				}
			})

			if err != nil {
				fmt.Println("failed to list segments", err.Error())
			}

			fmt.Println("total invalided path:", cnt)
		},
	}

	cmd.Flags().Bool("getKey", false, "flags indicating whether to execute get key command")
	cmd.Flags().Bool("getAndUpdate", false, "flags indicating whether to execute get and removed command")

	return cmd
}
