<template>
  <div class="app-footer">
    <div class="links">
      <a href="https://wpa1.qq.com/KziXGWJs?_type=wpa&qidian=true" target="_blank">技术支持</a> |
      <a href="https://bk.tencent.com/s-mart/community/" target="_blank">社区论坛</a> |
      <a href="https://bk.tencent.com/index/" target="_blank">产品官网</a>
    </div>
    <p>Copyright © 2012 Tencent BlueKing. All Rights Reserved. {{ version }}</p>
  </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue';
import { storeToRefs } from 'pinia';
import useGlobalStore from '../store/global';
const { bscpVersion } = storeToRefs(useGlobalStore());


export default defineComponent({
  name: 'AppFooter',
  setup() {
    return {
      version: bscpVersion
    };
  },
});
</script>
<style lang="scss" scoped>
.app-footer {
  font-size: 12px;
  color: #b7c0ca;
  width: 100%;
  text-align: center;
  line-height: 20px;
  padding-top: 25px;
  .links a {
    color: #3a84ff;
  }
}
</style>
