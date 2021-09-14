<template>
  <el-table :data="consumers" style="width: 100%">
    <el-table-column label="ID" prop="id">
      <template slot-scope="item">
        <el-button type="text">{{ item.row.id }}</el-button>
      </template>
    </el-table-column>
    <el-table-column label="CustomID" prop="custom_id" />
    <el-table-column label="Username" prop="username" />
    <el-table-column label="Tags" prop="tags">
      <template slot-scope="props">
        <el-tag v-for="tag in props.row.tags" :key="tag" size="mini">{{ tag }}</el-tag>
      </template>
    </el-table-column>
    <el-table-column label="CreatedAt" prop="created_at" />
  </el-table>
</template>

<script>
export default {
  name: 'ConsumerTable',
  props: {
    sid: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      consumers: [
        {
          id: 'ssssssss',
          custom_id: 'custom_id_hhhhh',
          username: 'kkk',
          created_at: 0,
          tags: ['a', 'n']
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
</style>
