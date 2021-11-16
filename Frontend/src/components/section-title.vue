<template>
    <div class="section-title"><slot></slot></div>
</template>

<script>
    export default {
        name: "section-title"
    }
</script>

<style scoped>
    div.section-title {
        color: #757575;
        font-size: 16px;
        font-weight: 400;
        padding-bottom: 15px;
        border-bottom: 1px dashed #ECECEC;
    }
</style>
