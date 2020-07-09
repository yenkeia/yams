package main

import (
	"time"

	"github.com/davyxu/cellnet"
	_ "github.com/davyxu/cellnet/codec/binary"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/tcp"
	"github.com/davyxu/cellnet/proc"
	"github.com/davyxu/cellnet/timer"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yenkeia/yams/game"
	_ "github.com/yenkeia/yams/game/mircodec"
	_ "github.com/yenkeia/yams/game/mirtcp"
)

func main() {
	env := game.NewEnviron(game.NewConfig("../../configs/yams.yaml"))

	queue := cellnet.NewEventQueue() // 这里用 cellnet 单线程模式。消息处理都在queue线程。无需再另开线程

	p := peer.NewGenericPeer("tcp.Acceptor", "server", "0.0.0.0:7000", queue)
	proc.BindProcessorHandler(p, "mir.server.tcp", env.HandleEvent)

	timer.NewLoop(queue, time.Second/time.Duration(60), func(*timer.Loop) {
		env.Update()
	}, nil).Start()

	env.Peer = p

	p.Start()         // 开始侦听
	queue.StartLoop() // 事件队列开始循环
	queue.Wait()      // 阻塞等待事件队列结束退出( 在另外的goroutine调用queue.StopLoop() )
}
