<template>
  <div>
    <div v-if="tableVisible">
      <el-table :data="routes" style="width: 100%">
        <el-table-column label="ID">
          <template slot-scope="item">
            <el-button type="text" @click="openTab(item.row.id)">{{ item.row.id }}</el-button>
          </template>
        </el-table-column>
        <el-table-column label="Name" prop="name" />
        <el-table-column label="Paths" prop="paths">
          <template slot-scope="item">
            <el-tag v-for="path in item.row.paths" :key="path" size="mini">{{ path }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="tags" prop="tags">
          <template slot-scope="props">
            <el-tag v-for="tag in props.row.tags" :key="tag" size="mini">{{ tag }}</el-tag>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div v-if="tabVisible">
      <el-page-header title="back" content="Route" @back="goBack" />
      <Tab style="margin-top: 20px" />
    </div>
  </div>
</template>

<script>
import Tab from '@/components/kong/route/tab'
export default {
  name: 'PluginInstance',
  components: { Tab },
  props: {
    serviceId: {
      type: Array,
      default: null
    }
  },
  data() {
    return {
      tableVisible: true,
      tabVisible: false,
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
      ]
    }
  },
  created() {
    console.log('create')
  },
  methods: {
    goBack() {
      this.tableVisible = true
      this.tabVisible = false
    },
    openTab(id) {
      this.sid = id
      this.tableVisible = false
      this.tabVisible = true
    }
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
