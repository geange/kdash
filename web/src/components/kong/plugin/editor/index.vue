<template>
  <el-form ref="plugin" :model="plugin" label-width="180px" label-position="left">
    <el-form-item label="ID">
      <el-tag>{{ plugin.id }}</el-tag>
    </el-form-item>
    <el-form-item label="Name">
      <el-input v-model="plugin.name" style="width: 300px" />
    </el-form-item>
    <el-form-item label="Enable">
      <el-switch v-model="plugin.enable" />
    </el-form-item>
    <el-form-item label="RunOn">
      <el-input v-model="plugin.run_on" style="width: 300px" />
    </el-form-item>
    <el-form-item label="Protocols">
      <el-tag v-for="protocol in plugin.protocols" :key="protocol" closable :disable-transitions="false" @close="handleProtocolsClose(protocol)">{{ protocol }}</el-tag>
      <el-input v-if="inputProtocolsVisible" ref="saveProtocolsInput" v-model="inputProtocolValue" class="input-new-tag" size="small" @keyup.enter.native="handleProtocolsInputConfirm" @blur="handleProtocolsInputConfirm" />
      <el-button v-else class="button-new-tag" size="small" @click="showProtocolsInput">+ New Tag</el-button>
    </el-form-item>
    <el-form-item label="Tags">
      <el-tag v-for="tag in plugin.tags" :key="tag" closable :disable-transitions="false" @close="handleTagsClose(tag)">{{ tag }}</el-tag>
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
      plugin: {
        created_at: 0,
        id: '222222222',
        name: 'dddd',
        config: {},
        enable: true,
        run_on: '',
        protocols: ['a', 'n'],
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
      this.plugin.tags.splice(this.plugin.tags.indexOf(key), 1)
    },
    showTagsInput() {
      this.inputTagsVisible = true
      this.$nextTick(_ => {
        this.$refs.saveTagsInput.$refs.input.focus()
      })
    },
    handleTagsInputConfirm() {
      if (this.inputTagValue) {
        if (this.plugin.tags === undefined) {
          this.plugin.tags = []
        }
        this.plugin.tags.push(this.inputTagValue)
      }
      this.inputTagsVisible = false
      this.inputTagValue = ''
    },
    handleProtocolsClose(key) {
      this.plugin.protocols.splice(this.plugin.protocols.indexOf(key), 1)
    },
    showProtocolsInput() {
      this.inputProtocolsVisible = true
      this.$nextTick(_ => {
        this.$refs.saveProtocolsInput.$refs.input.focus()
      })
    },
    handleProtocolsInputConfirm() {
      if (this.inputProtocolValue) {
        if (this.plugin.protocols === undefined) {
          this.plugin.protocols = []
        }
        this.plugin.protocols.push(this.inputProtocolValue)
      }
      this.inputProtocolsVisible = false
      this.inputProtocolValue = ''
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
