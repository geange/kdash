<template>
  <el-table :data="targets" style="width: 100%">
    <el-table-column label="ID">
      <template slot-scope="item">
        <el-button type="text">{{ item.row.id }}</el-button>
      </template>
    </el-table-column>
    <el-table-column label="Target" prop="target" />
    <el-table-column label="Weight" prop="weight" />
    <el-table-column label="Tags" prop="tags">
      <template slot-scope="props">
        <el-tag v-for="tag in props.row.tags" :key="tag" size="mini">{{ tag }}</el-tag>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: 'PluginInstance',
  props: {
    serviceId: {
      type: Array,
      default: null
    }
  },
  data() {
    return {
      targets: [
        {
          created_at: 0,
          id: 'sssssssss',
          target: 'ss',
          weight: 10,
          tags: ['a', 'b']
        }
      ]
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
