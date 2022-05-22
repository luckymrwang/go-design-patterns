package composite

/*
	设计思想：
		struct不依赖interface
		1. 包含角色：
			1). 共同的接口IOrganization，为root和leaf结构体共有的方法
			2). root结构体(包含leaf列表)和leaf结构体
			3). 将结构体中共同部分的数据抽离，使用匿名组合的方式实现2中的两类结构体
*/

// IOrganization 组织接口，都实现统计人数的功能
type IOrganization interface {
	Count() int
}

// Employee 员工
type Employee struct {
	Name string
}

// Count 人数统计
func (Employee) Count() int {
	return 1
}

// Department 部门
type Department struct {
	Name string

	SubOrganizations []IOrganization
}

// Count 人数统计
func (d Department) Count() int {
	c := 0
	for _, org := range d.SubOrganizations {
		c += org.Count()
	}
	return c
}

// AddSub 添加子节点
func (d *Department) AddSub(org IOrganization) {
	d.SubOrganizations = append(d.SubOrganizations, org)
}

// NewOrganization 构建组织架构 demo
func NewOrganization() IOrganization {
	root := &Department{Name: "root"}
	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{})
		root.AddSub(&Department{Name: "sub", SubOrganizations: []IOrganization{&Employee{}}})
	}
	return root
}
