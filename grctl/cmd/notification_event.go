package cmd

import (
	"github.com/urfave/cli"
	"github.com/goodrain/rainbond/grctl/clients"
	"fmt"
	"errors"
	"github.com/apcera/termtables"
	"time"
	"strconv"
)

//NewCmdNode NewCmdNode
func NewCmdNotificationEvent() cli.Command {
	c := cli.Command{
		Name:  "notification",
		Usage: "应用异常通知事件。grctl notification",
		Subcommands: []cli.Command{
			{
				Name:  "get",
				Usage: "get notification",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "StartTime,st",
						Value: "",
						Usage: "StartTime timestamp",
					},
					cli.StringFlag{
						Name:  "EndTime,et",
						Value: "",
						Usage: "EndTime timestamp",
					},
				},
				Action: func(c *cli.Context) error {
					Common(c)
					if c.IsSet("StartTime") {
						startTime := c.String("StartTime")
						EndTme := c.String("EndTime")
						if EndTme == "" {
							NowTime := time.Now().Unix()
							EndTme = strconv.FormatInt(NowTime, 10)
						}
						val, err := clients.RegionClient.Notification().GetNotification(startTime, EndTme)
						handleErr(err)
						serviceTable := termtables.CreateTable()
						serviceTable.AddHeaders("ServiceName", "TenantName", "Type", "Message", "Reason", "Count", "LastTime", "FirstTime", "IsHandle", "HandleMessage")
						for _, v := range val {
							serviceTable.AddRow(v.ServiceName, v.TenantName, v.Type, v.Message, v.Reason, v.Count, v.LastTime, v.FirstTime, v.IsHandle, v.HandleMessage)
						}
						fmt.Println(serviceTable.Render())
					}
					return errors.New("StartTime not null")
				},
			},
		},
	}
	return c
}
