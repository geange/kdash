<template>
  <div>
    <div v-if="tableVisible">
      <el-table :data="snis" style="width: 100%">
        <el-table-column label="ID">
          <template slot-scope="item">
            <el-button type="text" @click="openTab(item.row.id)">{{ item.row.id }}</el-button>
          </template>
        </el-table-column>
        <el-table-column label="Name" prop="name" />
        <el-table-column label="Tags" prop="tags">
          <template slot-scope="props">
            <el-tag v-for="tag in props.row.tags" :key="tag" size="mini">{{ tag }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="CreatedAt" prop="created_at" />
      </el-table>
    </div>
    <div v-if="tabVisible">
      <el-page-header title="back" content="SNI" @back="goBack" />
      <Tab />
    </div>
  </div>
</template>

<script>
import Tab from '@/components/kong/sni/tab'
export default {
  name: 'SNI',
  components: { Tab },
  data() {
    return {
      tableVisible: true,
      tabVisible: false,
      snis: [
        {
          created_at: 0,
          id: 'sssssssss',
          name: 'ss',
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
