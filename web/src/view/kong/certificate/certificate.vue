<template>
  <div>
    <div v-if="tableVisible">
      <el-table :data="certificates" style="width: 100%">
        <el-table-column label="ID">
          <template slot-scope="item">
            <el-button type="text" @click="openTab(item.row.id)">{{ item.row.id }}</el-button>
          </template>
        </el-table-column>
        <el-table-column label="Cert" prop="cert" />
        <el-table-column label="Key" prop="key" />
        <el-table-column label="Tags" prop="tags">
          <template slot-scope="props">
            <el-tag v-for="tag in props.row.tags" :key="tag" size="mini">{{ tag }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="CreatedAt" prop="created_at" />
      </el-table>
    </div>
    <div v-if="tabVisible">
      <el-page-header title="back" content="Certificate" @back="goBack" />
      <Tab style="margin-top: 20px" />
    </div>
  </div>
</template>

<script>
import Tab from '@/components/kong/certificate/tab'
export default {
  name: 'Certificate',
  components: { Tab },
  data() {
    return {
      tableVisible: true,
      tabVisible: false,
      certificates: [
        {
          created_at: 0,
          id: 'sssssssss',
          cert: 'cert',
          key: 'key',
          snis: ['a', 'b'],
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
  .el-tag + .el-tag {
    margin-left: 10px;
  }
</style>
