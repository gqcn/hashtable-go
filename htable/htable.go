package htable


// 哈希表
type Table struct {
    array  []*tableItem // 分区数组
}

// 数据项
type tableItem struct {
    key   interface{}   // 数据项键名
    value interface{}   // 数据项键值
    table *Table        // 溢出，指向另一个哈希表(链表)
}

// 创建指定大小的哈希表
func New(size int) *Table {
    return &Table{
        array : make([]*tableItem, size),
    }
}

// 设置键值对数据
func (t *Table) Set(key int, value interface{}) {
    hash := key%len(t.array)
    if t.array[hash] != nil {
        if t.array[hash].key != key {
            // 出现碰撞时，数据项的table参数指向新的哈希表，但是新哈希表的大小不会变化
            if t.array[hash].table == nil {
                t.array[hash].table = New(len(t.array))
            }
            t.array[hash].table.Set(key, value)
        } else {
            t.array[hash].value = value
        }
    } else {
        t.array[hash] = &tableItem{
            key   : key,
            value : value,
        }
    }
}

// 根据键名查询键值
func (t *Table) Get(key int) interface{} {
    hash := key%len(t.array)
    if t.array[hash] != nil {
        if t.array[hash].key != key {
            return t.array[hash].table.Get(key)
        } else {
            return t.array[hash].value
        }
    }
    return nil
}