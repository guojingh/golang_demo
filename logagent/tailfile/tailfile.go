package tailfile

import (
	"context"
	"strings"
	"time"

	"github.com/IBM/sarama"
	"github.com/guojinghu/logagent/kafka"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

// tail 相关操作
type tailTask struct {
	path   string
	topic  string
	tObj   *tail.Tail
	ctx    context.Context
	cancel context.CancelFunc
}

func newTailTask(path, topic string) tailTask {
	ctx, cancel := context.WithCancel(context.Background())
	tt := tailTask{
		path:   path,
		topic:  topic,
		ctx:    ctx,
		cancel: cancel,
	}
	return tt
}

func (t *tailTask) Init() (err error) {
	cfg := tail.Config{
		ReOpen:    true,                                 // 如果文件被重命名或移动，重新打开文件
		Follow:    true,                                 // 持续跟踪文件的新增内容
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件末尾开始读取
		MustExist: false,                                // 文件不存在时不报错
		Poll:      true,                                 // 使用轮询方式检查文件变化
	}

	t.tObj, err = tail.TailFile(t.path, cfg)
	return
}

func (t *tailTask) run() {
	// 读取日志，发往 Kafka
	logrus.Infof("collect for path:%s is running...", t.path)
	// 循环读数据
	for {
		select {
		case <-t.ctx.Done():
			logrus.Infof("path:%s is stopping...", t.path) // 只要调用 t.cancel() 就会收到信号
			return
		case line, ok := <-t.tObj.Lines: // chan
			if !ok {
				logrus.Warn("tailfile file close reopen, filename:%s\n", t.path)
				time.Sleep(time.Second)
				continue
			}

			// 如果是空行就略过
			if len(strings.Trim(line.Text, "\r")) == 0 {
				logrus.Info("出现空行，直接跳过...")
				continue
			}
			// 利用通道将同步的代码改为异步的
			// 把读出来的一行日志包装成kafka的msg类型，丢到通道中
			msg := &sarama.ProducerMessage{}
			msg.Topic = t.topic // 每个 tailObj 自己的topic
			msg.Value = sarama.StringEncoder(line.Text)
			kafka.MsgChan(msg)
		}
	}
}
