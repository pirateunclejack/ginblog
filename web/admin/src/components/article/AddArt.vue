<template>
  <div>
    <a-card>
      <h3>{{id? 'Edit article': 'Add article'}}</h3>
      <a-form-model :model="artInfo" ref="artInfoRef" :rules="artInfoRules">
        <a-form-model-item label="Article title" prop="title">
          <a-input style="width:300px" v-model="artInfo.title"></a-input>
        </a-form-model-item>
        <a-form-model-item label="Article category" prop="cid">
          <a-select
            style="width:200px"
            v-model="artInfo.cid"
            placeholder="Please select category"
            @change="cateChange"
          >
            <a-select-option
              v-for="item in Catelist"
              :key="item.id"
              :value="item.id"
            >
                {{item.name}}
            </a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item label="Article descirbe" prop="describe">
          <a-input type="textarea" v-model="artInfo.describe"></a-input>
        </a-form-model-item>
        <a-form-model-item label="Article image" prop="img">
          <a-upload
            name="file"
            :multiple="false"
            :action="upUrl"
            :headers="headers"
            @change="upChange"
            listType="picture"
          >
            <a-button>
              <a-icon type="upload"/>Upload
            </a-button>
            <template v-if="id">
              <img :src="artInfo.img" style="width: 120px; height: 100px; margin-left: 15px" />
            </template>
          </a-upload>
        </a-form-model-item>
        <a-form-model-item label="Article content" prop="content">
          <a-input type="textarea" v-model="artInfo.content"></a-input>
        </a-form-model-item>
        <a-form-model-item>
          <a-button
            type="danger"
            style="margin-right:15px"
            @click="artOk(artInfo.id)"
          >{{artInfo.id? 'Update': 'Submit'}}</a-button>
          <a-button type="primary" @click="addCancle">Cancel</a-button>
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>

<script>
import { Url } from '../../plugin/http'
export default {
  props: ['id'],
  data () {
    return {
      artInfo: {
        id: 0,
        title: '',
        cid: undefined,
        descrbe: '',
        content: '',
        img: ''
      },
      Catelist: [],
      upUrl: Url + '/upload',
      headers: [],
      artInfoRules: {
        title: [{ required: true, message: 'Please input article title', trigger: 'blur' }],
        cid: [{ required: true, message: 'Please select article category', trigger: 'change' }],
        describe: [{ required: true, message: 'Please input article describe', trigger: 'blur' },
          { max: 120, message: 'Article describe should no longer than 120 words', trigger: 'change' }],
        img: [{ required: true, message: 'Please select article image', trigger: 'blur' }],
        content: [{ required: true, message: 'Please input article content', trigger: 'blur' }]
      }
    }
  },
  created() {
    this.getCateList()
    this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
    if (this.id) {
      this.getArtInfo(this.id)
    }
  },
  methods: {
    // Get Article info
    async getArtInfo(id) {
      const { data: res } = await this.$http.get(`article/info/${id}`)
      if (res.status !== 200) return this.$message.error(res.message)
      this.artInfo = res.data
      this.artInfo.id = res.data.ID
    },
    // Get Category list
    async getCateList() {
      const { data: res } = await this.$http.get('category')
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.Catelist = res.data
    },
    cateChange(value) {
      this.artInfo.cid = value
    },
    upChange(info) {
      if (info.file.status === 'done') {
        this.$message.success('Image upload success')
        this.artInfo.img = info.file.response.url
      } else if (info.file.status === 'error') {
        this.$message.error(`${info.file.name} file upload failed.`)
      }
    },
    // Update or Submit article
    artOk(id) {
      this.$refs.artInfoRef.validate(async (valid) => {
        if (id === 0) {
          const { data: res } = await this.$http.post('article/add', this.artInfo)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$router.push('/admin/artlist')
          this.$message.success('Add article success')
        } else {
          const { data: res } = await this.$http.put(`article/${id}`, this.artInfo)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$router.push('/admin/artlist')
          this.$message.success('Update article success')
        }
      })
    },
    addCancle() {
      this.$refs.artInfoRef.resetFields()
    }
  }
}
</script>
