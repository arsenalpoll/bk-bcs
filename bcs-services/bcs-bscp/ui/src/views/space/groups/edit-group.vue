<template>
  <bk-dialog
    ext-cls="edit-group-dialog"
    confirm-text="提交"
    :width="640"
    :is-show="props.show"
    :esc-close="false"
    :quick-close="false"
    :close-icon="false"
    :is-loading="pending"
    @closed="handleClose"
    @confirm="handleConfirm"
  >
    <div class="group-edit-content">
      <section class="group-form-wrapper">
        <div class="dialog-title">编辑分组</div>
        <div class="group-edit-form">
          <group-edit-form ref="groupFormRef" :group="groupData" @change="updateData"></group-edit-form>
        </div>
      </section>
    </div>
  </bk-dialog>
</template>
<script setup lang="ts">
import { ref, watch } from 'vue';
import { storeToRefs } from 'pinia';
import useGlobalStore from '../../../store/global';
import { IGroupEditing, IGroupRuleItem, IGroupItem } from '../../../../types/group';
import { updateGroup } from '../../../api/group';
import groupEditForm from './components/group-edit-form.vue';

const { spaceId } = storeToRefs(useGlobalStore());

const props = defineProps<{
  show: boolean;
  group: IGroupItem;
}>();

const emits = defineEmits(['update:show', 'reload']);

const groupData = ref<IGroupEditing>({
  id: 0,
  name: '',
  public: true,
  bind_apps: [],
  rule_logic: 'AND',
  rules: [{ key: '', op: '', value: '' }],
});
const groupFormRef = ref();
const pending = ref(false);

watch(
  () => props.show,
  (val) => {
    if (val) {
      const { id, name, public: isPublic, bind_apps, selector } = props.group;
      groupData.value = {
        id,
        name,
        bind_apps: bind_apps.map(item => item.id),
        public: isPublic,
        rule_logic: selector.labels_and ? 'AND' : 'OR',
        rules: (selector.labels_and || selector.labels_or) as IGroupRuleItem[],
      };
    }
  },
);

// 修改分组信息
const updateData = (data: IGroupEditing) => {
  groupData.value = data;
};

// 保存
const handleConfirm = async () => {
  const result = await groupFormRef.value.validate();
  if (!result) {
    return;
  }
  pending.value = true;
  try {
    const { id, name, public: isPublic, bind_apps, rule_logic, rules } = groupData.value;
    const params = {
      name,
      public: isPublic,
      bind_apps: isPublic ? [] : bind_apps,
      selector: rule_logic === 'AND' ? { labels_and: rules } : { labels_or: rules },
    };
    await updateGroup(spaceId.value, id as number, params);
    handleClose();
    emits('reload');
  } catch (e) {
    console.error(e);
  } finally {
    pending.value = false;
  }
};

const handleClose = () => {
  emits('update:show', false);
};
</script>
<style lang="scss" scoped>
.group-edit-content {
  .group-edit-form {
    padding: 0 24px;
    min-height: 268px;
    max-height: 386px;
    overflow: auto;
  }
  .rule-detail-wrapper {
    width: 320px;
    background: #f5f7fa;
  }
  .dialog-title {
    margin: 16px 0 24px;
    padding: 0 24px;
    font-size: 20px;
    line-height: 28px;
    color: #313238;
  }
}
</style>
<style lang="scss">
.edit-group-dialog.bk-modal-wrapper {
  .bk-dialog-header {
    display: none;
  }
  .bk-modal-content {
    padding: 0;
  }
}
</style>
