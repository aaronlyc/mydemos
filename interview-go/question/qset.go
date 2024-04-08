package question

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/util/sets"
)

var (
	testyaml = `
update:
    - name: test1
      file: test1.tar
      md5: 1jd42
current:
    - name: test1
      file: test1.tar
      md5: 1jd42
    - name: test2
      file: test2.tar
      md5: 3j4dt
backup:
    - name: test5
      file: test5.tar
      md5: dctg6
    - name: test7
      file: test7.tar
      md5: drt56
`
)

type config struct {
	Update  []info `yaml:"update"`
	Current []info `yaml:"current"`
	Backup  []info `yaml:"backup"`
}

type info struct {
	Name string `yaml:"name"`
	File string `yaml:"file"`
	Md5  string `yaml:"md5"`
}

func do(testConfig *config) {
	setUpdate := sets.NewString()
	for _, info := range testConfig.Update {
		setUpdate.Insert(fmt.Sprintf("%s/%s", info.Name, info.Md5))
	}

	setCurrent := sets.NewString()
	for _, info := range testConfig.Current {
		setCurrent.Insert(fmt.Sprintf("%s/%s", info.Name, info.Md5))
	}

	setBackup := sets.NewString()
	for _, info := range testConfig.Backup {
		setBackup.Insert(fmt.Sprintf("%s/%s", info.Name, info.Md5))
	}

	if setUpdate.Equal(setCurrent) {
		//	没有变化，不更新
		fmt.Println("No changes")
		return
	}

	// 更新current值
	backupCurrent := testConfig.Current
	newCurrent := make([]info, len(testConfig.Update))
	for i, value := range testConfig.Update {
		newCurrent[i] = value
	}
	testConfig.Current = newCurrent

	toAdd := setUpdate.Difference(setCurrent)
	toDelete := setCurrent.Difference(setUpdate)
	if toDelete.Len() == 0 { // 没有删除的，则一定都是添加的（不能相等了）
		fmt.Printf("just add new item: %#v, no don't backup\n", toAdd.List())
		return
	}
	if toAdd.Len() > 0 {
		fmt.Printf("should add: %#v\n", toAdd.List())
	}

	// 这里一定有变化，需要再比较current和backup
	if setCurrent.Equal(setBackup) {
		//	没有变化，不更新
		fmt.Println("No changes to backup")
		return
	}

	// 更新current值
	testConfig.Backup = backupCurrent

	// 这里一定会删除了
	if toDelete := setBackup.Difference(setCurrent); toDelete.Len() > 0 {
		fmt.Printf("should delete: %#v\n", toDelete.List())
	}

	// 这里计算一定会backup的
	if tobackup := setCurrent.Difference(setBackup); tobackup.Len() > 0 {
		fmt.Printf("should to backup: %#v\n", tobackup.List())
	}
}

func Do() {
	var testConfig config
	err := yaml.Unmarshal([]byte(testyaml), &testConfig)
	if err != nil {
		return
	}
	fmt.Printf("before do config: %s\n", testyaml)
	do(&testConfig)

	date, err := yaml.Marshal(&testConfig)
	if err != nil {
		return
	}
	fmt.Printf("after do config: %s\n", string(date))
}
