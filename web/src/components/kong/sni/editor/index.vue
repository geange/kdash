<template>
  <el-form ref="sni" :model="sni" label-width="180px" label-position="left">
    <el-form-item label="ID">
      <el-tag>{{ sni.id }}</el-tag>
    </el-form-item>
    <el-form-item label="Name">
      <el-input v-model="sni.name" style="width: 300px" />
    </el-form-item>
    <el-form-item label="Tags">
      <el-tag v-for="tag in sni.tags" :key="tag" closable :disable-transitions="false" @close="handleTagsClose(tag)">{{ tag }}</el-tag>
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
  name: 'Tab',
  props: {
    sid: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      inputTagsVisible: false,
      inputTagValue: '',
      inputProtocolsVisible: false,
      inputProtocolValue: '',
      sni: {
        created_at: 0,
        id: 'sssssssss',
        name: 'ss',
        tags: ['a', 'b']
      }
    }
  },
  created() {
  },
  methods: {
    onSubmit() {
    },
    handleTagsClose(key) {
      this.sni.tags.splice(this.sni.tags.indexOf(key), 1)
    },
    showTagsInput() {
      this.inputTagsVisible = true
      this.$nextTick(_ => {
        this.$refs.saveTagsInput.$refs.input.focus()
      })
    },
    handleTagsInputConfirm() {
      if (this.inputTagValue) {
        if (this.sni.tags === undefined) {
          this.sni.tags = []
        }
        this.sni.tags.push(this.inputTagValue)
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
