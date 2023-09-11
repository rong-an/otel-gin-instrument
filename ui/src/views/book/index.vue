<template>
    <div class="app-container">
      <el-table
        :key="tableKey"
        v-loading="listLoading"
        :data="list"
        border
        fit
        highlight-current-row
        style="width: 100%;"
      >
        <el-table-column label="ID" prop="id" align="center" width="50">
          <template slot-scope="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column label="标题" min-width="110px">
          <template slot-scope="{row}">
            <span>{{ row.title }}</span>
          </template>
        </el-table-column>
        <el-table-column label="作者" min-width="120px">
          <template slot-scope="{row}">
            <span>{{ row.author }}</span>
          </template>
        </el-table-column>
        <!-- <el-table-column label="操作" align="center" width="230" class-name="small-padding fixed-width">
          <template slot-scope="{row,$index}">
            <router-link :to="'/service/service_stat/'+row.id">
              <el-button type="primary" size="mini">
                统计
              </el-button>
            </router-link>
            <router-link v-if="row.load_type===0" :to="'/service/service_edit_http/'+row.id">
              <el-button type="primary" size="mini">
                修改
              </el-button>
            </router-link>
            <router-link v-if="row.load_type===1" :to="'/service/service_edit_tcp/'+row.id">
              <el-button type="primary" size="mini">
                修改
              </el-button>
            </router-link>
            <router-link v-if="row.load_type===2" :to="'/service/service_edit_grpc/'+row.id">
              <el-button type="primary" size="mini">
                修改
              </el-button>
            </router-link>
            <el-button size="mini" type="danger" @click="handleDelete(row,$index)">
              删除
            </el-button>
          </template>
        </el-table-column> -->
      </el-table>
  
      <!-- <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" /> -->
    </div>
  </template>
  
  <script>
  import { bookList } from '@/api/book'
  import waves from '@/directive/waves' // waves directive
//   import Pagination from '@/components/Pagination' // secondary package based on el-pagination
  
  
//   const loadTypeKeyValue = loadTypeOptions.reduce((acc, cur) => {
//     acc[cur.key] = cur.display_name
//     return acc
//   }, {})
  
  export default {
    name: 'BookList',
    components: {},
    directives: { waves },
    filters: {},
    data() {
      return {
        tableKey: 0,
        list: null,
        total: 0,
        listLoading: false,
        listQuery: {
          page_no: 1,
          page_size: 20,
          info: ''
        },
        temp: {
          id: undefined
        }
      }
    },
    created() {
      this.getList()
    },
    methods: {
      getList() {
        // this.listLoading = true
        bookList().then(response => {
          this.list = response.data
        //   this.list = response.data.list

        //   this.listLoading = false
        })
      },
    //   handleFilter() {
    //     this.listQuery.page_no = 1
    //     this.getList()
    //   },
    //   handleDelete(row, index) {
    //     this.$confirm('此操作将删除该记录, 是否继续?', '提示', {
    //       confirmButtonText: '确定',
    //       cancelButtonText: '取消',
    //       type: 'warning'
    //     }).then(() => {
    //       const deleteQuery = {
    //         'id': row.id
    //       }
    //       serviceDelete(deleteQuery).then(response => {
    //         this.$notify({
    //           title: 'Success',
    //           message: '删除成功',
    //           type: 'success',
    //           duration: 2000
    //         })
    //         this.getList()
    //       })
    //     }).catch(() => {
    //       this.$notify({
    //         title: 'Success',
    //         message: '已取消删除',
    //         type: 'info',
    //         duration: 2000
    //       })
    //     })
  
    //     // this.list.splice(index, 1)
    //   }
    }
  }
  </script>
  