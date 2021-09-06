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
  name: 'TargetTable',
  props: {
    sid: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      targets: [
        {
          created_at: 0,
          id: '222222222',
          target: 'dddd:8080',
          weight: 80,
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
