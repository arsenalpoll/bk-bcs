<template>
  <div class="group-select-wrapper">
    <h3 class="title">选择上线实例范围</h3>
    <div class="select-group-radius">
      <bk-radio-group :model-value="type" @change="handleTypeChange">
        <bk-radio label="all">
          全部实例上线
        </bk-radio>
        <bk-radio label="select">
          选择分组实例上线
          <GroupTree
            v-if="type === 'select'"
            :group-list="props.groupList"
            :group-list-loading="props.groupListLoading"
            :version-list="props.versionList"
            :version-list-loading="props.versionListLoading"
            :disabled="props.disabled"
            :value="selectedGroup"
            @change="handleSelectGroup">
          </GroupTree>
        </bk-radio>
        <bk-radio label="exclude" :disabled="props.versionStatus === 'not_released'">
          排除分组实例上线
          <GroupTree
            v-if="type === 'exclude'"
            :group-list="props.groupList.filter(item => item.release_id !== 0)"
            :group-list-loading="groupListLoading"
            :version-list="versionList"
            :version-list-loading="versionListLoading"
            :value="selectedGroup"
            @change="handleSelectGroup">
          </GroupTree>
        </bk-radio>
      </bk-radio-group>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, computed } from 'vue';
import { IGroupToPublish } from '../../../../../../../../types/group';
import { IConfigVersion } from '../../../../../../../../types/config';
import GroupTree from './tree.vue';

const props = withDefaults(defineProps<{
    groupListLoading: boolean;
    groupList: IGroupToPublish[];
    versionListLoading: boolean;
    versionList: IConfigVersion[];
    versionStatus: string;
    disabled?: number[];
    releaseType: string;
    value: IGroupToPublish[];
  }>(), {
  groupList: () => [],
  versionList: () => [],
  disabled: () => [],
  value: () => [],
});

const emits = defineEmits(['releaseTypeChange', 'change']);

const type = ref(props.releaseType);

// 节点树中选中的分组节点
// 排除分组实例上线时，选中的分组节点为：非默认节点&&不在待上线分组列表中&&在其他版本中上线的分组
const selectedGroup = computed(() => {
  if (type.value === 'exclude') {
    return props.groupList.filter(group => group.id !== 0 && group.release_id > 0 && props.value.findIndex(item => item.id === group.id) === -1);
  }
  return props.value;
});

// 切换选择分组类型
const handleTypeChange = (val: string) => {
  type.value = val;
  if (val === 'all') {
    handleSelectGroup(props.groupList);
  } else if (val === 'select') {
    const list = props.groupList.filter(group => props.disabled.includes(group.id));
    handleSelectGroup(list);
  } else {
    handleSelectGroup([]);
  }
  emits('releaseTypeChange', val);
};

const handleSelectGroup = (val: IGroupToPublish[]) => {
  if (type.value === 'exclude') {
    // 排除分组实例上线时，实际需要上线的分组为：默认分组和未被排除且已上线的分组
    const list: IGroupToPublish[] = props.groupList.filter(group => {
      // 默认分组
      if (group.id === 0) {
        return true;
      }
      return group.release_id > 0 && val.findIndex(item => item.id === group.id) === -1;
    });
    emits('change', list);
  } else {
    emits('change', val);
  }
};

defineExpose({
  selectedGroup
});

</script>
<style lang="scss" scoped>
  .group-select-wrapper {
    height: 100%
  }
  .title {
    margin: 0 0 22px;
    line-height: 19px;
    font-size: 14px;
    font-weight: 700;
    color: #63656e;
  }
  .bk-radio-group {
    display: block;
  }
  .bk-radio {
    display: block;
    margin: 0 0 12px;
    &:last-of-type {
      margin-bottom: 0;
    }
  }
</style>
