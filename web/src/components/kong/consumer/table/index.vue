<template>
  <div>
    <div class="button-box clearflex">
      <el-button size="mini" type="primary" icon="el-icon-plus" @click="add">ADD</el-button>
    </div>
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
    <Dialog :visible="dialogFormVisible" @closeHandler="closeDialog" />
  </div>
</template>

<script>
import Dialog from '@/components/kong/consumer/dialog'
export default {
  name: 'ConsumerTable',
  components: { Dialog },
  props: {
    sid: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      dialogFormVisible: false,
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
  },
  methods: {
    add() {
      console.log('add new route')
      this.dialogFormVisible = true
    },
    closeDialog() {
      this.dialogFormVisible = false
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
