<template>
    <div>
        <!-- 文章头部 -->
        <header class="entry-header">
            <!-- 标题输出 -->
            <h1 class="entry-title">{{title}}</h1>
            <hr>
            <div class="breadcrumbs">
                <div id="crumbs">创建时间：{{publishTime}}</div>
            </div>
        </header>
        <!-- 正文输出 -->
        <div class="entry-content" v-html="content">
            {{content}}
        </div>
        <!-- 文章底部 -->
        <section-title>
            <footer class="post-footer">
                <!-- 阅读次数 -->
                <div class="post-like">
                    <i class="iconfont iconeyes"></i>
                    <span class="count">{{viewsCount}}</span>
                </div>
                <!-- 分享按钮 -->
                <!--                        <div class="post-share">-->
                <!--                            <ul class="sharehidden">-->
                <!--                                <li><a href="http://qr.liantu.com/api.php?text=https://zhebk.cn/Web/gongan.html"-->
                <!--                                       onclick="window.open(this.href, 'renren-share', 'width=490,height=700');return false;"-->
                <!--                                       class="s-weixin" target="_blank" rel="nofollow noopener noreferrer"><img src="https://cdn.zhebk.cn/usr/themes/Akina/images/wechat.png"></a></li>-->
                <!--                                <li>-->
                <!--                                    <a href="http://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?url=https://zhebk.cn/Web/gongan.html/&amp;title=公安联网备案记录"-->
                <!--                                       onclick="window.open(this.href, 'weibo-share', 'width=730,height=500');return false;"-->
                <!--                                       class="s-qq" target="_blank" rel="nofollow noopener noreferrer"><img src="https://cdn.zhebk.cn/usr/themes/Akina/images/qzone.png"></a></li>-->
                <!--                                <li>-->
                <!--                                    <a href="http://service.weibo.com/share/share.php?url=https://zhebk.cn/Web/gongan.html/&amp;title=公安联网备案记录"-->
                <!--                                       onclick="window.open(this.href, 'weibo-share', 'width=550,height=235');return false;"-->
                <!--                                       class="s-sina" target="_blank" rel="nofollow noopener noreferrer"><img src="https://cdn.zhebk.cn/usr/themes/Akina/images/sina.png"></a></li>-->
                <!--                                <li>-->
                <!--                                    <a href="http://shuo.douban.com/!service/share?https://zhebk.cn/Web/gongan.html/&amp;title=公安联网备案记录"-->
                <!--                                       onclick="window.open(this.href, 'renren-share', 'width=490,height=600');return false;"-->
                <!--                                       class="s-douban" target="_blank" rel="nofollow noopener noreferrer"><img src="https://cdn.zhebk.cn/usr/themes/Akina/images/douban.png"></a></li>-->
                <!--                            </ul>-->
                <!--                            <i class="iconfont show-share"></i>-->
                <!--                        </div>-->
                <!-- 赞助按钮 -->
                <!-- <div class="donate" @click="showDonate=!showDonate">
                    <span>赏</span>
                    <ul class="donate_inner" :class="{'show':showDonate}">
                        <li class="wedonate"><img src="http://cdn.fengziy.cn/gblog/wexin_pay.png"><p>微信</p></li>
                        <li class="alidonate"><img src="http://cdn.fengziy.cn/gblog/ali_pay.jpg"><p>支付宝</p></li>
                    </ul>
                </div> -->
                <!-- 文章标签 -->
                <div class="post-tags">
                    <i class="iconfont iconcategory"></i>
                    <router-link :to="`/category/${cate}`">标签：{{cate}}</router-link>
                </div>
            </footer>
        </section-title>
    </div>
</template>

<script>
    import sectionTitle from '@/components/section-title'
    import {fetchArticleContent,fetchArticleInfo} from '../api'
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
              publishTime:"",
              viewsCount:"",
              content:"",
          }
        },
        components: {
            sectionTitle,
        },
        methods: {
          getArticleContent(cate,title){
              fetchArticleContent(cate,title).then(res => {
                  this.content = res.data || ""
              }).catch(err => {
                  console.log(err)
              })
          },
          getArticleInfo(cate,title){
              fetchArticleInfo(cate,title).then(res => {
                  this.publishTime = res.data.pubTime
                  this.viewsCount = res.data.viewsCount
              }).catch(err => {
                  console.log(err)
              })
          },
        },
        mounted() {
            this.getArticleContent(this.cate,this.title)
            this.getArticleInfo(this.cate,this.title)
        }
    }
</script>

<style scoped lang="less">
    .entry-header {
        .entry-title {
                font-size: 23px;
                font-weight: 600;
                color: #737373;
                margin: 0.67em 0;

                &:before {
                    content: "#";
                    margin-right: 6px;
                    color: #d82e16;
                    font-size: 20px;
                    font-weight: 600;
                }
            }

            hr {
                height: 1px;
                border: 0;
                background: #EFEFEF;
                margin: 15px 0;
            }

            .breadcrumbs {
                font-size: 14px;
                color: #D2D2D2;
                text-decoration: none;
                margin-bottom: 30px;
            }
        }

        footer.post-footer {
            width: 100%;
            padding: 20px 10px;
            margin-top: 30px;
            height: 65px;
            position: relative;
            i{
                font-size: 18px;
                margin-right: 5px;
            }
            .post-like {
                float: right;
                margin: 7px 0 0 20px;
            }

            .post-share {
                float: right;
                list-style: none;
                margin-right: 20px;
            }

            .donate {
                float: left;
                line-height: 36px;
                border-radius: 100%;
                -webkit-border-radius: 100%;
                -moz-border-radius: 100%;
                border: 1px solid #2B2B2B;
                &:hover{
                    border: 1px solid goldenrod;
                    span{
                        color: goldenrod;
                    }
                }
                span {
                    color: #2B2B2B;
                    padding: 10px;
                    position: relative;
                    cursor: pointer;
                }

                .donate_inner {
                    display: none;
                    margin: 0;
                    list-style: none;
                    position: absolute;
                    left: 80px;
                    top: -40px;
                    background: #FFF;
                    padding: 10px;
                    border: 1px solid #ddd;
                    box-shadow: 0 2px 6px rgba(0, 0, 0, .08);
                    border-radius: 3px;
                    &.show{
                        display: block;
                    }
                    li {
                        float: left;
                    }

                    img {
                        width: 100px;
                    }
                    p {
                        text-align: center;
                        font-size: 15px;
                        color: #D2D2D2;
                        line-height: 1rem;
                    }
                }

                .donate_inner:after, .donate_inner:before {
                    content: "";
                    position: absolute;
                    left: 0;
                    bottom: 45%;
                    margin-left: -8px;
                    border-top: 8px solid transparent;
                    border-bottom: 8px solid transparent;
                    border-right: 8px solid #fff;
                }

                .donate_inner:before {
                    left: -1px;
                    border-right: 8px solid #ddd;
                }

            }

            .post-tags {
                margin: 7px 0 0 20px;
                float: left;
                a:hover{
                    color: #ff6d6d;
                }
            }
    }
</style>
