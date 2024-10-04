package tailfile

import (
	"github.com/guojinghu/logagent/common"
	"github.com/sirupsen/logrus"
)

// tailTask 的管理者
type tailTaskMgr struct {
	tailTaskMap      map[string]*tailTask       // 所有的tailTask 任务
	collectEntryList []common.CollectEntry      // 所有配置项
	confChan         chan []common.CollectEntry // 等待新配置的通道
}

var (
	ttMgr *tailTaskMgr
)

// main 函数调用
func Init(allConf []common.CollectEntry) (err error) {
	// allConf 里面存了存了若干个日志的收集项
	// 针对每一个日志收集项创建一个对应的 tailObj
	ttMgr = &tailTaskMgr{
		tailTaskMap:      make(map[string]*tailTask, 20),
		collectEntryList: allConf,
		confChan:         make(chan []common.CollectEntry),
	}
	for _, conf := range allConf {
		tt := newTailTask(conf.Path, conf.Topic)
		if err = tt.Init(); err != nil {
			logrus.Errorf("create tailObj for path:%s failed, err:%v", conf.Path, err)
			continue
		}
		logrus.Infof("create a tail task for path:%s success...", conf.Path)
		ttMgr.tailTaskMap[tt.path] = &tt
		// 起一个后台的goroutine去收集日志
		go tt.run()
	}
	go ttMgr.watch() // 在后台等新的配置来

	return
}

// 一直等 confChan 有值，有值就开始去管理之前的tailTask,管理分三种情况1.原来有就什么都不干 2.原来没有现在有就新建 3.原来有现在没有就停止
func (t *tailTaskMgr) watch() {
	for {
		// 派一个小弟等着新配置来
		newConf := <-t.confChan // 取到值说明新的配置来了
		// 新配置来了之后应该管理一下我之前启动的那些新配置
		logrus.Infof("get new conf from etcd, conf:%v, start manager tailTask...", newConf)
		for _, conf := range newConf {
			// 1.原来有存在的任务不需要动
			if t.isExist(conf) {
				continue
			}
			// 2.原来没有的需要新创建一个 tailTask 任务
			tt := newTailTask(conf.Path, conf.Topic)
			if err := tt.Init(); err != nil {
				logrus.Errorf("create tailObj for path:%s failed, err:%v", conf.Path, err)
				continue
			}
			logrus.Infof("create a tail task for path:%s success...", conf.Path)
			ttMgr.tailTaskMap[tt.path] = &tt
			// 起一个后台的goroutine去收集日志
			go tt.run()
		}
		// 3.原来有的现在没有的 tailTask 需要停掉
		// 找出 tailTaskMap 中存在，但是 newConf 不存在的那些 tailTask,把它们都关掉
		for key, task := range t.tailTaskMap {
			var found bool
			for _, conf := range newConf {
				if key == conf.Path {
					found = true
					break
				}
			}
			if !found {
				// 这个 tailTask 要停掉了
				logrus.Infof("the task collect path:%s need to stop.", task.path)
				delete(t.tailTaskMap, key) // 从管理类中删掉
				task.cancel()
			}
		}
	}
}

// 判断 tailTaskMap 中是否存在该收集项
func (t *tailTaskMgr) isExist(conf common.CollectEntry) bool {
	_, ok := t.tailTaskMap[conf.Path]
	return ok
}

// 把新的配置丢到了管理对象的 confChan 中
func SendNewConf(newConf []common.CollectEntry) {
	ttMgr.confChan <- newConf
}
