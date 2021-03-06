import request from '@/utils/request'

export function fetchList(params) {
    return request({
        url: '/articles',
        method: 'get',
        params: params
    })
}

export function fetchArticleContent(cate,title) {
    return request({
        url: '/article/'+cate+'/'+title,
        method: 'get',
        params: {}
    })
}

export function fetchArticleInfo(cate,title) {
    return request({
        url: '/article/info/'+cate+'/'+title,
        method: 'get',
        params: {}
    })
}

export function fetchFocus() {
    return request({
        url: '/focus/list',
        method: 'get',
        params: {}
    })
}

export function fetchCategory() {
    return request({
        url: '/category',
        method: 'get',
        params: {}
    })
}

export function fetchFriend() {
    return request({
        url: '/friend',
        method: 'get',
        params: {}
    })
}

export function fetchSocial() {
    return request({
        url: '/social',
        method: 'get',
        params: {}
    });
}

export function fetchSiteInfo() {
    return request({
        url: '/site',
        method: 'get',
        params: {}
    })
}

export function fetchComment() {
    return request({
        url: '/comment',
        method: 'get',
        params: {}
    })
}
