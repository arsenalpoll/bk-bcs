<template>
  <bcs-select
    class="cluster-select"
    v-model="localValue"
    :clearable="false"
    :searchable="searchable"
    :disabled="disabled"
    :popover-min-width="320"
    :remote-method="remoteMethod"
    :search-placeholder="$t('cluster.placeholder.searchCluster')"
    :size="size"
    @change="handleClusterChange">
    <bcs-option-group
      v-for="item, index in clusterData"
      :key="item.type"
      :name="item.title"
      :is-collapse="collapseList.includes(item.type)"
      :class="[
        'mt-[8px]',
        index === (clusterData.length - 1) ? 'mb-[4px]' : ''
      ]">
      <template #group-name>
        <CollapseTitle
          :title="`${item.title} (${item.list.length})`"
          :collapse="collapseList.includes(item.type)"
          @click="handleToggleCollapse(item.type)" />
      </template>
      <bcs-option
        v-for="cluster in item.list"
        :key="cluster.clusterID"
        :id="cluster.clusterID"
        :name="cluster.clusterName"
        :disabled="cluster.status !== 'RUNNING'">
        <div
          class="flex flex-col justify-center h-[50px] px-[12px]"
          v-bk-tooltips="{
            content: $t('cluster.tips.clusterStatus', [CLUSTER_MAP[cluster.status || '']]),
            disabled: cluster.status === 'RUNNING',
            placement: 'right'
          }">
          <span class="leading-6 bcs-ellipsis" v-bk-overflow-tips>{{ cluster.clusterName }}</span>
          <span
            :class="[
              'leading-4',
              {
                'text-[#979BA5]': cluster.status === 'RUNNING'
              }]">
            {{ cluster.clusterID }}
          </span>
        </div>
      </bcs-option>
    </bcs-option-group>
  </bcs-select>
</template>
<script lang="ts">
import {  defineComponent, PropType, toRefs, watch } from 'vue';

import CollapseTitle from './collapse-title.vue';
import useClusterSelector, { ClusterType } from './use-cluster-selector';

import { CLUSTER_MAP } from '@/common/constant';

export default defineComponent({
  name: 'ClusterSelect',
  components: { CollapseTitle },
  model: {
    prop: 'value',
    event: 'change',
  },
  props: {
    value: {
      type: String,
      default: '',
    },
    searchable: {
      type: Boolean,
      default: true,
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    clusterType: {
      type: [String, Array] as PropType<ClusterType|ClusterType[]>,
      default: () => ['independent', 'managed'],
    },
    size: {
      type: String,
      default: '',
    },
  },
  emits: ['change'],
  setup(props, ctx) {
    const { value, clusterType } = toRefs(props);

    const {
      localValue,
      keyword,
      collapseList,
      clusterData,
      handleToggleCollapse,
      handleClusterChange,
    } = useClusterSelector(ctx.emit, value.value, clusterType.value);

    watch(value, (v) => {
      localValue.value = v;
    });

    // 远程搜索
    const remoteMethod = (searhcKey) => {
      keyword.value = searhcKey;
    };

    return {
      localValue,
      collapseList,
      clusterData,
      CLUSTER_MAP,
      handleToggleCollapse,
      remoteMethod,
      handleClusterChange,
    };
  },
});
</script>
<style lang="postcss" scoped>
.cluster-select {
    width: 254px;
    &:not(.is-disabled) {
      background-color: #fff;
    }
}
/deep/ .bk-option-group-name {
  border-bottom: 0 !important;
}
.bk-options .bk-option:first-child {
  margin-top: 0;
}
.bk-options .bk-option:last-child {
  margin-bottom: 0;
}
</style>
