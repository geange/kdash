<template>
  <el-form ref="consumer" :model="consumer" label-width="80px">
    <el-form-item label="CustomID">
      <el-input v-model="consumer.custom_id" style="width: 300px" />
    </el-form-item>
    <el-form-item label="Username">
      <el-input v-model="consumer.username" style="width: 300px" />
    </el-form-item>
    <el-form-item label="Tags">
      <el-tag v-for="tag in consumer.tags" :key="tag" closable :disable-transitions="false" @close="handleTagsClose(tag)">
        {{ tag }}
      </el-tag>
      <el-input v-if="inputTagsVisible" ref="saveTagsInput" v-model="inputTagValue" class="input-new-tag" size="small" @keyup.enter.native="handleTagsInputConfirm" @blur="handleTagsInputConfirm" />
      <el-button v-else class="button-new-tag" size="small" @click="showTagsInput">+ New Tag</el-button>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">立即创建</el-button>
      <el-button>取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
export default {
  name: 'ConsumerEditor',
  data() {
    return {
      inputTagsVisible: false,
      inputTagValue: '',
      consumer: {
        id: 'xxxxxxxxxx',
        custom_id: 'custom_id',
        username: 'ssss',
        created_at: 1221,
        tags: ['a', 'b']
      }
    }
  },
  created() {
    console.log('create')
  },
  methods: {
    onSubmit() {
      console.log(this.consumer)
    },
    handleTagsClose(tag) {
      this.consumer.tags.splice(this.consumer.tags.indexOf(tag), 1)
    },
    showTagsInput() {
      this.inputTagsVisible = true
      this.$nextTick(_ => {
        this.$refs.saveTagsInput.$refs.input.focus()
      })
    },
    handleTagsInputConfirm() {
      if (this.inputTagValue) {
        if (this.consumer.tags === undefined) {
          this.consumer.tags = []
        }
        this.consumer.tags.push(this.inputTagValue)
      }
      this.inputTagsVisible = false
      this.inputTagValue = ''
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
  .button-new-tag {
    margin-left: 10px;
  }
  .input-new-tag {
    width: 120px;
    margin-left: 10px;
    vertical-align: bottom;
  }
</style>
