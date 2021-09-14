<template>
  <div>
    <div v-if="tableVisible">
      <el-table :data="plugins" style="width: 100%">
        <el-table-column label="ID">
          <template slot-scope="item">
            <el-button type="text" @click="openTab(item.row.id)">{{ item.row.id }}</el-button>
          </template>
        </el-table-column>
        <el-table-column label="Name" prop="name" />
        <el-table-column label="Enable" prop="desc">
          <template slot-scope="item">
            <el-switch v-model="item.row.enable" />
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
      <el-page-header title="back" content="Plugin" @back="goBack" />
      <Tab style="margin-top: 20px" />
    </div>
  </div>
</template>

<script>
import Tab from '@/components/kong/plugin/tab'
export default {
  name: 'Plugin',
  components: { Tab },
  data() {
    return {
      tableVisible: true,
      tabVisible: false,
      plugins: [
        {
          created_at: 0,
          id: '222222222',
          name: 'dddd',
          config: {},
          enable: true,
          run_on: '',
          protocols: ['a', 'n'],
          tags: ['a', 'b']
        }
      ]
    }
  },
  created() {
  },
  methods: {
    goBack() {
      this.tableVisible = true
      this.tabVisible = false
    },
    openTab(id) {
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
</style>
