<template>
  <el-form ref="ca_certificate" :model="ca_certificate" label-width="auto">
    <el-form-item label="ID">
      <el-tag>{{ ca_certificate.id }}</el-tag>
    </el-form-item>
    <el-form-item label="Cert">
      <el-input v-model="ca_certificate.cert" />
    </el-form-item>
    <el-form-item label="Tags">
      <el-tag v-for="tag in ca_certificate.tags" :key="tag" closable :disable-transitions="false" @close="handleTagsClose(tag)">{{ tag }}</el-tag>
      <el-input v-if="inputTagsVisible" ref="saveTagsInput" v-model="inputTagValue" class="input-new-tag" size="small" @keyup.enter.native="handleTagsInputConfirm" @blur="handleTagsInputConfirm" />
      <el-button v-else class="button-new-tag" size="small" @click="showTagsInput">+ New Tag</el-button>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">Update</el-button>
      <el-button>Cancel</el-button>
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
      inputSNIsVisible: false,
      inputSNIValue: '',
      ca_certificate: {
        created_at: 0,
        id: 'sssssssss',
        cert: 'cert',
        key: 'key',
        snis: ['a', 'b'],
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
      this.ca_certificate.tags.splice(this.ca_certificate.tags.indexOf(key), 1)
    },
    showTagsInput() {
      this.inputTagsVisible = true
      this.$nextTick(_ => {
        this.$refs.saveTagsInput.$refs.input.focus()
      })
    },
    handleTagsInputConfirm() {
      if (this.inputTagValue) {
        if (this.ca_certificate.tags === undefined) {
          this.ca_certificate.tags = []
        }
        this.ca_certificate.tags.push(this.inputTagValue)
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
