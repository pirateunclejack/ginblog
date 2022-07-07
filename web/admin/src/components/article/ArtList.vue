<template>
  <div>
    <h3>Article List</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search
            v-model="queryParam.title"
            placeholder="input title to search"
            enter-button
            allowClear
            @search="getArtList"
          />
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="$router.push('/admin/addart')">Add</a-button>
        </a-col>
        <a-col :span="4">
          <a-select
            defaultValue="Please selece Category"
            style="width: 200px"
            @change="CateChange"
          >
            <a-select-option
              v-for="item in Catelist"
              :key="item.id"
              :value="item.id"
              >{{ item.name }}</a-select-option
            >
          </a-select>
        </a-col>
      </a-row>
      <a-table
        rowKey="ID"
        :columns="columns"
        :pagination="pagination"
        :dataSource="Artlist"
        bordered
        @change="handleTableChange"
      >
        <span class="ArtImg" slot="img" slot-scope="img">
          <img :src="img" />
        </span>
        <template slot="action" slot-scope="text, data">
          <div class="actionSlot">
            <a-button
              size="small"
              type="primary"
              icon="edit"
              style="margin-right: 15px"
              @click="$router.push(`addart/${data.ID}`)"
              >Edit</a-button
            >
            <a-button
              size="small"
              type="danger"
              icon="delete"
              style="margin-right: 15px"
              @click="deleteArt(data.ID)"
              >Delete</a-button
            >
          </div>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script>
import day from 'dayjs'
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '5%',
    key: 'id',
    align: 'center'
  },
  {
    title: 'Updated at',
    dataIndex: 'UpdatedAt',
    width: '10%',
    key: 'UpdatedAt',
    align: 'center',
    customRender: (val) => {
      return val ? day(val).format('YYYY-MM-DD HH:mm') : 'NULL'
    }
  },
  {
    title: 'Category',
    dataIndex: 'Category.name',
    width: '5%',
    key: 'Category.name',
    align: 'center'
  },
  {
    title: 'title',
    dataIndex: 'title',
    width: '20%',
    key: 'title',
    align: 'center'
  },
  {
    title: 'describe',
    dataIndex: 'describe',
    width: '20%',
    key: 'describe',
    align: 'center'
  },
  {
    title: 'img',
    dataIndex: 'img',
    width: '10%',
    key: 'img',
    align: 'center',
    scopedSlots: { customRender: 'img' }
  },
  {
    title: 'Action',
    dataIndex: 'action',
    width: '10%',
    key: 'action',
    scopedSlots: { customRender: 'action' },
    align: 'center'
  }
]
export default {
  data() {
    return {
      pagination: {
        pageSizeOptions: ['2', '4', '8'],
        pageSize: 2,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `Total ${total}`
      },
      Artlist: [],
      Catelist: [],
      columns,
      queryParam: {
        title: '',
        pagesize: 2,
        pagenum: 1
      }
    }
  },
  created() {
    this.getArtList()
    this.getCateList()
  },
  computed: {
    IsAdmin: function () {
      if (this.ArtInfo.role === 1) {
        return true
      } else {
        return false
      }
    }
  },
  methods: {
    async getArtList() {
      const { data: res } = await this.$http.get('article', {
        params: {
          title: this.queryParam.title,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.Artlist = res.data
      this.pagination.total = res.total
    },
    // Get Category list
    async getCateList() {
      const { data: res } = await this.$http.get('category')
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.Catelist = res.data
      this.pagination.total = res.total
    },
    handleTableChange(pagination, filters, sorter) {
      const pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getArtList()
    },
    deleteArt(id) {
      this.$confirm({
        title: 'Do you want to delete this article?',
        content:
          'Are you sure you want to delete this article? Once it is deleted , it will be impossible to recover it!',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`article/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('Delete success')
          this.getArtList()
        },
        onCancel: () => {
          this.$message.info('Cancel delete')
        }
      })
    },
    CateChange(value) {
      this.getCateArt(value)
    },
    async getCateArt(id) {
      const { data: res } = await this.$http.get(`article/list/${id}`, {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.Artlist = res.data
      this.pagination.total = res.total
    }
  }
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
.ArtImg {
  height: 100%;
  width: 100%;
}
.ArtImg img {
  width: 100px;
  height: 80px;
}
</style>
