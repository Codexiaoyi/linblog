<template>
    <article class="post post-list">
        <div class="post-entry">
            <div class="feature">
                <router-link :to="`/article/${post.category}/${post.title}`">
                    <img :src="post.banner"/>
                </router-link>
            </div>
            <h1 class="entry-title">
                <router-link :to="`/article/${post.category}/${post.title}`"><span v-if="post.isTop" style="color:#ff6d6d;font-weight:600">[置顶] </span>{{post.title}}</router-link>
            </h1>
            <div class="p-time">
                <i class="iconfont iconmeditor-time"></i> {{post.pubTime}}<i v-if="post.isHot" class="iconfont iconfire" style="margin-left: 5px;color: #ff6d6d;"></i>
            </div>
            <p class="summary">{{post.summary}}</p>
            <footer class="entry-footer">
                <div class="post-more">
                    <router-link :to="`/article/${post.category}/${post.title}`">
                        <svg t="1636960677418" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1565" width="20" height="20"><path d="M506.545121 541.803147m-432.777817 0a432.777817 432.777817 0 1 0 865.555634 0 432.777817 432.777817 0 1 0-865.555634 0Z" fill="#E9EAEB" p-id="1566"></path><path d="M295.936217 162.634869h551.484654v658.703733H295.936217z" fill="#FFFFFF" p-id="1567"></path><path d="M862.737882 836.655613H341.88725V668.168489h30.634023v137.853101h459.582586V239.219925H372.521273v61.268045h-30.634023V208.585903h520.850632v628.06971z" fill="#2A5082" p-id="1568"></path><path d="M847.348621 821.338602L555.602907 928.55768h-14.594511V55.415791h14.594511l291.745714 107.219078v658.703733z" fill="#FFFFFF" p-id="1569"></path><path d="M558.348409 943.874691h-32.657024V40.098779h32.657024L862.737882 151.725111v680.089748zM556.325407 72.250053v839.690115l275.706202-101.150074V173.400127zM188.717138 806.02159h168.487124v30.634023H188.717138zM862.737882 806.02159h61.268045v30.634023h-61.268045zM954.639949 806.02159h45.951034v30.634023h-45.951034z" fill="#2A5082" p-id="1570"></path><path d="M464.42334 499.681366L265.302194 300.48797v107.291329H35.547026v183.804134h229.755168v107.219079l199.121146-199.121146z" fill="#A3D4FF" p-id="1571"></path><path d="M249.985183 735.794539V606.900445H20.157765V392.390037H249.985183V263.495943l236.113173 236.185423zM50.575037 576.266422h229.755168V661.810485l162.418119-162.129119-162.129118-162.201369v85.544063H50.575037zM616.003951 567.740916l-27.38277-13.65526 13.65526-27.45502V457.342835l-13.65526-27.45502 27.38277-13.65526 16.906513 33.813025V533.927891l-16.906513 33.813025zM127.449093 806.02159h30.634023v30.634023h-30.634023zM311.253228 269.853948h-30.634022V147.317858h260.38919v30.634022h-229.755168v91.902068zM280.619206 729.436534h30.634022v91.902068h-30.634022z" fill="#2A5082" p-id="1572"></path></svg>
                    </router-link>
                </div>
                <div class="info-meta">
                    <div class="comnum">
                        <span>
                            <i class="iconfont iconcomment"></i>
                            <a href="https://zhebk.cn/Web/Akina.html">{{post.commentsCount}} 条评论</a>
                        </span>
                    </div>
                    <div class="views">
                        <span><i class="iconfont iconeyes"></i>{{post.viewsCount}} 热度</span>
                    </div>
                </div>
            </footer>
        </div>
        <hr/>
    </article>
</template>

<script>

    export default {
        name: "post",
        props: {
            post: {
                type: Object
            }
        }
    }
</script>

<style scoped lang="less">
    .post {
        margin: 0 0 4% 0;
        position: relative;
    }

    .post-entry {
        .feature {
            position: absolute;
            margin-top: 10px;

            img {
                width: 100px;
                height: 100px;
                object-fit: cover;
                border-radius: 50%;
                padding: 2px;
                border: 1px solid #DADADA;
                position: relative;
                transition: all 0.2s linear;
                overflow: hidden;
                &:hover {
                    transform: scale(1.1, 1.1);
                    filter: contrast(130%);
                }
            }
        }

        .entry-title {
            font-size: 21px;
            font-weight: 600;
            line-height: 50px;
            margin: 0 0 0 17%;
            position: relative;
            z-index: 1;
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
            width: 70%;
            color: #737373;

            & a:hover {
                color: #8fd0cc;
            }
        }

        .p-time {
            position: absolute;
            right: 0;
            top: 16px;
            font-size: 12px;
            color: #989898;
            letter-spacing: 1px;
            font-family: din, 'Hiragino Sans GB', 'Microsoft Yahei', Arial, sans-serif;
            display: flex;
            align-items: center;
        }

        p.summary {
            min-height: 60px;
            margin: 0 0 0 17%;
            font-size: 15px;
            color: #6f6f6f;
            letter-spacing: 1px;
            line-height: 30px;
        }

        footer.entry-footer {
            margin: 0 0 0 17%;
            list-style: none;

            .post-more {
                margin-top: 10px;
                text-align: right;
                color: #A0DAD0;

                a:hover {
                    color: #737373;
                }
            }

            .info-meta {
                margin-top: 10px;
                font-family: din, 'Hiragino Sans GB', 'Microsoft Yahei', Arial, sans-serif;
                position: absolute;
                top: 20px;
                opacity: 0;
                padding-top: 8px;
                border-top: 1px solid #ddd;
                -webkit-transform: translate3d(-150px, 0, 0);
                transform: translate3d(-150px, 0, 0);
                visibility: hidden;
                transition: .7s all ease;
                -webkit-transition: .6s all ease;
                -moz-transition: .6s all linear;
                -o-transition: .6s all ease;
                -ms-transition: .6s all ease;

                span, a {
                    color: #B3B3B3;
                    font-size: 12px;
                }

                i {
                    margin-top: 3px;
                    margin-right: 10px;
                    float: left
                }
            }

            .comnum {
                float: left;
            }
        }

    }

    .post-entry:hover footer.entry-footer .info-meta {
        -webkit-transform: translate3d(-230px, 0, 0);
        transform: translate3d(-230px, 0, 0);
        opacity: 1;
        visibility: visible;
    }

    .post hr {
        width: 30%;
        height: 1px;
        margin: 0 auto;
        border: 0;
        background: #EFEFEF;
    }

    /*******/
    @media (max-width: 1060px) {
        .info-meta{
            display: none;
        }
    }
    @media (max-width: 800px) {
        .post-entry {
            .feature img {
                width: 50px;
                height: 50px;
            }
            .entry-title{
                font-size: 16px;
                line-height: 30px;
                margin: 0 0 0 65px;
            }
            .p-time {
                position: relative;
                margin: -15px 0 0 65px;
            }
            p.summary {
                margin: 20px 0 0 65px;
                font-size: 14px;
                height: 30px;
                overflow: hidden;
            }
            .post-more{
                display: none;
            }
        }
        .post-list hr {
            margin-top: 20px;
        }
    }

    /******/
</style>
