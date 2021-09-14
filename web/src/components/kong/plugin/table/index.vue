<template>
  <el-table :data="plugins" style="width: 100%">
    <el-table-column label="ID">
      <template slot-scope="item">
        <el-button type="text">{{ item.row.id }}</el-button>
      </template>
    </el-table-column>
    <el-table-column label="Name" prop="name" />
    <el-table-column label="Enable" prop="desc">
      <template slot-scope="item">
        <el-switch v-model="item.row.enable" />
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
  name: 'PluginTable',
  props: {
    sid: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
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
