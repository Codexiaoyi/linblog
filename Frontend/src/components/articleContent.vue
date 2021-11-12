<template>
   <div class="entry-content" v-html="content">
        {{content}}
    </div>
</template>

<script>
    import {fetchArticleContent} from '../api'
    export default {
        name: 'articleContent',
        props: {
            cate: {
                type: String
            },
            title: {
                type: String
            }
        },
        data(){
          return{
              content:"",
          }
        },
        methods: {
          getArticleContent(cate,title){
              fetchArticleContent(cate,title).then(res => {
                  this.content = res.data || ""
              }).catch(err => {
                  console.log(err)
              })
          },
        },
        created() {
            this.getArticleContent(this.cate,this.title)
        }
    }
</script>

<style scoped lang="less">

</style>
