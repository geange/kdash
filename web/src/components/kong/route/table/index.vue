<template>
  <el-table :data="routes" style="width: 100%">
    <el-table-column label="ID">
      <template slot-scope="item">
        <el-button type="text">{{ item.row.id }}</el-button>
      </template>
    </el-table-column>
    <el-table-column label="Name" prop="name" />
    <el-table-column label="Paths" prop="paths">
      <template slot-scope="props">
        <el-tag v-for="path in props.row.paths" :key="path" size="mini">{{ path }}</el-tag>
      </template>
    </el-table-column>
    <el-table-column label="Tags" prop="tags">
      <template slot-scope="props">
        <el-tag v-for="tag in props.row.tags" :key="tag" size="mini">{{ tag }}</el-tag>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: 'RouteTable',
  props: {
    serviceId: {
      type: Array,
      default: null
    }
  },
  data() {
    return {
      routes: [
        {
          created_at: 0,
          hosts: ['a.cn', 'b.cn'],
          headers: {
            'name': ['a', 'b']
          },
          id: 'sssssssss',
          name: 'ss',
          methods: ['GET'],
          paths: ['/abssssssssssssssssssssssssc9', '/abssssssssssssssssssssssssc'],
          path_handling: 'ss',
          preserve_host: 'sss.cn',
          protocols: ['http', 'https'],
          regex_priority: 0,
          service: {
            id: 'sssssffffff'
          },
          strip_path: false,
          updated_at: 0,
          snis: ['a1', 'a2'],
          sources: [{
            ip: '',
            port: 8000
          }],
          destinations: [],
          tags: ['a', 'b'],
          https_redirect_status_code: 302,
          request_buffering: false,
          response_buffering: false
        }
      ],
      tableData: [{
        id: '12987122',
        name: '好滋好味鸡蛋仔',
        category: '江浙小吃、小吃零食',
        desc: '荷兰优质淡奶，奶香浓而不腻',
        address: '上海市普陀区真北路',
        shop: '王小虎夫妻店',
        shopId: '10333'
      }, {
        id: '12987123',
        name: '好滋好味鸡蛋仔',
        category: '江浙小吃、小吃零食',
        desc: '荷兰优质淡奶，奶香浓而不腻',
        address: '上海市普陀区真北路',
        shop: '王小虎夫妻店',
        shopId: '10333'
      }, {
        id: '12987125',
        name: '好滋好味鸡蛋仔',
        category: '江浙小吃、小吃零食',
        desc: '荷兰优质淡奶，奶香浓而不腻',
        address: '上海市普陀区真北路',
        shop: '王小虎夫妻店',
        shopId: '10333'
      }, {
        id: '12987126',
        name: '好滋好味鸡蛋仔',
        category: '江浙小吃、小吃零食',
        desc: '荷兰优质淡奶，奶香浓而不腻',
        address: '上海市普陀区真北路',
        shop: '王小虎夫妻店',
        shopId: '10333'
      }]
    }
  },
  created() {
    console.log('create')
  }
}
</script>

<style>
  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
  .el-tag + .el-tag {
    margin-left: 10px;
  }
</style>
