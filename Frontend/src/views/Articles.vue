<template>
    <div class="articles">
        <banner></banner>
        <div class="site-content animate">
            <!-- 文章目录 -->
            <!-- <div id="article-menus">
                <menu-tree :menus="menus" child-label="child"></menu-tree>
            </div>-->
            <main class="site-main">
                <article class="hentry">
                    <!-- 正文输出 -->
                    <articleContent class="entry-content" :cate="cate" :title="title">
                    </articleContent>
                    <!--声明-->
                    <div class="open-message">
                        <p>声明：aaayi的博客|版权所有，违者必究|如未注明，均为原创|本网站采用<a href="https://creativecommons.org/licenses/by-nc-sa/3.0/" target="_blank">BY-NC-SA</a>协议进行授权</p>
                        <p>转载：转载请注明原文链接 - <a href="/">{{title}}</a></p>
                    </div>
                    <!--评论-->
                    <div class="comments">
                        <comment v-for="item in comments" :key="item.comment.id" :comment="item.comment">
                            <template v-if="item.reply.length">
                                <comment v-for="reply in item.reply" :key="reply.id" :comment="reply"></comment>
                            </template>
                        </comment>
                    </div>
                </article>
            </main>
        </div>
    </div>
</template>

<script>
    import Banner from '@/components/banner'
    import comment from '@/components/comment'
    import menuTree from '@/components/menu-tree'
    import articleContent from '@/components/articleContent'
    import {fetchComment} from '../api'
    export default {
        name: 'articles',
        data(){
          return{
              cate: "",
              title:"",
              showDonate: false,
              comments: [],
              menus: []
          }
        },
        components: {
            articleContent,
            Banner,
            comment,
            menuTree
        },
        methods: {
          getComment(){
              fetchComment().then(res => {
                  this.comments = res.data || []
              }).catch(err => {
                  console.log(err)
              })
          },
          fetchH(arr,left,right){
              if (right) {
                  return arr.filter(item => item.offsetTop > left && item.offsetTop < right)
              }else {
                  return arr.filter(item => item.offsetTop > left)
              }
          },
          // 生成目录
          createMenus(){
              let arr = []
              for(let i=6;i>0;i--){
                  let temp = []
                  let e = document.querySelector(".entry-content").querySelectorAll(`h${i}`)
                  console.log(e,i)
                  for (let j=0;j<e.length;j++){
                      let child = this.fetchH(arr,e[j].offsetTop,(j+1 === e.length)?undefined:e[j+1].offsetTop)
                      temp.push({
                          h: i,
                          title: e[j].innerText,
                          id: e[j].id,
                          offsetTop: e[j].offsetTop,
                          child
                      })
                  }
                  if (temp.length){
                      arr = temp
                  }
              }
              this.menus = arr
          }
        },
        mounted(){
            this.createMenus()
        },
        created() {
            this.getComment()
            this.title = this.$route.params.title
            this.cate = this.$route.params.cate
        }
    }
</script>
<style scoped lang="less">
    .site-content {
        position: relative;
        .site-main {
            padding: 80px 0 0 0;
        }
    }
    #article-menus{
        position: sticky;
        top: 0;
        box-shadow: 0 2px 6px rgba(0, 0, 0, .1);
        border-radius: 3px;
        padding: 15px;
        width: 300px;
        transform: translateX(-120%) translateY(150px);
        font-size: 14px;
    }
    article.hentry {
        .open-message {
            margin: 50px 0;
            position: relative;
            background: #2B2B2B;
            padding: 10px 30px;
            border-radius: 3px;
            font-size: 14px;
            color: #fff;
            &:after {
                content: "";
                border-left: 10px solid transparent;
                border-right: 10px solid transparent;
                border-bottom: 10px solid #2B2B2B;
                position: absolute;
                top: -8px;
                left: 48%;
            }
            p {
                margin: 10px 0;
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
            }
            a {
                color: #A0DAD0;
                padding: 0 5px;
            }
        }
    }
</style>
